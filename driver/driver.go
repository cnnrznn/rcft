package main

import (
    "fmt"
    "github.com/cnnrznn/rcft"
    "math/rand"
)

func main() {
    n := 7
    //f := 3
    replicas := []rcft.Replica{}

    for i := 0; i < n; i++ {
        replicas = append(replicas, rcft.NewReplica(rand.Intn(2)))
    }

    fmt.Println("I'm alive!!!")
    fmt.Println(replicas)
}

