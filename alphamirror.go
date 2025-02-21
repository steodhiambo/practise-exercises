package main

import (
        "fmt"
        "os"
)

func alphamirror(s string) string {
        result := ""
        for _, char := range s {
                if 'a' <= char && char <= 'z' {
                        result += string('z' - (char - 'a'))
                } else if 'A' <= char && char <= 'Z' {
                        result += string('Z' - (char - 'A'))
                } else {
                        result += string(char)
                }
        }
        return result
}

func main() {
        if len(os.Args) != 2 {
                fmt.Println()
        } else {
                fmt.Println(alphamirror(os.Args[1]))
        }
}