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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KubernetesDbaasSpec defines the desired state of KubernetesDbaas
type KubernetesDbaasSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DbName is the name of the database
	DbName string `json:"dbname,omitempty"`
	// DbStage is the stage of the database (e.g. test, prod)
	DbStage string `json:"dbstage,omitempty"`
	// DbmsType is the type of DBMS (e.g. mssql, psql...)
	DbmsType string `json:"dbtype,omitempty"`
}

// KubernetesDbaasStatus defines the observed state of KubernetesDbaas
type KubernetesDbaasStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KubernetesDbaas is the Schema for the kubernetesdbaas API
type KubernetesDbaas struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubernetesDbaasSpec   `json:"spec,omitempty"`
	Status KubernetesDbaasStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KubernetesDbaasList contains a list of KubernetesDbaas
type KubernetesDbaasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubernetesDbaas `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KubernetesDbaas{}, &KubernetesDbaasList{})
}