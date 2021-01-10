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

// CdnPublicKey represents TL type `cdnPublicKey#c982eaba`.
// Public key to use only during handshakes to CDN¹ DCs.
//
// Links:
//  1) https://core.telegram.org/cdn
//
// See https://core.telegram.org/constructor/cdnPublicKey for reference.
type CdnPublicKey struct {
	// CDN DC¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/cdn
	DCID int
	// RSA public key
	PublicKey string
}

// CdnPublicKeyTypeID is TL type id of CdnPublicKey.
const CdnPublicKeyTypeID = 0xc982eaba

func (c *CdnPublicKey) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.DCID == 0) {
		return false
	}
	if !(c.PublicKey == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CdnPublicKey) String() string {
	if c == nil {
		return "CdnPublicKey(nil)"
	}
	var sb strings.Builder
	sb.WriteString("CdnPublicKey")
	sb.WriteString("{\n")
	sb.WriteString("\tDCID: ")
	sb.WriteString(fmt.Sprint(c.DCID))
	sb.WriteString(",\n")
	sb.WriteString("\tPublicKey: ")
	sb.WriteString(fmt.Sprint(c.PublicKey))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *CdnPublicKey) TypeID() uint32 {
	return CdnPublicKeyTypeID
}

// Encode implements bin.Encoder.
func (c *CdnPublicKey) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode cdnPublicKey#c982eaba as nil")
	}
	b.PutID(CdnPublicKeyTypeID)
	b.PutInt(c.DCID)
	b.PutString(c.PublicKey)
	return nil
}

// GetDCID returns value of DCID field.
func (c *CdnPublicKey) GetDCID() (value int) {
	return c.DCID
}

// GetPublicKey returns value of PublicKey field.
func (c *CdnPublicKey) GetPublicKey() (value string) {
	return c.PublicKey
}

// Decode implements bin.Decoder.
func (c *CdnPublicKey) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode cdnPublicKey#c982eaba to nil")
	}
	if err := b.ConsumeID(CdnPublicKeyTypeID); err != nil {
		return fmt.Errorf("unable to decode cdnPublicKey#c982eaba: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode cdnPublicKey#c982eaba: field dc_id: %w", err)
		}
		c.DCID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode cdnPublicKey#c982eaba: field public_key: %w", err)
		}
		c.PublicKey = value
	}
	return nil
}

// Ensuring interfaces in compile-time for CdnPublicKey.
var (
	_ bin.Encoder = &CdnPublicKey{}
	_ bin.Decoder = &CdnPublicKey{}
)
