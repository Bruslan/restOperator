package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RestOperatorSpec defines the desired state of RestOperator
type RestOperatorSpec struct {
	SimpleString string `json:"word"`
}

// RestOperatorStatus defines the observed state of RestOperator
type RestOperatorStatus struct {
	Node string `json:"node"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RestOperator is the Schema for the restoperators API
// +k8s:openapi-gen=true
type RestOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestOperatorSpec   `json:"spec,omitempty"`
	Status RestOperatorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RestOperatorList contains a list of RestOperator
type RestOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestOperator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestOperator{}, &RestOperatorList{})
}
