// Copyright 2022 Red Hat, Inc.
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

// TODO:
// - Support testing the "just run Live" case - maybe try to figure out
//   how to have main `kola` tests apply?
// - Test `coreos-install iso embed` path

package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/coreos/mantle/platform/conf"
	"github.com/coreos/mantle/system"
	"github.com/coreos/mantle/util"
	"github.com/pkg/errors"

	"github.com/spf13/cobra"

	"github.com/coreos/mantle/kola"
	"github.com/coreos/mantle/platform"
)

var (
	cmdTestContainer = &cobra.Command{
		RunE:    runTestContainer,
		PreRunE: preRun,
		Use:     "testcontainer",
		Short:   "Test a CoreOS container image",

		SilenceUsage: true,
	}
)

func init() {
	cmdTestContainer.Args = cobra.ExactArgs(0)

	root.AddCommand(cmdTestContainer)
}

func runTestContainer(cmd *cobra.Command, args []string) error {
	if kola.CosaBuild == nil {
		return fmt.Errorf("Must provide --build")
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	// note this reassigns a *global*
	outputDir, err = kola.SetupOutputDir(outputDir, "testiso")
	if err != nil {
		return err
	}
	
	builder := platform.NewQemuBuilder()
	defer builder.Close()

	container := kola.CosaBuild.Meta.BaseOsContainer 
	if container == nil {
		return fmt.Errorf("build %s has no baseoscontainer", kola.CosaBuild.Meta.Name)
	}

	return nil
}
