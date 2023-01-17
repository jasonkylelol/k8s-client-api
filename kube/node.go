package kube

import (
	"fmt"
	"strings"

	apiCoreV1 "k8s.io/api/core/v1"
)

type NodeHandler struct{}

func InitNodeHandler() *NodeHandler {
	handler := &NodeHandler{}
	return handler
}

func (h *NodeHandler) OnAdd(obj interface{}) {
	node := obj.(*apiCoreV1.Node)
	fmt.Printf("\nNODE CREATED: %s\n", node.Name)
}

func (h *NodeHandler) OnUpdate(oldObj, newObj interface{}) {
	oldNode := oldObj.(*apiCoreV1.Node)
	newNode := newObj.(*apiCoreV1.Node)
	var bu strings.Builder
	fmt.Fprintf(&bu, "\nNODE UPDATED. %s", oldNode.Name)
	fmt.Fprintf(&bu, "\nCapacity:")
	for k, v := range newNode.Status.Capacity {
		fmt.Fprintf(&bu, " %v:%v", k, v.String())
	}
	fmt.Fprintf(&bu, "\nAllocatable:")
	for k, v := range newNode.Status.Allocatable {
		fmt.Fprintf(&bu, " %v:%v", k, v.String())
	}
	fmt.Printf("%s\n", bu.String())
}

func (h *NodeHandler) OnDelete(obj interface{}) {
	node := obj.(*apiCoreV1.Node)
	fmt.Printf("\nNODE DELETED: %s\n", node.Name)
}
