package client

import (
	"github.com/pkg/errors"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
)

func NewClientSet(configFlags *genericclioptions.ConfigFlags) (*kubernetes.Clientset, error) {
	config, err := configFlags.ToRESTConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read kubeconfig")
	}

	return kubernetes.NewForConfig(config)
}
