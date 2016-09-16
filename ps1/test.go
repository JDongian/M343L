package main

import (
    "os"
)

func main() {
    d1 := []byte("123")
    f, err := os.Create("output.txt")
    //err := ioutil.WriteFile("output.txt", d1, 0644)

    if (err != nil) {
        panic (err)
    }
    /*
    f, err := os.Create("/tmp/dat2")
    check(err)
    defer f.Close()
    */
}
