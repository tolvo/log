/*
Copyright © 2026 tolvo <tolvo.pedro@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// inCmd represents the in command
var inCmd = &cobra.Command{
	Use:   "in",
	Short: "It defines in which book the entry will be made.",
	Long: 
`The 'in' command allows you to specify the book where the log entry will be recorded.
You can choose from various predefined books such as 'personal', 'work', 'ideas', etc.
For example, to log an entry in the 'work' book, you would use:
	
	log in work "Completed the project documentation."
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please specify a book and an text to log in, like example above:")
			fmt.Println(`log in remember "don't be an asshole"`)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(inCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
