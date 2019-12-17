// This package implements the encoding and decoding of the remaining length in the MQTT protocol.
package remaininglength

import (
    "errors"
)

var ErrMalformedRemainingLength = errors.New("Malformed remaining length")

func CountBytes(length int) (int, error) {
    if length >= 0 && length < 128 {
        return 1, nil
    } else if length < 16384 {
        return 2, nil
    } else if length < 2097152 {
        return 3, nil
    } else if length < 268435456 {
        return 4, nil
    } else {
        return 0, ErrMalformedRemainingLength
    }
}

func Encode(length int) ([]byte, int, error) {
    buf := make([]byte, 0)
    nbytes := 0
    for {
        encodedByte := length % 128
        length = length / 128
        if length > 0 {
            encodedByte = encodedByte | 128
        }
        buf = append(buf, byte(encodedByte))
        nbytes++
        if length <= 0 {
            break
        }
    }
    if nbytes > 4 {
        return buf, nbytes, ErrMalformedRemainingLength
    } else {
        return buf, nbytes, nil
    }
}

func Decode(buf []byte) (int, int, error) {
    length := 0
    nbytes := 0
    multiplier := 1
    for i := 0; i < len(buf); i++ {
        encodedByte := int(buf[i])
        if (multiplier > 2097152) {
            return 0, 0, ErrMalformedRemainingLength
        }
        length += (encodedByte & 127) * multiplier
        multiplier *= 128
        nbytes++
        if encodedByte & 128 == 0 {
            return length, nbytes, nil
        }
    }
    return length, nbytes, ErrMalformedRemainingLength
}
