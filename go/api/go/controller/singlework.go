package controller

import (
	"context"
	

	"project/dao"
	"project/pkg"
	"project/service"
	"strconv"
	"time"
)

func Do() {
	go Push()
    go Pull()
	select{}
}
//入栈
func Push(){
	ticker:=time.NewTicker(1 *time.Minute)
    defer ticker.Stop()
	for {
		select{
	case<-ticker.C:
		now:=time.Now()
		result:=dao.QueryNextTime(now)
		for _,v:=range result{
			pkg.RDB.LPush(context.Background(),"workurl",v.ID)
		}
	}
	}
}
//出栈
func Pull(){
  for{
   pop,_:=pkg.RDB.BLPop(context.Background(),0,"workurl").Result()
    go func(id string){
      index,_:=strconv.Atoi(id)
	 
	_= service.WorkService(index)
	}(pop[1])
	}
}