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

// MessagesImportChatInviteRequest represents TL type `messages.importChatInvite#6c50051c`.
// Import a chat invite and join a private chat/supergroup/channel
//
// See https://core.telegram.org/method/messages.importChatInvite for reference.
type MessagesImportChatInviteRequest struct {
	// hash from t.me/joinchat/hash
	Hash string
}

// MessagesImportChatInviteRequestTypeID is TL type id of MessagesImportChatInviteRequest.
const MessagesImportChatInviteRequestTypeID = 0x6c50051c

// Ensuring interfaces in compile-time for MessagesImportChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesImportChatInviteRequest{}
	_ bin.Decoder     = &MessagesImportChatInviteRequest{}
	_ bin.BareEncoder = &MessagesImportChatInviteRequest{}
	_ bin.BareDecoder = &MessagesImportChatInviteRequest{}
)

func (i *MessagesImportChatInviteRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Hash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *MessagesImportChatInviteRequest) String() string {
	if i == nil {
		return "MessagesImportChatInviteRequest(nil)"
	}
	type Alias MessagesImportChatInviteRequest
	return fmt.Sprintf("MessagesImportChatInviteRequest%+v", Alias(*i))
}

// FillFrom fills MessagesImportChatInviteRequest from given interface.
func (i *MessagesImportChatInviteRequest) FillFrom(from interface {
	GetHash() (value string)
}) {
	i.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesImportChatInviteRequest) TypeID() uint32 {
	return MessagesImportChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesImportChatInviteRequest) TypeName() string {
	return "messages.importChatInvite"
}

// TypeInfo returns info about TL type.
func (i *MessagesImportChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.importChatInvite",
		ID:   MessagesImportChatInviteRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *MessagesImportChatInviteRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.importChatInvite#6c50051c as nil")
	}
	b.PutID(MessagesImportChatInviteRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *MessagesImportChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.importChatInvite#6c50051c as nil")
	}
	b.PutString(i.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (i *MessagesImportChatInviteRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.importChatInvite#6c50051c to nil")
	}
	if err := b.ConsumeID(MessagesImportChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.importChatInvite#6c50051c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *MessagesImportChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.importChatInvite#6c50051c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.importChatInvite#6c50051c: field hash: %w", err)
		}
		i.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (i *MessagesImportChatInviteRequest) GetHash() (value string) {
	return i.Hash
}

// MessagesImportChatInvite invokes method messages.importChatInvite#6c50051c returning error if any.
// Import a chat invite and join a private chat/supergroup/channel
//
// Possible errors:
//  400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups.
//  400 CHANNEL_INVALID: The provided channel is invalid.
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//  400 CHAT_INVALID: Invalid chat.
//  400 INVITE_HASH_EMPTY: The invite hash is empty.
//  400 INVITE_HASH_EXPIRED: The invite link has expired.
//  400 INVITE_HASH_INVALID: The invite hash is invalid.
//  400 MSG_ID_INVALID: Invalid message ID provided.
//  400 PEER_ID_INVALID: The provided peer id is invalid.
//  400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example).
//  400 USER_ALREADY_PARTICIPANT: The user is already in the group.
//  400 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups.
//
// See https://core.telegram.org/method/messages.importChatInvite for reference.
func (c *Client) MessagesImportChatInvite(ctx context.Context, hash string) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesImportChatInviteRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
