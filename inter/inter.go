package main

import (
        "fmt"
        "os"
)

func main() {
        if len(os.Args) != 3 {
                return
        }

        s1 := os.Args[1]
        s2 := os.Args[2]

        seen := make(map[rune]bool)
        result := ""
        s2Map := make(map[rune]bool)

        for _, r := range s2 {
                s2Map[r] = true
        }

        for _, r := range s1 {
                if !seen[r] && s2Map[r] {
                        result += string(r)
                        seen[r] = true
                }
        }

        fmt.Println(result)
}