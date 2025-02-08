package main

import (
    "bufio"
    "container/list"
    "fmt"
    "log"
    "os"
    "strings"

    agm "github.com/adammccartney/algorill/internal/anagram"
    "github.com/adammccartney/algorill/pkg/cli"
)

func main() {

    // UI - basically we would like to load the dictionary into memory,
    // then have a simple loop running that allows the user to query one
    // word at a time.
    // Once a query has been submitted, we can make the key and look up the
    // associated anagrams

    // For initial testing, simply load a file containing a plain text english
    // dictionary
    dpath := "/usr/share/dict/words"
    dict, frerr := os.Open(dpath)
    if frerr != nil {
        log.Fatal("failed open", dpath, frerr)
    }

    agMap := agm.AnagramDict{
        Map: make(map[string]*list.List),
    }

    lerr := agMap.Load(dict)
    if lerr != nil {
        log.Fatal("failed load: ", lerr)
    }

    var pfn cli.Prompt = func(args ...interface{}) []string {
        fmt.Println("")
        fmt.Println("Welcome anagram search")
        fmt.Println("----------------------")
        fmt.Println("Hit return twice to search for anagrams")
        fmt.Println("To exit send a signal (^c)")
        fmt.Println("Start by inputting one or more words")
        fmt.Print("words > ")

        var lines []string
        var words []string
        scanner := bufio.NewScanner(os.Stdin)
        for {
            scanner.Scan()
            line := scanner.Text()
            if len(line) == 0 {
                break
            }
            lines = append(lines, line)
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintln(os.Stderr, "reading input:", err)
        }
        for _, l := range lines {
            _words := strings.Split(l, " ")

            for _, w := range _words {
                words = append(words, w)
            }
        }
        return words
    }

    var tfn cli.Transformer = func(words []string) interface{} {

        for _, w := range words {
            t, err := agm.InitToken(w)
            if err != nil {
                log.Println("error in word: ", err)
            } else {
                agMap.Printer(t)
            }
        }
        return nil
    }
    cli.InteractivePrompt(pfn, tfn)
}
