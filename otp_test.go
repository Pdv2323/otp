package otp

import (
	"testing"
)

// TestGenerateOtp tests the GenerateOtp function
func TestGenerateOtp(t *testing.T) {
	// Test for 4 digit OTP
	otp := GenerateOtp(4)
	if otp < 1000 || otp > 9999 {
		t.Errorf("Expected a 4 digit OTP, got %d", otp)
	}

	// Test for 6 digit OTP
	otp = GenerateOtp(6)
	if otp < 100000 || otp > 999999 {
		t.Errorf("Expected a 6 digit OTP, got %d", otp)
	}

	// Test for invalid input (0 or negative number)
	otp = GenerateOtp(0)
	if otp != 0 {
		t.Errorf("Expected 0 for invalid input, got %d", otp)
	}

	otp = GenerateOtp(-1)
	if otp != 0 {
		t.Errorf("Expected 0 for invalid input, got %d", otp)
	}
}

// TestVerifyOtp tests the VerifyOtp function
func TestVerifyOtp(t *testing.T) {
	// Test case where OTPs match
	otp := 123456
	newOtp := 123456
	if !VerifyOtp(otp, newOtp) {
		t.Errorf("Expected OTPs to match, but they didn't")
	}

	// Test case where OTPs do not match
	otp = 123456
	newOtp = 654321
	if VerifyOtp(otp, newOtp) {
		t.Errorf("Expected OTPs to not match, but they did")
	}
}
