package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type HeaderInfo struct {
	SendCount          uint64
	L1BlockNumber      uint64
	ArbOSFormatVersion uint64
}

// Function to generate the mixDigest
func (info HeaderInfo) mixDigest() [32]byte {
	mixDigest := [32]byte{}
	binary.BigEndian.PutUint64(mixDigest[:8], info.SendCount)
	binary.BigEndian.PutUint64(mixDigest[8:16], info.L1BlockNumber)
	binary.BigEndian.PutUint64(mixDigest[16:24], info.ArbOSFormatVersion)
	return mixDigest
}

func main() {
	info := HeaderInfo{
		SendCount:          0,
		L1BlockNumber:      0,
		ArbOSFormatVersion: 40,
	}

	digest := info.mixDigest()
	fmt.Println("mixDigest (hex):", hex.EncodeToString(digest[:]))
	fmt.Println("mixDigest (bytes):", digest)
}
