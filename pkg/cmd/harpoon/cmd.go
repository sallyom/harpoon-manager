package harpoon

import (
	"context"

	"github.com/spf13/cobra"
	//	"k8s.io/apimachinery/pkg/version"
)

//var (
//	harpoonVersion = "0.0.0"
//)

// NewHarpoonCommand implements the harpoon command
func NewHarpoonStartCommand(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   name,
		Short: "Start the Harpoon GitOps controllers",
		Run: func(cmd *cobra.Command, args []string) {
			RunHarpoon(context.Background())
		},
	}
	return cmd
}
