/*
Copyright © 2022 poorops poorops@163.com

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/utils/strings/slices"
	"kubectl-check/pkg/base"
	"kubectl-check/pkg/check"
	"kubectl-check/pkg/client"
	"kubectl-check/pkg/table"
	"os"
	"strings"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
	field                 string
	selector              string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-check",
	Short: "check field from k8s resource definition",
	Long: `kubectl-check is a plugin for kubectl command.
the plugin is a tool to check definition for the  specified filed from k8s resource.
For Example:
	kubectl check -f nodeSelector`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		namespace, _ := cmd.Flags().GetString("namespace")
		if namespace == "" {
			namespace = "default"
		}

		f, _ := cmd.Flags().GetString("field")
		l, _ := cmd.Flags().GetString("selector")

		if !slices.Contains(base.ValidFields, f) {
			fmt.Printf("invalid filed, for now supported filed is: %s, current is %s\n", strings.Join(base.ValidFields, ", "), f)
			return
		}

		client, err := client.NewClientSet(KubernetesConfigFlags)
		if err != nil {
			fmt.Printf("error to set client: %v", err)
			return
		}

		plugin := check.NewPlugin(f, namespace, l, client)
		result, e := plugin.Value()
		if e != nil {
			fmt.Printf("error to get result: %v", e)
			return
		}
		table.GenTable(result)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubectl-check.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	KubernetesConfigFlags = genericclioptions.NewConfigFlags(true)
	rootCmd.Flags().StringVarP(&field, "field", "f", "image", "The field to check")
	rootCmd.Flags().StringVarP(&selector, "selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	KubernetesConfigFlags.AddFlags(rootCmd.Flags())
}
