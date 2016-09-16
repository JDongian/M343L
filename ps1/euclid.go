package main

import (
    "bufio"
    "io/ioutil"
    "log"
    "os"
    "strconv"
)

func gcd(a int, b int) int {
    x := b
    y := a

    if a > b {
        x = a
        y = b
    }

    for x > 1 && y > 1 {
        rem := x % y
        x = y
        y = rem
    }

    return x
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line1, err := strconv.Atoi(scanner.Text())
    scanner.Scan()
    line2, err := strconv.Atoi(scanner.Text())

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    output := gcd(line1, line2)

    err = ioutil.WriteFile("output.txt", []byte(strconv.Itoa(output)), 0644)
    if (err != nil) {
        panic (err)
    }
}
