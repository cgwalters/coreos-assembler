// Copyright 2016 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/url"
	"path"
	"strings"

	"github.com/coreos/mantle/Godeps/_workspace/src/github.com/spf13/pflag"

	"github.com/coreos/mantle/lang/maps"
	"github.com/coreos/mantle/sdk"
)

type storageSpec struct {
	BaseURL       string
	NamedPath     string // Copy to $BaseURL/$Board/$NamedPath
	VersionPath   bool   // Copy to $BaseURL/$Board/$Version
	DirectoryHTML bool
	IndexHTML     bool
}

type channelSpec struct {
	BaseURL      string // Copy from $BaseURL/$Board/$Version
	Destinations []storageSpec
}

var (
	specFlags   *pflag.FlagSet
	specBoard   string
	specChannel string
	specVersion string
	specs       = map[string]channelSpec{
		"alpha": channelSpec{
			BaseURL: "gs://builds.release.core-os.net/alpha/boards",
			Destinations: []storageSpec{storageSpec{
				BaseURL:     "gs://alpha.release.core-os.net",
				NamedPath:   "current",
				VersionPath: true,
				IndexHTML:   true,
			}, storageSpec{
				BaseURL:       "gs://coreos-alpha",
				NamedPath:     "current",
				VersionPath:   true,
				DirectoryHTML: true,
				IndexHTML:     true,
			}, storageSpec{
				BaseURL:     "gs://storage.core-os.net/coreos",
				NamedPath:   "alpha",
				VersionPath: true,
				IndexHTML:   true,
			}, storageSpec{
				BaseURL:       "gs://coreos-net-storage/coreos",
				NamedPath:     "alpha",
				VersionPath:   true,
				DirectoryHTML: true,
				IndexHTML:     true,
			}},
		},
		"beta": channelSpec{
			BaseURL: "gs://builds.release.core-os.net/beta/boards",
			Destinations: []storageSpec{storageSpec{
				BaseURL:     "gs://beta.release.core-os.net",
				NamedPath:   "current",
				VersionPath: true,
				IndexHTML:   true,
			}, storageSpec{
				BaseURL:       "gs://coreos-beta",
				NamedPath:     "current",
				VersionPath:   true,
				DirectoryHTML: true,
				IndexHTML:     true,
			}, storageSpec{
				BaseURL:   "gs://storage.core-os.net/coreos",
				NamedPath: "beta",
				IndexHTML: true,
			}, storageSpec{
				BaseURL:       "gs://coreos-net-storage/coreos",
				NamedPath:     "beta",
				DirectoryHTML: true,
				IndexHTML:     true,
			}},
		},
		"stable": channelSpec{
			BaseURL: "gs://builds.release.core-os.net/stable/boards",
			Destinations: []storageSpec{storageSpec{
				BaseURL:     "gs://stable.release.core-os.net",
				NamedPath:   "current",
				VersionPath: true,
				IndexHTML:   true,
			}, storageSpec{
				BaseURL:       "gs://coreos-stable",
				NamedPath:     "current",
				VersionPath:   true,
				DirectoryHTML: true,
				IndexHTML:     true,
			}},
		},
	}
)

func AddSpecFlags(flags *pflag.FlagSet) {
	board := sdk.DefaultBoard()
	channels := strings.Join(maps.SortedKeys(specs), " ")
	versions, _ := sdk.VersionsFromManifest()
	flags.StringVarP(&specBoard, "board", "B",
		board, "target board")
	flags.StringVarP(&specChannel, "channel", "C",
		"alpha", "channels: "+channels)
	flags.StringVarP(&specVersion, "version", "V",
		versions.VersionID, "release version")
}

func ChannelSpec() channelSpec {
	if specBoard == "" {
		plog.Fatal("--board is required")
	}
	if specChannel == "" {
		plog.Fatal("--channel is required")
	}
	if specVersion == "" {
		plog.Fatal("--version is required")
	}

	spec, ok := specs[specChannel]
	if !ok {
		plog.Fatalf("Unknown channel: %s", specChannel)
	}

	return spec
}

func (cs channelSpec) SourceURL() string {
	u, err := url.Parse(cs.BaseURL)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, specBoard, specVersion)
	return u.String()
}

func (ss storageSpec) ParentURL() string {
	u, err := url.Parse(ss.BaseURL)
	if err != nil {
		panic(err)
	}
	u.Path = path.Join(u.Path, specBoard)
	return u.String()
}

func (ss storageSpec) Prefixes() []string {
	u, err := url.Parse(ss.BaseURL)
	if err != nil {
		plog.Panic(err)
	}

	prefixes := []string{}
	if ss.VersionPath {
		prefixes = append(prefixes,
			path.Join(u.Path, specBoard, specVersion))
	}
	if ss.NamedPath != "" {
		prefixes = append(prefixes,
			path.Join(u.Path, specBoard, ss.NamedPath))
	}
	if len(prefixes) == 0 {
		plog.Panicf("Invalid destination: %#v", ss)
	}

	return prefixes
}
