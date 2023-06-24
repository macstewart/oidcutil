package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	setCmd = &cobra.Command{Use: "set", Run: runSet}
	email  *string
	key    *string
	values *[]string
)

func init() {
	email = setCmd.Flags().StringP("email", "e", "", "The email claim value in the returned token")
	key = setCmd.Flags().StringP("key", "k", "", "The key for the custom claim in the returned token")
	values = setCmd.Flags().StringSliceP("value", "v", []string{}, "The value for the custom claim in the returned token")
	rootCmd.AddCommand(setCmd)
}

func runSet(cmd *cobra.Command, args []string) {
	fmt.Println(*email, *key, *values)
}
