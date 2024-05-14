// Copyright 2024, Northwood Labs
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

package clihelpers

import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/northwood-labs/archstring"
	"github.com/spf13/cobra"
)

func VersionScreen() *cobra.Command {
	version := "dev"
	commit := VCS("vcs.revision", "unknown")
	buildDate := VCS("vcs.time", "unknown")
	dirty := VCS("vcs.modified", "unknown")
	pgoEnabled := VCS("-pgo", "false")

	return &cobra.Command{
		Use:   "version",
		Short: "Long-form version information",
		Long: LongHelpText(`
		Long-form version information, including the build commit hash, build date, Go
		version, and external dependencies.`),
		Run: func(cmd *cobra.Command, args []string) {
			t := table.New().
				Border(lipgloss.RoundedBorder()).
				BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
				BorderColumn(true).
				StyleFunc(func(row, col int) lipgloss.Style {
					return lipgloss.NewStyle().Padding(0, 1)
				}).
				Headers("BUILD INFO", "VALUE")

			t.Row("Version", version)
			t.Row("Go version", runtime.Version())
			t.Row("Git commit", commit)
			if dirty == "true" {
				t.Row("Dirty repo", dirty)
			}
			if !strings.Contains(pgoEnabled, "false") {
				t.Row("PGO", filepath.Base(pgoEnabled))
			}
			t.Row("Build date", buildDate)
			t.Row("OS/Arch", fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))
			t.Row("System", archstring.GetFriendlyName(runtime.GOOS, runtime.GOARCH))
			t.Row("CPU cores", fmt.Sprintf("%d", runtime.NumCPU()))

			fmt.Println(t.Render())

			//----------------------------------------------------------------------

			if buildInfo, ok := debug.ReadBuildInfo(); ok {
				td := table.New().
					Border(lipgloss.RoundedBorder()).
					BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
					BorderColumn(true).
					StyleFunc(func(row, col int) lipgloss.Style {
						return lipgloss.NewStyle().Padding(0, 1)
					}).
					Headers("DEPENDENCY", "VERSION")

				for i := range buildInfo.Deps {
					dependency := buildInfo.Deps[i]
					td.Row(dependency.Path, dependency.Version)
				}

				fmt.Println(td.Render())
			}

			fmt.Println("")
		},
	}
}
