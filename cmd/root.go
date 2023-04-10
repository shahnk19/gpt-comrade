package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gpt-comrade",
		Short: "CLI tool to fix command line errors with ChatGPT.",
		Long: `CLI tool to fix command line errors with ChatGPT. When you hit an error on
	your command line, get immediate help and fix from ChatGPT. For example:

		> ERROR: Keyboard not present, press any key.
		>
		> gpt-comrade fix --cheeky
		>
		> Hello comrade. Well, it seems like your computer has developed a sudden case of 
		keyboardophobia and is trying to distance itself from the keyboard! Maybe try gently
		reassuring your keyboard that it's still loved and needed, and see if that helps
		resolve the issue? Alternatively, you could try communicating with your computer
		via interpretive dance instead of using the keyboard. Who knows, it might just work!,
		>
		> gpt-comrade fix
		>
		> Hello comrade. It's possible that there is an issue with your keyboard or your 
		system's firmware. Restart your computer: As with many technical issues, a simple
		restart may resolve the issue.`,

		Run: func(cmd *cobra.Command, args []string) {
			apiKey, _ := cmd.Flags().GetString("api-key")

			os.Setenv("OPENAI_API_KEY", apiKey)
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("api-key", "k", os.Getenv("OPENAI_API_KEY"), "")
	rootCmd.MarkFlagRequired("api-key")
}
