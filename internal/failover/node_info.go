package failover

import (
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
	TowerFileSizeBytes             int64
	TowerFileBytes                 []byte
	TowerFileHash                  string
	VoteHistoryFile                string
	VoteHistoryFileSizeBytes       int64
	VoteHistoryFileBytes           []byte
	VoteHistoryFileHash            string
	SetIdentityCommand             string
	ClientVersion                  string
	SolanaValidatorFailoverVersion string
	RPCAddress                     string
}

// SetTowerFileBytes sets the tower file bytes
func (n *NodeInfo) SetTowerFileBytes() error {
	towerFileBytes, err := os.ReadFile(n.TowerFile)
	if err != nil {
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

// HasVoteHistoryFile reports whether this node has a vote history file on disk.
// Pre-Alpenglow clusters have none, so its absence is not an error.
func (n NodeInfo) HasVoteHistoryFile() bool {
	info, err := os.Stat(n.VoteHistoryFile)
	return err == nil && !info.IsDir()
}

// SetVoteHistoryFileBytes sets the vote history file bytes
func (n *NodeInfo) SetVoteHistoryFileBytes() error {
	voteHistoryFileBytes, err := os.ReadFile(n.VoteHistoryFile)
	if err != nil {
		return fmt.Errorf("failed to read vote history file: %w", err)
	}
	n.VoteHistoryFileBytes = voteHistoryFileBytes
	n.VoteHistoryFileSizeBytes = int64(len(voteHistoryFileBytes))
	n.VoteHistoryFileHash = n.ComputeTowerFileHashFromBytes(voteHistoryFileBytes)
	return nil
}
