package gh

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func GetLatestRelease() (string, error) {
	return gh_exec_clean("api", "/repos/{owner}/{repo}/releases/latest", "--jq", ".tag_name")
}

func CreateRelease(tag string) (string, error) {
	str, _, err := gh_exec("release", "create", "--generate-notes", tag)
	return str.String(), err

}

func resolveRepository() (string, error) {
	viewOut, err := gh_exec_clean("repo", "view")
	if err != nil {
		return "", err
	}
	repo := strings.TrimSpace(strings.Split(viewOut, ":")[1])
	return repo, nil

}

func gh_exec_clean(args ...string) (string, error) {
	sout, eout, err := gh_exec(args...)
	if err != nil {
		if strings.Contains(eout.String(), "not a git repository") {
			return "", errors.New("Try running this command from inside a git repository or with the -R flag")
		}
		return "", err
	}
	viewOut := strings.Split(sout.String(), "\n")[0]

	return viewOut, nil
}

// From https://github.com/vilmibm/gh-contribute

// gh shells out to gh, returning STDOUT/STDERR and any error
func gh_exec(args ...string) (sout, eout bytes.Buffer, err error) {
	ghBin, err := exec.LookPath("gh")
	if err != nil {
		err = fmt.Errorf("could not find gh. Is it installed? error: %w", err)
		return
	}

	cmd := exec.Command(ghBin, args...)
	cmd.Stderr = &eout
	cmd.Stdout = &sout

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("failed to run gh. error: %w, stderr: %s", err, eout.String())
		return
	}

	return
}
