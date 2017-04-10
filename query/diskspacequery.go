package query

import (
	"log"
	"os/exec"
	"strings"
)

type DiskSpace struct {
}

// df -H | grep -vE '^Filesystem|tmpfs|cdrom|map|devfs' | awk '{ print $5 " " $1 }'
func (d DiskSpace) Get() QueryResult {
	out, err := exec.Command("df", "-H").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("output: %s", out)
	stdOut := string(out)
	return QueryResult{strings.Count(stdOut, "%") > 0, stdOut}
}
