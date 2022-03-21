package main
import "fmt"
import "bufio"
import "os"

type Person struct{
  name string
  skill string
}

func (p *Person)set_name(name string){
  p.name=name
}

func (p *Person)set_skill(skill string){
  p.skill=skill
}

func (p Person)display(){
   fmt.Println(p.name+"  "+p.skill)
}

func main(){
  person:=make([]Person,4,10)
  scanner:=bufio.NewScanner(os.Stdin)
  for i:=0;i<len(person);i++{
     fmt.Println("Enter the name of the person:")
     scanner.Scan()
     person[i].set_name(scanner.Text())
     fmt.Println("Enter the skill of the person:")
     scanner.Scan()
     person[i].set_skill(scanner.Text())
  }
  for i:=0;i<len(person);i++{
      person[i].display()
  }
  fmt.Println(len(person))
  fmt.Println(cap(person))
}

