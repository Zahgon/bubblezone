// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	tea "charm.land/bubbletea/v2"
)

type history struct {
	id     string
	height int
	width  int

	active string
	items  []string
}

func (m *history) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m *history) Update(msg tea.Msg) tea.Cmd {
	_ = "STUB: not implemented" //nolint:unparam
	return *new(tea.Cmd)
}

// Check each item to see if it's in bounds.

func (m *history) View() string { _ = "STUB: not implemented"; return "" }

// Customize the active item.

// Make sure to mark all zones.
