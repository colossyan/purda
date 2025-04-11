package common

import (
	"k8s.io/apimachinery/pkg/types"
)

type NamespacedName struct {
	// The namespace of the referenced resource
	Namespace string `json:"namespace"`

	// The name of the referenced resource
	Name string `json:"name"`
}

func (n *NamespacedName) ToNamespacedName() types.NamespacedName {
	return types.NamespacedName{
		Namespace: n.Namespace,
		Name:      n.Name,
	}
}
