package detect

import (
	"encoding/xml"
	"testing"
)

func TestDetect(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		Args    args
		want    string
		wantErr bool
	}{
		{
      name:"test",
	  Args:args{url:"https://wiwi.blog/blog/rss.xml"},
	  want:"",
	  wantErr:false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Detect(tt.Args.url)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Detect() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if got != tt.want {
				t.Errorf("Detect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXmlVersion_Do(t *testing.T) {
	type fields struct {
		XMLName xml.Name
		Version string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &XmlVersion{
				XMLName: tt.fields.XMLName,
				Version: tt.fields.Version,
			}
			if got := v.Do(); got != tt.want {
				t.Errorf("XmlVersion.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
