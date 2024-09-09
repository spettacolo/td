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

// ChatBoostStatus represents TL type `chatBoostStatus#401e753c`.
type ChatBoostStatus struct {
	// An HTTP URL, which can be used to boost the chat
	BoostURL string
	// Identifiers of boost slots of the current user applied to the chat
	AppliedSlotIDs []int32
	// Current boost level of the chat
	Level int32
	// The number of boosts received by the chat from created Telegram Premium gift codes and
	// giveaways; always 0 if the current user isn't an administrator in the chat
	GiftCodeBoostCount int32
	// The number of boosts received by the chat
	BoostCount int32
	// The number of boosts added to reach the current level
	CurrentLevelBoostCount int32
	// The number of boosts needed to reach the next level; 0 if the next level isn't
	// available
	NextLevelBoostCount int32
	// Approximate number of Telegram Premium subscribers joined the chat; always 0 if the
	// current user isn't an administrator in the chat
	PremiumMemberCount int32
	// A percentage of Telegram Premium subscribers joined the chat; always 0 if the current
	// user isn't an administrator in the chat
	PremiumMemberPercentage float64
	// The list of prepaid giveaways available for the chat; only for chat administrators
	PrepaidGiveaways []PrepaidGiveaway
}

// ChatBoostStatusTypeID is TL type id of ChatBoostStatus.
const ChatBoostStatusTypeID = 0x401e753c

// Ensuring interfaces in compile-time for ChatBoostStatus.
var (
	_ bin.Encoder     = &ChatBoostStatus{}
	_ bin.Decoder     = &ChatBoostStatus{}
	_ bin.BareEncoder = &ChatBoostStatus{}
	_ bin.BareDecoder = &ChatBoostStatus{}
)

