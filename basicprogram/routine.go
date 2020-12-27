package main
import "fmt"  
func display() {
for i:=0; i<5; i++ {
fmt.Println("sathyabama")
}
}
func main() {
go display()
for i:=0; i<5; i++ {
fmt.Println("sathya")
}
}