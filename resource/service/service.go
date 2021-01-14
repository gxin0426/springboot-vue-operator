package service

import(
	devopsv1beta1 "gxin0426/springboot-vue-operator/api/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func NewBackendService(bf *devopsv1beta1.BackAndFront) *corev1.Service {
	return &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind: "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: bf.Name,
			Namespace: bf.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(bf, schema.GroupVersionKind{
					Group:   devopsv1beta1.GroupVersion.Group,
					Version: devopsv1beta1.GroupVersion.Version,
					Kind:    "BackAndFront",
				}),
			},
		},
		Spec:       corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port: bf.Spec.BackPort,
					Name: bf.Spec.BackName,
				},
			},
			Type: "NodePort",
			//ClusterIP: string(corev1.DNSClusterFirst),
			Selector: map[string]string{
				"app": bf.Name,
			},
		},

	}
}


func NewFrontendService(bf *devopsv1beta1.BackAndFront) *corev1.Service {
	return &corev1.Service{
		TypeMeta:   metav1.TypeMeta{
			Kind: "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: bf.Name,
			Namespace: bf.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(bf, schema.GroupVersionKind{
					Group:   devopsv1beta1.GroupVersion.Group,
					Version: devopsv1beta1.GroupVersion.Version,
					Kind:    "BackAndFront",
				}),
			},
		},
		Spec:       corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port: bf.Spec.BackPort,
					Name: bf.Spec.BackName,
				},
			},
			Type: "NodePort",
			//ClusterIP: string(corev1.DNSClusterFirst),
			Selector: map[string]string{
				"app": bf.Name,
			},
		},

	}
}