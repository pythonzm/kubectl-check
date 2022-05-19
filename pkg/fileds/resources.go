package fileds

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"kubectl-check/pkg/base"
)

type ResourcesField struct {
	base.Field
}

func (r *ResourcesField) Value() ([][]string, error) {
	result := make([][]string, 0)

	resource, err := r.GetResources()
	if err != nil {
		return nil, err
	}

	if len(resource.Deployments) != 0 {
		for _, deployment := range resource.Deployments {
			containers := deployment.Spec.Template.Spec.Containers
			var resources string
			for _, container := range containers {
				resources += fmt.Sprintf("{container: \"%v\", resources: %v}\n", container.Name, getResourcesString(container.Resources))
			}
			result = append(result, []string{"Deployment", deployment.Name, "resources", resources})
		}
	}

	if len(resource.StatefulSets) != 0 {
		for _, s := range resource.StatefulSets {
			containers := s.Spec.Template.Spec.Containers
			var resources string
			for _, container := range containers {
				resources += fmt.Sprintf("{container: \"%v\", resources: %v}\n", container.Name, getResourcesString(container.Resources))
			}
			result = append(result, []string{"StatefulSet", s.Name, "resources", resources})
		}
	}

	if len(resource.DaemonSets) != 0 {
		for _, s := range resource.DaemonSets {
			containers := s.Spec.Template.Spec.Containers
			var resources string
			for _, container := range containers {
				resources += fmt.Sprintf("{container: \"%v\", resources: %v}\n", container.Name, getResourcesString(container.Resources))
			}
			result = append(result, []string{"DaemonSet", s.Name, "resources", resources})
		}
	}

	if len(resource.CronJobs) != 0 {
		for _, cronjob := range resource.CronJobs {
			containers := cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers
			var resources string
			for _, container := range containers {
				resources += fmt.Sprintf("{container: \"%v\", resources: %v}\n", container.Name, getResourcesString(container.Resources))
			}
			result = append(result, []string{"CronJob", cronjob.Name, "resources", resources})
		}
	}

	return result, nil
}

func getResourcesString(r v1.ResourceRequirements) string {
	requests := r.Requests
	requestCPU := fmt.Sprintf("%dm", requests.Cpu().MilliValue())
	requestMem := fmt.Sprintf("%dMi", requests.Memory().Value()/1024/1024)
	limits := r.Limits
	limitCPU := fmt.Sprintf("%dm", limits.Cpu().MilliValue())
	limitMem := fmt.Sprintf("%dMi", limits.Memory().Value()/1024/1024)

	jsonFormat := `{"limits":{"cpu":"%s","memory":"%s"},"requests":{"cpu":"%s","memory":"%s"}}`

	return fmt.Sprintf(jsonFormat, limitCPU, limitMem, requestCPU, requestMem)
}
