package pubsub

import (
    "testing"
    "sync"
)

func TestPubsub(t *testing.T) {
    ps := New()
    if ps.NumTopic() != 0 {
        t.Errorf("ps.NumTopic() != 0")
    }
    ch1 := make(chan TopicMessage)
    ch2 := make(chan TopicMessage)
    ch3 := make(chan TopicMessage)
    ps.Sub(ch1, "test")
    ps.Sub(ch2, "test")
    ps.Sub(ch3, "test")
    if ps.NumTopic() != 1 {
        t.Errorf("ps.NumTopic() != 1")
    }
    if ps.NumSub("test") != 3 {
        t.Errorf("ps.NumSub() != 3")
    }
    msg1, msg2, msg3 := TopicMessage{}, TopicMessage{}, TopicMessage{}

    var wg sync.WaitGroup
    wg.Add(3)

    go func() {
        msg1 = <- ch1
        wg.Done()
    }()

    go func() {
        msg2 = <- ch2
        wg.Done()
    }()

    go func() {
        msg3 = <- ch3
        wg.Done()
    }()

    go func(ps *PubSub) {
        n := ps.Pub("hello", "test")
        if n != 3 {
            t.Errorf("n != 3")
        }
    }(ps)

    wg.Wait()

    t.Logf("%+v, %+v, %+v, %+v", ps.Topics(), msg1, msg2, msg3)
    if msg1 != msg2 || msg2 != msg3 {
        t.Errorf("msg1 != msg2 || msg2 != msg3")
    }
    ps.RemoveTopics("test")
    if ps.NumTopic() != 0 {
        t.Errorf("ps.NumTopic() != 0")
    }
    if ps.NumSub("test") != 0 {
        t.Errorf("ps.NumSub() != 0")
    }
}
