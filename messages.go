// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package zone

import (
	tea "charm.land/bubbletea/v2"
)

// MsgZoneInBounds is a message sent when the manager detects that a zone is within
// bounds of a mouse event.
type MsgZoneInBounds struct {
	Zone *ZoneInfo // The zone that is in bounds.

	Event tea.MouseMsg // The mouse event that caused the zone to be in bounds.
}

func (m *Manager) findInBounds(mouse tea.MouseMsg) []*ZoneInfo {
	_ = "STUB: not implemented"
	return nil
}

// AnyInBoundsAndUpdate is the same as AnyInBounds; except the results of the calls
// to Update() are carried through and returned.
//
// The tea.Cmd's that comd off the calls to Update() are wrapped in tea.Batch().
func (m *Manager) AnyInBoundsAndUpdate(model tea.Model, mouse tea.MouseMsg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// AnyInBounds sends a MsgZoneInBounds message to the provided model for each zone
// that is in the bounds of the provided mouse event. The results of the call to
// Update() are discarded.
//
// Note that if multiple zones are within bounds, each one will be sent as an event
// in alphabetical sorted order of the ID.
func (m *Manager) AnyInBounds(model tea.Model, mouse tea.MouseMsg) {
	_ = "STUB: not implemented"
	return
}
