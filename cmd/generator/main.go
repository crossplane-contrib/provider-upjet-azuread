// Copyright 2022 Upbound Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/upbound/upjet/pkg/pipeline"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/upbound/provider-azuread/config"
)

func main() {
	var (
		app                 = kingpin.New("generator", "Run Upjet code generation pipelines for provider-azuread").DefaultEnvars()
		repoRoot            = app.Arg("repo-root", "Root directory for the provider repository").Required().String()
		skippedResourcesCSV = app.Flag("skipped-resources-csv", "File path where a list of skipped (not-generated) Terraform resource names will be stored as a CSV").Envar("SKIPPED_RESOURCES_CSV").String()
	)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	absRootDir, err := filepath.Abs(*repoRoot)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", *repoRoot))
	}
	p := config.GetProvider()
	pipeline.Run(p, absRootDir)
	if len(*skippedResourcesCSV) != 0 {
		skippedCount := len(p.GetSkippedResourceNames())
		totalCount := skippedCount + len(p.Resources)
		summaryLine := fmt.Sprintf("Skipped, total, coverage: %d, %d, %.1f%%", skippedCount, totalCount, (1.0-float64(skippedCount)/float64(totalCount))*100)
		if err := os.WriteFile(*skippedResourcesCSV, []byte(strings.Join(append([]string{summaryLine}, p.GetSkippedResourceNames()...), "\n")), 0o600); err != nil {
			panic(fmt.Sprintf("cannot write skipped resources CSV to file %s: %s", *skippedResourcesCSV, err.Error()))
		}
	}
}
