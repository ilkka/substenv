package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "regexp"
    "gopkg.in/alecthomas/kingpin.v1"
)

var (
    app = kingpin.New("substenv", "Substitute environment variables into templates")
    input = app.Arg("input", "Input file or stdin if not given").File()
    re = regexp.MustCompile(`\$(?:\{([A-Z][A-Z0-9_]*)\}|([A-Z][A-Z0-9_]*)\b)`)
)

func main() {
    kingpin.Version("0.0.1")
    kingpin.Parse()

    var bio = bufio.NewReader(os.Stdin)
    for {
        var line, err = bio.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
        matches := re.FindAllStringSubmatchIndex(line, -1)
        fmt.Printf("%s -> %v\n", line, matches)
        for i := 0; i < len(matches); i++ {
            fmt.Printf("-- %v\n", matches[i])
        }
    }
}
