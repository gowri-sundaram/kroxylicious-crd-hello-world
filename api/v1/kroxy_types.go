package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type Kroxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KroxySpec `json:"spec,omitempty"`
}

type KroxySpec struct {
	Image   string       `json:"image"`
	Volumes []VolumeSpec `json:"volumes,omitempty"`
	Ports   []PortSpec   `json:"ports,omitempty"`
}

type VolumeSpec struct {
	Name      string `json:"name"`
	HostPath  string `json:"hostPath"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
}

type PortSpec struct {
	ContainerPort int32 `json:"containerPort"`
	HostPort      int32 `json:"hostPort"`
}

type KroxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Kroxy `json:"items"`
}

var GroupVersion = schema.GroupVersion{Group: "example.com", Version: "v1"}

func (in *Kroxy) DeepCopyObject() runtime.Object {
	out := *in
	return &out
}

func (in *KroxyList) DeepCopyObject() runtime.Object {
	out := *in
	return &out
}
