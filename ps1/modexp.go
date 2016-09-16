package main

import (
    "bufio"
    "io/ioutil"
    "log"
    "os"
    "strconv"
)

func exp(modulus int, base int, power int) int {
    if modulus == 1 {
        return 0;
    }
    result := 1
    base %= modulus
    for power > 0 {
        if power % 2 == 1 {
            result = (result * base) % modulus
        }
        power /= 2
        base = (base * base) % modulus
    }
    return result
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
    scanner.Scan()
    line3, err := strconv.Atoi(scanner.Text())

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    output := exp(line1, line2, line3)

    err = ioutil.WriteFile("output.txt", []byte(strconv.Itoa(output)), 0644)
    if (err != nil) {
        panic (err)
    }
}
