package config

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		data []CsvData
		length int
	}{
	  {
      name: "test",
	  args:args{
		path:"/workspaces/update-crawler/go/input.csv" },
		wantErr:false,
         length:2,
},      
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _,err := ReadFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			 data,_:=ReadFile(tt.args.path)
           if len(data)==tt.length{
			 fmt.Printf("rightcsvsdata%v",data)
		   }else{
			t.Errorf("csvDatalengt erro%v",len(data))
		   } 
		})
	}
}
