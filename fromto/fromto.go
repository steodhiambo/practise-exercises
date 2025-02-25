package main

import "fmt"

func FromTo(from int, to int) string {
        if from > 99 || from < 0 || to > 99 || to < 0 {
                return "Invalid\n"
        }

        result := ""

        if from <= to {
                for i := from; i <= to; i++ {
                        if i < 10 {
                                result += "0"
                        }
                        result += convertIntToString(i)
                        if i < to {
                                result += ", "
                        }
                }
        } else {
                for i := from; i >= to; i-- {
                        if i < 10 {
                                result += "0"
                        }
                        result += convertIntToString(i)
                        if i > to {
                                result += ", "
                        }
                }
        }

        result += "\n"
        return result
}

func convertIntToString(n int) string {
    if n == 0 {
        return "0"
    }
    s := ""
    for n > 0 {
        digit := n % 10
        s = string(rune('0'+digit)) + s
        n /= 10
    }
    return s
}
func main() {
    fmt.Print(FromTo(1, 10))
    fmt.Print(FromTo(10, 1))
    fmt.Print(FromTo(10, 10))
    fmt.Print(FromTo(100, 10))
}
