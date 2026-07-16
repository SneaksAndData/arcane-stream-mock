package v2

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CronJobBackend represents the backend configuration for batch processing, including the cron schedule and a reference
// to the job template.
type CronJobBackend struct {

	// Schedule represents the cron schedule for batch processing.
	Schedule string `json:"schedule"`

	// JobTemplateRef represents a reference to the job template.
	JobTemplateRef v1.ObjectReference `json:"jobTemplateRef"`
}

// BatchJobBackend represents the backend configuration for real-time streaming, including the change capture interval
// and a reference to the job template.
type BatchJobBackend struct {
	// JobTemplateRef represents a reference to the job template.
	JobTemplateRef v1.ObjectReference `json:"jobTemplateRef"`

	// BackfillJobTemplateRef represents a reference to the job template.
	BackfillJobTemplateRef v1.ObjectReference `json:"backfillJobTemplateRef"`
}

// StreamingBackend represents the backend configuration for streaming, including both real-time and batch processing options.
type StreamingBackend struct {
	// BatchJobBackend represents the backend configuration for real-time streaming.
	BatchJobBackend *BatchJobBackend `json:"changeCapture,omitempty"`

	// CronJobBackend represents the backend configuration for batch processing.
	CronJobBackend *CronJobBackend `json:"batch,omitempty"`
}

// ExecutionSettings represents the execution settings for a stream, including suspension status and backend configuration.
type ExecutionSettings struct {
	// LayoutVersion represents the layout version of the execution settings.
	LayoutVersion string `json:"layoutVersion"`

	// Suspended indicates whether the stream is suspended.
	Suspended bool `json:"suspended"`

	// StreamingBackend represents the backend configuration for streaming.
	StreamingBackend StreamingBackend `json:"streamingBackend"`
}

// TestsStreamDefinitionSpec is a mock implementation of the StreamDefinitionSpec for testing purposes.
type TestsStreamDefinitionSpec struct {
	// Source represents the source of the stream.
	Source string `json:"source"`

	// Destination represents the destination of the stream.
	Destination string `json:"destination"`

	// RunDuration represents the duration for which the stream should run.
	// +kubebuilder:default="15s"
	RunDuration string `json:"runDuration"`

	// ShouldFail indicates whether the stream should simulate a failure.
	// +kubebuilder:default=false
	ShouldFail bool `json:"shouldFail"`

	// TestSecretRef represents a reference to a secret for testing purposes.
	TestSecretRef *v1.LocalObjectReference `json:"testSecretRef,omitempty"`

	// ExecutionSettings represents the execution settings for the stream.
	ExecutionSettings ExecutionSettings `json:"execution"`

	// ChangeCaptureInterval represents the interval at which changes are captured for real-time processing.
	ChangeCaptureInterval string `json:"changeCaptureInterval"`
}

type TestStreamDefinitionStatus struct {
	// Phase represents the current phase of the stream.
	Phase string `json:"phase"`

	// ConfigurationHash represents the hash of the current configuration.
	ConfigurationHash string `json:"configurationHash"`

	// Conditions represent the current conditions of the stream.
	Conditions []metav1.Condition `json:"conditions"`
}

// TestStreamDefinitionV2 is a mock implementation of the StreamDefinition for testing purposes.
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Source",type=string,JSONPath=`.spec.source`
// +kubebuilder:printcolumn:name="Destination",type=string,JSONPath=`.spec.destination`
// +kubebuilder:printcolumn:name="Run Duration",type=string,JSONPath=`.spec.runDuration`
// +kubebuilder:printcolumn:name="Should Fail",type=string,JSONPath=`.spec.shouldFail`
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
type TestStreamDefinitionV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TestsStreamDefinitionSpec  `json:"spec,omitempty"`
	Status TestStreamDefinitionStatus `json:"status,omitempty"`
}

// TestStreamDefinitionV2List contains a list of TestStreamDefinition resources
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type TestStreamDefinitionV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TestStreamDefinitionV2 `json:"items"`
}
