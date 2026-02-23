package controller

import (
	"context"
	"fmt"
	"project/Utils"
	"project/dao"
	"project/pkg"
	"project/service"
	"strconv"
	"time"
)
func Do(){
	now:= time.Now()
	result:=dao.QueryNextTime(now )
	//压栈
   for _,v:=range result{
	pkg.RDB.LPush(context.Background(),"workurl",v.ID)
   }
   
   for {
   pop,_:=pkg.RDB.BLPop(context.Background(),0,"workurl").Result()
   fmt.Println(pop[1])
    index,_:=strconv.Atoi(pop[1])
     q:=dao.QueryWork(index)
	deal,_:= service.WorkService(q.Url)
	var  newtime int
	var next time.Time
	 if deal==true{
		newtime=int(float64(q.DoMinute)*0.9)
		next= Utils.DealTime(newtime)
	 }else {
		newtime=int(float64(q.DoMinute)*1.1)
		next=Utils.DealTime(newtime)
	 }
	 
   }
     
}