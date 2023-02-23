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
	AuditTimestamp Time        `json:"auditTimestamp,omitempty"`
	Violations     []Violation `json:"violations,omitempty"`
	// How many times a particular constraint has been violated
	Stats map[string]int `json:"stats,omitempty"`
}

type Violation struct {
	kmapi.ObjectInfo
	EnforcementAction string `json:"enforcementAction"`
	Message           string `json:"message,omitempty"`
	ConstraintName    string `json:"constraintName,omitempty"`
}
