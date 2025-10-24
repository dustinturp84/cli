package cmd

import (
	"github.com/spf13/cobra"
)

var instanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Manage CloudAMQP instances",
	Long:  `Create, list, update, and delete CloudAMQP instances.`,
}

func init() {
	instanceCmd.AddCommand(instanceCreateCmd)
	instanceCmd.AddCommand(instanceListCmd)
	instanceCmd.AddCommand(instanceGetCmd)
	instanceCmd.AddCommand(instanceUpdateCmd)
	instanceCmd.AddCommand(instanceDeleteCmd)
	instanceCmd.AddCommand(instanceResizeCmd)
	instanceCmd.AddCommand(instanceConfigCmd)
	instanceCmd.AddCommand(instanceNodesCmd)
	instanceCmd.AddCommand(instancePluginsCmd)
	// Action commands (flattened from actions subcommand)
	instanceCmd.AddCommand(restartRabbitMQCmd)
	instanceCmd.AddCommand(restartClusterCmd)
	instanceCmd.AddCommand(restartManagementCmd)
	instanceCmd.AddCommand(stopCmd)
	instanceCmd.AddCommand(startCmd)
	instanceCmd.AddCommand(rebootCmd)
	instanceCmd.AddCommand(stopClusterCmd)
	instanceCmd.AddCommand(startClusterCmd)
	instanceCmd.AddCommand(upgradeErlangCmd)
	instanceCmd.AddCommand(upgradeRabbitMQCmd)
	instanceCmd.AddCommand(upgradeRabbitMQErlangCmd)
	instanceCmd.AddCommand(upgradeVersionsCmd)
}
