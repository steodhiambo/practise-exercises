package main 
import (
     "fmt"
//     "os"
//     "strconv"
// 
)



func lastWord(s string) string {
        word := ""
        words := []string{}
        for _, r := range s {
                if r == ' ' {
                        if word != "" { // Only append if word is not empty
                                words = append(words, word)
                                word = ""
                        }
                } else {
                        word += string(r)
                }

        }
        if word != "" { // Append the last word (if any)
                words = append(words, word)
        }

        if len(words) == 0 { // Handle the case where there are no words
                return ""
        }

    for i := len(words) -1; i >=0; i-- {
        if len(words[i]) > 0 {
            return words[i]
        }
    }
    return ""

}

func main() {
        fmt.Println(lastWord("this        ...       is sparta, then again, maybe    not"))
        fmt.Println(lastWord(" lorem,ipsum "))
        fmt.Println(lastWord(" "))
    fmt.Println(lastWord(""))
    fmt.Println(lastWord("       "))

}