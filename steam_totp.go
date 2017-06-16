// Package inspired by https://github.com/DoctorMcKay/node-steam-totp
package steam_totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"time"
)

// ErrInvalidSharedSecret can be returned when shared secret isn't base64
// encoded string.
var ErrInvalidSharedSecret = errors.New("invalid base64 shared secret")

// GenerateAuthCode generate and returns 5 symbols authentication code.
// Shared Secret must be valid base64 string otherwise error will be returned.
func GenerateAuthCode(sharedSecret string, t time.Time) (string, error) {
	key, err := base64.StdEncoding.DecodeString(sharedSecret)
	if err != nil {
		return "", ErrInvalidSharedSecret
	}

	// Converting time for any reason
	// 00 00 00 00 00 00 00 00
	// 00 00 00 00 xx xx xx xx
	ut := uint64(t.Unix()) / 30
	tb := make([]byte, 8)
	binary.BigEndian.PutUint64(tb, ut)

	// Evaluate hash code for `tb` by key
	mac := hmac.New(sha1.New, key)
	mac.Write(tb)
	hashcode := mac.Sum(nil)

	// Last 4 bits provide initial position
	// len(hashcode) = 20 bytes
	start := hashcode[19] & 0xf

	// Extract 4 bytes at `start` and drop first bit
	fc32 := binary.BigEndian.Uint32(hashcode[start : start+4])
	fc32 &= 1<<31 - 1
	fullcode := int(fc32)

	// Range of possible chars for auth code.
	var chars = "23456789BCDFGHJKMNPQRTVWXY"
	var charsLen = len(chars)

	// Generate auth code
	code := make([]byte, 5)
	for i := range code {
		code[i] = chars[fullcode%charsLen]
		fullcode /= charsLen
	}

	return string(code[:]), nil
}
