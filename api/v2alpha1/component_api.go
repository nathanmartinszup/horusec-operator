package v2alpha1

import (
	"fmt"
)

func (h *HorusecPlatform) GetAPIComponent() ExposableComponent {
	return h.Spec.Components.API
}

func (h *HorusecPlatform) GetAPIAutoscaling() Autoscaling {
	return h.GetAPIComponent().Pod.Autoscaling
}

func (h *HorusecPlatform) GetAPIName() string {
	name := h.GetAPIComponent().Name
	if name == "" {
		return fmt.Sprintf("%s-api", h.GetName())
	}
	return name
}

func (h *HorusecPlatform) GetAPIPath() string {
	path := h.GetAPIComponent().Ingress.Path
	if path == "" {
		return "/api"
	}
	return path
}

func (h *HorusecPlatform) GetAPIPortHTTP() int {
	port := h.GetAPIComponent().Port.HTTP
	if port == 0 {
		return 8000
	}
	return port
}

func (h *HorusecPlatform) GetApiLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       h.GetName(),
		"app.kubernetes.io/component":  "api",
		"app.kubernetes.io/managed-by": "horusec",
	}
}

func (h *HorusecPlatform) GetAPIReplicaCount() *int32 {
	if !h.GetAPIAutoscaling().Enabled {
		count := h.GetAPIComponent().ReplicaCount
		return &count
	}
	return nil
}

func (h *HorusecPlatform) GetAPIDefaultURL() string {
	return fmt.Sprintf("http://%s:%v", h.GetAPIName(), h.GetAPIPortHTTP())
}

func (h *HorusecPlatform) GetAPIRegistry() string {
	registry := h.GetAPIComponent().Container.Image.Registry
	if registry == "" {
		return "docker.io/"
	}
	return registry
}

func (h *HorusecPlatform) GetAPIRepository() string {
	repository := h.GetAPIComponent().Container.Image.Repository
	if repository == "" {
		return "horuszup/horusec-api"
	}
	return repository
}

func (h *HorusecPlatform) GetAPITag() string {
	tag := h.GetAPIComponent().Container.Image.Tag
	if tag == "" {
		return h.GetLatestVersion()
	}
	return tag
}

func (h *HorusecPlatform) GetAPIImage() string {
	return fmt.Sprintf("%s%s:%s", h.GetAPIRegistry(), h.GetAPIRepository(), h.GetAPITag())
}

func (h *HorusecPlatform) GetAPIHost() string {
	host := h.Spec.Components.API.Ingress.Host
	if host == "" {
		return "api.local"
	}

	return host
}

func (h *HorusecPlatform) IsAPIIngressEnabled() bool {
	enabled := h.Spec.Components.API.Ingress.Enabled
	if enabled == nil {
		return true
	}

	return *enabled
}
