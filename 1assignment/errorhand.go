package main

import (
"database/sql"
"fmt"
)
func foo() error {
return sql.ErrNoRows
}    
func bar() error {
return foo()
}    
func main() {
err := bar()
if err == sql.ErrNoRows {
fmt.Printf("data not found, %+v\n", err)
return
}
if err != nil {
}
}