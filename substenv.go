package main

import (
    "fmt"
    "os"
    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "substenv"
    app.Usage = "substitute environment variables into template"
    app.Action = func(c *cli.Context) {
        fmt.Printf("Hello, World!\n")
    }

    app.Run(os.Args)
}
