package addon

import (
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/spf13/cobra"
)

var List []*Addon

type Addon struct {
	Name        string
	Version     string
	Author      string
	Repository  string
	Description string

	// This command should be implemented by the addon side to register command in the active cli
	RegisterCommands func(rootCmd *cobra.Command) error
}

// RegisterAddons should not be exported but for simplicity it is.
func RegisterAddons(rootCmd *cobra.Command) {
	_ = filepath.Walk("addons", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		openedAddon, err := plugin.Open(path)
		if err != nil {
			log.Println("addon error:", err)
			return nil
		}

		addon, err := openedAddon.Lookup("Addon")
		if err != nil {
			log.Println("addon error:", err)
			return nil
		}

		if err := addon.(*Addon).RegisterCommands(rootCmd); err != nil {
			log.Println("addon error:", err)
			return nil
		}

		List = append(List, addon.(*Addon))

		return nil
	})
}
