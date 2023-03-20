package entity

type CreateApplicationEntity struct {
	ScholarshipItemId int64  `json:"scholarship_item_id"` // 奖学金子项id
	ScholarshipId     int64  `json:"scholarship_id"`      // 奖学金id
	Deadline          string `json:"deadline"`
}

type GetUserApplicationEntity struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ApplicationEntity struct {
	Id                  int64  `json:"id"`                    // 申请id
	ScholarshipItemId   int64  `json:"scholarship_item_id"`   // 奖学金子项id
	ScholarshipItemName string `json:"scholarship_item_name"` // 奖学金子项name
	ScholarshipId       int64  `json:"scholarship_id"`        // 奖学金id
	ScholarshipName     string `json:"scholarship_name"`      // 奖学金名称
	StudentId           int64  `json:"student_id"`            // 申请学生id
	Status              string `json:"status"`                // 申请状态，APPROVE-通过|PROCESS-处理中|FAILURE-驳回
	Deadline            string `json:"deadline"`              //过期时间
}
