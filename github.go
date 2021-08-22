package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/go-git/go-git/v5"
)

var lastCommitDate time.Time = time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

func updateAndApplyFromGithub() {
	var projectResponse ProjectResponse
	projectUrl := "https://api.github.com/repos/natery2000/local-kube"

	getResponse(projectUrl, &projectResponse)
	if projectResponse.DefaultBranch == "" {
		fmt.Println("Failed to get project default branch")
		return
	}

	var branchResponse BranchResponse
	branchUrl := projectUrl + "/branches/" + projectResponse.DefaultBranch

	getResponse(branchUrl, &branchResponse)
	if branchResponse.Commit.Commit.Committer.Date == "" {
		fmt.Println("Failed to get last commit date")
		return
	}

	date, _ := time.Parse(time.RFC3339, branchResponse.Commit.Commit.Committer.Date)
	fmt.Println(date)

	if date.After(lastCommitDate) {
		repo, _ := git.PlainOpen(projectUrl)

		repo, err := git.PlainClone("c:/Temp/local-kube", false, &git.CloneOptions{URL: projectUrl, Progress: os.Stdout})
		if err != nil {
			fmt.Println(err)
			return
		}

		work, _ := repo.Worktree()

		err = work.Pull(&git.PullOptions{RemoteName: "origin"})

		cmd := exec.Command("kubectl")
		cmd.Args = []string{"apply -f . --recursive"}
		cmd.Dir = "c:/Temp/local-kube"
		out, ee := cmd.Output()

		fmt.Println(string(out))
		fmt.Println(ee)

		lastCommitDate = date
	} else {
		fmt.Println("Environment up to date")
	}
}
