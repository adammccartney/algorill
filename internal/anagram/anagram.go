package anagram

import (
    "bufio"
    "container/list"
    "fmt"
    "io"
    "os"
    "slices"
    "sort"
    "strings"
)



type ByAlpha []rune

func (a ByAlpha) Len() int      { return len(a) }
func (a ByAlpha) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAlpha) Less(i, j int) bool { return a[i] < a[j] }


func sortRunes(w string) []rune {
    res := []rune(w)
    sort.Sort(ByAlpha(res))
    return res
}


type Token struct {
    Word string
    Key string
}

func InitToken (word string) (Token, error) {
    // check that word is in fact a word, not two or more words seperated by
    // whitespace
    words := strings.Fields(word)

    if len(words) != 1 {
        return Token{}, fmt.Errorf("word contains more than one word: %s", word)
    }

    // sort runes
    runes := sortRunes(word)

    // make map of runes, {char: num_occur}
    m := make(map[rune]int)
    for _, r := range runes {
        m[r] += 1
    }

    var res string
    // deduplicate our sorted chars, then use as keys below
    seq := slices.Compact(runes)
    // convert map to string using our sorted runes
    for _, c := range seq {
        res += fmt.Sprintf("%c", c)
        val := m[c]
        if val > 1 {
            res += fmt.Sprintf("%d", val)
        }
    }

    return Token{Word: word, Key: res}, nil
}

type AnagramDict struct {
    Map map[string]*list.List
}

// Load values from reader into map
func (a *AnagramDict) Load (reader io.Reader) (error) {
    scanner := bufio.NewScanner(reader)
    // pointer used to manage insertion of words into list for each key in map
    var l *list.List
    for scanner.Scan() {
        t, terr := InitToken(scanner.Text())
        if terr != nil {
            return terr
        }
        // use the Token's key to create a.Mapap entry
        if a.Map[t.Key] == nil {  // associate first value
            l = list.New()
        } else {  // Key already exists, found and anagram
            l = a.Map[t.Key]
        }
        l.PushBack(t.Word)
        a.Map[t.Key] = l
    }
    var err error
    if err = scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading: ", err)
    }

    return err
}

// Print anagrams associated with a specific key to stdout, one per line
func (a *AnagramDict) Printer (t Token) {
    fmt.Printf("Anagrams for %s are:\n", t.Word)
    l := a.Map[t.Key]
    for e := l.Front(); e != nil; e = e.Next() {
        // do something with e.Value
        fmt.Printf("\t\t%s\n", e.Value)
    }
}
