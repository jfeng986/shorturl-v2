package util

import (
	"os"
	"testing"

	qrcode "github.com/skip2/go-qrcode"
)

func TestGenerateQRCode(t *testing.T) {
	content := "https://github.com/"
	filename := "test_qr.png"

	err := GenerateQRCode(content, qrcode.Medium, 256, filename)
	if err != nil {
		t.Fatalf("Failed to generate QR code: %s", err)
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Fatalf("File %s was not generated", filename)
	}

	os.Remove(filename)
}
