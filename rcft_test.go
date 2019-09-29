package rcft

import (
    "fmt"
    "testing"
)

func TestReplicaInit(t *testing.T) {
    r := new_replica(0)
    fmt.Println(r)
}

func TestReplicaConsensus(t *testing.T) {
    r := new_replica(0)
    fmt.Println(r)

    ch := make(chan Event)

    r.Consensus(9, 4, ch)
}
