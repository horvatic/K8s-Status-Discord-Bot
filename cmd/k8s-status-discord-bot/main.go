package main

import (
	"github.com/horvatic/k8s-status-discord-bot/pkg/report"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {

	//To run local uncomment this follow, and comment out in-cluster config
	// uses the current context in kubeconfig
	// path-to-kubeconfig -- for example, /root/.kube/config
	//config, _ := clientcmd.BuildConfigFromFlags("", "<path-to-kubeconfig>")
	//clientset, _ := kubernetes.NewForConfig(config)

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	report.ProcessReport(clientset)

}
