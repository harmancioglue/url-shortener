package services

import (
	"testing"
	"time"

	"harmancioglue/url-shortener/internal/common/utils"
)

func TestSnowflakeIDGenerator(t *testing.T) {
	generator, err := NewSnowflakeIDGenerator(1)
	if err != nil {
		t.Fatalf("Failed to create Snowflake ID generator: %v", err)
	}

	ids := make(map[int64]bool)
	for i := 0; i < 1000; i++ {
		id, err := generator.GenerateID()
		if err != nil {
			t.Fatalf("Failed to generate ID: %v", err)
		}

		if ids[id] {
			t.Fatalf("Duplicate ID generated: %d", id)
		}
		ids[id] = true

		if id <= 0 {
			t.Fatalf("Generated ID should be positive, got: %d", id)
		}
	}
}

func TestSnowflakeIDGeneratorInvalidWorkerID(t *testing.T) {
	_, err := NewSnowflakeIDGenerator(1024)
	if err == nil {
		t.Fatal("Expected error for worker ID 1024, got nil")
	}

	_, err = NewSnowflakeIDGenerator(-1)
	if err == nil {
		t.Fatal("Expected error for negative worker ID, got nil")
	}
}

func TestBase62Encoding(t *testing.T) {
	testCases := []struct {
		input    int64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{10, "a"},
		{61, "Z"},
		{62, "10"},
		{3844, "100"}, // 62^2
		{123456789, "8m0Kx"},
	}

	for _, tc := range testCases {
		result := utils.Encode(tc.input)
		if result != tc.expected {
			t.Errorf("Encode(%d) = %s, expected %s", tc.input, result, tc.expected)
		}

		// Test decoding
		decoded, err := utils.Decode(result)
		if err != nil {
			t.Errorf("Decode(%s) failed: %v", result, err)
		}
		if decoded != tc.input {
			t.Errorf("Decode(%s) = %d, expected %d", result, decoded, tc.input)
		}
	}
}

func TestSnowflakeIDGeneratorConcurrency(t *testing.T) {
	generator, err := NewSnowflakeIDGenerator(1)
	if err != nil {
		t.Fatalf("Failed to create Snowflake ID generator: %v", err)
	}

	numGoroutines := 10
	numIDsPerGoroutine := 100
	idChan := make(chan int64, numGoroutines*numIDsPerGoroutine)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			for j := 0; j < numIDsPerGoroutine; j++ {
				id, err := generator.GenerateID()
				if err != nil {
					t.Errorf("Failed to generate ID: %v", err)
					return
				}
				idChan <- id
			}
		}()
	}

	ids := make(map[int64]bool)
	for i := 0; i < numGoroutines*numIDsPerGoroutine; i++ {
		id := <-idChan
		if ids[id] {
			t.Fatalf("Duplicate ID generated in concurrent test: %d", id)
		}
		ids[id] = true
	}
}

func TestSnowflakeIDGeneratorPerformance(t *testing.T) {
	generator, err := NewSnowflakeIDGenerator(1)
	if err != nil {
		t.Fatalf("Failed to create Snowflake ID generator: %v", err)
	}

	numIDs := 100000
	start := time.Now()

	for i := 0; i < numIDs; i++ {
		_, err := generator.GenerateID()
		if err != nil {
			t.Fatalf("Failed to generate ID: %v", err)
		}
	}

	duration := time.Since(start)
	idsPerSecond := float64(numIDs) / duration.Seconds()

	t.Logf("Generated %d IDs in %v (%.0f IDs/second)", numIDs, duration, idsPerSecond)

	if idsPerSecond < 100000 {
		t.Errorf("Performance too slow: %.0f IDs/second, expected at least 100,000", idsPerSecond)
	}
}
