package payloadlength

import (
    "testing"
)

func TestPayloadlength(t *testing.T) {
    length := uint64(0)
    for ; length < max64Bits + 1; length++ {
        cnb, err := CountBytes64(length)
        if err != nil {
            t.Errorf("CountBytes64 failed: %s", err)
        }
        buf, enb, err := Encode64(length)
        if err != nil {
            t.Errorf("Encode64 failed: %s", err)
        }
        lengthdec, dnb, err := Decode64(buf)
        if err != nil {
            t.Errorf("Decode64 failed: %s", err)
        }
        if length != lengthdec {
            t.Errorf("%d != %d", length, lengthdec)
        }
        if cnb != enb {
            t.Errorf("cnb != enb")
        }
        if enb != dnb {
            t.Errorf("enb != dnb")
        }
    }
}
