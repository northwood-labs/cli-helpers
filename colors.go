// Copyright 2024-2026, Northwood Labs
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
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/log/v2"
)

var (
	// Understand background colors
	ClrHasDarkBG = lipgloss.HasDarkBackground(os.Stdin, os.Stdout)
	ClrLightDark = lipgloss.LightDark(ClrHasDarkBG)

	// Base colors
	ClrBlack  = lipgloss.Color("#171e21")
	ClrBlue   = lipgloss.Color("#0087ff")
	ClrGreen  = lipgloss.Color("#009900")
	ClrOrange = lipgloss.Color("#ffa500")
	ClrPurple = lipgloss.Color("#8700ff")
	ClrRed    = lipgloss.Color("#cc0000")
	ClrWhite  = lipgloss.Color("#ffffff")
	ClrYellow = lipgloss.Color("#ffff00")

	// Meanings (logging)
	ClrFatal = ClrRed
	ClrError = ClrRed
	ClrWarn  = ClrOrange
	ClrInfo  = ClrBlue
	ClrDebug = ClrPurple

	// Meanings (success)
	ClrFailure = ClrRed
	ClrSuccess = ClrGreen

	// Inline styles change based on background
	StyleInlineHighlight = lipgloss.NewStyle().
				Foreground(ClrLightDark(ClrBlue, ClrYellow))

	// Base styles for backgrounds
	StyleFailure = lipgloss.NewStyle().
			Foreground(ClrWhite).
			Background(ClrFailure)
	StyleSuccess = lipgloss.NewStyle().
			Foreground(ClrWhite).
			Background(ClrSuccess)

	StyleDebug = lipgloss.NewStyle().
			Foreground(ClrWhite).
			Background(ClrDebug)
	StyleInfo = lipgloss.NewStyle().
			Foreground(ClrWhite).
			Background(ClrInfo)
	StyleWarn = lipgloss.NewStyle().
			Foreground(ClrBlack).
			Background(ClrWarn)
	StyleError = lipgloss.NewStyle().
			Foreground(ClrWhite).
			Background(ClrError)
)

func GetLoggerStyles() *log.Styles {
	styles := log.DefaultStyles()

	// DEBUG
	styles.Levels[log.DebugLevel] = StyleDebug.
		SetString(strings.ToUpper(log.DebugLevel.String()[0:4])).
		Padding(0, 1, 0, 1).
		Bold(true)

	// INFO
	styles.Levels[log.InfoLevel] = StyleInfo.
		SetString(strings.ToUpper(log.InfoLevel.String()[0:4])).
		Padding(0, 1, 0, 1).
		Bold(true)

	// WARNING
	styles.Levels[log.WarnLevel] = StyleWarn.
		SetString(strings.ToUpper(log.WarnLevel.String()[0:4])).
		Padding(0, 1, 0, 1).
		Bold(true)

	// ERROR
	styles.Levels[log.ErrorLevel] = StyleError.
		SetString(strings.ToUpper(log.ErrorLevel.String()[0:4])).
		Padding(0, 1, 0, 1).
		Bold(true)

	// FATAL
	styles.Levels[log.FatalLevel] = StyleError.
		SetString(strings.ToUpper(log.FatalLevel.String()[0:4])).
		Padding(0, 1, 0, 1).
		Bold(true)

	// Add a custom style for key `err`
	styles.Keys["err"] = lipgloss.NewStyle().Foreground(ClrError)

	return styles
}
