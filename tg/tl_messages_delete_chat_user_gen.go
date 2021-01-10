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

// MessagesDeleteChatUserRequest represents TL type `messages.deleteChatUser#e0611f16`.
// Deletes a user from a chat and sends a service message on it.
//
// See https://core.telegram.org/method/messages.deleteChatUser for reference.
type MessagesDeleteChatUserRequest struct {
	// Chat ID
	ChatID int
	// User ID to be deleted
	UserID InputUserClass
}

// MessagesDeleteChatUserRequestTypeID is TL type id of MessagesDeleteChatUserRequest.
const MessagesDeleteChatUserRequestTypeID = 0xe0611f16

func (d *MessagesDeleteChatUserRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteChatUserRequest) String() string {
	if d == nil {
		return "MessagesDeleteChatUserRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesDeleteChatUserRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(d.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(d.UserID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *MessagesDeleteChatUserRequest) TypeID() uint32 {
	return MessagesDeleteChatUserRequestTypeID
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteChatUserRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteChatUser#e0611f16 as nil")
	}
	b.PutID(MessagesDeleteChatUserRequestTypeID)
	b.PutInt(d.ChatID)
	if d.UserID == nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#e0611f16: field user_id is nil")
	}
	if err := d.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteChatUser#e0611f16: field user_id: %w", err)
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (d *MessagesDeleteChatUserRequest) GetChatID() (value int) {
	return d.ChatID
}

// GetUserID returns value of UserID field.
func (d *MessagesDeleteChatUserRequest) GetUserID() (value InputUserClass) {
	return d.UserID
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteChatUserRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteChatUser#e0611f16 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteChatUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteChatUser#e0611f16: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#e0611f16: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteChatUser#e0611f16: field user_id: %w", err)
		}
		d.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesDeleteChatUserRequest.
var (
	_ bin.Encoder = &MessagesDeleteChatUserRequest{}
	_ bin.Decoder = &MessagesDeleteChatUserRequest{}
)

// MessagesDeleteChatUser invokes method messages.deleteChatUser#e0611f16 returning error if any.
// Deletes a user from a chat and sends a service message on it.
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//  400 USER_NOT_PARTICIPANT: You're not a member of this supergroup/channel
//
// See https://core.telegram.org/method/messages.deleteChatUser for reference.
// Can be used by bots.
func (c *Client) MessagesDeleteChatUser(ctx context.Context, request *MessagesDeleteChatUserRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
