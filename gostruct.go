package main
import "strconv"
import "bufio"
import "fmt"
import "os"

type Book struct{
  name string
  pages int64
}

func get_book_name()string{
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter the book name")
  scanner.Scan()
  book:=scanner.Text()
  return book
}

func get_book_pages()int64{
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter the number of pages")
  scanner.Scan()
  pages,err:=strconv.ParseInt(scanner.Text(),10,64)
  if err!=nil{
    fmt.Println(err)
    return -1
  }
  return pages
}

func main(){
  book_name:=get_book_name()
  pages:=get_book_pages()
  book:=Book{book_name,pages}
  fmt.Println(book)
  file,_:=os.OpenFile("hello.txt",os.O_WRONLY|os.O_CREATE,0666)
  fmt.Fprintf(file,"%s book has %d pages",book_name,pages)
  defer file.Close()
}