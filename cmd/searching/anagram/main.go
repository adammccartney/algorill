package main

import (
    "container/list"
    "fmt"

    agm "github.com/adammccartney/algorill/internal/anagram"
)

func main() {

    // allocate a pointer to the main storage
    m := make(map[string]*list.List)

    // read words from dict @ /usr/share/dict/words
    words := []string{"hello", "race", "care"}

    for _, w := range words {
        // create a Token for each word
        var t agm.Token
        t = agm.InitToken(w)

        // pointer to the list containing words
        var l *list.List
        // use the Token's key to create a map entry
        if m[t.Key] == nil {  // associate first value
            l = list.New()
        } else {
            l = m[t.Key]
        }
        l.PushBack(t.Word)
        m[t.Key] = l
    }

    // Just a sanity check
    var lres *list.List
    for k, _ := range m {
        lres = m[k]
        for e := lres.Front(); e != nil; e = e.Next() {
            fmt.Println(e.Value)
        }
    }

    // UI - basically we would like to load the dictionary into memory,
    // then have a simple loop running that allows the user to query one
    // word at a time.
    // Once a query has been submitted, we can make the key and look up the
    // associated anagrams
}
