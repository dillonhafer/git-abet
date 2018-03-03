package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

var NumberOfFiles int

var rootCmd = &cobra.Command{
	Use:   "git-abet",
	Short: "git-abet shows related commited files",
	Long: `Show files abetting in a commit
                Complete documentation is available at https://github.com/dillonhafer/git-abet`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		commits, err := findCommits(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "File does not exist or is outside the current git directory")
			os.Exit(1)
		}

		var files []string
		for _, c := range commits {
			foundFiles, err := findFiles(c)
			if err != nil {
				println("Could not find files for commit:", c)
				os.Exit(1)
			}

			for _, f := range foundFiles {
				if _, err := os.Stat(f); err == nil {
					if f != "" && f != file {
						files = append(files, f)
					}
				}
			}
		}

		groups := groupCount(files)
		type ranking struct {
			File  string
			Count int
		}

		var rankings []ranking
		for fileName, fileCount := range groups {
			rankings = append(rankings, ranking{fileName, fileCount})
		}
		sort.Slice(rankings, func(i, j int) bool {
			return rankings[i].Count > rankings[j].Count
		})

		for i := 0; i < NumberOfFiles; i++ {
			if len(rankings) > i {
				println(rankings[i].File)
			} else {
				return
			}
		}
	},
}

func groupCount(list []string) map[string]int {
	frequency := make(map[string]int)
	for _, item := range list {
		_, exist := frequency[item]
		if exist {
			frequency[item] += 1
		} else {
			frequency[item] = 1
		}
	}

	return frequency
}

func Execute() {
	rootCmd.PersistentFlags().IntVarP(&NumberOfFiles, "number", "n", 5, "Number of files to return")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func gitCommand(cmdArgs []string) ([]string, error) {
	cmdOut, err := exec.Command("git", cmdArgs...).CombinedOutput()
	if err != nil {
		println(string(cmdOut))
		return []string{}, err
	}

	lines := strings.Split(string(cmdOut), "\n")
	return lines, nil
}

func findFiles(commit string) ([]string, error) {
	cmdArgs := []string{"show", "-r", "--no-commit-id", "--pretty=", "--name-only", commit}
	return gitCommand(cmdArgs)
}

func findCommits(absFilePath string) ([]string, error) {
	cmdArgs := []string{"log", `--pretty=format:%H`, absFilePath}
	return gitCommand(cmdArgs)
}
