package kube

import (
	"fmt"

	apiCoreV1 "k8s.io/api/core/v1"
)

type PodHandler struct{}

func InitPodHandler() *PodHandler {
	handler := &PodHandler{}
	return handler
}

func (h *PodHandler) OnAdd(obj interface{}) {
	pod := obj.(*apiCoreV1.Pod)
	fmt.Printf("\nPOD CREATED: %s/%s\n", pod.Namespace, pod.Name)
}

func (h *PodHandler) OnUpdate(oldObj, newObj interface{}) {
	oldPod := oldObj.(*apiCoreV1.Pod)
	newPod := newObj.(*apiCoreV1.Pod)
	fmt.Printf("\nPOD UPDATED. %s/%s %s\n",
		oldPod.Namespace, oldPod.Name, newPod.Status.Phase)
}

func (h *PodHandler) OnDelete(obj interface{}) {
	pod := obj.(*apiCoreV1.Pod)
	fmt.Printf("\nPOD DELETED: %s/%s\n", pod.Namespace, pod.Name)
}

func createPod() error {
	return nil
}

func deletePod() error {
	return nil
}
