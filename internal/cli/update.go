package cli

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	fofaxUpdateStore "github.com/tj/go-update/stores/github"

	"fofax/internal/printer"
	"github.com/pkg/errors"
	"github.com/tj/go-update"
	"github.com/tj/go-update/progress"
)

func updateTips(tagName string, isDown bool) error {
	tagName = fmt.Sprintf("v%s", tagName)
	if strings.HasSuffix(tagName, "-next") {
		printer.Debug("Self-compiled versions do not check for updates")
		return nil
	}
	latest, err := updateFoFaXVersionToLatest(isDown)
	if err != nil {
		return err
	}
	if !isDown {
		bannerSite(fmt.Sprintf("New:\n\nVersion:%s\n\n%s\n", latest.Version, latest.Notes))
		bannerSite("Please Use [./fofax -update] to download\n\n")
	}
	return nil
}

// updateFoFaXVersionToLatest from nuclei.
func updateFoFaXVersionToLatest(isDown bool) (*update.Release, error) {
	var command string
	switch runtime.GOOS {
	case "windows":
		command = "fofax.exe"
	default:
		command = "fofax"
	}
	m := &update.Manager{
		Command: command,
		Store: &fofaxUpdateStore.Store{
			Owner:   "xiecat",
			Repo:    "fofax",
			Version: FoFaXVersion,
		},
	}
	releases, err := m.LatestReleases()
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch latest release")
	}
	if len(releases) == 0 {
		return nil, errors.Wrap(err, "No new updates found for fofax engine!")
	}
	latest := releases[0]
	if isDown {
		currentOS := runtime.GOOS
		var final *update.Asset
		switch runtime.GOOS {
		case "windows":
			final = latest.FindZip(currentOS, runtime.GOARCH)
		default:
			final = latest.FindTarball(currentOS, runtime.GOARCH)
		}
		if final == nil {
			return nil, fmt.Errorf("no compatible binary found for %s/%s", currentOS, runtime.GOARCH)
		}
		printer.Info("Download will start soon...")
		tarball, err := final.DownloadProxy(progress.Reader)
		if err != nil {
			return nil, errors.Wrap(err, "could not download latest release")
		}
		currentPath, _ := os.Getwd()
		if err := m.InstallTo(tarball, currentPath); err != nil {
			return nil, errors.Wrap(err, "could not install latest release")
		}
		printer.Successf("Successfully updated to fofax %s", latest.Version)
		printer.Success("Use [./fofax -v] to view")
	}
	return latest, nil
}
