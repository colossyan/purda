/*
Copyright 2025.

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
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PurdaSpec defines the desired state of Purda.
type PurdaSpec struct {
	// CommonSpec is the common settings for all targets
	CommonSpec CommonSpec `json:"commonSpec"`

	// RepositorySpec is the configuration for the git repository
	RepositorySpec RepositorySpec `json:"repositorySpec"`

	// Target is the target resource to be reconciled
	Target Target `json:"target"`
}

type RepositorySpec struct {
	// Owner is the owner of the repository
	Owner string `json:"owner"`

	// Name is the name of the repository
	Name string `json:"name"`
}

type CommonSpec struct{}

type Target struct {
	// HelmDeploy is the configuration on how to deploy a Helm chart
	HelmDeploy *HelmDeploy `json:"helmDeploy,omitempty"`
}

type HelmDeploy struct {
	// SourcePath is the path to the helm chart in the repository.
	// This is usually the /charts directory in the repository.
	SourcePath *string `json:"sourcePath,omitempty"`

	// ValuesOverride is a JSON patch to be applied to the HelmRelease values
	ValuesOverride *apiextensionsv1.JSON `json:"valuesOverride,omitempty"`
}

// PurdaStatus defines the observed state of Purda.
type PurdaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Purda is the Schema for the purda API.
type Purda struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PurdaSpec   `json:"spec,omitempty"`
	Status PurdaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PurdaList contains a list of Purda.
type PurdaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Purda `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Purda{}, &PurdaList{})
}
