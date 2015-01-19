package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "gopkg.in/alecthomas/kingpin.v1"
)

var (
    app = kingpin.New("substenv", "Substitute environment variables into templates")
    input = kingpin.Arg("input", "Input file or stdin if not given").File()
)

func main() {
    kingpin.Version("0.0.1")
    kingpin.Parse()

    var bio = bufio.NewReader(os.Stdin)
    for {
        var line, _, err = bio.ReadLine()
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
        fmt.Printf("%s\n", line)
    }
}
