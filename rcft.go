package rcft

import (
    "fmt"
    "github.com/cnnrznn/dsdriver"
    "github.com/cnnrznn/util"
    "sync"
)

type Replica struct {
    value int
    cardinality int
    phaseno int
    witness_count [2]int
    message_count [2]int
}

type Message struct {
    value int
    cardinality int
    phaseno int
}

type Event struct {
    Pid int
    m Message
}

func (e Event) Dest() int {
    return e.Pid
}

func (e Event) String() string {
    return fmt.Sprintf("%v, %v", e.Pid, e.m)
}

func (r Replica) String() string {
    return fmt.Sprintf("v:%v, c:%v, p:%v, ws:%v, ms:%v",
                            r.value,
                            r.cardinality,
                            r.phaseno,
                            r.witness_count,
                            r.message_count)
}

func NewReplica(value int) Replica {
    return Replica { value : value,
                     cardinality : 1,
                     phaseno : 0,
                     witness_count : [2]int{},
                     message_count : [2]int{} }
}

func (r Replica) Consensus(n, f int, sendChan, recvChan chan dsdriver.Dester, wg *sync.WaitGroup) {
    defer wg.Done()

    decision := 0

    for r.witness_count[0] <= f && r.witness_count[1] <= f {
        r.message_count = [2]int{}
        r.witness_count = [2]int{}

        sendm := Message { value : r.value,
                           cardinality : r.cardinality,
                           phaseno : r.phaseno }

        for i := 0; i < n; i++ {
            sendChan <- Event {i, sendm}
        }

        for util.Sum(r.message_count[:]) < n - f {
            d := <-recvChan
            e := d.(Event)
            msg := e.m
            if msg.phaseno == r.phaseno {
                r.message_count[msg.value]++
                if msg.cardinality > n/2 {
                    r.witness_count[msg.value]++
                }
            } else if msg.phaseno > r.phaseno {
                sendChan <- e
            }
        }

        if r.witness_count[0] > 0 {
            r.value = 0
        } else if r.witness_count[1] > 0 {
            r.value = 1
        } else if r.message_count[0] > r.message_count[1] {
            r.value = 0
        } else {
            r.value = 1
        }

        r.cardinality = r.message_count[r.value]
        r.phaseno = r.phaseno + 1
    }

    if r.witness_count[0] > f {
        decision = 0
    } else {
        decision = 1
    }

    for i := 0; i < n; i++ {
        sendChan <- Event { i,
                      Message { value : r.value,
                                cardinality : n - f,
                                phaseno : r.phaseno } }
        sendChan <- Event { i,
                      Message { value : r.value,
                                cardinality : n - f,
                                phaseno : r.phaseno + 1 } }
    }

    fmt.Println("Decided: ", decision)
}

