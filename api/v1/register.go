package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func AddToScheme(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(GroupVersion, &Kroxy{}, &KroxyList{})
	metav1.AddToGroupVersion(scheme, GroupVersion)
	return nil
}
