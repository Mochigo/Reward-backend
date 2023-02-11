package config

import "testing"

func TestInit(t *testing.T) {
	t.Run("解析", func(t *testing.T) {
		if err := Init(""); err != nil {
			t.Errorf("Init() error = %v", err.Error())
		}
	})
}
