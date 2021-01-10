// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// Null represents TL type `null#56730bcc`.
// Corresponds to an arbitrary empty object.
//
// See https://core.telegram.org/constructor/null for reference.
type Null struct {
}

// NullTypeID is TL type id of Null.
const NullTypeID = 0x56730bcc

func (n *Null) Zero() bool {
	if n == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (n *Null) String() string {
	if n == nil {
		return "Null(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Null")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (n *Null) TypeID() uint32 {
	return NullTypeID
}

// Encode implements bin.Encoder.
func (n *Null) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode null#56730bcc as nil")
	}
	b.PutID(NullTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (n *Null) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode null#56730bcc to nil")
	}
	if err := b.ConsumeID(NullTypeID); err != nil {
		return fmt.Errorf("unable to decode null#56730bcc: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for Null.
var (
	_ bin.Encoder = &Null{}
	_ bin.Decoder = &Null{}
)
