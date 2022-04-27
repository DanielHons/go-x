package gitx

import "os/exec"

func ReadHeadCommitMessage() ([]byte, error) {
	strings := []string{"log", "--pretty=%s", "-n1"}

	return exec.Command("git", strings...).Output()

}
