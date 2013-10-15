package main

import (
  "crypto/sha256"
  "encoding/hex"
)

func FmtName(packageStr string) string {
  bytes := []byte(packageStr)
  hasher := sha256.New()
  hasher.Write(bytes)
  firstSix := hasher.Sum(nil)[0:3]
  return hex.EncodeToString(firstSix)
}
