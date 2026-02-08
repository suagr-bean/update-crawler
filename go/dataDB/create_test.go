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
		{name: "text", fields: fields{
			Name:       "wiwi",
			Url:        "http://wiwi.blog/blog/xml",
			Version:    "Rss2.0",
			LastUpdate: "",
		}, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = InitDB()
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

func TestDetail_SaveInfo(t *testing.T) {
	_ = InitDB()
	type fields struct {
		Model    gorm.Model
		DBDataId uint
		Title    string
		Link     string
		Path     string
	}
	type args struct {
		g   *gorm.DB
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "test",
			fields:  fields{Title: "35"},
			args:    args{g: db, url: "http://wiwi.blog/blog/xml"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			d := &Detail{
				Model:    tt.fields.Model,
				DBDataId: tt.fields.DBDataId,
				Title:    tt.fields.Title,
				Link:     tt.fields.Link,
				Path:     tt.fields.Path,
			}
			if err := d.SaveInfo(tt.args.g, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("Detail.SaveInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
