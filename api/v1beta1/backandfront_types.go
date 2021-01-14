/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BackAndFrontSpec defines the desired state of BackAndFront
type BackAndFrontSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of BackAndFront. Edit BackAndFront_types.go to remove/update
	//Foo string `json:"foo,omitempty"`
	//back service name
	BackName string `json:"backname"`
	//front service name
	FrontName string `json:"frontname"`
	//后端服务镜像
	BackImage string `json:"backimage"`
	//前端服务镜像
	FrontImage string `json:"frontimage"`
	//后端暴露服务的端口
	BackPort int32 `json:"backport"`
	//前端暴露服务的端口
	FrontPort int32 `json:"front_port,omitempty"`
	//pod个数
	Repicas *int32 `json:"repicas,omitempty"`
}

// BackAndFrontStatus defines the observed state of BackAndFront
type BackAndFrontStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	appsv1.DeploymentStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// BackAndFront is the Schema for the backandfronts API
type BackAndFront struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackAndFrontSpec   `json:"spec,omitempty"`
	Status BackAndFrontStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackAndFrontList contains a list of BackAndFront
type BackAndFrontList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackAndFront `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackAndFront{}, &BackAndFrontList{})
}
