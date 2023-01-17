package kube

import (
	"fmt"

	apiCoreV1 "k8s.io/api/core/v1"
)

type QuotaHandler struct{}

func InitQuotaHandler() *QuotaHandler {
	handler := &QuotaHandler{}
	return handler
}

func (h *QuotaHandler) OnAdd(obj interface{}) {
	quota := obj.(*apiCoreV1.ResourceQuota)
	fmt.Printf("\nQUOTA CREATED: %s/%s\n", quota.Namespace, quota.Name)
}

func (h *QuotaHandler) OnUpdate(oldObj, newObj interface{}) {
	oldQuota := oldObj.(*apiCoreV1.ResourceQuota)
	newQuota := newObj.(*apiCoreV1.ResourceQuota)
	fmt.Printf("\nQUOTA UPDATED. %s/%s %s\n",
		oldQuota.Namespace, oldQuota.Name, newQuota.Status.String())
}

func (h *QuotaHandler) OnDelete(obj interface{}) {
	quota := obj.(*apiCoreV1.ResourceQuota)
	fmt.Printf("\nQUOTA DELETED: %s/%s\n", quota.Namespace, quota.Name)
}
