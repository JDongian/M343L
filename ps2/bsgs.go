package main

import (
    "bufio"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "math"
    "math/big"
)

func exp(modulus int, base int, power int) int {
    m := big.NewInt(10)
    result := m.Exp(big.NewInt(int64(base)),
                    big.NewInt(int64(power)),
                    big.NewInt(int64(modulus)))
    return int(result.Int64())

}

func in_map_values(pairs map[int]int, value int) int {
    //fmt.Println(value)
    for k, v := range pairs {
        //fmt.Println(k, v)
        if value == v {
            return k;
        }
    }
    return -1;
}

func prime_inv(a int, m int, n int) int {
    return exp(n, a, (n - 1) - m)
}

// a is a primitive root in (Z/b)*
func bsgs(a int, b int, n int) int {
    //fmt.Println(a, b, n)
    m := int(math.Ceil(math.Sqrt(float64(n))))
    // n is prime, so a**(n-2) is the inverse of m
    // assert base > 2

    results := make(map[int]int)

    for j := 0; j < m; j++ {
        results[j] = exp(n, a, j)
    }

    inv := prime_inv(a, m, n)
    //fmt.Println(a, m, inv)

    y := b

    for i := 0; i < m; i++ {
        j := in_map_values(results, y)

        if j != -1 {
            return i * m + j
        } else {
            y = (y * inv) % n
        }
    }

    return -1 // probably panic or something because math died
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
    h, err := strconv.Atoi(scanner.Text())

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    output := bsgs(g, h, p)
    //fmt.Println(output)

    err = ioutil.WriteFile("output.txt", []byte(strconv.Itoa(output)), 0644)
    if (err != nil) {
        panic (err)
    }
}
