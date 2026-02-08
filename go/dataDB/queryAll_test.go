package dataDB

import "testing"

func TestQueryAll(t *testing.T) {
	err := InitDB()
	if err != nil {
		t.Errorf("数据库初始错误%v", err)
	}
	result, err := QueryAll()
	if err != nil {
		t.Errorf("查询时出错误%v", err)
	}
	for _, v := range result {
		t.Logf("查询结果%v,%v", v.Url, v.Name)
	}
}
