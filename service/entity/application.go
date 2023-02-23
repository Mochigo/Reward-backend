package entity

type CreateApplicationEntity struct {
	ScholarshipItemId int64 `json:"scholarship_item_id"` // 奖学金子项id
	ScholarshipId     int64 `json:"scholarship_id"`      // 奖学金id
}
