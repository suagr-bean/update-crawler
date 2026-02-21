package model
type Show struct{
	Size int 
	Url string 
	Start int
	SeachTitle string 
	
}
type InfoResult struct {
   Name string `json:"name"`
   Url string `json:"url"`
   LastUpdate string `json:"last"`
}