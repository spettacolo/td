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

// MessagesSendScreenshotNotificationRequest represents TL type `messages.sendScreenshotNotification#c97df020`.
// Notify the other user in a private chat that a screenshot of the chat was taken
//
// See https://core.telegram.org/method/messages.sendScreenshotNotification for reference.
type MessagesSendScreenshotNotificationRequest struct {
	// Other user
	Peer InputPeerClass
	// ID of message that was screenshotted, can be 0
	ReplyToMsgID int
	// Random ID to avoid message resending
	RandomID int64
}

// MessagesSendScreenshotNotificationRequestTypeID is TL type id of MessagesSendScreenshotNotificationRequest.
const MessagesSendScreenshotNotificationRequestTypeID = 0xc97df020

func (s *MessagesSendScreenshotNotificationRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.ReplyToMsgID == 0) {
		return false
	}
	if !(s.RandomID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSendScreenshotNotificationRequest) String() string {
	if s == nil {
		return "MessagesSendScreenshotNotificationRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSendScreenshotNotificationRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tReplyToMsgID: ")
	sb.WriteString(fmt.Sprint(s.ReplyToMsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(s.RandomID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSendScreenshotNotificationRequest) TypeID() uint32 {
	return MessagesSendScreenshotNotificationRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSendScreenshotNotificationRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sendScreenshotNotification#c97df020 as nil")
	}
	b.PutID(MessagesSendScreenshotNotificationRequestTypeID)
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.sendScreenshotNotification#c97df020: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.sendScreenshotNotification#c97df020: field peer: %w", err)
	}
	b.PutInt(s.ReplyToMsgID)
	b.PutLong(s.RandomID)
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSendScreenshotNotificationRequest) GetPeer() (value InputPeerClass) {
	return s.Peer
}

// GetReplyToMsgID returns value of ReplyToMsgID field.
func (s *MessagesSendScreenshotNotificationRequest) GetReplyToMsgID() (value int) {
	return s.ReplyToMsgID
}

// GetRandomID returns value of RandomID field.
func (s *MessagesSendScreenshotNotificationRequest) GetRandomID() (value int64) {
	return s.RandomID
}

// Decode implements bin.Decoder.
func (s *MessagesSendScreenshotNotificationRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sendScreenshotNotification#c97df020 to nil")
	}
	if err := b.ConsumeID(MessagesSendScreenshotNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sendScreenshotNotification#c97df020: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendScreenshotNotification#c97df020: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendScreenshotNotification#c97df020: field reply_to_msg_id: %w", err)
		}
		s.ReplyToMsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sendScreenshotNotification#c97df020: field random_id: %w", err)
		}
		s.RandomID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSendScreenshotNotificationRequest.
var (
	_ bin.Encoder = &MessagesSendScreenshotNotificationRequest{}
	_ bin.Decoder = &MessagesSendScreenshotNotificationRequest{}
)

// MessagesSendScreenshotNotification invokes method messages.sendScreenshotNotification#c97df020 returning error if any.
// Notify the other user in a private chat that a screenshot of the chat was taken
//
// Possible errors:
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.sendScreenshotNotification for reference.
func (c *Client) MessagesSendScreenshotNotification(ctx context.Context, request *MessagesSendScreenshotNotificationRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
