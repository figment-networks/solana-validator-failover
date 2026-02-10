package identities

// Config holds the configuration for the identities this validator can assume
// depending on the role it is assigned.
//
// Each identity can be provided as either a keypair file path (active/passive)
// or a base58 public key string (active_pubkey/passive_pubkey). When a keypair
// file is provided it takes precedence. When only a public key is supplied, the
// daemon operates in pubkey-only mode for that identity.
type Config struct {
	Active        string `mapstructure:"active"`
	ActivePubkey  string `mapstructure:"active_pubkey"`
	Passive       string `mapstructure:"passive"`
	PassivePubkey string `mapstructure:"passive_pubkey"`
}
