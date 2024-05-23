package rand

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

func GenerateTraceId() (string, error) {
	uid, err := Math(16)

	if err != nil {
		return "00000000000000000000000000000000", err
	}
	return hex.EncodeToString(uid), nil
}

func Math(bufLen int) ([]byte, error) {
	buf := make([]byte, bufLen)
	n, err := rand.Read(buf)
	if n != len(buf) || err != nil {
		return nil, errors.New("math rand error")
	}
	return buf, nil
}
