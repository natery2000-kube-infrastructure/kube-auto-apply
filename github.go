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
	localRepositoryDirectory := "/repo"
	repositoryUrl := "https://github.com/natery2000-kube-infrastructure/local-kube"

	fmt.Println("Cloning repository")
	repo, err := git.PlainClone(localRepositoryDirectory, false, &git.CloneOptions{URL: repositoryUrl, Progress: os.Stdout})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opening repository")
	repo, err = git.PlainOpen(localRepositoryDirectory)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Getting Worktree")
	work, err := repo.Worktree()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Pulling latest")
	err = work.Pull(&git.PullOptions{RemoteName: "origin"})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Applying new yaml files")
	cmd := exec.Command("kubectl", "apply", "-f", ".", "--recursive")
	cmd.Dir = localRepositoryDirectory
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
