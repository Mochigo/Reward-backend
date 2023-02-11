package model

import "testing"

func TestInitDatabase(t *testing.T) {
	t.Run("连接成功", func(t *testing.T) {
		if err := InitDatabase(); err != nil {
			t.Errorf("数据库连接失败, err=%v", err.Error())
		}
	})
}
