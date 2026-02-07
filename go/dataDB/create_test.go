package dataDB

import (
	"testing"

	"gorm.io/gorm"
)

func Test_dbData_Create(t *testing.T) {
	type fields struct {
		Model      gorm.Model
		Name       string
		Url        string
		Version    string
		LastUpdate string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
	   {name:"text",fields:fields{
		Name:"wiwi",
		Url:"http://wiwi.blog/blog/xml",
		Version:"Rss2.0",
		LastUpdate: "",
	   },wantErr:false,
},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_=InitDB()
			d := &DBData{
				Model:      tt.fields.Model,
				Name:       tt.fields.Name,
				Url:        tt.fields.Url,
				Version:    tt.fields.Version,
				LastUpdate: tt.fields.LastUpdate,
			}
			if err := d.Create(); (err != nil) != tt.wantErr {
				t.Errorf("dbData.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
