package detect

import "testing"

func TestReq(t *testing.T) {
	type data struct {
		Url string
	}
	test := []struct {
		name string
		data data
	}{
		{
			name: "test",
			data: data{
				Url: "https://wiwi.blog/blog/rss.xml",
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CrawlerReq(tt.data.Url)
			if err != nil {
				t.Errorf("err%v", err)
			}
			t.Logf("result%v", result)

		})
	}
}
