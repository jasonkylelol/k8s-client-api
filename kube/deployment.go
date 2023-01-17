package kube

import (
	"fmt"

	apiAppsV1 "k8s.io/api/apps/v1"
)

type DeploymentHandler struct{}

func InitDeploymentHandler() *DeploymentHandler {
	handler := &DeploymentHandler{}
	return handler
}

func (h *DeploymentHandler) OnAdd(obj interface{}) {
	deployment := obj.(*apiAppsV1.Deployment)
	fmt.Printf("\nDEPLOYMENT CREATED: %s/%s\n", deployment.Namespace, deployment.Name)
}

func (h *DeploymentHandler) OnUpdate(oldObj, newObj interface{}) {
	oldDeployment := oldObj.(*apiAppsV1.Deployment)
	newDeployment := newObj.(*apiAppsV1.Deployment)
	fmt.Printf("\nDEPLOYMENT UPDATED. %s/%s updated: %v available: %v\n",
		oldDeployment.Namespace, oldDeployment.Name,
		newDeployment.Status.UpdatedReplicas, newDeployment.Status.AvailableReplicas)
}

func (h *DeploymentHandler) OnDelete(obj interface{}) {
	deployment := obj.(*apiAppsV1.Deployment)
	fmt.Printf("\nDEPLOYMENT DELETED: %s/%s\n", deployment.Namespace, deployment.Name)
}

func createDeployment() error {
	return nil
}

func deleteDeployment() error {
	return nil
}
