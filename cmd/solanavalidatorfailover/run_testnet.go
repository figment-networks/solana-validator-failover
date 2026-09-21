//go:build testnet

package solanavalidatorfailover

var skipTowerFileCheck bool

func init() {
	runCmd.Flags().BoolVar(&skipTowerFileCheck, "skip-tower-file-check", false, "skip checking that the tower file exists and is non-empty before demoting")
}
