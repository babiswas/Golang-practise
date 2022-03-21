package main
import "fmt"
import "flag"

func main(){
   wordptr:=flag.String("word","foo","a string")
   numptr:=flag.Int("num",3,"an int")
   boolptr:=flag.Bool("test",false,"a bool")
   var stringFlag string
   flag.StringVar(&stringFlag,"stringvalue","Hello There!","A string flag")
   flag.Parse()
   fmt.Println(*wordptr)
   fmt.Println(*numptr)
   fmt.Println(*boolptr)
   fmt.Println(stringFlag)
}