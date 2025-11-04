/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// DTypeParameters are the configurable fields of a DType.
type DTypeParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// DTypeObservation are the observable fields of a DType.
type DTypeObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A DTypeSpec defines the desired state of a DType.
type DTypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DTypeParameters `json:"forProvider"`
}

// A DTypeStatus represents the observed state of a DType.
type DTypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A DType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type DType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DTypeSpec   `json:"spec"`
	Status DTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DTypeList contains a list of DType
type DTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DType `json:"items"`
}

// DType type metadata.
var (
	DTypeKind             = reflect.TypeOf(DType{}).Name()
	DTypeGroupKind        = schema.GroupKind{Group: Group, Kind: DTypeKind}.String()
	DTypeKindAPIVersion   = DTypeKind + "." + SchemeGroupVersion.String()
	DTypeGroupVersionKind = SchemeGroupVersion.WithKind(DTypeKind)
)

func init() {
	SchemeBuilder.Register(&DType{}, &DTypeList{})
}
