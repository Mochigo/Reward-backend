package entity

type AddCertificateEntity struct {
	Name          string `json:"name"`
	Level         string `json:"level"`
	Url           string `json:"url"`
	ApplicationId int64  `json:"application_id"`
}

type GetCertificatesEntity struct {
	ApplicationId int64 `json:"application_id"`
}

type CertificateEntity struct {
	Id             int64  `json:"id"`
	ApplicationId  int64  `json:"application_id"`  // 申请id
	Name           string `json:"name"`            // 荣誉名称
	Level          string `json:"level"`           // 荣誉级别, 校级-01|省级-02|国家级|-03
	Status         string `json:"status"`          // 申请状态，APPROVED-通过|PROCESS-待处理|REJECTED-驳回
	RejectedReason string `json:"rejected_reason"` // 驳回理由
	Url            string `json:"url"`             // 证明文件连接
}

type RemoveCertificateEntity struct {
	CertificateId int64 `json:"certificate_id"`
}

type AuditCertificateEntity struct {
	CertificateId  int64  `json:"certificate_id"`
	RejectedReason string `json:"rejected_reason"` // 驳回理由
	Operation      string `json:"operation"`
}
