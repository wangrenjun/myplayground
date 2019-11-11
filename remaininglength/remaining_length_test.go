package remaininglength

import (
    "testing"
)

func TestRemainLength(t *testing.T) {
    for length := 0; length < 268435456; length++ {
        buf, nbytes, err := Encode(length)
        if err != nil {
            t.Fail()
        }
        nbytes2, err := CountBytes(length)
        if err != nil {
            t.Fail()
        }
        length2, nbytes3, err := Decode(buf)
        if err != nil {
            t.Fail()
        }
        if length != length2 {
            t.Fail()
        }
        if nbytes != nbytes2 {
            t.Fail()
        }
        if nbytes != nbytes3 {
            t.Fail()
        }
    }
}
