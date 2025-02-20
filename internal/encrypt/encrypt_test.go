package encrypt

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestCalculateChecksum(t *testing.T) {
	testData := []byte("123456abcABC")

	tests := []struct {
		name      string
		algorithm string
		expected  func([]byte) string
	}{
		{
			name:      "MD5",
			algorithm: "MD5",
			expected: func(data []byte) string {
				hash := md5.Sum(data)
				return hex.EncodeToString(hash[:])
			},
		},
		{
			name:      "SHA256",
			algorithm: "SHA256",
			expected: func(data []byte) string {
				hash := sha256.Sum256(data)
				return hex.EncodeToString(hash[:])
			},
		},
		{
			name:      "SHA512",
			algorithm: "SHA512",
			expected: func(data []byte) string {
				hash := sha512.Sum512(data)
				return hex.EncodeToString(hash[:])
			},
		},
		{
			name:      "BCRYPT",
			algorithm: "BCRYPT",
			expected: func(data []byte) string {
				hash, _ := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
				return hex.EncodeToString(hash)
			},
		},
		{
			name:      "BASE64 (default case)",
			algorithm: "UNKNOWN",
			expected: func(data []byte) string {
				return base64.StdEncoding.EncodeToString(data)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := CalculateChecksum(testData, tc.algorithm)

			if tc.algorithm == "BCRYPT" {
				hashBytes, err := hex.DecodeString(result)
				if err != nil {
					t.Errorf("Failed to decode BCRYPT hash: %v", err)
				}
				if err := bcrypt.CompareHashAndPassword(hashBytes, testData); err != nil {
					t.Errorf("BCRYPT hash does not match input password")
				}
			} else {
				expected := tc.expected(testData)
				if result != expected {
					t.Errorf("Expected %s, got %s", expected, result)
				}
			}
		})
	}
}
