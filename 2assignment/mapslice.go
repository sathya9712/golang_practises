package main

import "fmt"

func main()
{
m := map[string][]string{
"cat": {"orange", "grey"},
"dog": {"black"},
}

r := append(m["dog"], "brown")
fmt.Println(r)

m["fish"] = []string{"orange", "red"}

fmt.Println(m["fish"])

for i := range m["fish"] {
fmt.Println(i, m["fish"][i])
}
}