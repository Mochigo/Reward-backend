package entity

// attahcment
type AddAttachmentEntity struct {
	ScholarshipId int64  `json:"scholarship_id"`
	Url           string `json:"url"`
}

type GetAttachmentsEntity struct {
	ScholarshipId int64 `json:"scholarship_id"`
}

type RemoveAttachmentEntity struct {
	AttachmentId int64 `json:"attachment_id"`
}

type AttachmentEntity struct {
	Id  int64  `json:"attachment_id"`
	Url string `json:"url"`
}

// scholarship
type ScholarshipEntity struct {
	Id        int64  `json:"scholarship_id"`
	Name      string `json:"name"`       // 奖学金名称
	CollegeId int64  `json:"college_id"` // 学院id
	StartTime string `json:"start_time"` // 开始时间
	EndTime   string `json:"end_time"`   // 结束时间
}

type CreateScholarshipEntity struct {
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type GetScholarshipsEntity struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type GetScholarshipInfoEntity struct {
	ScholarshipId int64 `json:"scholarship_id"`
}

// scholarshipItem
type AddScholarshipItemEntity struct {
	ScholarshipId int64  `json:"scholarship_id"`
	Name          string `json:"name"`
}

type GetScholarshipItemsEntity struct {
	ScholarshipId int64 `json:"scholarship_id"`
}

type ScholarshipItemEntity struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"` // 奖学金子项名称
	ScholarshipId int64  `json:"scholarship_id"`
}

type RemoveScholarshipItemEntity struct {
	ScholarshipItemId int64 `json:"scholarship_item_id"`
}
