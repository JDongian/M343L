package main

import (
    "bufio"
    "io/ioutil"
    "log"
    "os"
    "strconv"
)

func compute(order int, op string, a int, b int) int {
    if op == "*" {
        return (a * b) % order
    } else if op == "+" {
        return (a + b) % order
    } else {
        panic("Invalid operation used")
    }
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
    line2 := scanner.Text()
    scanner.Scan()
    line3, err := strconv.Atoi(scanner.Text())
    scanner.Scan()
    line4, err := strconv.Atoi(scanner.Text())

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    output := compute(line1, line2, line3, line4)

    err = ioutil.WriteFile("output.txt", []byte(strconv.Itoa(output)), 0644)
    if (err != nil) {
        panic (err)
    }
}
