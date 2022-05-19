package fileds

import (
	"fmt"
	"kubectl-check/pkg/base"
)

type ImageField struct {
	base.Field
}

func (i *ImageField) Value() ([][]string, error) {
	result := make([][]string, 0)

	resource, err := i.GetResources()
	if err != nil {
		return nil, err
	}

	if len(resource.Deployments) != 0 {
		for _, deployment := range resource.Deployments {
			containers := deployment.Spec.Template.Spec.Containers
			var images string
			for _, container := range containers {
				images += fmt.Sprintf("{container: \"%v\", image: \"%v\"}\n", container.Name, container.Image)
			}
			result = append(result, []string{"Deployment", deployment.Name, "image", images})
		}
	}

	if len(resource.StatefulSets) != 0 {
		for _, s := range resource.StatefulSets {
			containers := s.Spec.Template.Spec.Containers
			var images string
			for _, container := range containers {
				images += fmt.Sprintf("{container: \"%v\", image: \"%v\"}\n", container.Name, container.Image)
			}
			result = append(result, []string{"StatefulSet", s.Name, "image", images})
		}
	}

	if len(resource.DaemonSets) != 0 {
		for _, s := range resource.DaemonSets {
			containers := s.Spec.Template.Spec.Containers
			var images string
			for _, container := range containers {
				images += fmt.Sprintf("{container: \"%v\", image: \"%v\"}\n", container.Name, container.Image)
			}
			result = append(result, []string{"DaemonSet", s.Name, "image", images})
		}
	}

	if len(resource.CronJobs) != 0 {
		for _, cronjob := range resource.CronJobs {
			containers := cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers
			var images string
			for _, container := range containers {
				images += fmt.Sprintf("{container: \"%v\", image: \"%v\"}\n", container.Name, container.Image)
			}
			result = append(result, []string{"CronJob", cronjob.Name, "image", images})
		}
	}

	return result, nil
}
