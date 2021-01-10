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

// TestVectorStringObject represents TL type `testVectorStringObject#e5ecc0d`.
//
// See https://localhost:80/doc/constructor/testVectorStringObject for reference.
type TestVectorStringObject struct {
	// Vector of objects
	Value []TestString
}

// TestVectorStringObjectTypeID is TL type id of TestVectorStringObject.
const TestVectorStringObjectTypeID = 0xe5ecc0d

func (t *TestVectorStringObject) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorStringObject) String() string {
	if t == nil {
		return "TestVectorStringObject(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TestVectorStringObject")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range t.Value {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TestVectorStringObject) TypeID() uint32 {
	return TestVectorStringObjectTypeID
}

// Encode implements bin.Encoder.
func (t *TestVectorStringObject) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorStringObject#e5ecc0d as nil")
	}
	b.PutID(TestVectorStringObjectTypeID)
	b.PutVectorHeader(len(t.Value))
	for idx, v := range t.Value {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode testVectorStringObject#e5ecc0d: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorStringObject) GetValue() (value []TestString) {
	return t.Value
}

// Decode implements bin.Decoder.
func (t *TestVectorStringObject) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorStringObject#e5ecc0d to nil")
	}
	if err := b.ConsumeID(TestVectorStringObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorStringObject#e5ecc0d: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorStringObject#e5ecc0d: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestString
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode testVectorStringObject#e5ecc0d: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorStringObject.
var (
	_ bin.Encoder = &TestVectorStringObject{}
	_ bin.Decoder = &TestVectorStringObject{}
)
