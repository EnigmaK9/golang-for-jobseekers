package main

import (
    "fmt" // Importa el paquete fmt para imprimir mensajes en la consola
)

// Define una estructura llamada Stack que contiene un slice de enteros
type Stack struct {
    stack []int
}

// NewStack inicializa una nueva instancia de Stack
func NewStack() Stack {
    return Stack{stack: []int{}} // Devuelve una nueva pila vacía
}

// AddToStack agrega un nuevo elemento a la pila
func (s *Stack) AddToStack(item int) {
    s.stack = append(s.stack, item) // Añade el elemento al final del slice
}

// RemoveItemFromStack elimina el último elemento de la pila y lo devuelve
func (s *Stack) RemoveItemFromStack() (int, error) {
    if len(s.stack) == 0 { // Comprueba si la pila está vacía
        return 0, fmt.Errorf("stack is empty") // Devuelve un error si la pila está vacía
    }
    temp := s.stack[len(s.stack)-1] // Guarda el último elemento del slice
    s.stack = s.stack[0 : len(s.stack)-1] // Elimina el último elemento del slice
    return temp, nil // Devuelve el elemento eliminado
}

func main() {
    aa := NewStack() // Crea una nueva pila llamada aa

    aa.AddToStack(1) // Añade el número 1 a la pila
    aa.AddToStack(2) // Añade el número 2 a la pila

    fmt.Println(aa.stack) // Imprime el contenido de la pila

    // Intenta eliminar un elemento de la pila y manejar el error si ocurre
    item, err := aa.RemoveItemFromStack()
    if err != nil {
        fmt.Println(err) // Imprime el error si la pila está vacía
    } else {
        fmt.Println("Removed item:", item) // Imprime el elemento eliminado
    }

    // Intenta eliminar otro elemento de la pila y manejar el error si ocurre
    item, err = aa.RemoveItemFromStack()
    if err != nil {
        fmt.Println(err) // Imprime el error si la pila está vacía
    } else {
        fmt.Println("Removed item:", item) // Imprime el elemento eliminado
    }
}

