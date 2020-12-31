
package main
 
import (
	"fmt"
	"math/rand"
	)
 
type Node struct {
	num int
	next *Node
}
 
type List struct {
	length int
	start *Node
}
 
 
 
func main() {
	items := &List{}
	size := 10
    //rand_number := make([]int, size, size)
 	for i := 0; i < size; i++ {
        node := Node{num: rand.Intn(100)}
        if node.num == 0 {
        	i= i -1
        	continue;
 
        }
        items.insertNode(&node)
        fmt.Printf("%v Length: %v and number is %v\n", i, items.length, node.num)
    }
    items.Display()
}
 
func (l *List) Display() {
	list := l.start
    for list != nil {
        fmt.Printf("%v ->", list.num)
       list = list.next
    }
    fmt.Println()
}
 
func (l *List)insertNode(newNode *Node) {
	if l.length == 0 {
		l.start = newNode
	} else {
		currentNode :=  l.start
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
 
	l.length++
}