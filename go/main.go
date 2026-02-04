package main

import "time"

func main() {
	db := DataBase()
	task := readInfo()
	tasks := HeaderTry(task) //请求头决定怎么爬
	Insert(tasks, db)        //初始插入
	for {
		ReadData(db)
		time.Sleep(2 * time.Minute)
	}

}
