package main

import (
    "bufio"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "math/big"
    "math/rand"
    "fmt"
)

func exp(modulus int, base int, power int) int {
    m := big.NewInt(10)
    result := m.Exp(big.NewInt(int64(base)),
                    big.NewInt(int64(power)),
                    big.NewInt(int64(modulus)))
    return int(result.Int64())

}

func elgamal(p int, g int, m int, h int) (int, int) {
    y := (rand.Int() % (p - 1)) + 1
    c1 := exp(p, g, y)
    s := exp(p, h, y)
    // ??
    m %= p
    c2 := (m * s) % p
    return c1, c2
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    p, err := strconv.Atoi(scanner.Text())
    scanner.Scan()
    g, err := strconv.Atoi(scanner.Text())
    scanner.Scan()
    m, err := strconv.Atoi(scanner.Text())
    scanner.Scan()
    h, err := strconv.Atoi(scanner.Text())

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    output1, output2 := elgamal(p, g, m, h)
    //fmt.Println(output)
    output := fmt.Sprintf("%v %v", output1, output2)

    err = ioutil.WriteFile("output.txt", []byte(output), 0644)
    if (err != nil) {
        panic (err)
    }
}
