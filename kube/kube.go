package kube

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/informers"
	appsV1 "k8s.io/client-go/informers/apps/v1"
	coreV1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClient struct {
	clientSet       *kubernetes.Clientset
	informerFactory informers.SharedInformerFactory

	deploymentInformer appsV1.DeploymentInformer
	podInformer        coreV1.PodInformer
	nodeInformer       coreV1.NodeInformer
	quotaInformer      coreV1.ResourceQuotaInformer
}

var kubeCli *KubeClient

func GetKubeClient() *KubeClient {
	return kubeCli
}

func InitKubeClient() error {
	kubeCli = &KubeClient{}
	kubeCli.initClientSet()
	kubeCli.initInformer()
	go kubeCli.run()
	return nil
}

var (
	kubeCfg = "/root/.kube/config"
)

func (k *KubeClient) initClientSet() {

	fmt.Printf("[initClientSet] kubeCfg: %v\n", kubeCfg)
	clientCfg, err := clientcmd.BuildConfigFromFlags("", kubeCfg)
	if err != nil {
		fmt.Printf("clientcmd.BuildConfigFromFlags err:%v\n", err)
		panic(err.Error())
	}
	k.clientSet, err = kubernetes.NewForConfig(clientCfg)
	if err != nil {
		fmt.Printf("kubernetes.NewForConfig err:%v\n", err)
		panic(err.Error())
	}
	fmt.Printf("[initClientSet] succeed\n")
}

func (k *KubeClient) initInformer() {
	k.informerFactory = informers.NewSharedInformerFactory(k.clientSet, 10*time.Second)
	// init deployment informer
	k.deploymentInformer = k.informerFactory.Apps().V1().Deployments()
	k.deploymentInformer.Informer().AddEventHandler(InitDeploymentHandler())
	// init pod informer
	k.podInformer = k.informerFactory.Core().V1().Pods()
	k.podInformer.Informer().AddEventHandler(InitPodHandler())
	// init node informer
	k.nodeInformer = k.informerFactory.Core().V1().Nodes()
	k.nodeInformer.Informer().AddEventHandler(InitNodeHandler())
	// init quota informer
	k.quotaInformer = k.informerFactory.Core().V1().ResourceQuotas()
	k.quotaInformer.Informer().AddEventHandler(InitQuotaHandler())
}

func (k *KubeClient) run() {
	ctx := context.Background()
	// Starts all the shared informers that have been created by the factory so far.
	k.informerFactory.Start(ctx.Done())
	// wait for the initial synchronization of the local cache.
	if !cache.WaitForCacheSync(ctx.Done(), k.deploymentInformer.Informer().HasSynced) {
		fmt.Printf("cache.WaitForCacheSync deploymentInformer failed\n")
	}
	if !cache.WaitForCacheSync(ctx.Done(), k.podInformer.Informer().HasSynced) {
		fmt.Printf("cache.WaitForCacheSync podInformer failed\n")
	}
	if !cache.WaitForCacheSync(ctx.Done(), k.nodeInformer.Informer().HasSynced) {
		fmt.Printf("cache.WaitForCacheSync nodeInformer failed\n")
	}
	if !cache.WaitForCacheSync(ctx.Done(), k.quotaInformer.Informer().HasSynced) {
		fmt.Printf("cache.WaitForCacheSync quotaInformer failed\n")
	}
}
