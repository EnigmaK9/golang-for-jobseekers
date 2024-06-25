package main

import "fmt"

// BubbleSort ordena una lista de enteros utilizando el algoritmo de ordenamiento de burbuja
func BubbleSort(items []int) {
    // Si la longitud de items es 1 o menos, no hay nada que ordenar, así que retornamos de inmediato.
    if len(items) <= 1 {
        return
    }

    // Este bucle se repetirá hasta que no haya más elementos para intercambiar.
    for {
        sortHappened := false // Bandera que indica si ocurrió un intercambio en esta pasada.

        // Recorremos todos los elementos de la lista excepto el último.
        for i := 0; i < len(items)-1; i++ {
            // Si el elemento actual es mayor que el siguiente, los intercambiamos.
            if items[i] > items[i+1] {
                temp := items[i]         // Guardamos el valor del elemento actual en una variable temporal.
                items[i] = items[i+1]    // Colocamos el valor del siguiente elemento en la posición actual.
                items[i+1] = temp        // Colocamos el valor temporal en la posición del siguiente elemento.
                sortHappened = true      // Indicamos que ocurrió un intercambio.
            }
        }

        // Si no ocurrió ningún intercambio en esta pasada, significa que la lista ya está ordenada.
        if sortHappened == false {
            break // Salimos del bucle principal.
        }
    }
}

func main() {
    // Creamos un slice de enteros con los valores 4, 3, 2 y 1.
    values := []int{4, 3, 2, 1}

    // Imprimimos el slice antes de ordenar.
    fmt.Println("Antes de ordenar:", values)

    // Llamamos a la función BubbleSort para ordenar el slice.
    BubbleSort(values)

    // Imprimimos el slice después de ordenar.
    fmt.Println("Después de ordenar:", values)
}

