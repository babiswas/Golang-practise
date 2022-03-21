package main
import "fmt"
func main(){
  sum:=0
  arr:=[5]int{5,6,7,8,9}
  for i:=0;i<len(arr);i++{
    sum+=arr[i]
  }
  fmt.Println(sum)
}