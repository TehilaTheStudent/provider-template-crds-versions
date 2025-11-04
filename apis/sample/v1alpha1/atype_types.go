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

// ATypeParameters are the configurable fields of a AType.
type ATypeParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// ATypeObservation are the observable fields of a AType.
type ATypeObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A ATypeSpec defines the desired state of a AType.
type ATypeSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ATypeParameters `json:"forProvider"`
}

// A ATypeStatus represents the observed state of a AType.
type ATypeStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ATypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A AType is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type AType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ATypeSpec   `json:"spec"`
	Status ATypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ATypeList contains a list of AType
type ATypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AType `json:"items"`
}

// AType type metadata.
var (
	ATypeKind             = reflect.TypeOf(AType{}).Name()
	ATypeGroupKind        = schema.GroupKind{Group: Group, Kind: ATypeKind}.String()
	ATypeKindAPIVersion   = ATypeKind + "." + SchemeGroupVersion.String()
	ATypeGroupVersionKind = SchemeGroupVersion.WithKind(ATypeKind)
)

func init() {
	SchemeBuilder.Register(&AType{}, &ATypeList{})
}
