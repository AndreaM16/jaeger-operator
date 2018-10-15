package service

import (
	"fmt"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jaegertracing/jaeger-operator/pkg/apis/io/v1alpha1"
)

// NewCollectorService returns a new Kubernetes service for Jaeger Collector backed by the pods matching the selector
func NewCollectorService(jaeger *v1alpha1.Jaeger, selector map[string]string) *v1.Service {
	trueVar := true

	return &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       metaKind,
			APIVersion: metaAPIVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      GetNameForCollectorService(jaeger),
			Namespace: jaeger.Namespace,
			Labels:    selector,
			OwnerReferences: []metav1.OwnerReference{
				metav1.OwnerReference{
					APIVersion: jaeger.APIVersion,
					Kind:       jaeger.Kind,
					Name:       jaeger.Name,
					UID:        jaeger.UID,
					Controller: &trueVar,
				},
			},
		},
		Spec: v1.ServiceSpec{
			Selector:  selector,
			ClusterIP: clusterIP,
			Ports: []v1.ServicePort{
				{
					Name: "zipkin",
					Port: 9411,
				},
				{
					Name: "c-tchan-trft",
					Port: 14267,
				},
				{
					Name: "c-binary-trft",
					Port: 14268,
				},
			},
		},
	}
}

// GetNameForCollectorService returns the service name for the collector in this Jaeger instance
func GetNameForCollectorService(jaeger *v1alpha1.Jaeger) string {
	return fmt.Sprintf("%s-collector", jaeger.Name)
}
