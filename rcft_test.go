package rcft

import (
    "fmt"
    "github.com/cnnrznn/dsdriver"
    "math/rand"
    "sync"
    "testing"
    "time"
)

func TestReplicaInit(t *testing.T) {
    r := NewReplica(0)
    fmt.Println(r)
}

func TestReplicaConsensus(t *testing.T) {
    n := 101
    f := 50
    replicas := make([]Replica, n)

    rand.Seed(time.Now().UnixNano())

    for i := range replicas {
        replicas[i] = NewReplica(rand.Intn(2))
    }

    fr, tos := dsdriver.Local(n, dsdriver.ReorderHub)

    var wg sync.WaitGroup
    wg.Add(n)
    defer wg.Wait()

    for i, r := range replicas {
        go r.Consensus(n, f, fr, tos[i], &wg)
    }
}
