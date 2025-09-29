// internal/util/id.go
// Генерация уникальных ID с помощью crypto/rand.

package util

import (
    "crypto/rand"
    "encoding/hex"
)

func NewID() string {
    b := make([]byte, 16)
    rand.Read(b)
    return hex.EncodeToString(b)
}
