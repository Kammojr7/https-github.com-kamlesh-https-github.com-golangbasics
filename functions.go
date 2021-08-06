package main  
import "fmt"  
type Employee struct {  
   fname string  
   lname string  
}  
func (emp Employee) fullname(){  
   fmt.Println(emp.fname+" "+emp.lname)  
}  
func main() {  
   e1 := Employee{ "Kamlesh","Kumar"}  
   e1.fullname()  
}  
