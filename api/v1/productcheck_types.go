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

package v1

import (
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProductCheckSpec defines the desired state of ProductCheck
type ProductCheckSpec struct {
	// Add fields specific to your CLI tool's functionality
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ProductCheckStatus defines the observed state of ProductCheck
type ProductCheckStatus struct {
	IsEOL    bool        `json:"isEOL"`
	Message  string      `json:"message,omitempty"`
	CheckdAt metav1.Time `json:"checkedAt,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProductCheck is the Schema for the productchecks API
type ProductCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProductCheckSpec   `json:"spec,omitempty"`
	Status ProductCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProductCheckList contains a list of ProductCheck
type ProductCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProductCheck{}, &ProductCheckList{})

	if err := scheme.AddToScheme(scheme.Scheme); err != nil {
		log.Printf("%t Unable to add ProductCheck API to scheme", err)
		os.Exit(1)
	}
}
