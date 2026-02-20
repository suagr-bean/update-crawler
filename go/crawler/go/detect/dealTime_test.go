package detect

import "testing"

func TestTime(t *testing.T){
	timestr:="Fri, 20 Feb 2026 07:40:00 GMT"
     time:=DealTime(timestr)
	 t.Logf("时间%v",time)
}