package models

type ReviewStatus string

const (
	ReviewStatusPending    ReviewStatus = "pending"
	ReviewStatusInit       ReviewStatus = "init"
	ReviewStatusPrechecked ReviewStatus = "prechecked"
	ReviewStatusQueued     ReviewStatus = "queued"
	ReviewStatusCompleted  ReviewStatus = "completed"
	ReviewStatusOnHold     ReviewStatus = "onHold"
)

type ApplicantStatusResponse struct {
	CreateDate   string       `json:"createDate"`
	StartDate    string       `json:"startDate,omitempty"`
	AttemptCnt   int          `json:"attemptCnt,omitempty"`
	LevelName    string       `json:"levelName,omitempty"`
	ReviewDate   string       `json:"reviewDate,omitempty"`
	ReviewResult ReviewResult `json:"reviewResult,omitempty"`
	ReviewStatus ReviewStatus `json:"reviewStatus"`
}

type ReviewResult struct {
	ReviewAnswer      string   `json:"reviewAnswer"`
	ModerationComment string   `json:"moderationComment,omitempty"`
	ClientComment     string   `json:"clientComment,omitempty"`
	RejectLabels      []string `json:"rejectLabels,omitempty"`
	ReviewRejectType  string   `json:"reviewRejectType,omitempty"`
	ButtonIds         []string `json:"buttonIds,omitempty"`
}
