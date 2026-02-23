package controller

import (
	"project/dao"
	"project/model"
	"project/service"
	"sync"
	"time"
	
)

// 开始
func WorkStart(times time.Duration) {
	ticker := time.NewTicker(times)
	defer ticker.Stop()
	 WorkController()
	for {
		select {
		case <-ticker.C:
			WorkController()
		}
	}

}

// 日常的轮询
func WorkController() {
	//先分页查询数据库的URL
	var show model.Show
	var wg sync.WaitGroup
	need := make(chan string, 50)
    
	//数据库工人
	wg.Add(1)
	show.Start=1 
	go func() {
		defer wg.Done()
		defer close(need)
		show.Size = 50
		for {
			query, err := dao.QueryUrl(show)
			
			if err != nil {
				return
			}
			if len(query.Info) == 0 {
				return
			}
			
			for _, v := range query.Info {
                 
				need <- v.Url
			}
			show.Start++
		}

	}()

	//执行 工作
	const worker = 5
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				url, ok := <-need
				if !ok {
					break
				}
				service.WorkService(url)
			}
		}()
	}

	wg.Wait()
   

}
