package harpoon

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"os"
	"time"

	kcpserver "github.com/kcp-dev/kcp/pkg/server"
	kcpserveroptions "github.com/kcp-dev/kcp/pkg/server/options"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	//genericapiserver "k8s.io/apiserver/pkg/server"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	//"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type HarpoonOpts struct {
	Client *kubernetes.Clientset
	KubeConfig string
}

func NewHarpoonOpts() *HarpoonOpts {
	return &HarpoonOpts{ KubeConfig: ""}
}

// type StartFunc func(context.Context, *ControllerContext) error
func RunHarpoon(ctx context.Context) error {
	o := NewHarpoonOpts()
	go o.startKCPServer(ctx)

	go o.startHarpoonController(ctx)

	go o.startHarpoonEngine(ctx)

	<-ctx.Done()
	return nil
}

// startKCPServer starts the APIserver and returns a kubeconfig
func (o *HarpoonOpts) startKCPServer(ctx context.Context) {
	// start KCP
	//return "/tmp/admin.kubeconfig", nil
	// Server manages the configuration and kcp api-server. It allows callers to easily use kcp
	// as a library rather than as a single binary. Using its constructor function, you can easily
	// setup a new api-server and start it:
	//
	klog.Infof("Starting KCP Server")
		klog.Infof("Started KCP Server")
		opts, err := kcpserveroptions.NewOptions().Complete()
		if err != nil {
			log.Fatalf(err.Error())
		}
		srv, err := kcpserver.NewServer(opts)
		srv.Run(ctx)
		if err = o.waitForClient(ctx); err != nil {
			log.Fatalf("error getting KubeClient: %w", err)
		}
}

func (o *HarpoonOpts) waitForClient(ctx context.Context) error {
	    if o.Client != nil {
			return nil
		}
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		kcPath := filepath.Join(dir, ".kcp", "admin.kubeconfig")
	    if err := wait.PollUntilWithContext(ctx, 30 * time.Second, func(ctx context.Context) (bool, error) {
			if _, err := os.Stat(kcPath); err != nil {
		        klog.Errorf("waiting for Kubeconfig file %s", kcPath)
				return false, nil // keep going
			}
			return true, nil
		}); err != nil {
			// nolint:nilerr
			return fmt.Errorf("did not find Kubeconfig at %s", kcPath)
		}
	    restConfig, err := clientcmd.BuildConfigFromFlags("", kcPath)
	    if err != nil {
			return err
		}
	    o.Client = kubernetes.NewForConfigOrDie(restConfig)
	    if err := wait.PollUntilWithContext(ctx, 30 * time.Second, func(ctx context.Context) (bool, error) {
	        if o.Client == nil {
		        klog.Errorf("waiting for Kubeclient")
				return false, nil // keep going
			}
			return true, nil
		}); err != nil {
			// nolint:nilerr
			return fmt.Errorf("KubeClient is nil")
		}
		return nil
	}


// startHarpoonController launches the Harpoon CustomResource controller
func (o *HarpoonOpts) startHarpoonController(ctx context.Context) {
	klog.Infof("Started Harpoon Controllers")
	if err := o.waitForClient(ctx); err != nil {
		log.Fatalf(err.Error())
	}
	nsList, err := o.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, n := range nsList.Items {
	klog.Infof("NAMESPACE: %v", n)
    }
}

// startHarpoonEngine launches the Harpoon GitOps Engine
func (o *HarpoonOpts) startHarpoonEngine(ctx context.Context) {
	klog.Infof("Started Harpoon Engine")
	if err := o.waitForClient(ctx); err != nil {
		log.Fatalf(err.Error())
	}
}

