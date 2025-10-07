//go:build !windows

package updater

func hideFile(_ string) error {
	return nil
}
