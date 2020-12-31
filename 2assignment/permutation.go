package main

import "fmt"

func main() {
x:= "abc"
y := []rune(x)
generatePermutation(y, 0, len(y)-1)
}

func generatePermutation(y []rune, left, right int) {
if left == right {
fmt.Println(string(y))
} else {
for i := left; i <= right; i++ {
y[left], y[i] = y[i], y[left]
generatePermutation(y, left+1, right)
y[left], y[i] = y[i], y[left]
}
}
}