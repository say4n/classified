package main


import (
    "./classified"
    "fmt"
)


func main() {
    myStr := "Hi there!"

    fmt.Println("original  : ", myStr)
    fmt.Println("redacted  : ", classified.Redact(myStr))
    fmt.Println("sanitized : ", classified.Sanitize(myStr))
}