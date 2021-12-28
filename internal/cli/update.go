package cli

import (
	"context"
	"fmt"
	"fofax/internal/printer"
	"github.com/google/go-github/github"
	"github.com/hashicorp/go-version"
	"strings"
)

func UpdateTips(tagName string) error {
	if !strings.HasPrefix(tagName, "v") {
		return fmt.Errorf("%s", "this file does not have a version number")
	}

	if strings.HasSuffix(tagName, "-next") {
		printer.Debug("Self-compiled versions do not check for updates")
		return nil
	}

	bVersion, err := version.NewVersion(tagName)
	if err != nil {
		return err
	}

	gtag, ginfo, err := getGithubVersionInfo("xiecat", "fofax")
	if err != nil {
		return err
	}

	gVersion, err := version.NewVersion(gtag)
	if err != nil {
		return err
	}

	if gVersion.GreaterThan(bVersion) {
		bannerSite(fmt.Sprintf("New:\n\nVersion:%s\n\n%s\n", gtag, ginfo))
		bannerSite("Please go to https://github.com/xiecat/fofax/releases to download\n\n")
	}
	return nil
}

func getGithubVersionInfo(owner, repo string) (tag, info string, err error) {
	client := github.NewClient(nil)
	opt := &github.ListOptions{Page: 1, PerPage: 1}
	ctx := context.Background()

	releases, rsp, err := client.Repositories.ListReleases(ctx, owner, repo, opt)

	if err != nil {
		fmt.Println(err)
	}
	if rsp.StatusCode != 200 {
		printer.Infof("update status code err :%d", rsp.StatusCode)
		return "", "", fmt.Errorf("update status code err :%d", rsp.StatusCode)
	}
	if len(releases) != 1 {
		printer.Infof("releases err")
		return "", "", fmt.Errorf("releases err")
	}
	return *releases[0].TagName, *releases[0].Body, nil
}
