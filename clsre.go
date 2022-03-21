package main
import "fmt"

func sample(myfunc func(x int)int)func(x int)int{
   return myfunc
}

func main(){
  test:=func(x int)int{
     return x*2*2
  }
  func2:=sample(test)
  fmt.Println(func2(4))
}