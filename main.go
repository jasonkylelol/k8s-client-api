package main

import (
	"time"

	"github.com/jasonkylelol/k8s-client-api/kube"
)

func main() {
	// init kube client
	kube.InitKubeClient()

	for {
		time.Sleep(time.Second)
	}
}
