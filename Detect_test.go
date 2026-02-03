package main
import "testing"
func TestDetect(t *testing.T){
	data,err:=ReadSQL()
	defer data.Close()
	if err!=nil{

	}
   url:="https://weather.yahoo.co.jp/weather/jp/13/4410"
	state:=Detect(url)
	t.Log("测试",state)
	
}