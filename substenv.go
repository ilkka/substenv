package main

import (
    "fmt"
    "gopkg.in/alecthomas/kingpin.v1"
)

var (
    app = kingpin.New("substenv", "Substitute environment variables into templates")
    input = kingpin.Arg("input", "Input file or stdin if not given").File()
)

func main() {
    kingpin.Version("0.0.1")
    kingpin.Parse()
    fmt.Printf("Hello, World!\n")
}
