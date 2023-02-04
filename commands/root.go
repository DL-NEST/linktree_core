package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"linktree_core/commands/flag"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:     "linkTree",
		Long:    "一个私有化部署的智能家居系统",
		Version: "",
	}
	ConfigPath string
	LogPath    string
	Mode       string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "./", "设置配置文件的路径")
	rootCmd.PersistentFlags().StringVarP(&LogPath, "log", "l", "./logs", "设置日记文件的路径")
	rootCmd.AddCommand(&cobra.Command{
		Use:   flag.Start,
		Short: "Start linktree server",
		Run: func(cmd *cobra.Command, args []string) {
			Mode = cmd.Use
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   flag.StartBg,
		Short: "Start the linktree service in the background",
		Run: func(cmd *cobra.Command, args []string) {
			Mode = cmd.Use
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   flag.Stop,
		Short: "Stop linktree server",
		Run: func(cmd *cobra.Command, args []string) {
			Mode = cmd.Use
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   flag.Update,
		Short: "Update linktree server",
		Run: func(cmd *cobra.Command, args []string) {
			Mode = cmd.Use
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   flag.Reboot,
		Short: "Reboot linktree server",
		Run: func(cmd *cobra.Command, args []string) {
			Mode = cmd.Use
		},
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   flag.Pwd,
		Short: "linktree server temporary password",
		Run: func(cmd *cobra.Command, args []string) {
			Mode = cmd.Use
		},
	})
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
