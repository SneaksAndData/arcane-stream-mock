package v1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestsStreamDefinitionSpec is a mock implementation of the StreamDefinitionSpec for testing purposes.
type TestsStreamDefinitionSpec struct {
	// Source represents the source of the stream.
	Source string `json:"source"`

	// Destination represents the destination of the stream.
	Destination string `json:"destination"`

	// Suspended indicates whether the stream is suspended.
	Suspended bool `json:"suspended"`

	// JobTemplateRef represents a reference to the job template.
	JobTemplateRef v1.ObjectReference `json:"jobTemplateRef"`

	// BackfillJobTemplateRef represents a reference to the job template.
	BackfillJobTemplateRef v1.ObjectReference `json:"backfillJobTemplateRef"`

	// RunDuration represents the duration for which the stream should run.
	// +kubebuilder:default="15s"
	RunDuration string `json:"runDuration"`

	// ShouldFail indicates whether the stream should simulate a failure.
	// +kubebuilder:default=false
	ShouldFail bool `json:"shouldFail"`
}

type TestStreamDefinitionStatus struct {
	// Phase represents the current phase of the stream.
	Phase string `json:"phase"`

	// ConfigurationHash represents the hash of the current configuration.
	ConfigurationHash string `json:"configurationHash"`

	// Conditions represent the current conditions of the stream.
	Conditions []metav1.Condition `json:"conditions"`
}

// TestStreamDefinition is a mock implementation of the StreamDefinition for testing purposes.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Source",type=string,JSONPath=`.spec.source`
// +kubebuilder:printcolumn:name="Destination",type=string,JSONPath=`.spec.destination`
// +kubebuilder:printcolumn:name="Run Duration",type=string,JSONPath=`.spec.runDuration`
// +kubebuilder:printcolumn:name="Should Fail",type=string,JSONPath=`.spec.shouldFail`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:object:root=true
type TestStreamDefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestsStreamDefinitionSpec  `json:"spec,omitempty"`
	Status TestStreamDefinitionStatus `json:"status,omitempty"`
}

// TestStreamDefinitionList contains a list of TestStreamDefinition resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TestStreamDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestStreamDefinition `json:"items"`
}
