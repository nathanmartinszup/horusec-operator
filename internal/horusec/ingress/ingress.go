package ingress

import (
	"k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
	"github.com/ZupIT/horusec-operator/internal/horusec/analytic"
	"github.com/ZupIT/horusec-operator/internal/horusec/api"
	"github.com/ZupIT/horusec-operator/internal/horusec/auth"
	"github.com/ZupIT/horusec-operator/internal/horusec/core"
	"github.com/ZupIT/horusec-operator/internal/horusec/manager"
	"github.com/ZupIT/horusec-operator/internal/horusec/messages"
	"github.com/ZupIT/horusec-operator/internal/horusec/vulnerability"
	"github.com/ZupIT/horusec-operator/internal/horusec/webhook"
)

//nolint:funlen // improve in the future
func NewIngress(resource *v2alpha1.HorusecPlatform) *v1beta1.Ingress {
	pathType := v1beta1.PathTypePrefix

	return &v1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resource.GetName(),
			Namespace: resource.GetNamespace(),
			Labels:    resource.GetDefaultLabel(),
		},
		Spec: v1beta1.IngressSpec{
			Rules: []v1beta1.IngressRule{ // TODO: single rule for each host
				analytic.NewIngressRule(resource, pathType),
				api.NewIngressRule(resource, pathType),
				auth.NewIngressRule(resource, pathType),
				core.NewIngressRule(resource, pathType),
				manager.NewIngressRule(resource, pathType),
				messages.NewIngressRule(resource, pathType),
				vulnerability.NewIngressRule(resource, pathType),
				webhook.NewIngressRule(resource, pathType),
			},
			TLS: NewIngressTLS(resource),
		},
	}
}

func NewIngressTLS(resource *v2alpha1.HorusecPlatform) []v1beta1.IngressTLS {
	if !resource.Spec.Components.Analytic.Ingress.Enabled {
		return []v1beta1.IngressTLS{}
	}

	var ingressList []v1beta1.IngressTLS
	for key, value := range mapSecretsAndHosts(resource) {
		ingress := v1beta1.IngressTLS{
			Hosts:      value,
			SecretName: key,
		}

		ingressList = append(ingressList, ingress)
	}

	return ingressList
}

func mapSecretsAndHosts(resource *v2alpha1.HorusecPlatform) map[string][]string {
	ingressConfig := newIngressConfigList(resource)

	tlsMap := map[string][]string{}
	for index := range ingressConfig {
		if value, ok := tlsMap[ingressConfig[index].TLS.SecretName]; ok {
			tlsMap[ingressConfig[index].TLS.SecretName] = append(value, ingressConfig[index].Host)
		} else {
			tlsMap[ingressConfig[index].TLS.SecretName] = []string{ingressConfig[index].Host}
		}
	}

	return tlsMap
}

func newIngressConfigList(resource *v2alpha1.HorusecPlatform) []v2alpha1.Ingress {
	return []v2alpha1.Ingress{
		resource.Spec.Components.Analytic.Ingress,
		resource.Spec.Components.Api.Ingress,
		resource.Spec.Components.Auth.Ingress,
		resource.Spec.Components.Core.Ingress,
		resource.Spec.Components.Manager.Ingress,
		resource.Spec.Components.Messages.Ingress,
		resource.Spec.Components.Vulnerability.Ingress,
		resource.Spec.Components.Webhook.Ingress,
	}
}
