package cli

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	fofaxUpdateStore "github.com/tj/go-update/stores/github"

	"github.com/xiecat/fofax/internal/printer"

	"github.com/pkg/errors"
	"github.com/tj/go-update"
	"github.com/tj/go-update/progress"
)

func updateTips(tagName string) error {
	tagName = fmt.Sprintf("v%s", strings.TrimPrefix(tagName, "v"))
	if strings.HasSuffix(tagName, "-next") {
		printer.Debugf("Self-compiled versions: %s do not check for updates", tagName)
		return nil
	}
	latest, err := updateFoFaXVersionToLatest()
	if latest == nil {
		return errors.New("latest version is nil")
	}

	if err != nil {
		return err
	}
	if !args.Update {
		bannerSite(fmt.Sprintf("New:\n\nVersion:%s\n\n%s\n", latest.Version, latest.Notes))
		bannerSite("Please Use [./fofax -update] to download\n\n")
	}
	return nil
}

// updateFoFaXVersionToLatest from nuclei.
func updateFoFaXVersionToLatest() (*update.Release, error) {
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
		return nil, errors.New("No new updates found for fofax engine!")
	}
	latest := releases[0]
	if args.Update {
		bannerSite(fmt.Sprintf("New:\n\nVersion:%s\n\n%s\n", latest.Version, latest.Notes))
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
