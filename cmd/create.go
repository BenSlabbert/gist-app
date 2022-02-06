package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var useEditor *bool

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create gist",
	Long:  `Create a gist from a file, all files in a directory, or from vim/nano`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if *useEditor {
			return createWithEditor()
		}

		fmt.Println("can only create with the --editor/-e flag")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	useEditor = createCmd.Flags().BoolP("editor", "e", false, "Launches vim in the terminal so you can create your gist in a temp file.")
}

func createWithEditor() error {
	file, err := ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		return err
	}
	// delete the file when we are done
	defer os.Remove(file.Name())

	// create a tmp file instead of test.txt
	// read the file and then delete?
	cmd := exec.Command("vi", file.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err = cmd.Run()
	if err != nil {
		return err
	}

	// go to the start of the file
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	gist, err := api.PutGist("from-vim-example.txt", string(all), "example description", true)
	if err != nil {
		return err
	}

	fmt.Printf("created gist with id: %s\n", gist.Id)
	return nil
}