func (c *ChatBoostStatus) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.BoostURL == "") {
		return false
	}
	if !(c.AppliedSlotIDs == nil) {
		return false
	}
	if !(c.Level == 0) {
		return false
	}
	if !(c.GiftCodeBoostCount == 0) {
		return false
	}
	if !(c.BoostCount == 0) {
		return false
	}
	if !(c.CurrentLevelBoostCount == 0) {
		return false
	}
	if !(c.NextLevelBoostCount == 0) {
		return false
	}
	if !(c.PremiumMemberCount == 0) {
		return false
	}
	if !(c.PremiumMemberPercentage == 0) {
		return false
	}
	if !(c.PrepaidGiveaways == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatBoostStatus) String() string {
	if c == nil {
		return "ChatBoostStatus(nil)"
	}
	type Alias ChatBoostStatus
	return fmt.Sprintf("ChatBoostStatus%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatBoostStatus) TypeID() uint32 {
	return ChatBoostStatusTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatBoostStatus) TypeName() string {
	return "chatBoostStatus"
}

// TypeInfo returns info about TL type.
func (c *ChatBoostStatus) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatBoostStatus",
		ID:   ChatBoostStatusTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BoostURL",
			SchemaName: "boost_url",
		},
		{
			Name:       "AppliedSlotIDs",
			SchemaName: "applied_slot_ids",
		},
		{
			Name:       "Level",
			SchemaName: "level",
		},
		{
			Name:       "GiftCodeBoostCount",
			SchemaName: "gift_code_boost_count",
		},
		{
			Name:       "BoostCount",
			SchemaName: "boost_count",
		},
		{
			Name:       "CurrentLevelBoostCount",
			SchemaName: "current_level_boost_count",
		},
		{
			Name:       "NextLevelBoostCount",
			SchemaName: "next_level_boost_count",
		},
		{
			Name:       "PremiumMemberCount",
			SchemaName: "premium_member_count",
		},
		{
			Name:       "PremiumMemberPercentage",
			SchemaName: "premium_member_percentage",
		},
		{
			Name:       "PrepaidGiveaways",
			SchemaName: "prepaid_giveaways",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatBoostStatus) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostStatus#401e753c as nil")
	}
	b.PutID(ChatBoostStatusTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatBoostStatus) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostStatus#401e753c as nil")
	}
	b.PutString(c.BoostURL)
	b.PutInt(len(c.AppliedSlotIDs))
	for _, v := range c.AppliedSlotIDs {
		b.PutInt32(v)
	}
	b.PutInt32(c.Level)
	b.PutInt32(c.GiftCodeBoostCount)
	b.PutInt32(c.BoostCount)
	b.PutInt32(c.CurrentLevelBoostCount)
	b.PutInt32(c.NextLevelBoostCount)
	b.PutInt32(c.PremiumMemberCount)
	b.PutDouble(c.PremiumMemberPercentage)
	b.PutInt(len(c.PrepaidGiveaways))
	for idx, v := range c.PrepaidGiveaways {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatBoostStatus#401e753c: field prepaid_giveaways element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatBoostStatus) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostStatus#401e753c to nil")
	}
	if err := b.ConsumeID(ChatBoostStatusTypeID); err != nil {
		return fmt.Errorf("unable to decode chatBoostStatus#401e753c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatBoostStatus) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostStatus#401e753c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field boost_url: %w", err)
		}
		c.BoostURL = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field applied_slot_ids: %w", err)
		}

		if headerLen > 0 {
			c.AppliedSlotIDs = make([]int32, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field applied_slot_ids: %w", err)
			}
			c.AppliedSlotIDs = append(c.AppliedSlotIDs, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field level: %w", err)
		}
		c.Level = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field gift_code_boost_count: %w", err)
		}
		c.GiftCodeBoostCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field boost_count: %w", err)
		}
		c.BoostCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field current_level_boost_count: %w", err)
		}
		c.CurrentLevelBoostCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field next_level_boost_count: %w", err)
		}
		c.NextLevelBoostCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field premium_member_count: %w", err)
		}
		c.PremiumMemberCount = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field premium_member_percentage: %w", err)
		}
		c.PremiumMemberPercentage = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field prepaid_giveaways: %w", err)
		}

		if headerLen > 0 {
			c.PrepaidGiveaways = make([]PrepaidGiveaway, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PrepaidGiveaway
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatBoostStatus#401e753c: field prepaid_giveaways: %w", err)
			}
			c.PrepaidGiveaways = append(c.PrepaidGiveaways, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatBoostStatus) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostStatus#401e753c as nil")
	}
	b.ObjStart()
	b.PutID("chatBoostStatus")
	b.Comma()
	b.FieldStart("boost_url")
	b.PutString(c.BoostURL)
	b.Comma()
	b.FieldStart("applied_slot_ids")
	b.ArrStart()
	for _, v := range c.AppliedSlotIDs {
		b.PutInt32(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("level")
	b.PutInt32(c.Level)
	b.Comma()
	b.FieldStart("gift_code_boost_count")
	b.PutInt32(c.GiftCodeBoostCount)
	b.Comma()
	b.FieldStart("boost_count")
	b.PutInt32(c.BoostCount)
	b.Comma()
	b.FieldStart("current_level_boost_count")
	b.PutInt32(c.CurrentLevelBoostCount)
	b.Comma()
	b.FieldStart("next_level_boost_count")
	b.PutInt32(c.NextLevelBoostCount)
	b.Comma()
	b.FieldStart("premium_member_count")
	b.PutInt32(c.PremiumMemberCount)
	b.Comma()
	b.FieldStart("premium_member_percentage")
	b.PutDouble(c.PremiumMemberPercentage)
	b.Comma()
	b.FieldStart("prepaid_giveaways")
	b.ArrStart()
	for idx, v := range c.PrepaidGiveaways {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatBoostStatus#401e753c: field prepaid_giveaways element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatBoostStatus) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostStatus#401e753c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatBoostStatus"); err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: %w", err)
			}
		case "boost_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field boost_url: %w", err)
			}
			c.BoostURL = value
		case "applied_slot_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int32()
				if err != nil {
					return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field applied_slot_ids: %w", err)
				}
				c.AppliedSlotIDs = append(c.AppliedSlotIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field applied_slot_ids: %w", err)
			}
		case "level":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field level: %w", err)
			}
			c.Level = value
		case "gift_code_boost_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field gift_code_boost_count: %w", err)
			}
			c.GiftCodeBoostCount = value
		case "boost_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field boost_count: %w", err)
			}
			c.BoostCount = value
		case "current_level_boost_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field current_level_boost_count: %w", err)
			}
			c.CurrentLevelBoostCount = value
		case "next_level_boost_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field next_level_boost_count: %w", err)
			}
			c.NextLevelBoostCount = value
		case "premium_member_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field premium_member_count: %w", err)
			}
			c.PremiumMemberCount = value
		case "premium_member_percentage":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field premium_member_percentage: %w", err)
			}
			c.PremiumMemberPercentage = value
		case "prepaid_giveaways":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value PrepaidGiveaway
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field prepaid_giveaways: %w", err)
				}
				c.PrepaidGiveaways = append(c.PrepaidGiveaways, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode chatBoostStatus#401e753c: field prepaid_giveaways: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBoostURL returns value of BoostURL field.
func (c *ChatBoostStatus) GetBoostURL() (value string) {
	if c == nil {
		return
	}
	return c.BoostURL
}

// GetAppliedSlotIDs returns value of AppliedSlotIDs field.
func (c *ChatBoostStatus) GetAppliedSlotIDs() (value []int32) {
	if c == nil {
		return
	}
	return c.AppliedSlotIDs
}

// GetLevel returns value of Level field.
func (c *ChatBoostStatus) GetLevel() (value int32) {
	if c == nil {
		return
	}
	return c.Level
}

// GetGiftCodeBoostCount returns value of GiftCodeBoostCount field.
func (c *ChatBoostStatus) GetGiftCodeBoostCount() (value int32) {
	if c == nil {
		return
	}
	return c.GiftCodeBoostCount
}

// GetBoostCount returns value of BoostCount field.
func (c *ChatBoostStatus) GetBoostCount() (value int32) {
	if c == nil {
		return
	}
	return c.BoostCount
}

// GetCurrentLevelBoostCount returns value of CurrentLevelBoostCount field.
func (c *ChatBoostStatus) GetCurrentLevelBoostCount() (value int32) {
	if c == nil {
		return
	}
	return c.CurrentLevelBoostCount
}

// GetNextLevelBoostCount returns value of NextLevelBoostCount field.
func (c *ChatBoostStatus) GetNextLevelBoostCount() (value int32) {
	if c == nil {
		return
	}
	return c.NextLevelBoostCount
}

// GetPremiumMemberCount returns value of PremiumMemberCount field.
func (c *ChatBoostStatus) GetPremiumMemberCount() (value int32) {
	if c == nil {
		return
	}
	return c.PremiumMemberCount
}

// GetPremiumMemberPercentage returns value of PremiumMemberPercentage field.
func (c *ChatBoostStatus) GetPremiumMemberPercentage() (value float64) {
	if c == nil {
		return
	}
	return c.PremiumMemberPercentage
}

// GetPrepaidGiveaways returns value of PrepaidGiveaways field.
func (c *ChatBoostStatus) GetPrepaidGiveaways() (value []PrepaidGiveaway) {
	if c == nil {
		return
	}
	return c.PrepaidGiveaways
}
