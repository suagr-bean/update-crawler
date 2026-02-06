package dataDB

import "testing"

func TestQuery(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	 { name:"text",args:args{
		name:"example",
	 }, want:false,
	 },
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_= InitDB()
			if got,_:= Query(tt.args.name); got != tt.want {
				t.Errorf("Query() = %v, want %v", got, tt.want)
			}
			_,err:=Query(tt.args.name)
			if err!=nil{
				t.Errorf("错误%v",err)
			}
		})
	}
}
