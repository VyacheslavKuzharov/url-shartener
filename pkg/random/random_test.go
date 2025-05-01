package random

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenShortKey(t *testing.T) {
	testCases := []struct {
		name        string
		keyLen      int
		expectedKey string
	}{
		{
			name:        "shortKey length is equal 6",
			keyLen:      6,
			expectedKey: GenShortKey(),
		},
		{
			name:        "expectedKey not equal new Key",
			keyLen:      6,
			expectedKey: GenShortKey(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newKey := GenShortKey()

			require.Equal(t, tc.keyLen, len(newKey))
			assert.NotEqual(t, tc.expectedKey, newKey)
		})
	}
}
