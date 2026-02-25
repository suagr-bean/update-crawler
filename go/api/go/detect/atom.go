package detect

import (
	"encoding/xml"
	"project/model"
)

type AtomProcess struct {
	Data *model.Atom
}

func(atom*AtomProcess)Deal(data []byte )(model.DealData,error){
   if len(data)==0{
	return model.DealData{},nil
   }
    var deal model.DealData
    err:=xml.Unmarshal(data,&atom.Data)
	if err!=nil{
		return deal,err
	}
	deal=model.DealData{
		Name: atom.Data.Title,
		LastUpdate:atom.Data.Last(),
		Version:"Atom",
	}
	return deal,nil

}