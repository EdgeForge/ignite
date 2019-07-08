package imgcmd

import (
	"io"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"github.com/weaveworks/ignite/cmd/ignite/run"
	"github.com/weaveworks/ignite/pkg/errutils"
)

// NewCmdImport imports  a new VM image
func NewCmdImport(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import <OCI image>",
		Short: "Import a new base image for VMs",
		Long: dedent.Dedent(`
			Import a base image from an OCI image for VMs, takes in a Docker image as the source.
			This importing is done automatically when the run or create commands are run. This step
			is essentially a cache to be used later when running VMs.
		`),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			errutils.Check(func() error {
				_, err := run.ImportImage(args[0])
				return err
			}())
		},
	}
	return cmd
}
