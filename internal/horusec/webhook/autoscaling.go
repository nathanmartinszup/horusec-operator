package webhook

import (
	autoScalingV2beta2 "k8s.io/api/autoscaling/v2beta2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
)

// nolint:funlen // constructor is required all data
func NewAutoscaling(resource *v2alpha1.HorusecPlatform) autoScalingV2beta2.HorizontalPodAutoscaler {
	autoScaling := resource.GetWebhookAutoscaling()
	if !autoScaling.Enabled {
		return autoScalingV2beta2.HorizontalPodAutoscaler{}
	}
	return autoScalingV2beta2.HorizontalPodAutoscaler{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resource.GetWebhookName(),
			Namespace: resource.GetNamespace(),
			Labels:    resource.GetWebhookLabels(),
		},
		Spec: autoScalingV2beta2.HorizontalPodAutoscalerSpec{
			MinReplicas: autoScaling.MinReplicas,
			MaxReplicas: autoScaling.MaxReplicas,
			ScaleTargetRef: autoScalingV2beta2.CrossVersionObjectReference{
				Kind:       "Deployment",
				Name:       "webhook",
				APIVersion: "apps/v1",
			},
			Metrics: []autoScalingV2beta2.MetricSpec{
				{
					Type: "Resource",
					Resource: &autoScalingV2beta2.ResourceMetricSource{
						Name: "cpu",
						Target: autoScalingV2beta2.MetricTarget{
							AverageUtilization: autoScaling.TargetCPU,
						},
					},
				},
				{
					Type: "Resource",
					Resource: &autoScalingV2beta2.ResourceMetricSource{
						Name: "memory",
						Target: autoScalingV2beta2.MetricTarget{
							AverageUtilization: autoScaling.TargetMemory,
						},
					},
				},
			},
		},
	}
}
