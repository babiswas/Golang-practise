package main
import "fmt"
import "bufio"
import "os"

func main(){
  scan1:=bufio.NewScanner(os.Stdin)
  scan1.Scan()
  fmt.Println(scan1.Text())
}