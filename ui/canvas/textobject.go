package canvas

import "image/color"

import "github.com/fyne-io/fyne/ui"
import "github.com/fyne-io/fyne/ui/theme"

// TextObject describes a text primitive in a Fyne canvas
type TextObject struct {
	Size     ui.Size     // The current size of the TextObject - the font will not scale to this Size
	Position ui.Position // The current position of the TextObject
	Color    color.RGBA  // The main text draw colour

	Text     string // The string content of this TextObject
	FontSize int    // Size of the font - if the Canvas scale is 1.0 this will be equivalent to point size
	Bold     bool   // Should the text be bold
	Italic   bool   // Should the text be italic

	minSize ui.Size
}

// CurrentSize gets the current size of this text object
func (t *TextObject) CurrentSize() ui.Size {
	return t.Size
}

// Resize sets a new size for the text object
func (t *TextObject) Resize(size ui.Size) {
	t.Size = size
}

// CurrentPosition gets the current position of this text object, relative to it's parent / canvas
func (t *TextObject) CurrentPosition() ui.Position {
	return t.Position
}

// Move the text object to a new position, relative to it's parent / canvas
func (t *TextObject) Move(pos ui.Position) {
	t.Position = pos
}

// SetMinSize is used by the render implementation to specify the space the text requires.
// This should not be called by application code.
func (t *TextObject) SetMinSize(size ui.Size) {
	t.minSize = size
}

// MinSize returns the minimum size of this text objet based on it's font size and content.
// This is normally determined by the render implementation.
func (t *TextObject) MinSize() ui.Size {
	return t.minSize
}

// NewText returns a new TextObject implementation
func NewText(text string) *TextObject {
	return &TextObject{
		Color:    theme.TextColor(),
		Text:     text,
		FontSize: theme.TextSize(),
	}
}
