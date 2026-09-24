package failover

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasVoteHistoryFile_Missing(t *testing.T) {
	n := NodeInfo{VoteHistoryFile: filepath.Join(t.TempDir(), "vote_history-nope.bin")}
	assert.False(t, n.HasVoteHistoryFile(), "pre-Alpenglow clusters have no vote history file")
}

func TestHasVoteHistoryFile_Directory(t *testing.T) {
	n := NodeInfo{VoteHistoryFile: t.TempDir()}
	assert.False(t, n.HasVoteHistoryFile(), "a directory must not be treated as a vote history file")
}

func TestSetVoteHistoryFileBytes(t *testing.T) {
	path := filepath.Join(t.TempDir(), "vote_history-abc.bin")
	want := []byte("alpenglow vote history")
	require.NoError(t, os.WriteFile(path, want, 0o600))

	n := NodeInfo{VoteHistoryFile: path}
	require.True(t, n.HasVoteHistoryFile())
	require.NoError(t, n.SetVoteHistoryFileBytes())

	assert.Equal(t, want, n.VoteHistoryFileBytes)
	assert.Equal(t, int64(len(want)), n.VoteHistoryFileSizeBytes)
	assert.Equal(t, n.ComputeTowerFileHashFromBytes(want), n.VoteHistoryFileHash)
}

func TestSetVoteHistoryFileBytes_Missing(t *testing.T) {
	n := NodeInfo{VoteHistoryFile: filepath.Join(t.TempDir(), "absent.bin")}
	assert.Error(t, n.SetVoteHistoryFileBytes())
}

// The receiving side recomputes the hash to detect a corrupted transfer before it sets the
// identity, so a mutated payload must not match.
func TestVoteHistoryHashDetectsCorruption(t *testing.T) {
	path := filepath.Join(t.TempDir(), "vote_history-abc.bin")
	require.NoError(t, os.WriteFile(path, []byte("original"), 0o600))

	n := NodeInfo{VoteHistoryFile: path}
	require.NoError(t, n.SetVoteHistoryFileBytes())

	assert.NotEqual(t, n.VoteHistoryFileHash, n.ComputeTowerFileHashFromBytes([]byte("tampered")))
}
