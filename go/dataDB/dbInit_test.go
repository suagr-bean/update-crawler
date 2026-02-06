package dataDB

import "testing"

func TestOpenDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err:=InitDB()
			if err!=nil{
				t.Errorf("Opendb%v",err)
			}
		})
	}
}
