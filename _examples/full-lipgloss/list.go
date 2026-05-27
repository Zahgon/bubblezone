// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	listStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			MarginRight(2)

	listHeader = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			MarginRight(2)

	listItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	checkMark = lipgloss.NewStyle().
			SetString("✓").
			PaddingRight(1)

	listDoneStyle = lipgloss.NewStyle().
			Strikethrough(true)
)

type listItem struct {
	name string
	done bool
}

type list struct {
	id    string
	title string
	items []listItem
}

func (m *list) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *list) GetHeight() int { _ = "STUB: not implemented"; return 0 }

func (m *list) Update(msg tea.Msg) tea.Cmd {
	_ = "STUB: not implemented" //nolint:unparam
	return *new(tea.Cmd)
}

// Check each item to see if it's in bounds.

func (m *list) View() string { _ = "STUB: not implemented"; return "" }
