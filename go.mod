module github.com/cert-manager/cert-manager

go 1.17

require (
	github.com/Azure/azure-sdk-for-go v56.2.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.19
	github.com/Azure/go-autorest/autorest/adal v0.9.14
	github.com/Azure/go-autorest/autorest/to v0.4.0
	github.com/Venafi/vcert/v4 v4.14.3
	github.com/akamai/AkamaiOPEN-edgegrid-golang v1.1.1
	github.com/aws/aws-sdk-go v1.40.21
	github.com/cloudflare/cloudflare-go v0.20.0
	github.com/cpu/goacmedns v0.1.1
	github.com/digitalocean/godo v1.65.0
	github.com/go-logr/logr v1.2.0
	github.com/google/gofuzz v1.2.0
	github.com/googleapis/gnostic v0.5.5
	github.com/hashicorp/vault/api v1.1.1
	github.com/hashicorp/vault/sdk v0.2.1
	github.com/kr/pretty v0.3.0
	github.com/miekg/dns v1.1.34
	github.com/mitchellh/go-homedir v1.1.0
	github.com/munnerz/crd-schema-fuzz v1.0.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.17.0
	github.com/pavel-v-chernykh/keystore-go/v4 v4.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/sergi/go-diff v1.2.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	gomodules.xyz/jsonpatch/v2 v2.2.0
	google.golang.org/api v0.53.0
	helm.sh/helm/v3 v3.7.1
	k8s.io/api v0.23.1
	k8s.io/apiextensions-apiserver v0.23.1
	k8s.io/apimachinery v0.23.1
	k8s.io/apiserver v0.23.1
	k8s.io/cli-runtime v0.23.1
	k8s.io/client-go v0.23.1
	k8s.io/code-generator v0.23.1
	k8s.io/component-base v0.23.1
	k8s.io/klog/v2 v2.30.0
	k8s.io/kube-aggregator v0.23.1
	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65
	k8s.io/kubectl v0.23.1
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-tools v0.7.0
	sigs.k8s.io/gateway-api v0.3.0
	sigs.k8s.io/structured-merge-diff/v4 v4.2.0
	sigs.k8s.io/yaml v1.3.0
	software.sslmate.com/src/go-pkcs12 v0.0.0-20210415151418-c5206de65a78
)

require (
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Microsoft/hcsshim v0.9.2 // indirect
	github.com/Nvveen/Gotty v0.0.0-20120604004816-cd527374f1e5 // indirect
	github.com/containerd/cgroups v1.0.2 // indirect
	github.com/containerd/containerd v1.5.9 // indirect
	github.com/containerd/continuity v0.2.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.1 // indirect
	github.com/docker/docker v17.12.1-ce+incompatible // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gotestyourself/gotestyourself v2.2.0+incompatible // indirect
	github.com/klauspost/compress v1.14.1 // indirect
	github.com/opencontainers/runc v1.1.0 // indirect
	google.golang.org/genproto v0.0.0-20220118154757-00ab72f36ad5 // indirect
	google.golang.org/grpc v1.43.0 // indirect
)

replace (
	golang.org/x/net => golang.org/x/net v0.0.0-20210224082022-3d97a244fca7

	// Update gengo to ensure we have the --trim-path-prefix feature in code-generator tools.
	k8s.io/gengo => k8s.io/gengo v0.0.0-20211115164449-b448ea381d54
)
