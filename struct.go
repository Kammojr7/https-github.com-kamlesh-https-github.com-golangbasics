package main  
import (  
   "fmt"  
)  
type person struct {  
   firstName string  
   lastName  string  
   age       int  
}  
func main() {  
   x := person{age: 24, firstName: "Kamlesh", lastName: "Kumar", }  
   fmt.Println(x)  
   fmt.Println(x.firstName)  
}  
