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

func (p Person)get_name()string{
     return p.name
}

func (p Person)get_skill()string{
     return p.skill
}

func (p Person)display(){
    skill:=p.get_skill()
    name:=p.get_name()
    fmt.Println(skill+" "+name)
}

func main(){
   var persons[5] Person
   scanner:=bufio.NewScanner(os.Stdin)
   for i:=0;i<len(persons);i++{
      fmt.Println("Enter the person name:")
      scanner.Scan()
      persons[i].set_name(scanner.Text())
      fmt.Println("Enter the skill of the person")
      scanner.Scan()
      persons[i].set_skill(scanner.Text())
   }
   for i:=0;i<len(persons);i++{
      persons[i].display()
   }
}