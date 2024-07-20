package main

import "fmt"

func main() {
    abcd := []int{12, 33, 44, 21, 55, 88, 24, 43, 2, 7, 78, 90, 42}
    fmt.Println(BubbleSort(abcd))
}

func BubbleSort(abcd []int) []int {
    for i := 0; i < len(abcd)-1; i++ {
        for j := 0; j < len(abcd)-i-1; j++ {
            if abcd[j] > abcd[j+1] {
                abcd[j], abcd[j+1] = abcd[j+1], abcd[j]
            }
        }
    }
    return abcd
}

