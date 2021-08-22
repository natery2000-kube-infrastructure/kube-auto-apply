package main

type ProjectResponse struct {
	DefaultBranch string `json:"default_branch"`
}

type BranchResponse struct {
	Commit Commit `json:"commit"`
}

type Commit struct {
	Commit CommitCommit `json:"commit"`
}

type CommitCommit struct {
	Committer Committer `json:"committer"`
}

type Committer struct {
	Date string `json:"date"`
}
