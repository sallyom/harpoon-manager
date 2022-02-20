module github.com/containers/harpoon

go 1.16

require (
	github.com/containers/common v0.47.4
	github.com/containers/podman/v4 v4.0.0
	github.com/go-co-op/gocron v1.12.0
	github.com/go-git/go-git/v5 v5.4.2
	github.com/kcp-dev/kcp v0.0.0-20220217152528-4feeabd4b0f0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.18.1
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	k8s.io/apimachinery v0.23.0
	k8s.io/apiserver v0.23.0
	k8s.io/client-go v0.23.0
	k8s.io/component-base v0.23.0
	k8s.io/klog/v2 v2.30.0
	k8s.io/kube-aggregator v0.23.0 // indirect
	sigs.k8s.io/controller-runtime v0.11.1
)

replace (
	k8s.io/api => github.com/kcp-dev/kubernetes/staging/src/k8s.io/api v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/apiextensions-apiserver => github.com/kcp-dev/kubernetes/staging/src/k8s.io/apiextensions-apiserver v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/apimachinery => github.com/kcp-dev/kubernetes/staging/src/k8s.io/apimachinery v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/apiserver => github.com/kcp-dev/kubernetes/staging/src/k8s.io/apiserver v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/cli-runtime => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cli-runtime v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/client-go => github.com/kcp-dev/kubernetes/staging/src/k8s.io/client-go v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/cloud-provider => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cloud-provider v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/cluster-bootstrap => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cluster-bootstrap v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/code-generator => github.com/kcp-dev/kubernetes/staging/src/k8s.io/code-generator v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/component-base => github.com/kcp-dev/kubernetes/staging/src/k8s.io/component-base v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/component-helpers => github.com/kcp-dev/kubernetes/staging/src/k8s.io/component-helpers v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/controller-manager => github.com/kcp-dev/kubernetes/staging/src/k8s.io/controller-manager v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/cri-api => github.com/kcp-dev/kubernetes/staging/src/k8s.io/cri-api v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/csi-translation-lib => github.com/kcp-dev/kubernetes/staging/src/k8s.io/csi-translation-lib v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kube-aggregator => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-aggregator v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kube-controller-manager => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-controller-manager v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kube-proxy => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-proxy v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kube-scheduler => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kube-scheduler v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kubectl => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kubectl v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kubelet => github.com/kcp-dev/kubernetes/staging/src/k8s.io/kubelet v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/kubernetes => github.com/kcp-dev/kubernetes v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/legacy-cloud-providers => github.com/kcp-dev/kubernetes/staging/src/k8s.io/legacy-cloud-providers v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/metrics => github.com/kcp-dev/kubernetes/staging/src/k8s.io/metrics v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/mount-utils => github.com/kcp-dev/kubernetes/staging/src/k8s.io/mount-utils v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/pod-security-admission => github.com/kcp-dev/kubernetes/staging/src/k8s.io/pod-security-admission v0.0.0-20220210121228-50b2affb6ca1
	k8s.io/sample-apiserver => github.com/kcp-dev/kubernetes/staging/src/k8s.io/sample-apiserver v0.0.0-20220210121228-50b2affb6ca1
)
