package check

import (
	"k8s.io/client-go/kubernetes"
	"kubectl-check/pkg/base"
	"kubectl-check/pkg/fileds"
)

type Plugin interface {
	Value() ([][]string, error)
}

func NewPlugin(name, namespace, label string, client *kubernetes.Clientset) Plugin {
	switch name {
	case "image":
		return &fileds.ImageField{
			Field: base.Field{
				Name:      name,
				Namespace: namespace,
				Label:     label,
				Client:    client,
			},
		}
	case "resources":
		return &fileds.ResourcesField{
			Field: base.Field{
				Name:      name,
				Namespace: namespace,
				Label:     label,
				Client:    client,
			},
		}
	}
	return nil
}
