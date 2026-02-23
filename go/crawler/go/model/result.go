package model
type Result struct {
	Code int  `json:"code"`
	Message string `json:"message"`
	Info []InfoResult `json:"info"`
	Data  [] DetailData `json:"data"`
	HaseMore bool `json:"hashmore"`
}
type DetailData struct{ 
	Count int64 `json:"count"`
    Title string `json:"title"`
	Link string `json:"link"`
	AutoLink string `json:"auto_link"`
	PublishedTime int64 `json:"published_time"`
}

type QueryResult struct{
	Info []Info
	Count int64
	Details []Detail
}