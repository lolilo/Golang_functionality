// https://gobyexample.com/environment-variables


package main

import "os"
import "strings"
import "fmt"

func main() {
	os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR")) // os.Getevn here returns nothing

	fmt.Println()
    for id_num, value := range os.Environ() {
    	fmt.Println("\n" + value)
        pair := strings.Split(value, "=")
        fmt.Println(id_num, pair[0])
    }
}