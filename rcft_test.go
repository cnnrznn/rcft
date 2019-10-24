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

//func TestReplicaConsensus(t *testing.T) {
//    n := 101
//    f := 50
//    replicas := make([]Replica, n)
//
//    rand.Seed(time.Now().UnixNano())
//
//    for i := range replicas {
//        replicas[i] = NewReplica(rand.Intn(2))
//    }
//
//    fr, tos := dsdriver.Local(n, dsdriver.ReorderHub)
//
//    var wg sync.WaitGroup
//    wg.Add(n)
//    defer wg.Wait()
//
//    for i, r := range replicas {
//        go r.Consensus(n, f, fr, tos[i], &wg)
//    }
//}

func runNode(replica Replica, n, f, i int, wg *sync.WaitGroup, nodes []string) {
	fr, to := dsdriver.Remote(i, nodes)
	replica.Consensus(n, f, fr, to, wg)
}

func TestReplicaRemote(t *testing.T) {
	n := 5
	f := 2
	replicas := make([]Replica, n)

	var wg sync.WaitGroup
	wg.Add(n)
	defer wg.Wait()

	rand.Seed(time.Now().UnixNano())

	nodes := []string{"127.0.0.1:3330",
		"127.0.0.1:3331",
		"127.0.0.1:3332",
		"127.0.0.1:3333",
		"127.0.0.1:3334"}

	for i := range replicas {
		replicas[i] = NewReplica(rand.Intn(2))
		go runNode(replicas[i], n, f, i, &wg, nodes)
	}
}
