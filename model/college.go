package model

type College struct {
	Id   int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"` // 学院名称
}

func (m *College) TableName() string {
	return "`college`"
}
