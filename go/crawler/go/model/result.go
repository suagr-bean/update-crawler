package model
type Result struct {
	Code int  `json:"code"`
	Message string `json:"message"`
	Info []InfoResult `json:"info"`
	Data  [] DetailData `json:"data"`
}
type DetailData struct{
    Title string `json:"title"`
	Link string `json:"link"`
	AutoLink string `json:"auto_link"`
	PublishedTime int64 `json:"published_time"`
}