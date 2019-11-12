package mdtojson

import (
	"encoding/json"
	"fmt"

	"../../model"
	"../parse"
)

// ProcessRepo clones and parses a directory of markdown files with YAML
// frontmatter and returns a JSON string
func ProcessRepo(repoURL string, tmpDIR string) (string, error) {
	if repoURL == "" || tmpDIR == "" {
		return "", fmt.Errorf("repoURL and tmpDIR are required")
	}
	/* _, err := download.RepoToDisk(repoURL, tmpDIR)
	if err != nil {
		return "", err
	} */

	md, err := parse.Files(tmpDIR)
	if err != nil {
		return "", err
	}

	var vrecords model.EventosRecords
	vrecords.Result = "OK"
	vrecords.TotalRecordCount = len(md) - 1
	vrecords.Records = md
	// create json response from struct
	j, err := json.Marshal(vrecords)

	if err != nil {
		return "", err
	}

	return string(j), nil
}
