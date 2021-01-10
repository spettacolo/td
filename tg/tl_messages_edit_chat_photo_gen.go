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

// MessagesEditChatPhotoRequest represents TL type `messages.editChatPhoto#ca4c79d8`.
// Changes chat photo and sends a service message on it
//
// See https://core.telegram.org/method/messages.editChatPhoto for reference.
type MessagesEditChatPhotoRequest struct {
	// Chat ID
	ChatID int
	// Photo to be set
	Photo InputChatPhotoClass
}

// MessagesEditChatPhotoRequestTypeID is TL type id of MessagesEditChatPhotoRequest.
const MessagesEditChatPhotoRequestTypeID = 0xca4c79d8

func (e *MessagesEditChatPhotoRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatPhotoRequest) String() string {
	if e == nil {
		return "MessagesEditChatPhotoRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesEditChatPhotoRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(e.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("\tPhoto: ")
	sb.WriteString(fmt.Sprint(e.Photo))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *MessagesEditChatPhotoRequest) TypeID() uint32 {
	return MessagesEditChatPhotoRequestTypeID
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatPhotoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatPhoto#ca4c79d8 as nil")
	}
	b.PutID(MessagesEditChatPhotoRequestTypeID)
	b.PutInt(e.ChatID)
	if e.Photo == nil {
		return fmt.Errorf("unable to encode messages.editChatPhoto#ca4c79d8: field photo is nil")
	}
	if err := e.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatPhoto#ca4c79d8: field photo: %w", err)
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (e *MessagesEditChatPhotoRequest) GetChatID() (value int) {
	return e.ChatID
}

// GetPhoto returns value of Photo field.
func (e *MessagesEditChatPhotoRequest) GetPhoto() (value InputChatPhotoClass) {
	return e.Photo
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatPhotoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatPhoto#ca4c79d8 to nil")
	}
	if err := b.ConsumeID(MessagesEditChatPhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatPhoto#ca4c79d8: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatPhoto#ca4c79d8: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatPhoto#ca4c79d8: field photo: %w", err)
		}
		e.Photo = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditChatPhotoRequest.
var (
	_ bin.Encoder = &MessagesEditChatPhotoRequest{}
	_ bin.Decoder = &MessagesEditChatPhotoRequest{}
)

// MessagesEditChatPhoto invokes method messages.editChatPhoto#ca4c79d8 returning error if any.
// Changes chat photo and sends a service message on it
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  400 INPUT_CONSTRUCTOR_INVALID: The provided constructor is invalid
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 PHOTO_CROP_SIZE_SMALL: Photo is too small
//  400 PHOTO_EXT_INVALID: The extension of the photo is invalid
//  400 PHOTO_INVALID: Photo invalid
//
// See https://core.telegram.org/method/messages.editChatPhoto for reference.
// Can be used by bots.
func (c *Client) MessagesEditChatPhoto(ctx context.Context, request *MessagesEditChatPhotoRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
