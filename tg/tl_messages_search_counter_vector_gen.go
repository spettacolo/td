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

// MessagesSearchCounterVector is a box for Vector<messages.SearchCounter>
type MessagesSearchCounterVector struct {
	// Elements of Vector<messages.SearchCounter>
	Elems []MessagesSearchCounter
}

// MessagesSearchCounterVectorTypeID is TL type id of MessagesSearchCounterVector.
const MessagesSearchCounterVectorTypeID = bin.TypeVector

func (vec *MessagesSearchCounterVector) Zero() bool {
	if vec == nil {
		return true
	}
	if !(vec.Elems == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (vec *MessagesSearchCounterVector) String() string {
	if vec == nil {
		return "MessagesSearchCounterVector(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSearchCounterVector")
	sb.WriteByte('[')
	for _, e := range vec.Elems {
		sb.WriteString(fmt.Sprint(e) + ",\n")
	}
	sb.WriteByte(']')
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (vec *MessagesSearchCounterVector) TypeID() uint32 {
	return MessagesSearchCounterVectorTypeID
}

// Encode implements bin.Encoder.
func (vec *MessagesSearchCounterVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<messages.SearchCounter> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<messages.SearchCounter>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetElems returns value of Elems field.
func (vec *MessagesSearchCounterVector) GetElems() (value []MessagesSearchCounter) {
	return vec.Elems
}

// Decode implements bin.Decoder.
func (vec *MessagesSearchCounterVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<messages.SearchCounter> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<messages.SearchCounter>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessagesSearchCounter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<messages.SearchCounter>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSearchCounterVector.
var (
	_ bin.Encoder = &MessagesSearchCounterVector{}
	_ bin.Decoder = &MessagesSearchCounterVector{}
)
