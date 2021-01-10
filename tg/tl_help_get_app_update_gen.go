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

// HelpGetAppUpdateRequest represents TL type `help.getAppUpdate#522d5a7d`.
// Returns information on update availability for the current application.
//
// See https://core.telegram.org/method/help.getAppUpdate for reference.
type HelpGetAppUpdateRequest struct {
	// Source
	Source string
}

// HelpGetAppUpdateRequestTypeID is TL type id of HelpGetAppUpdateRequest.
const HelpGetAppUpdateRequestTypeID = 0x522d5a7d

func (g *HelpGetAppUpdateRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Source == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetAppUpdateRequest) String() string {
	if g == nil {
		return "HelpGetAppUpdateRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpGetAppUpdateRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tSource: ")
	sb.WriteString(fmt.Sprint(g.Source))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetAppUpdateRequest) TypeID() uint32 {
	return HelpGetAppUpdateRequestTypeID
}

// Encode implements bin.Encoder.
func (g *HelpGetAppUpdateRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getAppUpdate#522d5a7d as nil")
	}
	b.PutID(HelpGetAppUpdateRequestTypeID)
	b.PutString(g.Source)
	return nil
}

// GetSource returns value of Source field.
func (g *HelpGetAppUpdateRequest) GetSource() (value string) {
	return g.Source
}

// Decode implements bin.Decoder.
func (g *HelpGetAppUpdateRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getAppUpdate#522d5a7d to nil")
	}
	if err := b.ConsumeID(HelpGetAppUpdateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getAppUpdate#522d5a7d: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getAppUpdate#522d5a7d: field source: %w", err)
		}
		g.Source = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetAppUpdateRequest.
var (
	_ bin.Encoder = &HelpGetAppUpdateRequest{}
	_ bin.Decoder = &HelpGetAppUpdateRequest{}
)

// HelpGetAppUpdate invokes method help.getAppUpdate#522d5a7d returning error if any.
// Returns information on update availability for the current application.
//
// See https://core.telegram.org/method/help.getAppUpdate for reference.
func (c *Client) HelpGetAppUpdate(ctx context.Context, source string) (HelpAppUpdateClass, error) {
	var result HelpAppUpdateBox

	request := &HelpGetAppUpdateRequest{
		Source: source,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.AppUpdate, nil
}
