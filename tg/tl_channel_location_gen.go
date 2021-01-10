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

// ChannelLocationEmpty represents TL type `channelLocationEmpty#bfb5ad8b`.
// No location (normal supergroup)
//
// See https://core.telegram.org/constructor/channelLocationEmpty for reference.
type ChannelLocationEmpty struct {
}

// ChannelLocationEmptyTypeID is TL type id of ChannelLocationEmpty.
const ChannelLocationEmptyTypeID = 0xbfb5ad8b

func (c *ChannelLocationEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelLocationEmpty) String() string {
	if c == nil {
		return "ChannelLocationEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelLocationEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelLocationEmpty) TypeID() uint32 {
	return ChannelLocationEmptyTypeID
}

// Encode implements bin.Encoder.
func (c *ChannelLocationEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocationEmpty#bfb5ad8b as nil")
	}
	b.PutID(ChannelLocationEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelLocationEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocationEmpty#bfb5ad8b to nil")
	}
	if err := b.ConsumeID(ChannelLocationEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode channelLocationEmpty#bfb5ad8b: %w", err)
	}
	return nil
}

// construct implements constructor of ChannelLocationClass.
func (c ChannelLocationEmpty) construct() ChannelLocationClass { return &c }

// Ensuring interfaces in compile-time for ChannelLocationEmpty.
var (
	_ bin.Encoder = &ChannelLocationEmpty{}
	_ bin.Decoder = &ChannelLocationEmpty{}

	_ ChannelLocationClass = &ChannelLocationEmpty{}
)

// ChannelLocation represents TL type `channelLocation#209b82db`.
// Geographical location of supergroup (geogroups)
//
// See https://core.telegram.org/constructor/channelLocation for reference.
type ChannelLocation struct {
	// Geographical location of supergrup
	GeoPoint GeoPointClass
	// Textual description of the address
	Address string
}

// ChannelLocationTypeID is TL type id of ChannelLocation.
const ChannelLocationTypeID = 0x209b82db

func (c *ChannelLocation) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.GeoPoint == nil) {
		return false
	}
	if !(c.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelLocation) String() string {
	if c == nil {
		return "ChannelLocation(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelLocation")
	sb.WriteString("{\n")
	sb.WriteString("\tGeoPoint: ")
	sb.WriteString(fmt.Sprint(c.GeoPoint))
	sb.WriteString(",\n")
	sb.WriteString("\tAddress: ")
	sb.WriteString(fmt.Sprint(c.Address))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChannelLocation) TypeID() uint32 {
	return ChannelLocationTypeID
}

// Encode implements bin.Encoder.
func (c *ChannelLocation) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelLocation#209b82db as nil")
	}
	b.PutID(ChannelLocationTypeID)
	if c.GeoPoint == nil {
		return fmt.Errorf("unable to encode channelLocation#209b82db: field geo_point is nil")
	}
	if err := c.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelLocation#209b82db: field geo_point: %w", err)
	}
	b.PutString(c.Address)
	return nil
}

// GetGeoPoint returns value of GeoPoint field.
func (c *ChannelLocation) GetGeoPoint() (value GeoPointClass) {
	return c.GeoPoint
}

// GetAddress returns value of Address field.
func (c *ChannelLocation) GetAddress() (value string) {
	return c.Address
}

// Decode implements bin.Decoder.
func (c *ChannelLocation) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelLocation#209b82db to nil")
	}
	if err := b.ConsumeID(ChannelLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode channelLocation#209b82db: %w", err)
	}
	{
		value, err := DecodeGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode channelLocation#209b82db: field geo_point: %w", err)
		}
		c.GeoPoint = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channelLocation#209b82db: field address: %w", err)
		}
		c.Address = value
	}
	return nil
}

// construct implements constructor of ChannelLocationClass.
func (c ChannelLocation) construct() ChannelLocationClass { return &c }

// Ensuring interfaces in compile-time for ChannelLocation.
var (
	_ bin.Encoder = &ChannelLocation{}
	_ bin.Decoder = &ChannelLocation{}

	_ ChannelLocationClass = &ChannelLocation{}
)

// ChannelLocationClass represents ChannelLocation generic type.
//
// See https://core.telegram.org/type/ChannelLocation for reference.
//
// Example:
//  g, err := DecodeChannelLocation(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChannelLocationEmpty: // channelLocationEmpty#bfb5ad8b
//  case *ChannelLocation: // channelLocation#209b82db
//  default: panic(v)
//  }
type ChannelLocationClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChannelLocationClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeChannelLocation implements binary de-serialization for ChannelLocationClass.
func DecodeChannelLocation(buf *bin.Buffer) (ChannelLocationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelLocationEmptyTypeID:
		// Decoding channelLocationEmpty#bfb5ad8b.
		v := ChannelLocationEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", err)
		}
		return &v, nil
	case ChannelLocationTypeID:
		// Decoding channelLocation#209b82db.
		v := ChannelLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelLocationClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelLocation boxes the ChannelLocationClass providing a helper.
type ChannelLocationBox struct {
	ChannelLocation ChannelLocationClass
}

// Decode implements bin.Decoder for ChannelLocationBox.
func (b *ChannelLocationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelLocationBox to nil")
	}
	v, err := DecodeChannelLocation(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelLocation = v
	return nil
}

// Encode implements bin.Encode for ChannelLocationBox.
func (b *ChannelLocationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelLocation == nil {
		return fmt.Errorf("unable to encode ChannelLocationClass as nil")
	}
	return b.ChannelLocation.Encode(buf)
}
