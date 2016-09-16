package main

import (
    "io/ioutil"
    "log"
//    "fmt"
//    "sort"
)

/* DEBUG ONLY
type sortedMap struct {
    m map[string]float64
    s []string
}

func (sm *sortedMap) Len() int {
    return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
    return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
    sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]float64) []string {
    sm := new(sortedMap)
    sm.m = m
    sm.s = make([]string, len(m))
    i := 0
    for key, _ := range m {
        sm.s[i] = key
        i++
    }
    sort.Sort(sm)
    return sm.s
}
*/


func decode(text string, key int) string {
    // Assert 0 <= key < 26
    shift, offset := rune(key), rune(26)
    runes := []rune(text)

    for index, char := range runes {
        if 'a' <= char && char <= 'z' - shift ||
           'A' <= char && char <= 'Z' - shift {
            char = char + shift
        } else if 'z' - shift < char && char <= 'z' ||
                  'Z' - shift < char && char <= 'Z' {
            char = char + shift - offset
        }

        runes[index] = char
    }

    return string(runes)
}

func ngram(runes []rune, n int) map[string]int {
    counts := make(map[string]int)

    for index, _ := range runes {
        if index + n >= len(runes) {
            break;
        }

        token := string(runes[index:index + n])

        _, ok := counts[token]

        if ok == true {
            counts[token] += 1
        } else {
            counts[token] = 1
        }
    }

    return counts
}

func sumSqDiff(basis map[string]float64,
               test map[string]float64) float64{
    sum := 0.

    for token, freq := range basis {
        testFreq := 0.
        _, ok := test[token]

        if ok == true {
            testFreq = test[token]
        }

        difference := freq - testFreq
        sum += difference * difference
    }

    //fmt.Println("BASIS:")
    //fmt.Println(sortedKeys(basis))
    //fmt.Println("TEST:")
    //fmt.Println(sortedKeys(test))

    //fmt.Println("SQSUM:")
    //fmt.Println(sum)

    return sum
}


func isAlpha(s string) bool {
    runes := []rune(s)

    for _, char := range runes {
        if (char < 'a' || 'z' < char) &&
           (char < 'A' || 'Z' < char) {
            return false
        }
    }

    return true
}


func normalize(counts map[string]int) map[string]float64 {
    totalCount := 0.
    freqs := make(map[string]float64)

    for token, count := range counts {
        if isAlpha(token) {
            totalCount += float64(count)
        }
    }

    for token, count := range counts {
        if isAlpha(token) {
            freqs[token] = float64(count) / totalCount
        }
    }

    return freqs
}


func residual(freq map[string]int, tokenSize int) float64 {
    f1 := map[string]float64{
        "A": 0.0016072497221509788,
        "B": 0.001094297683166624,
        "C": 0.0009489612721210567,
        "D": 0.0010088056766692315,
        "E": 0.00024792681884243823,
        "F": 0.001265281696161409,
        "G": 0.0007095836539283577,
        "H": 0.0021458493630845516,
        "I": 0.007600239377618193,
        "J": 0.0004189108318372232,
        "K": 5.12952038984355e-05,
        "L": 0.00047875523638539794,
        "M": 0.0014619133111054116,
        "N": 0.001060100880567667,
        "O": 0.0007010344532786184,
        "P": 0.0007694280584765325,
        "Q": 3.4196802598957e-05,
        "R": 0.0007608788578267932,
        "S": 0.0015559545182525433,
        "T": 0.002769941010515517,
        "U": 0.0002821236214413952,
        "V": 0.00011113960844661024,
        "W": 0.0014875609130546295,
        "X": 0.0,
        "Y": 0.0005642472428827904,
        "Z": 2.564760194921775e-05,
        "a": 0.07958450884842268,
        "b": 0.015542446781225955,
        "c": 0.023176882961443105,
        "d": 0.04692656236641874,
        "e": 0.11941523467555784,
        "f": 0.020458237154826024,
        "g": 0.02425408224331025,
        "h": 0.06019492177481405,
        "i": 0.0629734119859793,
        "j": 0.0008634692656236642,
        "k": 0.011190903650508677,
        "l": 0.041643156364879884,
        "m": 0.024219885440711294,
        "n": 0.06918013165769,
        "o": 0.07333504317346329,
        "p": 0.017115499700777977,
        "q": 0.0010259040779687098,
        "r": 0.05628793707788322,
        "s": 0.060511242198854405,
        "t": 0.08623578695391981,
        "u": 0.028084124134393436,
        "v": 0.007241172950329144,
        "w": 0.020039326322988802,
        "x": 0.00136787210395828,
        "y": 0.019124561853466702,
        "z": 0.0008463708643241857}

    normalizedFreqs := normalize(freq)

    if tokenSize == 1 {
        return sumSqDiff(f1, normalizedFreqs)
    } else {
        return -100
    }
}

func score(text string) float64 {
    runes := []rune(text)

    freqOneGram := ngram(runes, 1)

    return residual(freqOneGram, 1)
}

func unshift(ciphertext string) string {
    bestScore := 1.
    bestShift := 0


    for shift := 0; shift < 26; shift++ {
        //fmt.Println("***")
        trial := decode(ciphertext, shift)
        trialScore := score(trial)

        if trialScore < bestScore {
            bestScore = trialScore
            bestShift = shift
        }

        //fmt.Println("SHIFT:")
        //fmt.Println(shift)
    }

    //fmt.Println("BEST SHIFT:")
    //fmt.Println(bestShift)

    return decode(ciphertext, bestShift)
}

func main() {
    dat, err := ioutil.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    output := unshift(string(dat))
    //fmt.Println("DECODED SAMPLE:")
    //fmt.Println(output)

    err = ioutil.WriteFile("output.txt", []byte(output), 0644)
    if (err != nil) {
        panic (err)
    }
}
