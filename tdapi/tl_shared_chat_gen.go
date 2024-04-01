// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// SharedChat represents TL type `sharedChat#4a87b01a`.
type SharedChat struct {
	// Chat identifier
	ChatID int64
	// Title of the chat; for bots only
	Title string
	// Username of the chat; for bots only
	Username string
	// Photo of the chat; for bots only; may be null
	Photo Photo
}

// SharedChatTypeID is TL type id of SharedChat.
const SharedChatTypeID = 0x4a87b01a

// Ensuring interfaces in compile-time for SharedChat.
var (
	_ bin.Encoder     = &SharedChat{}
	_ bin.Decoder     = &SharedChat{}
	_ bin.BareEncoder = &SharedChat{}
	_ bin.BareDecoder = &SharedChat{}
)

func (s *SharedChat) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Username == "") {
		return false
	}
	if !(s.Photo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SharedChat) String() string {
	if s == nil {
		return "SharedChat(nil)"
	}
	type Alias SharedChat
	return fmt.Sprintf("SharedChat%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SharedChat) TypeID() uint32 {
	return SharedChatTypeID
}

// TypeName returns name of type in TL schema.
func (*SharedChat) TypeName() string {
	return "sharedChat"
}

// TypeInfo returns info about TL type.
func (s *SharedChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sharedChat",
		ID:   SharedChatTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SharedChat) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sharedChat#4a87b01a as nil")
	}
	b.PutID(SharedChatTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SharedChat) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sharedChat#4a87b01a as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutString(s.Title)
	b.PutString(s.Username)
	if err := s.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sharedChat#4a87b01a: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SharedChat) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sharedChat#4a87b01a to nil")
	}
	if err := b.ConsumeID(SharedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode sharedChat#4a87b01a: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SharedChat) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sharedChat#4a87b01a to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sharedChat#4a87b01a: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sharedChat#4a87b01a: field title: %w", err)
		}
		s.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sharedChat#4a87b01a: field username: %w", err)
		}
		s.Username = value
	}
	{
		if err := s.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sharedChat#4a87b01a: field photo: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SharedChat) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sharedChat#4a87b01a as nil")
	}
	b.ObjStart()
	b.PutID("sharedChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("title")
	b.PutString(s.Title)
	b.Comma()
	b.FieldStart("username")
	b.PutString(s.Username)
	b.Comma()
	b.FieldStart("photo")
	if err := s.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sharedChat#4a87b01a: field photo: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SharedChat) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sharedChat#4a87b01a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sharedChat"); err != nil {
				return fmt.Errorf("unable to decode sharedChat#4a87b01a: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sharedChat#4a87b01a: field chat_id: %w", err)
			}
			s.ChatID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sharedChat#4a87b01a: field title: %w", err)
			}
			s.Title = value
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sharedChat#4a87b01a: field username: %w", err)
			}
			s.Username = value
		case "photo":
			if err := s.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sharedChat#4a87b01a: field photo: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SharedChat) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetTitle returns value of Title field.
func (s *SharedChat) GetTitle() (value string) {
	if s == nil {
		return
	}
	return s.Title
}

// GetUsername returns value of Username field.
func (s *SharedChat) GetUsername() (value string) {
	if s == nil {
		return
	}
	return s.Username
}

// GetPhoto returns value of Photo field.
func (s *SharedChat) GetPhoto() (value Photo) {
	if s == nil {
		return
	}
	return s.Photo
}
