package main

import "fmt"

// MergeSort ordena una lista de enteros utilizando el algoritmo de ordenamiento por mezcla
func MergeSort(items []int) []int {
    // Si la longitud de items es 1 o menos, no hay nada que ordenar, así que retornamos de inmediato.
    if len(items) <= 1 {
        return items
    }

    // Llama recursivamente a MergeSort en la primera mitad del slice items.
    leftSide := MergeSort(items[0 : len(items)/2])

    // Llama recursivamente a MergeSort en la segunda mitad del slice items.
    rightSide := MergeSort(items[len(items)/2 : len(items)])

    // Inicializa los índices i y j a 0.
    i, j := 0, 0

    // Inicializa un nuevo slice combined que contendrá los elementos ordenados de leftSide y rightSide.
    combined := []int{}

    // Un bucle que continúa mientras haya elementos en leftSide o rightSide para fusionar.
    for i < len(leftSide) || j < len(rightSide) {
        // Verifica si todos los elementos de leftSide ya han sido fusionados.
        if i >= len(leftSide) {
            combined = append(combined, rightSide[j:]...)
            j = len(rightSide)
            continue
        }

        // Verifica si todos los elementos de rightSide ya han sido fusionados.
        if j >= len(rightSide) {
            combined = append(combined, leftSide[i:]...)
            i = len(leftSide)
            continue
        }

        // Compara los elementos actuales de leftSide y rightSide.
        if leftSide[i] <= rightSide[j] {
            combined = append(combined, leftSide[i])
            i++
        } else {
            combined = append(combined, rightSide[j])
            j++
        }
    }

    // Devuelve el slice combined que contiene todos los elementos ordenados.
    return combined
}

func main() {
    // Inicializa un slice de enteros values con los valores [4, 3, 2, 1].
    values := []int{4, 3, 2, 1}

    // Llama a la función MergeSort para ordenar el slice values y almacena el resultado en sorted.
    sorted := MergeSort(values)

    // Imprime el slice sorted que contiene los valores ordenados.
    fmt.Println(sorted)
}

