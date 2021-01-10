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

// MessagesGetDhConfigRequest represents TL type `messages.getDhConfig#26cf8950`.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a random sequence of bytes of required length.
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
type MessagesGetDhConfigRequest struct {
	// Value of the version parameter from messages.dhConfig¹, avialable at the client
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.dhConfig
	Version int
	// Length of the required random sequence
	RandomLength int
}

// MessagesGetDhConfigRequestTypeID is TL type id of MessagesGetDhConfigRequest.
const MessagesGetDhConfigRequestTypeID = 0x26cf8950

func (g *MessagesGetDhConfigRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Version == 0) {
		return false
	}
	if !(g.RandomLength == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDhConfigRequest) String() string {
	if g == nil {
		return "MessagesGetDhConfigRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetDhConfigRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tVersion: ")
	sb.WriteString(fmt.Sprint(g.Version))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomLength: ")
	sb.WriteString(fmt.Sprint(g.RandomLength))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *MessagesGetDhConfigRequest) TypeID() uint32 {
	return MessagesGetDhConfigRequestTypeID
}

// Encode implements bin.Encoder.
func (g *MessagesGetDhConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDhConfig#26cf8950 as nil")
	}
	b.PutID(MessagesGetDhConfigRequestTypeID)
	b.PutInt(g.Version)
	b.PutInt(g.RandomLength)
	return nil
}

// GetVersion returns value of Version field.
func (g *MessagesGetDhConfigRequest) GetVersion() (value int) {
	return g.Version
}

// GetRandomLength returns value of RandomLength field.
func (g *MessagesGetDhConfigRequest) GetRandomLength() (value int) {
	return g.RandomLength
}

// Decode implements bin.Decoder.
func (g *MessagesGetDhConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDhConfig#26cf8950 to nil")
	}
	if err := b.ConsumeID(MessagesGetDhConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field version: %w", err)
		}
		g.Version = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDhConfig#26cf8950: field random_length: %w", err)
		}
		g.RandomLength = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetDhConfigRequest.
var (
	_ bin.Encoder = &MessagesGetDhConfigRequest{}
	_ bin.Decoder = &MessagesGetDhConfigRequest{}
)

// MessagesGetDhConfig invokes method messages.getDhConfig#26cf8950 returning error if any.
// Returns configuration parameters for Diffie-Hellman key generation. Can also return a random sequence of bytes of required length.
//
// Possible errors:
//  400 RANDOM_LENGTH_INVALID: Random length invalid
//
// See https://core.telegram.org/method/messages.getDhConfig for reference.
func (c *Client) MessagesGetDhConfig(ctx context.Context, request *MessagesGetDhConfigRequest) (MessagesDhConfigClass, error) {
	var result MessagesDhConfigBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.DhConfig, nil
}
