package main
import "fmt"
import "os"
import "bufio"
import "strconv"

func main(){
   scan1:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the first number")
   scan1.Scan()
   a,_:=strconv.ParseInt(scan1.Text(),10,64)
   fmt.Println("Enter the second number")
   scan1.Scan()
   b,_:=strconv.ParseInt(scan1.Text(),10,64)
   fmt.Println("The sum of the number is:")
   fmt.Println(a+b)
}