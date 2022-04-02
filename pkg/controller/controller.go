/*
Copyright 2020 The cert-manager Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	logf "github.com/cert-manager/cert-manager/pkg/logs"
	"github.com/cert-manager/cert-manager/pkg/metrics"
)

type runFunc func(context.Context)

type runDurationFunc struct {
	fn       runFunc
	duration time.Duration
}

type queueingController interface {
	Register(*Context) (workqueue.RateLimitingInterface, []cache.InformerSynced, error)
	ProcessItem(ctx context.Context, key string) error
}

func NewController(
	ctx context.Context,
	name string,
	metrics *metrics.Metrics,
	syncFunc func(ctx context.Context, key string) error,
	mustSync []cache.InformerSynced,
	runDurationFuncs []runDurationFunc,
	queue workqueue.RateLimitingInterface,
) Interface {
	return &controller{
		ctx:              ctx,
		name:             name,
		metrics:          metrics,
		syncHandler:      syncFunc,   // 注册controller 队列的处理函数
		mustSync:         mustSync,
		runDurationFuncs: runDurationFuncs,
		queue:            queue,
	}
}

type controller struct {
	// ctx is the root golang context for the controller
	ctx context.Context

	// name is the name for this controller
	name string

	// the function that should be called when an item is popped
	// off the workqueue
	// 当有事件从工作队列出栈时调用该同步函数
	syncHandler func(ctx context.Context, key string) error

	// mustSync is a slice of informers that must have synced before
	// this controller can start
	// 在 controller 启动前必须完成同步的一组informers
	mustSync []cache.InformerSynced

	// a set of functions that will be called just after controller initialisation, once.
	// 在 controller 启动前必须执行的一组informer函数
	runFirstFuncs []runFunc

	// a set of functions that should be called every duration.
	// 以固定间隔时间循环执行的一组任务
	runDurationFuncs []runDurationFunc

	// queue is a reference to the queue used to enqueue resources
	// to be processed
	// controller 的工作队列
	queue workqueue.RateLimitingInterface

	// metrics is used to expose Prometheus, shared by all controllers
	metrics *metrics.Metrics
}

// Run starts the controller loop
// 给每个controller 实例单独启动一个协程用于执行controller 注册的启动Run 函数
func (c *controller) Run(workers int, stopCh <-chan struct{}) error {
	ctx, cancel := context.WithCancel(c.ctx)
	defer cancel()
	log := logf.FromContext(ctx)

	log.V(logf.DebugLevel).Info("starting control loop")
	// wait for all the informer caches we depend on are synced
	// 等待所有mustSync 队列中的informers 完成同步
	if !cache.WaitForCacheSync(stopCh, c.mustSync...) {
		return fmt.Errorf("error waiting for informer caches to sync")
	}

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// worker 启动执行，在loop 中进行工作队列的调谐处理
			c.worker(ctx)
		}()
	}

	// 执行构建器with 方法中注册的定时任务
	for _, f := range c.runFirstFuncs {
		f(ctx)
	}

	for _, f := range c.runDurationFuncs {
		go wait.Until(func() { f.fn(ctx) }, f.duration, stopCh)
	}

	<-stopCh
	log.V(logf.InfoLevel).Info("shutting down queue as workqueue signaled shutdown")
	c.queue.ShutDown()
	log.V(logf.DebugLevel).Info("waiting for workers to exit...")
	wg.Wait()
	log.V(logf.DebugLevel).Info("workers exited")
	return nil
}

// 在一个loop 循环中，worker 会从controller 对应的工作队列中不断取出事件对应的index 标签，然后调用之前controller 注册过的syncHandler方法进行相应的事件处理
// 如果处理失败，worker 会重新将index 对象压入工作队列再次进行处理；如果处理成功，就将其从队列中删除
func (c *controller) worker(ctx context.Context) {
	log := logf.FromContext(c.ctx)

	log.V(logf.DebugLevel).Info("starting worker")
	for {
		obj, shutdown := c.queue.Get()
		if shutdown {
			break
		}

		var key string
		// use an inlined function so we can use defer
		func() {
			defer c.queue.Done(obj)
			var ok bool
			if key, ok = obj.(string); !ok {
				return
			}
			log := log.WithValues("key", key)
			log.V(logf.DebugLevel).Info("syncing item")

			// Increase sync count for this controller
			c.metrics.IncrementSyncCallCount(c.name)

			// 由各controller 实现的ProcessItem方法
			err := c.syncHandler(ctx, key)
			if err != nil {
				if strings.Contains(err.Error(), genericregistry.OptimisticLockErrorMsg) {
					log.Info("re-queuing item due to optimistic locking on resource", "error", err.Error())
				} else {
					log.Error(err, "re-queuing item due to error processing")
				}

				c.queue.AddRateLimited(obj)
				return
			}
			log.V(logf.DebugLevel).Info("finished processing work item")
			c.queue.Forget(obj)
		}()
	}
	log.V(logf.DebugLevel).Info("exiting worker loop")
}
