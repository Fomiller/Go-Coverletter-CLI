package cmd

import (
	"fmt"

	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	// KeyEnter is the default key for submission/selection.
	KeyEnter rune = readline.CharEnter

	// KeyBackspace is the default key for deleting input text.
	KeyBackspace rune = readline.CharBackspace

	// KeyPrev is the default key to go up during selection.
	KeyPrev        rune = readline.CharPrev
	KeyPrevDisplay      = "↑"
	up                  = promptui.Key{Code: KeyPrev}

	// KeyNext is the default key to go down during selection.
	KeyNext        rune = readline.CharNext
	KeyNextDisplay      = "↓"
	down                = promptui.Key{Code: KeyNext}

	// KeyBackward is the default key to page up during selection.
	KeyBackward        rune = readline.CharBackward
	KeyBackwardDisplay      = "←"
	left                    = promptui.Key{Code: KeyBackward}

	// KeyForward is the default key to page down during selection.
	KeyForward        rune = readline.CharForward
	KeyForwardDisplay      = "→"
	right                  = promptui.Key{Code: KeyForward}
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmdList := []string{}
		rcmds := rootCmd.Commands()
		for _, v := range rcmds {
			cmdList = append(cmdList, v.Use)
		}
		prompt := promptui.Select{
			Keys:  &promptui.SelectKeys{Next: down, Prev: up, PageUp: left, PageDown: right},
			Label: "Select a command",
			Items: cmdList,
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose %q\n", result)
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// promptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// promptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
