/*
    Copyright © 2023, Vitalii Tereshchuk | DOTOCA.NET All rights reserved.
    Homepage: https://dotoca.net/shuffle-files

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/

package cmd

import (
	"os"
    "fmt"
    "io/ioutil"
    "math/rand"
    "path/filepath"
    "time"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var outputPath string
var debug = false
var test = false


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shuffle-files <path_to_files>",
    Version: version,
    Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "This is a CLI tool which shuffles the files in the directory",
	Long: `Copyright © 2023, Vitalii Tereshchuk | DOTOCA.NET All rights reserved | https://dotoca.net/shuffle-files

This is a CLI tool which shuffles the files in the directory, their content, but without changing the file names.

Example: shuffle-files ./some_path_to_files
`,
    Run: mainApp,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.shuffle-files.yaml)")

    rootCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Output directory for files")

    rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "set log level to debug")
    rootCmd.PersistentFlags().BoolVar(&test, "test", false, "nothing will be done")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}




func mainApp(cmd *cobra.Command, args []string) {
    // Check for the correct number of command-line arguments
    if len(args) < 1 {
        cmd.Help()
        os.Exit(1)
    }

    dir     := args[0]   // input path
    dir_out := ""        // output path

    // Get the directory path from the command-line argument if there's
    if outputPath == "" {
        dir_out = dir
    } else {
        dir_out = outputPath
    }


    // List all files in the specified directory
    fileList, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Seed the random number generator
    rand.Seed(time.Now().Unix())

    // Shuffle the file names
    shuffledFileNames := shuffleFileNames(fileList)

    // Rename the files with shuffled names
    for i, file := range fileList {
        oldPath := filepath.Join(dir, file.Name())
        newName := shuffledFileNames[i]
        newPath := filepath.Join(dir, newName)

        if debug {
            fmt.Printf("%s -> %s \n", oldPath, newPath + ".tmp")
        }

        if !test {
           if err := os.Rename(oldPath, newPath + ".tmp"); err != nil {
                fmt.Printf("Error renaming %s to %s: %v\n", oldPath, newPath, err)
            }
        }
    }

    for i := range shuffledFileNames {
        oldPath := filepath.Join(dir, shuffledFileNames[i])
        newName := shuffledFileNames[i]
        //newPath := filepath.Join(dir, newName)

        newPath := filepath.Join(dir_out, newName)

        if debug {
            fmt.Printf("%s -> %s \n", oldPath + ".tmp", newPath)
        }

        if !test {
            if err := os.Rename(newPath + ".tmp", newPath); err != nil {
                fmt.Printf("Error renaming %s to %s: %v\n", newPath + ".tmp", newPath, err)
            }
        }
    }

    fmt.Printf("Shuffled %d file(s)\n", len(shuffledFileNames))
}


// Shuffle filenames array
func shuffleFileNames(fileList []os.FileInfo) []string {
    names := make([]string, len(fileList))
    for i, file := range fileList {
        names[i] = file.Name()
    }

    // Shuffle the file names using the Fisher-Yates algorithm
    for i := len(names) - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        names[i], names[j] = names[j], names[i]
    }
    return names
}

