package dataDB

import "testing"

func TestInit(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Errorf("wrong info %v", err)
	}
}
