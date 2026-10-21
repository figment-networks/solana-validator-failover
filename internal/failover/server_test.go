package failover

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidateActiveGossipIdentity(t *testing.T) {
	tests := []struct {
		name           string
		actualIP       string
		actualPubkey   string
		expectedIP     string
		expectedPubkey string
		errorContains  string
	}{
		{
			name:           "matching IP and pubkey",
			actualIP:       "79.127.227.20",
			actualPubkey:   "active-pubkey",
			expectedIP:     "79.127.227.20",
			expectedPubkey: "active-pubkey",
		},
		{
			name:           "IP mismatch",
			actualIP:       "79.127.227.21",
			actualPubkey:   "active-pubkey",
			expectedIP:     "79.127.227.20",
			expectedPubkey: "active-pubkey",
			errorContains:  "does not match expected IP",
		},
		{
			name:           "pubkey mismatch",
			actualIP:       "79.127.227.20",
			actualPubkey:   "stale-pubkey",
			expectedIP:     "79.127.227.20",
			expectedPubkey: "active-pubkey",
			errorContains:  "does not match expected pubkey",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateActiveGossipIdentity(tt.actualIP, tt.actualPubkey, tt.expectedIP, tt.expectedPubkey)
			if tt.errorContains == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errorContains)
		})
	}
}
