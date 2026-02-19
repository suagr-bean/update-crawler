package model
type Show struct{
	Size int 
	Url string 
	Start int
}
type ShowResult struct {
   Name string `json:"name"`
   Url string `json:"url"`
   LastUpdate string `json:"last"`

}