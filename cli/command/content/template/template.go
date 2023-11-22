package template

import (
	"github.com/Tinyblargon/proxmox-go-sdk/cli/command/content"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "With this command you can manage Lxc container templates in proxmox",
}

func init() {
	content.ContentCmd.AddCommand(templateCmd)
}
