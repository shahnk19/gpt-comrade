package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	// "strings"

	"github.com/shahnk19/gpt-comrade/pkg"
	"github.com/spf13/cobra"
)

var shellName = "bash"
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Find a fix for the previous command line error.",
	Long: `If you run into an error on your command line, use this command to 
				find a fix for the error from ChatGPT.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			fmt.Println(`ChatGPT API Key need to be set, get from
			https://platform.openai.com/account/api-keys.\nSet by 
			running 'export OPENAI_API_KEY=<your key>'`)
			return
		}

		lastCommand, _ := cmd.Flags().GetString("last-command")

		fmt.Printf("gpt-comrade finding a fix for: %v\r\n", lastCommand)

		if lastCommand == "" {
			fmt.Println("All is well!")
			return
		}

		_, stdErr, err := Shellout(lastCommand)
		if err == nil {
			fmt.Println("All is well!")
			return
		}

		cheeky, _ := cmd.Flags().GetBool("cheeky")

		pkg.FetchFix(lastCommand, err, stdErr, cheeky)
	},
}

func init() {
	FindShell()
	rootCmd.AddCommand(fixCmd)

	fixCmd.Flags().BoolP("cheeky", "c", false, "Add some humor to the fix.")
	fixCmd.Flags().StringP("last-command", "k", "", "The last command to fix.")
	rootCmd.MarkFlagRequired("last-command")
}

func FindShell() {
	shellName, _, err := Shellout("echo $SHELL")
	if err != nil {
		fmt.Printf("FindShell error: %v\n", err)
	}
	fmt.Printf("Current Shell: %v\n", shellName)
}

func Shellout(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(shellName, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
