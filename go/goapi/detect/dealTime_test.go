package detect

import(
"testing")


func TestTime(t *testing.T){
	timestr:="Thu, 19 Feb 2026 05:05:00 +0000"
     time:=DealTime(timestr)
	 t.Logf("时间%v",time)
	 													
}