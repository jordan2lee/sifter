
package get

import (
  "log"
	"github.com/spf13/cobra"
)

var siftPath = "./"
var siftRepo = "github.com/bmeg/sifter-tools"

// Cmd is the declaration of the command line
var Cmd = &cobra.Command{
	Use:   "get",
	Short: "Get sifter importer",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

    

    log.Printf("Getting: %s", args[0])
    return nil
  },
}


func init() {
	flags := Cmd.Flags()
	flags.StringVar(&siftPath, "sift", siftPath, "SIFT Path")
}
