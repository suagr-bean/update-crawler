package dataDB
import "testing"
func TestQuery(t*testing.T){
  type inputs struct{
	name string 
  }
test:=[]struct{
	name  string
	input inputs
	 
  }{
	{
	name:"test1",
	input:inputs{
		name:"wiwi",
	},
	
   },
  }
  for _,tt:=range test{
	t.Run(tt.name,func(t*testing.T){
		_=InitDB()
		result,err:=Query(tt.input.name)
		if err!=nil{
          t.Errorf("not you want%v",err)
		}
		if result==(DBData{}){
			t.Errorf("no data%v",result)
		}else{
			t.Logf("data%v",result.Name)
		}
	})
  }
}