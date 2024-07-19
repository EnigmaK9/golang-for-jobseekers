package main

import "fmt"

type Node struct {
    data  string
    left  *Node
    right *Node
}

func main() {
    A := &Node{data: "A"}
    B := &Node{data: "B"}
    C := &Node{data: "C"}
    D := &Node{data: "D"}
    E := &Node{data: "E"}
    F := &Node{data: "F"}
    G := &Node{data: "G"}

    A.left = B
    A.right = C
    B.left = D
    B.right = E
    C.left = F
    C.right = G

    fmt.Println("Inorder traversal:")
    inorder(A)
    fmt.Println()

    fmt.Println("Preorder traversal:")
    preorder(A)
    fmt.Println()

    fmt.Println("Postorder traversal:")
    postorder(A)
    fmt.Println()

    fmt.Println("Level order traversal:")
    levelOrder(A)
    fmt.Println()
}

func inorder(node *Node) {
    if node != nil {
        inorder(node.left)
        fmt.Printf("%s ", node.data)
        inorder(node.right)
    }
}

func preorder(node *Node) {
    if node != nil {
        fmt.Printf("%s ", node.data)
        preorder(node.left)
        preorder(node.right)
    }
}

func postorder(node *Node) {
    if node != nil {
        postorder(node.left)
        postorder(node.right)
        fmt.Printf("%s ", node.data)
    }
}

func levelOrder(root *Node) {
    if root == nil {
        return
    }
    queue := []*Node{root}
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        fmt.Printf("%s ", current.data)
        if current.left != nil {
            queue = append(queue, current.left)
        }
        if current.right != nil {
            queue = append(queue, current.right)
        }
    }
}

