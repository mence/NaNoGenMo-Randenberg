package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "os"
    "math/rand"
)

var sentences []string
var linecount int = 0

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
            sentences = append(sentences, string(line))
            linecount += 1
        }
    }
}

func printSlice(slice []string) {
    for i := 0; i < len(slice); i++ {
        r := rand.Intn(linecount)
        fmt.Printf("%s\r\n", slice[r])
    }
}
func main() {
    rand.Seed(12345)

    fileRead("resources/01_pride-and-prejudice.txt")
    fileRead("resources/02_adventures-of-huckleberry-finn.txt")
    fileRead("resources/03_the-yellow-wallpaper.txt")
    fileRead("resources/04_alices-adventures-in-wonderland.txt")
    fileRead("resources/05_the-adventures-of-sherlock-holmes.txt")
    fileRead("resources/16_moby-dick.txt")

    printSlice(sentences)
}
