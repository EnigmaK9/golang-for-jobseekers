package main

import "fmt"

// MergeSort ordena una lista de enteros utilizando el algoritmo de ordenamiento por mezcla
func MergeSort(items []int) []int {
    if len(items) <= 1 {
        return items
    }

    leftSide := MergeSort(items[0 : len(items)/2])
    rightSide := MergeSort(items[len(items)/2 : len(items)])

    i, j := 0, 0
    combined := []int{}

    for i < len(leftSide) || j < len(rightSide) {
        if i >= len(leftSide) {
            combined = append(combined, rightSide[j:]...)
            j = len(rightSide)
            continue
        }
        if j >= len(rightSide) {
            combined = append(combined, leftSide[i:]...)
            i = len(leftSide)
            continue
        }
        if leftSide[i] <= rightSide[j] {
            combined = append(combined, leftSide[i])
            i++
        } else {
            combined = append(combined, rightSide[j])
            j++
        }
    }

    return combined
}

func main() {
    values := []int{4, 3, 2, 1}
    sorted := MergeSort(values)
    fmt.Println(sorted)
}

