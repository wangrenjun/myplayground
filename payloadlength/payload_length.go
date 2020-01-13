// 
package payloadlength

import (
    "errors"
    "encoding/binary"
)

var ErrMalformedPayloadLength = errors.New("Malformed payload length")
var ErrPayloadLengthTooLarge = errors.New("Payload length too large")

const (
    max8Bits  = 1 << 6 - 1
    max16Bits = 1 << 14 - 1
    max32Bits = 1 << 30 - 1
    max64Bits = 1 << 62 - 1
)

const (
    mark8Bits = 0
    mark16Bits = 1 << 6
    mark32Bits = 2 << 6
    mark64Bits = 3 << 6
    mask = uint8(3 << 6)
)

func CountBytes16(length uint16) (int, error) {
    switch {
    case length <= max8Bits:
        return 1, nil
    case length <= max16Bits:
        return 2, nil
    }
    return 0, ErrPayloadLengthTooLarge
}

func Encode16(length uint16) (buf []byte, nbytes int, err error) {
    switch {
    case length <= max8Bits:
        nbytes = 1
        buf = append(buf, byte(length))
    case length <= max16Bits:
        nbytes = 2
        buf = make([]byte, nbytes)
        binary.BigEndian.PutUint16(buf, uint16(length))
        buf[0] |= mark16Bits
    default:
        err = ErrPayloadLengthTooLarge
    }
    return
}

func Decode16(buf []byte) (length uint16, nbytes int, err error) {
    if len(buf) == 0 {
        err = ErrMalformedPayloadLength
        return
    }
    switch buf[0] & mask {
    case mark8Bits:
        nbytes = 1
        length = uint16(buf[0])
    case mark16Bits:
        nbytes = 2
        if len(buf) < nbytes {
            err = ErrMalformedPayloadLength
            return
        }
        unmarkedbuf := append([]byte{}, buf...)
        unmarkedbuf[0] &= ^uint8(mark16Bits)
        length = uint16(binary.BigEndian.Uint16(unmarkedbuf))
    default:
        err = ErrMalformedPayloadLength
        return
    }
    return
}

func CountBytes64(length uint64) (int, error) {
    switch {
    case length <= max8Bits:
        return 1, nil
    case length <= max16Bits:
        return 2, nil
    case length <= max32Bits:
        return 4, nil
    case length <= max64Bits:
        return 8, nil
    }
    return 0, ErrPayloadLengthTooLarge
}

func Encode64(length uint64) (buf []byte, nbytes int, err error) {
    switch {
    case length <= max8Bits:
        nbytes = 1
        buf = append(buf, byte(length))
    case length <= max16Bits:
        nbytes = 2
        buf = make([]byte, nbytes)
        binary.BigEndian.PutUint16(buf, uint16(length))
        buf[0] |= mark16Bits
    case length <= max32Bits:
        nbytes = 4
        buf = make([]byte, nbytes)
        binary.BigEndian.PutUint32(buf, uint32(length))
        buf[0] |= mark32Bits
    case length <= max64Bits:
        nbytes = 8
        buf = make([]byte, nbytes)
        binary.BigEndian.PutUint64(buf, uint64(length))
        buf[0] |= mark64Bits
    default:
        err = ErrPayloadLengthTooLarge
    }
    return
}

func Decode64(buf []byte) (length uint64, nbytes int, err error) {
    if len(buf) == 0 {
        err = ErrMalformedPayloadLength
        return
    }
    switch buf[0] & mask {
    case mark8Bits:
        nbytes = 1
        length = uint64(buf[0])
    case mark16Bits:
        nbytes = 2
        if len(buf) < nbytes {
            err = ErrMalformedPayloadLength
            return
        }
        unmarkedbuf := append([]byte{}, buf...)
        unmarkedbuf[0] &= ^uint8(mark16Bits)
        length = uint64(binary.BigEndian.Uint16(unmarkedbuf))
    case mark32Bits:
        nbytes = 4
        if len(buf) < nbytes {
            err = ErrMalformedPayloadLength
            return
        }
        unmarkedbuf := append([]byte{}, buf...)
        unmarkedbuf[0] &= ^uint8(mark32Bits)
        length = uint64(binary.BigEndian.Uint32(unmarkedbuf))
    case mark64Bits:
        nbytes = 8
        if len(buf) < nbytes {
            err = ErrMalformedPayloadLength
            return
        }
        unmarkedbuf := append([]byte{}, buf...)
        unmarkedbuf[0] &= ^uint8(mark64Bits)
        length = uint64(binary.BigEndian.Uint64(unmarkedbuf))
    default:
        err = ErrMalformedPayloadLength
        return
    }
    return
}
