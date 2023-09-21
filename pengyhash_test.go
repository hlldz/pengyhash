package pengyhash

import (
	"testing"
)

func TestPengyHash(t *testing.T) {

	// Test case 1: Byte Array (Size 3751), Seed is Size
	var input1 [3751]byte
	expected1 := uint64(0x158dd38813559756)
	hash1 := pengyhash(input1[:], len(input1), uint64(len(input1)))

	if hash1 != expected1 {
		t.Errorf("Test-1 failed: expected %x, got %x", expected1, hash1)
	}

	// Test case 2: String, Seed is Size
	input2 := "Lorem ipsum"
	expected2 := uint64(0xe5426b567142f054)
	hash2 := pengyhash([]byte(input2), len(input2), uint64(len(input2)))

	if hash2 != expected2 {
		t.Errorf("Test-2 failed: expected %x, got %x", expected2, hash2)
	}

	// Test case 3: String, Seed is 12345
	input3 := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam aliquet purus eget interdum aliquet."
	expected3 := uint64(0x8d29571475067b7c)
	hash3 := pengyhash([]byte(input3), len(input3), uint64(12345))

	if hash3 != expected3 {
		t.Errorf("Test-3 failed: expected %x, got %x", expected3, hash3)
	}

}
