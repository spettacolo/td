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

// HelpGetNearestDcRequest represents TL type `help.getNearestDc#1fb33026`.
// Returns info on data centre nearest to the user.
//
// See https://core.telegram.org/method/help.getNearestDc for reference.
type HelpGetNearestDcRequest struct {
}

// HelpGetNearestDcRequestTypeID is TL type id of HelpGetNearestDcRequest.
const HelpGetNearestDcRequestTypeID = 0x1fb33026

func (g *HelpGetNearestDcRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetNearestDcRequest) String() string {
	if g == nil {
		return "HelpGetNearestDcRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpGetNearestDcRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetNearestDcRequest) TypeID() uint32 {
	return HelpGetNearestDcRequestTypeID
}

// Encode implements bin.Encoder.
func (g *HelpGetNearestDcRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getNearestDc#1fb33026 as nil")
	}
	b.PutID(HelpGetNearestDcRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetNearestDcRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getNearestDc#1fb33026 to nil")
	}
	if err := b.ConsumeID(HelpGetNearestDcRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getNearestDc#1fb33026: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetNearestDcRequest.
var (
	_ bin.Encoder = &HelpGetNearestDcRequest{}
	_ bin.Decoder = &HelpGetNearestDcRequest{}
)

// HelpGetNearestDc invokes method help.getNearestDc#1fb33026 returning error if any.
// Returns info on data centre nearest to the user.
//
// See https://core.telegram.org/method/help.getNearestDc for reference.
func (c *Client) HelpGetNearestDc(ctx context.Context) (*NearestDc, error) {
	var result NearestDc

	request := &HelpGetNearestDcRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
