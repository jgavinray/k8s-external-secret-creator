/*
Copyright 2021.

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

type SecretDescription struct {
	// Name of the secret to be stored
	Name string `json:"name"`

	// Type of secret to created (string, ed25519, rsa, etc)
	Type string `json:"type"`
}

// ExternalSecretCreateSpec defines the desired state of ExternalSecretCreate
type ExternalSecretCreateSpec struct {
	// What backend will this connect to (i.e. aws, gcp, azure)
	Backend string `json:"backend"`

	// List of Secrets
	Secrets []SecretDescription `json:"secrets"`
}

// ExternalSecretCreateStatus defines the observed state of ExternalSecretCreate
type ExternalSecretCreateStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ExternalSecretCreate is the Schema for the externalsecretcreates API
type ExternalSecretCreate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExternalSecretCreateSpec   `json:"spec,omitempty"`
	Status ExternalSecretCreateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExternalSecretCreateList contains a list of ExternalSecretCreate
type ExternalSecretCreateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExternalSecretCreate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExternalSecretCreate{}, &ExternalSecretCreateList{})
}
