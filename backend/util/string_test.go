package util

import "testing"

func TestGenerateRandomString(t *testing.T) {
	t.Run("Generate random string", func(t *testing.T) {
		randomString, err := GenerateRandomString(32)
		if err != nil {
			t.Errorf("Error when generating random string: %s", err)
		}

		if len(randomString) != 32 {
			t.Errorf("The length of random string is not 32")
		}
	})

	t.Run("Generate random string with length 0", func(t *testing.T) {
		randomString, err := GenerateRandomString(0)
		if err != nil {
			t.Errorf("Error when generating random string: %s", err)
		}

		if len(randomString) != 0 {
			t.Errorf("The length of random string is not 0")
		}
	})
}

func TestToPointerString(t *testing.T) {
	t.Run("Convert string to pointer string", func(t *testing.T) {
		pointerString := ToPointerString("test")
		if *pointerString != "test" {
			t.Errorf("The value of pointer string is not 'test'")
		}
	})

	t.Run("Convert empty string to pointer string", func(t *testing.T) {
		pointerString := ToPointerString("")
		if *pointerString != "" {
			t.Errorf("The value of pointer string is not ''")
		}
	})
}
