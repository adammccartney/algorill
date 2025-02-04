package anagram

import (
    "fmt"
    "sort"
)



type ByAlpha []rune

func (a ByAlpha) Len() int      { return len(a) }
func (a ByAlpha) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByAlpha) Less(i, j int) bool { return a[i] < a[j] }


func sortChars(w string) []rune {
    res := []rune(w)
    sort.Sort(ByAlpha(res))
    return res
}


type Token struct {
    Word string
    Key string
}

func InitToken (word string) (t Token) {
    // sort chars
    chars := sortChars(word)

    // make map of chars, {char: num_occur}
    m := make(map[rune]int)
    for _, r := range chars {
        m[r] += 1
    }

    var res string
    // convert map to string
    for key, val := range m {
        res += fmt.Sprintf("%c", key)
        if val > 1 {
            res += fmt.Sprintf("%d", val)
        }
    }

    return Token{Word: word, Key: res}
}



