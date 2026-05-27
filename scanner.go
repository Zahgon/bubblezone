// Copyright (c) Liam Stanley <liam@liam.sh>. All rights reserved. Use of
// this source code is governed by the MIT license that can be found in
// the LICENSE file.

package zone

const (
	eof = 1
)

type stateFn func(*scanner) stateFn

type scanner struct {
	manager   *Manager
	enabled   bool
	iteration int

	input string // Source input.
	pos   int    // Current position in the input.
	start int    // Start position of the current marker.
	width int    // Width of the current rune.

	// Used for width and height tracking.
	newlines    int
	lastNewline int

	// tracked is the temporary location for starting markers.
	tracked map[string]*ZoneInfo
}

func newScanner(m *Manager, input string, iteration int) *scanner {
	_ = "STUB: not implemented"
	return nil
}

// run initializes the scanner and starts the state machine.
func (s *scanner) run() { _ = "STUB: not implemented"; return }

// emit adds the current marker to the tracked map. If two markers are received,
// it is sent back to the manager.
func (s *scanner) emit() {
	_ = "STUB: not implemented"

	// If the manager is disabled, we don't need to track anything, just strip
	// the markers from the resulting output.
	return
}

// The end should be - 1, because it's the end of the encapsulation of the
// zone, and isn't actually taking up another space.

// next returns the next rune in the input, incrementing the position.
func (s *scanner) next() (r rune) { _ = "STUB: not implemented"; return 0 }

// backup steps back one rune. Can only be called once per call of next.
func (s *scanner) backup() {
	_ = "STUB: not implemented"

	// peek steps forward one rune, reads, and backs up again.
	return
}

func (s *scanner) peek() rune { _ = "STUB: not implemented"; return 0 }

// scanMain is the entrypoint into the state machine.
func scanMain(s *scanner) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

// scanID scans forward, matching a marker, otherwise cancelling and returning
// to scanMain if a valid marker isn't found.
func scanID(s *scanner) stateFn { _ = "STUB: not implemented"; return *new(stateFn) }

func isNumber(r rune) bool { _ = "STUB: not implemented"; return false }

// printableRuneWidth returns the printable cell width of the given string.
func printableRuneWidth(s string) int { _ = "STUB: not implemented"; return 0 }

// Start of ANSI escape sequence.

// Check if at the end of an ANSI escape sequence (terminator).
