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

// ChatJoinRequestsInfo represents TL type `chatJoinRequestsInfo#6aff5df5`.
type ChatJoinRequestsInfo struct {
	// Total number of pending join requests
	TotalCount int32
	// Identifiers of at most 3 users sent the newest pending join requests
	UserIDs []int64
}

// ChatJoinRequestsInfoTypeID is TL type id of ChatJoinRequestsInfo.
const ChatJoinRequestsInfoTypeID = 0x6aff5df5

// Ensuring interfaces in compile-time for ChatJoinRequestsInfo.
var (
	_ bin.Encoder     = &ChatJoinRequestsInfo{}
	_ bin.Decoder     = &ChatJoinRequestsInfo{}
	_ bin.BareEncoder = &ChatJoinRequestsInfo{}
	_ bin.BareDecoder = &ChatJoinRequestsInfo{}
)

func (c *ChatJoinRequestsInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.TotalCount == 0) {
		return false
	}
	if !(c.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatJoinRequestsInfo) String() string {
	if c == nil {
		return "ChatJoinRequestsInfo(nil)"
	}
	type Alias ChatJoinRequestsInfo
	return fmt.Sprintf("ChatJoinRequestsInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatJoinRequestsInfo) TypeID() uint32 {
	return ChatJoinRequestsInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatJoinRequestsInfo) TypeName() string {
	return "chatJoinRequestsInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatJoinRequestsInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatJoinRequestsInfo",
		ID:   ChatJoinRequestsInfoTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatJoinRequestsInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatJoinRequestsInfo#6aff5df5 as nil")
	}
	b.PutID(ChatJoinRequestsInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatJoinRequestsInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatJoinRequestsInfo#6aff5df5 as nil")
	}
	b.PutInt32(c.TotalCount)
	b.PutInt(len(c.UserIDs))
	for _, v := range c.UserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatJoinRequestsInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatJoinRequestsInfo#6aff5df5 to nil")
	}
	if err := b.ConsumeID(ChatJoinRequestsInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatJoinRequestsInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatJoinRequestsInfo#6aff5df5 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: field total_count: %w", err)
		}
		c.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: field user_ids: %w", err)
		}

		if headerLen > 0 {
			c.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: field user_ids: %w", err)
			}
			c.UserIDs = append(c.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatJoinRequestsInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatJoinRequestsInfo#6aff5df5 as nil")
	}
	b.ObjStart()
	b.PutID("chatJoinRequestsInfo")
	b.FieldStart("total_count")
	b.PutInt32(c.TotalCount)
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range c.UserIDs {
		b.PutInt53(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatJoinRequestsInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatJoinRequestsInfo#6aff5df5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatJoinRequestsInfo"); err != nil {
				return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: field total_count: %w", err)
			}
			c.TotalCount = value
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: field user_ids: %w", err)
				}
				c.UserIDs = append(c.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatJoinRequestsInfo#6aff5df5: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (c *ChatJoinRequestsInfo) GetTotalCount() (value int32) {
	return c.TotalCount
}

// GetUserIDs returns value of UserIDs field.
func (c *ChatJoinRequestsInfo) GetUserIDs() (value []int64) {
	return c.UserIDs
}
