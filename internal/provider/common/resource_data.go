package common

import "k8s.io/client-go/dynamic"

type ResourceData struct {
	Clientset *dynamic.DynamicClient
	Namespace string
}
