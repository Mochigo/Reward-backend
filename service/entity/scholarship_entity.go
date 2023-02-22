package entity

type AddAttachmentEntity struct {
	ScholarshipId int64  `json:"scholarship_id"`
	Url           string `json:"url"`
}

type GetAttachmentsEntity struct {
	ScholarshipId int64 `json:"scholarship_id"`
}

type AttachmentEntity struct {
	Url string `json:"url"`
}

type CreateScholarshipEntity struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
