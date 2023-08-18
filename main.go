package main

import (
    "fmt"
    "flag"
    "os"
    "regexp"
    "bufio"
    "path/filepath"
    "sync"
)

func parseFlags() (bool, bool) {
    r := flag.Bool("r", false, "recursive")
    n := flag.Bool("n", false, "show line numbers")

    rn := flag.Bool("rn", false, "show line numbers and go recursive")
    nr := flag.Bool("nr", false, "show line numbers and go recursive")

    flag.Parse()

    if *rn || *nr {
        *r, *n = true, true
    }

    return *r, *n
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func grepFile(re *regexp.Regexp, name string, n bool, full bool) {
    f, err := os.Open(name)
    check(err)

    scanner := bufio.NewScanner(f)

    i := 1
    for scanner.Scan() {
        line := scanner.Text()

        if re.MatchString(line) {
            if n {
                if full {
                    fmt.Printf("./%s:%d:%s\n", name, i, line)
                } else {
                    fmt.Printf("%d:%s\n", i, line)
                }
            } else {
                fmt.Printf("%s\n", line)
            }
        }

        i++
    }
}

// grep "pattern" file
// grep -rn "pattern"

func main() {
    r, n := parseFlags()

    pattern := flag.Arg(0)

    re, err := regexp.Compile(pattern)
    check(err)

    file := flag.Arg(1)

    if file != "" && r {
        panic("Can't specify a file and recursive.")
    }

    if r {
        var wg sync.WaitGroup

        filepath.Walk(".", func(p string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }

            wg.Add(1)
            go func() {
                defer wg.Done()
                grepFile(re, p, n, true)
            }()

            return nil
        })

        wg.Wait()
    } else {
        grepFile(re, file, n, false)
    }
}
