package update

import (
	"github.com/Tinyblargon/proxmox-go-sdk/cli"
	"github.com/Tinyblargon/proxmox-go-sdk/proxmox"
	"github.com/spf13/cobra"
)

var update_snapshotCmd = &cobra.Command{
	Use:   "snapshot GUESTID SNAPSHOTNAME [DESCRIPTION]",
	Short: "Updates the description on the specified snapshot",
	Args:  cobra.RangeArgs(2, 3),
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		id := cli.ValidateIntIDset(args, "GuestID")
		snapName := cli.RequiredIDset(args, 1, "SnapshotName")
		des := cli.OptionalIDset(args, 2)
		err = proxmox.UpdateSnapshotDescription(cli.NewClient(), proxmox.NewVmRef(id), proxmox.SnapshotName(snapName), des)
		if err != nil {
			return
		}
		cli.PrintItemUpdated(updateCmd.OutOrStdout(), snapName, "Snapshot")
		return
	},
}

func init() {
	updateCmd.AddCommand(update_snapshotCmd)
}
