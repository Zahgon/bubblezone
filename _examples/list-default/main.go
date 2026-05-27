// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	zone "github.com/lrstanley/bubblezone/v2"
)

// This is a modified version of this example, to support mouse click zones and
// scrolling events:
// 	https://github.com/charmbracelet/bubbletea/tree/master/examples/list-default

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	id    string
	title string
	desc  string
}

func (i item) Title() string       { _ = "STUB: not implemented"; return "" }
func (i item) Description() string { _ = "STUB: not implemented"; return "" }
func (i item) FilterValue() string { _ = "STUB: not implemented"; return "" }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Check each item to see if it's in bounds.

// If so, select it in the list.

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }

// Wrap the main models view in [zone.Scan].

func main() {
	// Initialize a global zone manager, so we don't have to pass around the manager
	// throughout components.
	zone.NewGlobal()

	items := []list.Item{
		// an ID field has been added here, however it's not required. You could use
		// any text field as long as it's unique for the zone.
		item{id: "item_1", title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{id: "item_2", title: "Nutella", desc: "It's good on toast"},
		item{id: "item_3", title: "Bitter melon", desc: "It cools you down"},
		item{id: "item_4", title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{id: "item_5", title: "Eight hours of sleep", desc: "I had this once"},
		item{id: "item_6", title: "Cats", desc: "Usually"},
		item{id: "item_7", title: "Plantasia, the album", desc: "My plants love it too"},
		item{id: "item_8", title: "Pour over coffee", desc: "It takes forever to make though"},
		item{id: "item_9", title: "VR", desc: "Virtual reality...what is there to say?"},
		item{id: "item_10", title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		item{id: "item_11", title: "Linux", desc: "Pretty much the best OS"},
		item{id: "item_12", title: "Business school", desc: "Just kidding"},
		item{id: "item_13", title: "Pottery", desc: "Wet clay is a great feeling"},
		item{id: "item_14", title: "Shampoo", desc: "Nothing like clean hair"},
		item{id: "item_15", title: "Table tennis", desc: "It’s surprisingly exhausting"},
		item{id: "item_16", title: "Milk crates", desc: "Great for packing in your extra stuff"},
		item{id: "item_17", title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		item{id: "item_18", title: "Stickers", desc: "The thicker the vinyl the better"},
		item{id: "item_19", title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		item{id: "item_20", title: "Warm light", desc: "Like around 2700 Kelvin"},
		item{id: "item_21", title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		item{id: "item_22", title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		item{id: "item_23", title: "Terrycloth", desc: "In other words, towel fabric"},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Left click on an items title to select it"

	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err) //nolint:forbidigo
		os.Exit(1)
	}
}
