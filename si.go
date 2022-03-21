package main
import "fmt"
import "os"
import "strconv"
import "bufio"

func get_principal()int64{
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter the principal")
  scanner.Scan()
  principal,err:=strconv.ParseInt(scanner.Text(),10,64)
  if err!=nil{
     return int64(0)
  }
  return principal
}

func get_time()int64{
  scanner:=bufio.NewScanner(os.Stdin)
  fmt.Println("Enter the time")
  scanner.Scan()
  time,err:=strconv.ParseInt(scanner.Text(),10,64)
  if err!=nil{
    return int64(0)
  }
  return time
}

func get_roi()float64{
   scanner:=bufio.NewScanner(os.Stdin)
   fmt.Println("Enter the rate of intrest")
   scanner.Scan()
   rate,err:=strconv.ParseFloat(scanner.Text(),64)
   if err!=nil{
      return float64(0.0)
   }
   return rate
}

func main(){
   principal:=get_principal()
   rate:=get_roi()
   time:=get_time()
   intrest:=float64(principal)*float64(time)*rate/float64(100)
   fmt.Println(intrest)
}