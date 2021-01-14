package deployment

import(
	devopsv1beta1 "gxin0426/springboot-vue-operator/api/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	corev1 "k8s.io/api/core/v1"
)

func NewBackDeploy(bf *devopsv1beta1.BackAndFront) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta:   metav1.TypeMeta{
			Kind: "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: bf.Name,
			Namespace: bf.Namespace,
			Labels: map[string]string{"app": bf.Name},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(bf, schema.GroupVersionKind{
					Group: devopsv1beta1.GroupVersion.Group,
					Version: devopsv1beta1.GroupVersion.Version,
					Kind: "BackAndFront",
				}),
			},
		},
		Spec:       appsv1.DeploymentSpec{
			Replicas: bf.Spec.Repicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": bf.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: bf.Name,
					Labels: map[string]string{"app": bf.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						corev1.Container{
							Name: "backend",
							Image: bf.Spec.BackImage,
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									ContainerPort: bf.Spec.BackPort,
								},
							},
						},
					},
				},
			},
		},
		//Status:     appsv1.DeploymentStatus{},
	}
}


func NewFrontDeploy(bf *devopsv1beta1.BackAndFront) *appsv1.Deployment {
	return &appsv1.Deployment{
		TypeMeta:   metav1.TypeMeta{
			Kind: "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: bf.Name,
			Namespace: bf.Namespace,
			Labels: map[string]string{"app": bf.Name},
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(bf, schema.GroupVersionKind{
					Group:   devopsv1beta1.GroupVersion.Group,
					Version: devopsv1beta1.GroupVersion.Version,
					Kind:    "BackAndFront",
				}),
			},

		},
		Spec:       appsv1.DeploymentSpec{
			Replicas: bf.Spec.Repicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": bf.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: bf.Name,
					Labels: map[string]string{"app": bf.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						corev1.Container{
							Name: "frontend",
							Image: bf.Spec.FrontImage,
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									ContainerPort: bf.Spec.BackPort,
								},
							},
						},
					},
				},
			},
		},
		Status:     appsv1.DeploymentStatus{},
	}
}