// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), true).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1).
			MarginRight(2)

	activeButtonStyle = buttonStyle.
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)
)

type dialog struct {
	id       string
	active   string
	question string
}

func (m *dialog) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *dialog) GetHeight() int { _ = "STUB: not implemented"; return 0 }

func (m *dialog) Update(msg tea.Msg) tea.Cmd {
	_ = "STUB: not implemented" //nolint:unparam
	return *new(tea.Cmd)
}

func (m *dialog) View() string { _ = "STUB: not implemented"; return "" }
