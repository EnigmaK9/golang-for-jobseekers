package main

import (
    "fmt"
)

type Node struct {
    data  int
    left  *Node
    right *Node
}

func InorderPrint(root *Node) {
    if root == nil {
        return
    }
    if root.left != nil {
        InorderPrint(root.left)
    }
    fmt.Println(root.data)
    if root.right != nil {
        InorderPrint(root.right)
    }
}

func main() {
    root := &Node{data: 1}
    root.left = &Node{data: 2}
    root.right = &Node{data: 3}
    root.left.left = &Node{data: 4}
    root.left.right = &Node{data: 5}
    InorderPrint(root)
}

