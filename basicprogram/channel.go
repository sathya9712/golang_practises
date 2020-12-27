package main
import "fmt"
import "time" 
func display()
{
time.Sleep(5 * time.Second)
fmt.Println("sathya")
}
func main() 
{
go display()
fmt.Println("Inside sathya")
}