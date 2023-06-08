package gh

import (
	"errors"
	"strings"

	"github.com/cli/go-gh"
)

func GetLatestRelease() (string, error) {
	return gh_exec_clean("api", "/repos/{owner}/{repo}/releases/latest", "--jq", ".tag_name")
}

func CreateRelease(tag, title string) (string, error) {
	str, _, err := gh.Exec("release", "create", "--generate-notes", "--title", title, tag)
	return str.String(), err

}

func gh_exec_clean(args ...string) (string, error) {
	sout, eout, err := gh.Exec(args...)
	if err != nil {
		if strings.Contains(eout.String(), "not a git repository") {
			return "", errors.New("try running this command from inside a git repository or with the -R flag")
		}
		return "", err
	}
	viewOut := strings.Split(sout.String(), "\n")[0]

	return viewOut, nil
}
