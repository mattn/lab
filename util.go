package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func SearchBrowserLauncher(goos string) (browser string) {
	switch goos {
	case "darwin":
		browser = "open"
	case "windows":
		browser = "cmd /c start"
	default:
		candidates := []string{
			"xdg-open",
			"cygstart",
			"x-www-browser",
			"firefox",
			"opera",
			"mozilla",
			"netscape",
		}
		for _, b := range candidates {
			path, err := exec.LookPath(b)
			if err == nil {
				browser = path
				break
			}
		}
	}
	return browser
}

type BrowseType int

const (
	Issue BrowseType = iota
	MergeRequest
)

var BrowseTypePrefix = map[string]BrowseType{
	"#": Issue,
	"i": Issue,
	"I": Issue,
	"!": MergeRequest,
	"m": MergeRequest,
	"M": MergeRequest,
}

func splitPrefixAndNumber(arg string) (BrowseType, int, error) {
	for k, v := range BrowseTypePrefix {
		if strings.HasPrefix(arg, k) {
			numberStr := strings.TrimPrefix(arg, k)
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				return 0, 0, errors.New(fmt.Sprintf("Invalid browsing number: %s", arg))
			}
			return v, number, nil
		}
	}
	return 0, 0, errors.New(fmt.Sprintf("Invalid arg: %s", arg))
}
