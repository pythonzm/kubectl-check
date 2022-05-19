package base

import (
	"context"
	"github.com/pkg/errors"
	v1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Field struct {
	Name      string
	Namespace string
	Label     string
	Client    *kubernetes.Clientset
}

type Resources struct {
	Deployments  []v1.Deployment
	StatefulSets []v1.StatefulSet
	DaemonSets   []v1.DaemonSet
	CronJobs     []batchv1.CronJob
}

var ValidFields = []string{"image", "resources"}

func (f *Field) genListOptions() metav1.ListOptions {
	return metav1.ListOptions{
		LabelSelector: f.Label,
	}
}

func (f *Field) GetDeployments() ([]v1.Deployment, error) {
	list, e := f.Client.AppsV1().Deployments(f.Namespace).List(context.TODO(), f.genListOptions())
	if e != nil {
		return nil, errors.Wrap(e, "fail to list deployments")
	}
	return list.Items, nil
}

func (f *Field) GetStatefulSets() ([]v1.StatefulSet, error) {
	list, e := f.Client.AppsV1().StatefulSets(f.Namespace).List(context.TODO(), f.genListOptions())
	if e != nil {
		return nil, errors.Wrap(e, "fail to list sts")
	}
	return list.Items, nil
}

func (f *Field) GetDaemonSets() ([]v1.DaemonSet, error) {
	list, e := f.Client.AppsV1().DaemonSets(f.Namespace).List(context.TODO(), f.genListOptions())
	if e != nil {
		return nil, errors.Wrap(e, "fail to list ds")
	}
	return list.Items, nil
}

func (f *Field) GetCronJobs() ([]batchv1.CronJob, error) {
	list, err := f.Client.BatchV1().CronJobs(f.Namespace).List(context.TODO(), f.genListOptions())
	if err != nil {
		return nil, errors.Wrap(err, "fail to list cronjob")
	}
	return list.Items, nil
}

func (f *Field) GetResources() (Resources, error) {
	deployments, err := f.GetDeployments()
	if err != nil {
		return Resources{}, err
	}

	sts, err := f.GetStatefulSets()
	if err != nil {
		return Resources{}, err
	}

	daemonsets, err := f.GetDaemonSets()
	if err != nil {
		return Resources{}, err
	}

	cronjobs, err := f.GetCronJobs()
	if err != nil {
		return Resources{}, err
	}
	return Resources{
		Deployments:  deployments,
		StatefulSets: sts,
		DaemonSets:   daemonsets,
		CronJobs:     cronjobs,
	}, nil
}
