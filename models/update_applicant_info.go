package models

type UpdateApplicantInfoResponse struct {
	ApplicantID    string `json:"applicantId"`
	InspectionID   string `json:"inspectionId"`
	CorrelationID  string `json:"correlationId"`
	ExternalUserID string `json:"externalUserId"`
	LevelName      string `json:"levelName"`
	Type           string `json:"type"`
	ReviewResult   struct {
		ReviewAnswer string `json:"reviewAnswer"`
	} `json:"reviewResult"`
	ReviewStatus string `json:"reviewStatus"`
	CreatedAtMs  string `json:"createdAtMs"`
}
