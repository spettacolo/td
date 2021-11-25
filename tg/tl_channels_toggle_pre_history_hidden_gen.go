// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// ChannelsTogglePreHistoryHiddenRequest represents TL type `channels.togglePreHistoryHidden#eabbb94c`.
// Hide/unhide message history for new channel/supergroup users
//
// See https://core.telegram.org/method/channels.togglePreHistoryHidden for reference.
type ChannelsTogglePreHistoryHiddenRequest struct {
	// Channel/supergroup
	Channel InputChannelClass
	// Hide/unhide
	Enabled bool
}

// ChannelsTogglePreHistoryHiddenRequestTypeID is TL type id of ChannelsTogglePreHistoryHiddenRequest.
const ChannelsTogglePreHistoryHiddenRequestTypeID = 0xeabbb94c

// Ensuring interfaces in compile-time for ChannelsTogglePreHistoryHiddenRequest.
var (
	_ bin.Encoder     = &ChannelsTogglePreHistoryHiddenRequest{}
	_ bin.Decoder     = &ChannelsTogglePreHistoryHiddenRequest{}
	_ bin.BareEncoder = &ChannelsTogglePreHistoryHiddenRequest{}
	_ bin.BareDecoder = &ChannelsTogglePreHistoryHiddenRequest{}
)

func (t *ChannelsTogglePreHistoryHiddenRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Channel == nil) {
		return false
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ChannelsTogglePreHistoryHiddenRequest) String() string {
	if t == nil {
		return "ChannelsTogglePreHistoryHiddenRequest(nil)"
	}
	type Alias ChannelsTogglePreHistoryHiddenRequest
	return fmt.Sprintf("ChannelsTogglePreHistoryHiddenRequest%+v", Alias(*t))
}

// FillFrom fills ChannelsTogglePreHistoryHiddenRequest from given interface.
func (t *ChannelsTogglePreHistoryHiddenRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetEnabled() (value bool)
}) {
	t.Channel = from.GetChannel()
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsTogglePreHistoryHiddenRequest) TypeID() uint32 {
	return ChannelsTogglePreHistoryHiddenRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsTogglePreHistoryHiddenRequest) TypeName() string {
	return "channels.togglePreHistoryHidden"
}

// TypeInfo returns info about TL type.
func (t *ChannelsTogglePreHistoryHiddenRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.togglePreHistoryHidden",
		ID:   ChannelsTogglePreHistoryHiddenRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ChannelsTogglePreHistoryHiddenRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.togglePreHistoryHidden#eabbb94c as nil")
	}
	b.PutID(ChannelsTogglePreHistoryHiddenRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ChannelsTogglePreHistoryHiddenRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode channels.togglePreHistoryHidden#eabbb94c as nil")
	}
	if t.Channel == nil {
		return fmt.Errorf("unable to encode channels.togglePreHistoryHidden#eabbb94c: field channel is nil")
	}
	if err := t.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.togglePreHistoryHidden#eabbb94c: field channel: %w", err)
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ChannelsTogglePreHistoryHiddenRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.togglePreHistoryHidden#eabbb94c to nil")
	}
	if err := b.ConsumeID(ChannelsTogglePreHistoryHiddenRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.togglePreHistoryHidden#eabbb94c: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ChannelsTogglePreHistoryHiddenRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode channels.togglePreHistoryHidden#eabbb94c to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.togglePreHistoryHidden#eabbb94c: field channel: %w", err)
		}
		t.Channel = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode channels.togglePreHistoryHidden#eabbb94c: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (t *ChannelsTogglePreHistoryHiddenRequest) GetChannel() (value InputChannelClass) {
	return t.Channel
}

// GetEnabled returns value of Enabled field.
func (t *ChannelsTogglePreHistoryHiddenRequest) GetEnabled() (value bool) {
	return t.Enabled
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (t *ChannelsTogglePreHistoryHiddenRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return t.Channel.AsNotEmpty()
}

// ChannelsTogglePreHistoryHidden invokes method channels.togglePreHistoryHidden#eabbb94c returning error if any.
// Hide/unhide message history for new channel/supergroup users
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//  400 CHAT_ID_INVALID: The provided chat id is invalid.
//  400 CHAT_LINK_EXISTS: The chat is public, you can't hide the history to new users.
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified.
//
// See https://core.telegram.org/method/channels.togglePreHistoryHidden for reference.
func (c *Client) ChannelsTogglePreHistoryHidden(ctx context.Context, request *ChannelsTogglePreHistoryHiddenRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
