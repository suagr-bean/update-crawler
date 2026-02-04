package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Task struct {
	Name string
	Url  string
	Etag string
	Last string
	Hash string
}

// 读配置文件  返回结构体 切片
func readInfo() []Task {
	file, err := os.Open("url.text")
	if err != nil {
		log.Println("配置文件有问题")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file) //流式读
	var task []Task
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if scanner.Scan() {
			url := strings.TrimSpace(scanner.Text())
			task = append(task, Task{Name: name, Url: url})
		}
	}
	return task
}
