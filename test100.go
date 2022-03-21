package main
import "fmt"
import "encoding/json"


type Person struct{
  Name string `json:"Name"`
  Id int64     `json:"Id"`
}

func main(){
  person:=Person{"bapan",24}
  data,err:=json.Marshal(person)
  if err!=nil{
     fmt.Println("Error occured")
     return
  }
  fmt.Println(string(data))
  data1:=json.RawMessage(string(data))
  bte,err:=data1.MarshalJSON()
  if err!=nil{
     fmt.Println(err)
     return
  }
  p1:=Person{}
  err=json.Unmarshal(bte,&p1)
  if err!=nil{
      fmt.Println(err)
      return
  }
  fmt.Println(p1)
}