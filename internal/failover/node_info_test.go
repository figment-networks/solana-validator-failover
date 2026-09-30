package failover

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetTowerFileBytes(t *testing.T) {
	dir := t.TempDir()
	present := filepath.Join(dir, "tower.bin")
	require.NoError(t, os.WriteFile(present, []byte("tower"), 0600))
	missing := filepath.Join(dir, "missing.bin")

	t.Run("present file is read", func(t *testing.T) {
		n := &NodeInfo{TowerFile: present}
		require.NoError(t, n.SetTowerFileBytes(false))
		assert.Equal(t, []byte("tower"), n.TowerFileBytes)
		assert.Equal(t, n.ComputeTowerFileHashFromBytes([]byte("tower")), n.TowerFileHash)
	})

	t.Run("missing file errors by default", func(t *testing.T) {
		n := &NodeInfo{TowerFile: missing}
		assert.ErrorContains(t, n.SetTowerFileBytes(false), "failed to read tower file")
	})

	t.Run("missing file allowed sends an empty tower with a matching hash", func(t *testing.T) {
		n := &NodeInfo{TowerFile: missing}
		require.NoError(t, n.SetTowerFileBytes(true))
		assert.Empty(t, n.TowerFileBytes)
		assert.Equal(t, n.ComputeTowerFileHashFromBytes(n.TowerFileBytes), n.TowerFileHash)
	})

	t.Run("allowMissing does not hide other read errors", func(t *testing.T) {
		n := &NodeInfo{TowerFile: dir}
		assert.ErrorContains(t, n.SetTowerFileBytes(true), "failed to read tower file")
	})
}
