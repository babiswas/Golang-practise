package main
import "fmt"
import "bufio"
import "os"

func main(){
  var books [5]string
  scanner:=bufio.NewScanner(os.Stdin)
  for i:=0;i<len(books);i++{
     fmt.Println("Enter the book name")
     scanner.Scan()
     books[i]=scanner.Text()
  }
  fmt.Println("Displaying all the books in an array")
  for i:=0;i<len(books);i++{
      fmt.Println(books[i])
  }
}