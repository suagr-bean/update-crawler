package service

import (
	"project/dao"
	"testing"
)

func TestAddService(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test",
			args: args{
				url: "https://wiwi.blog/blog/rss.xml",
			},
		},
	}
	path := "../data/test.db"
	dao.Init(path)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddService(tt.args.url)
		})
	}
}
