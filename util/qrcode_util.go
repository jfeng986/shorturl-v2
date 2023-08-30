package util

import "github.com/skip2/go-qrcode"

const (
	// Level L: 7% error recovery.
	Low = iota

	// Level M: 15% error recovery. Good default choice.
	Medium

	// Level Q: 25% error recovery.
	High

	// Level H: 30% error recovery.
	Highest
)

func GenerateQRCode(content string, level qrcode.RecoveryLevel, size int, filename string) error {
	return qrcode.WriteFile(content, level, size, filename)
}
