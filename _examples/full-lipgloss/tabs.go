// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
)

type tabs struct {
	id     string
	width  int
	active string
	items  []string
}

func (m *tabs) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *tabs) GetHeight() int { _ = "STUB: not implemented"; return 0 }

func (m *tabs) Update(msg tea.Msg) tea.Cmd {
	_ = "STUB: not implemented" //nolint:unparam
	return *new(tea.Cmd)
}

// Check each item to see if it's in bounds.

func (m *tabs) View() string { _ = "STUB: not implemented"; return "" }

// Make sure to mark each tab when rendering.
