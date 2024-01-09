package main

import (
    "context"
    "errors"
    "github.com/google/go-github/v57/github"
    "regexp"
    "strings"
)

func getVersionFromGithub(path string) (string, error) {
    a := strings.Split(path, "/")
    if !(len(a) == 2) {
        return "", errors.New("invalid path")
    }

    client := github.NewClient(nil)
    repo, _, err := client.Repositories.GetLatestRelease(context.TODO(), a[0], a[1])

    if err != nil {
        return "", err
    }

    r := regexp.MustCompile(`^(v?)(.*)`)
    s := r.ReplaceAllString(*repo.TagName, "$2")

    return s, err
}
