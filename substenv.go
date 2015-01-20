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

func template(line string) string {
    var output = ""
    var nextIncludedIndex = 0
    matches := re.FindAllStringSubmatchIndex(line, -1)
    for i := 0; i < len(matches); i++ {
        match := matches[i]
        // find indices of range to replace in line plus the range that
        // gives the name of the variable
        replaceRange := match[0:2]
        var nameRange = match[2:4]
        if nameRange[0] == -1 {
            nameRange = match[4:6]
        }
        name := line[nameRange[0]:nameRange[1]]
        output += line[nextIncludedIndex:replaceRange[0]] + os.Getenv(name)
        nextIncludedIndex = replaceRange[1]
    }
    output += line[nextIncludedIndex:len(line)]
    return output
}

func main() {
    app.Version("0.0.1")
    app.Parse(os.Args[1:])

    var bio = bufio.NewReader(os.Stdin)
    for {
        var line, err = bio.ReadString('\n')
        if err != nil {
            if err != io.EOF {
                log.Fatal(err)
            }
            break
        }
        fmt.Printf("%s", template(line))
    }
}
