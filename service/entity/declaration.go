package entity

type AddDeclarationEntity struct {
	Name          string `json:"name"`
	Level         string `json:"level"`
	Url           string `json:"url"`
	ApplicationId int64  `json:"application_id"`
}

type GetDeclarationsEntity struct {
	ApplicationId int64 `json:"application_id"`
}

type DeclarationEntity struct {
	Id             int64  `json:"id"`
	ApplicationId  int64  `json:"application_id"`
	StudentId      int64  `json:"student_id"`
	Name           string `json:"name"`
	Level          string `json:"level"`
	Status         string `json:"status"`
	RejectedReason string `json:"rejected_reason"`
	Url            string `json:"url"`
}

type RemoveDeclarationEntity struct {
	DeclarationId int64 `json:"declaration_id"`
}

type AuditDeclarationEntity struct {
	DeclarationId  int64  `json:"declaration_id"`
	RejectedReason string `json:"rejected_reason"`
	Operation      string `json:"operation"`
}
