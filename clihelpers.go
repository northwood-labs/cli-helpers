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
	"runtime/debug"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/lithammer/dedent"
)

// -----------------------------------------------------------------------------
// Cobra helpers

// LongHelpText returns a styled string with a rounded border and padding. It
// also allows you to ignore indentation. This is the standardized style for
// Northwood Labs CLI help text.
func LongHelpText(text string) string {
	helpText := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("99")).
		Padding(1, 2) // lint:allow_raw_number

	renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
	if err != nil {
		return helpText.Render(
			strings.TrimSpace(
				dedent.Dedent(text),
			),
		)
	}

	glowed, err := renderer.Render(
		strings.TrimSpace(
			dedent.Dedent(text),
		),
	)
	if err != nil {
		return helpText.Render(
			strings.TrimSpace(
				dedent.Dedent(text),
			),
		)
	}

	return helpText.Render(glowed)
}

// -----------------------------------------------------------------------------
// Bubbletea helpers

// DefaultTableStyles sets the default styles for a Bubbletea table. Set the
// values to the table with `t.SetStyles(s)`.
func DefaultTableStyles() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	return s
}

// BaseOuterTableStyle returns a lipgloss style for the outer border of a table.
func BaseOuterTableStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240"))
}

// -----------------------------------------------------------------------------
// Build/version information

// VCS reads data from the Go build info and returns the value of a specific
// key. Useful for pre-setting things like the VCS URL, build date, etc.
func VCS(key, fallback string) string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for i := range info.Settings {
			setting := info.Settings[i]

			if setting.Key == key {
				return setting.Value
			}
		}
	}

	return fallback
}
