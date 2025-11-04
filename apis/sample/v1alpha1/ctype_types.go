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

// CTypeParameters are the configurable fields of a CType.
type CTypeParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// CTypeObservation are the observable fields of a CType.
type CTypeObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A CTypeSpec defines the desired state of a CType.
type CTypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CTypeParameters `json:"forProvider"`
}

// A CTypeStatus represents the observed state of a CType.
type CTypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A CType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type CType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CTypeSpec   `json:"spec"`
	Status CTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CTypeList contains a list of CType
type CTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CType `json:"items"`
}

// CType type metadata.
var (
	CTypeKind             = reflect.TypeOf(CType{}).Name()
	CTypeGroupKind        = schema.GroupKind{Group: Group, Kind: CTypeKind}.String()
	CTypeKindAPIVersion   = CTypeKind + "." + SchemeGroupVersion.String()
	CTypeGroupVersionKind = SchemeGroupVersion.WithKind(CTypeKind)
)

func init() {
	SchemeBuilder.Register(&CType{}, &CTypeList{})
}
