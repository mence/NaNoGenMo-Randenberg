package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "math/rand"
)

var bytefile []string
var linecount int = 0
var sentences []string

func fileRead(file string) {
    f, err := os.Open(file)
    if err != nil {
        log.Fatal(err)
    }

    bf := bufio.NewReader(f)

    for {
        line, isPrefix, err := bf.ReadLine()

        if err == io.EOF {
            break
        }

        if err != nil {
            log.Fatal(err)
        }

        if isPrefix {
            log.Fatal("Error: Unexpected long line reading", f.Name())
        }

        if string(line) != "" {
            bytefile = append(bytefile, string(line))
            linecount += 1
        }
    }
}

func PrintSlice(slice []string) {
    rand.Seed(42)
    for i := 0; i < len(slice); i++ {
        r := rand.Intn(linecount)
        fmt.Printf("%s\r\n", slice[r])
    }
}
func main() {
    fileRead("resources/01_pride-and-prejudice.txt")
    fileRead("resources/02_adventures-of-huckleberry-finn.txt")
    fileRead("resources/03_the-yellow-wallpaper.txt")
    fileRead("resources/04_alices-adventures-in-wonderland.txt")
    fileRead("resources/05_the-adventures-of-sherlock-holmes.txt")

    PrintSlice(bytefile)
}