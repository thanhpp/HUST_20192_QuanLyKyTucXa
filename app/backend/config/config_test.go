package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init()
}

// func TestGetConfig(t *testing.T) {
// 	Init()
// 	cf := GetConfig()
// 	val := cf.GetString("TEST")
// 	if val != "TEST" {
// 		t.Errorf("Error read from config : %+v", val)
// 	}
// }
