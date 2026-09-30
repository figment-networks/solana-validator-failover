package failover

import (
	"errors"
	"fmt"
	"os"

	"github.com/sol-strategies/solana-validator-failover/internal/identities"
	"github.com/zeebo/xxh3"
)

// NodeInfo represents the information about a node that is needed to perform a failover
type NodeInfo struct {
	PublicIP                       string
	Hostname                       string
	Identities                     *identities.Identities
	TowerFile                      string
	TowerFileBytes                 []byte
	TowerFileHash                  string
	SetIdentityCommand             string
	ClientVersion                  string
	SolanaValidatorFailoverVersion string
	RPCAddress                     string
}

// SetTowerFileBytes sets the tower file bytes; with allowMissing a missing file sends an empty tower
func (n *NodeInfo) SetTowerFileBytes(allowMissing bool) error {
	towerFileBytes, err := os.ReadFile(n.TowerFile)
	if err != nil && !(allowMissing && errors.Is(err, os.ErrNotExist)) {
		return fmt.Errorf("failed to read tower file: %w", err)
	}
	n.TowerFileBytes = towerFileBytes
	n.setTowerFileHash()
	return nil
}

// SetTowerFileHash sets the tower file hash
func (n *NodeInfo) setTowerFileHash() {
	n.TowerFileHash = n.ComputeTowerFileHashFromBytes(n.TowerFileBytes)
}

// ComputeTowerFileHashFromBytes computes the tower file hash from the tower file bytes
func (n NodeInfo) ComputeTowerFileHashFromBytes(towerFileBytes []byte) string {
	hash := xxh3.Hash(towerFileBytes)
	return fmt.Sprintf("xxh3:%x", hash)
}
