package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime/schema"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	apiv1 "kroxylicious-operator/api/v1"
)

type KroxyReconciler struct {
	client.Client
}

func (r *KroxyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var kroxy apiv1.Kroxy
	kroxy.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "example.com",
		Version: "v1",
		Kind:    "Kroxy",
	})
	if err := r.Get(ctx, req.NamespacedName, &kroxy); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	var volumeMounts []corev1.VolumeMount
	var volumes []corev1.Volume
	for _, v := range kroxy.Spec.Volumes {
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      v.Name,
			MountPath: v.MountPath,
			ReadOnly:  v.ReadOnly,
		})
		volumes = append(volumes, corev1.Volume{
			Name: v.Name,
			VolumeSource: corev1.VolumeSource{
				HostPath: &corev1.HostPathVolumeSource{Path: v.HostPath},
			},
		})
	}

	var ports []corev1.ContainerPort
	for _, p := range kroxy.Spec.Ports {
		ports = append(ports, corev1.ContainerPort{
			ContainerPort: p.ContainerPort,
			HostPort:      p.HostPort,
		})
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      kroxy.Name,
			Namespace: req.Namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:         kroxy.Name,
					Image:        kroxy.Spec.Image,
					VolumeMounts: volumeMounts,
					Ports:        ports,
				},
			},
			Volumes: volumes,
		},
	}

	if err := r.Create(ctx, pod); err != nil {
		return ctrl.Result{}, err
	}
	logger.Info("Created pod for Kroxy", "pod", pod.Name)
	return ctrl.Result{}, nil
}

func SetupKroxyReconciler(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&apiv1.Kroxy{}).
		Complete(&KroxyReconciler{Client: mgr.GetClient()})
}
