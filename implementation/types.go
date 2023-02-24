package implementation

type PolicyReport struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	Request *PolicyReportRequest `json:"request,omitempty"`
	// +optional
	Response *PolicyReportResponse `json:"response,omitempty"`
}

type PolicyReportRequest struct {
	kmapi.ObjectInfo `json:",inline"`
}

type PolicyReportResponse struct {
	AuditTimestamp Time         `json:"auditTimestamp,omitempty"`
	Constraints    []Constraint `json:"constraints,omitempty"`
}

type Constraint struct {
	Name              string `json:"name,omitempty"`
	EnforcementAction string `json:"enforcementAction,omitempty"`

	Violations []Violation `json:"violations,omitempty"`
}

type Violation struct {
	kmapi.ObjectInfo `json:",inline"`
	Message          string `json:"message,omitempty"`
}
