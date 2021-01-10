// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// Theme represents TL type `theme#28f1114`.
//
// See https://localhost:80/doc/constructor/theme for reference.
type Theme struct {
	// Name field of Theme.
	Name string
}

// ThemeTypeID is TL type id of Theme.
const ThemeTypeID = 0x28f1114

func (t *Theme) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Theme) String() string {
	if t == nil {
		return "Theme(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Theme")
	sb.WriteString("{\n")
	sb.WriteString("\tName: ")
	sb.WriteString(fmt.Sprint(t.Name))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *Theme) TypeID() uint32 {
	return ThemeTypeID
}

// Encode implements bin.Encoder.
func (t *Theme) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#28f1114 as nil")
	}
	b.PutID(ThemeTypeID)
	b.PutString(t.Name)
	return nil
}

// GetName returns value of Name field.
func (t *Theme) GetName() (value string) {
	return t.Name
}

// Decode implements bin.Decoder.
func (t *Theme) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#28f1114 to nil")
	}
	if err := b.ConsumeID(ThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode theme#28f1114: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#28f1114: field name: %w", err)
		}
		t.Name = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Theme.
var (
	_ bin.Encoder = &Theme{}
	_ bin.Decoder = &Theme{}
)
