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

	fmt.Println("Getting default branch name.")
	getResponse(projectUrl, &projectResponse)
	if projectResponse.DefaultBranch == "" {
		fmt.Println("Failed to get project default branch")
		//return
	}

	fmt.Println("Getting last commit to default branch date")
	var branchResponse BranchResponse
	branchUrl := projectUrl + "/branches/" + projectResponse.DefaultBranch

	getResponse(branchUrl, &branchResponse)
	if branchResponse.Commit.Commit.Committer.Date == "" {
		fmt.Println("Failed to get last commit date")
		//return
	}

	date, _ := time.Parse(time.RFC3339, branchResponse.Commit.Commit.Committer.Date)
	fmt.Println(date)

	if date.After(lastCommitDate) {
		localRepositoryDirectory := "/repo"
		repositoryUrl := "https://github.com/natery2000/local-kube"

		fmt.Println("Cloning repository")
		repo, err := git.PlainClone(localRepositoryDirectory, false, &git.CloneOptions{URL: repositoryUrl, Progress: os.Stdout})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Opening repository")
		repo, err = git.PlainOpen(localRepositoryDirectory)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Getting Worktree")
		work, err := repo.Worktree()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Pulling latest")
		err = work.Pull(&git.PullOptions{RemoteName: "origin"})
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Applying new yaml files")
		cmd := exec.Command("kubectl", "apply", "-f", ".", "--recursive")
		cmd.Dir = localRepositoryDirectory
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

		lastCommitDate = date
	} else {
		fmt.Println("Environment up to date")
	}
}
