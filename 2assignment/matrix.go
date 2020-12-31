package main
 
import (
"fmt"
)
 
func pattren(n int) []int 
{
zz := make([]int, n*n)
i := 0
n2 := n * 2
for a := 1; a <= n2; a++ {
x := a - n
if x < 0 
{
x = 0
}
y := a - 1
if y > n-1 
{
y = n - 1
}
j := n2 - a
if j > a 
{
j = a
}
}
for k := 0; k < j; k++
{
if a&1 == 0 
{
zz[(x+k)*n+y-k] = i
} 
else {
zz[(y-k)*n+x+k] = i
}
i++
}
}

return zz
}
 
func main() 
{
num := 5
len := 2
for i, d := range pattren(num) 
{
fmt.Printf("%*d ", len, d)
if i%num == num-1 
{
fmt.Println("")
}
}
}