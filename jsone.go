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
}