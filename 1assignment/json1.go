package main

import (
"encoding/json"
"fmt"
)

func main()
{
type Price struct {
Gold int
Silver int
Platinum int
}

a := `{"gold": 1271,"silver": 1284,"platinum": 1270}`
var obj Price

err := json.Unmarshal([]byte(a), &obj)
if err != nil {
fmt.Println("error:", err)
}

fmt.Printf("Gold : $ %d\n", obj.Gold)
fmt.Printf("Silver : $ %d\n", obj.Silver)
fmt.Printf("Platinum : $ %d\n", obj.Platinum)
}
