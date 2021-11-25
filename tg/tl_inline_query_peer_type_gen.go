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

// InlineQueryPeerTypeSameBotPM represents TL type `inlineQueryPeerTypeSameBotPM#3081ed9d`.
// The inline query was sent in a private chat with the bot itself
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeSameBotPM for reference.
type InlineQueryPeerTypeSameBotPM struct {
}

// InlineQueryPeerTypeSameBotPMTypeID is TL type id of InlineQueryPeerTypeSameBotPM.
const InlineQueryPeerTypeSameBotPMTypeID = 0x3081ed9d

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeSameBotPM) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeSameBotPM.
var (
	_ bin.Encoder     = &InlineQueryPeerTypeSameBotPM{}
	_ bin.Decoder     = &InlineQueryPeerTypeSameBotPM{}
	_ bin.BareEncoder = &InlineQueryPeerTypeSameBotPM{}
	_ bin.BareDecoder = &InlineQueryPeerTypeSameBotPM{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeSameBotPM{}
)

func (i *InlineQueryPeerTypeSameBotPM) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeSameBotPM) String() string {
	if i == nil {
		return "InlineQueryPeerTypeSameBotPM(nil)"
	}
	type Alias InlineQueryPeerTypeSameBotPM
	return fmt.Sprintf("InlineQueryPeerTypeSameBotPM%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryPeerTypeSameBotPM) TypeID() uint32 {
	return InlineQueryPeerTypeSameBotPMTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryPeerTypeSameBotPM) TypeName() string {
	return "inlineQueryPeerTypeSameBotPM"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryPeerTypeSameBotPM) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryPeerTypeSameBotPM",
		ID:   InlineQueryPeerTypeSameBotPMTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeSameBotPM) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeSameBotPM#3081ed9d as nil")
	}
	b.PutID(InlineQueryPeerTypeSameBotPMTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryPeerTypeSameBotPM) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeSameBotPM#3081ed9d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeSameBotPM) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeSameBotPM#3081ed9d to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeSameBotPMTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeSameBotPM#3081ed9d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryPeerTypeSameBotPM) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeSameBotPM#3081ed9d to nil")
	}
	return nil
}

// InlineQueryPeerTypePM represents TL type `inlineQueryPeerTypePM#833c0fac`.
// The inline query was sent in a private chat
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypePM for reference.
type InlineQueryPeerTypePM struct {
}

// InlineQueryPeerTypePMTypeID is TL type id of InlineQueryPeerTypePM.
const InlineQueryPeerTypePMTypeID = 0x833c0fac

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypePM) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypePM.
var (
	_ bin.Encoder     = &InlineQueryPeerTypePM{}
	_ bin.Decoder     = &InlineQueryPeerTypePM{}
	_ bin.BareEncoder = &InlineQueryPeerTypePM{}
	_ bin.BareDecoder = &InlineQueryPeerTypePM{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypePM{}
)

func (i *InlineQueryPeerTypePM) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypePM) String() string {
	if i == nil {
		return "InlineQueryPeerTypePM(nil)"
	}
	type Alias InlineQueryPeerTypePM
	return fmt.Sprintf("InlineQueryPeerTypePM%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryPeerTypePM) TypeID() uint32 {
	return InlineQueryPeerTypePMTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryPeerTypePM) TypeName() string {
	return "inlineQueryPeerTypePM"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryPeerTypePM) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryPeerTypePM",
		ID:   InlineQueryPeerTypePMTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypePM) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypePM#833c0fac as nil")
	}
	b.PutID(InlineQueryPeerTypePMTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryPeerTypePM) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypePM#833c0fac as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypePM) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypePM#833c0fac to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypePMTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypePM#833c0fac: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryPeerTypePM) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypePM#833c0fac to nil")
	}
	return nil
}

// InlineQueryPeerTypeChat represents TL type `inlineQueryPeerTypeChat#d766c50a`.
// The inline query was sent in a chat¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeChat for reference.
type InlineQueryPeerTypeChat struct {
}

// InlineQueryPeerTypeChatTypeID is TL type id of InlineQueryPeerTypeChat.
const InlineQueryPeerTypeChatTypeID = 0xd766c50a

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeChat) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeChat.
var (
	_ bin.Encoder     = &InlineQueryPeerTypeChat{}
	_ bin.Decoder     = &InlineQueryPeerTypeChat{}
	_ bin.BareEncoder = &InlineQueryPeerTypeChat{}
	_ bin.BareDecoder = &InlineQueryPeerTypeChat{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeChat{}
)

func (i *InlineQueryPeerTypeChat) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeChat) String() string {
	if i == nil {
		return "InlineQueryPeerTypeChat(nil)"
	}
	type Alias InlineQueryPeerTypeChat
	return fmt.Sprintf("InlineQueryPeerTypeChat%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryPeerTypeChat) TypeID() uint32 {
	return InlineQueryPeerTypeChatTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryPeerTypeChat) TypeName() string {
	return "inlineQueryPeerTypeChat"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryPeerTypeChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryPeerTypeChat",
		ID:   InlineQueryPeerTypeChatTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeChat#d766c50a as nil")
	}
	b.PutID(InlineQueryPeerTypeChatTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryPeerTypeChat) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeChat#d766c50a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeChat#d766c50a to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeChat#d766c50a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryPeerTypeChat) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeChat#d766c50a to nil")
	}
	return nil
}

// InlineQueryPeerTypeMegagroup represents TL type `inlineQueryPeerTypeMegagroup#5ec4be43`.
// The inline query was sent in a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeMegagroup for reference.
type InlineQueryPeerTypeMegagroup struct {
}

// InlineQueryPeerTypeMegagroupTypeID is TL type id of InlineQueryPeerTypeMegagroup.
const InlineQueryPeerTypeMegagroupTypeID = 0x5ec4be43

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeMegagroup) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeMegagroup.
var (
	_ bin.Encoder     = &InlineQueryPeerTypeMegagroup{}
	_ bin.Decoder     = &InlineQueryPeerTypeMegagroup{}
	_ bin.BareEncoder = &InlineQueryPeerTypeMegagroup{}
	_ bin.BareDecoder = &InlineQueryPeerTypeMegagroup{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeMegagroup{}
)

func (i *InlineQueryPeerTypeMegagroup) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeMegagroup) String() string {
	if i == nil {
		return "InlineQueryPeerTypeMegagroup(nil)"
	}
	type Alias InlineQueryPeerTypeMegagroup
	return fmt.Sprintf("InlineQueryPeerTypeMegagroup%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryPeerTypeMegagroup) TypeID() uint32 {
	return InlineQueryPeerTypeMegagroupTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryPeerTypeMegagroup) TypeName() string {
	return "inlineQueryPeerTypeMegagroup"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryPeerTypeMegagroup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryPeerTypeMegagroup",
		ID:   InlineQueryPeerTypeMegagroupTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeMegagroup) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeMegagroup#5ec4be43 as nil")
	}
	b.PutID(InlineQueryPeerTypeMegagroupTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryPeerTypeMegagroup) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeMegagroup#5ec4be43 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeMegagroup) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeMegagroup#5ec4be43 to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeMegagroupTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeMegagroup#5ec4be43: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryPeerTypeMegagroup) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeMegagroup#5ec4be43 to nil")
	}
	return nil
}

// InlineQueryPeerTypeBroadcast represents TL type `inlineQueryPeerTypeBroadcast#6334ee9a`.
// The inline query was sent in a channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeBroadcast for reference.
type InlineQueryPeerTypeBroadcast struct {
}

// InlineQueryPeerTypeBroadcastTypeID is TL type id of InlineQueryPeerTypeBroadcast.
const InlineQueryPeerTypeBroadcastTypeID = 0x6334ee9a

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeBroadcast) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeBroadcast.
var (
	_ bin.Encoder     = &InlineQueryPeerTypeBroadcast{}
	_ bin.Decoder     = &InlineQueryPeerTypeBroadcast{}
	_ bin.BareEncoder = &InlineQueryPeerTypeBroadcast{}
	_ bin.BareDecoder = &InlineQueryPeerTypeBroadcast{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeBroadcast{}
)

func (i *InlineQueryPeerTypeBroadcast) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeBroadcast) String() string {
	if i == nil {
		return "InlineQueryPeerTypeBroadcast(nil)"
	}
	type Alias InlineQueryPeerTypeBroadcast
	return fmt.Sprintf("InlineQueryPeerTypeBroadcast%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InlineQueryPeerTypeBroadcast) TypeID() uint32 {
	return InlineQueryPeerTypeBroadcastTypeID
}

// TypeName returns name of type in TL schema.
func (*InlineQueryPeerTypeBroadcast) TypeName() string {
	return "inlineQueryPeerTypeBroadcast"
}

// TypeInfo returns info about TL type.
func (i *InlineQueryPeerTypeBroadcast) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inlineQueryPeerTypeBroadcast",
		ID:   InlineQueryPeerTypeBroadcastTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeBroadcast) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeBroadcast#6334ee9a as nil")
	}
	b.PutID(InlineQueryPeerTypeBroadcastTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InlineQueryPeerTypeBroadcast) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeBroadcast#6334ee9a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeBroadcast) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeBroadcast#6334ee9a to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeBroadcastTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeBroadcast#6334ee9a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InlineQueryPeerTypeBroadcast) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeBroadcast#6334ee9a to nil")
	}
	return nil
}

// InlineQueryPeerTypeClass represents InlineQueryPeerType generic type.
//
// See https://core.telegram.org/type/InlineQueryPeerType for reference.
//
// Example:
//  g, err := tg.DecodeInlineQueryPeerType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InlineQueryPeerTypeSameBotPM: // inlineQueryPeerTypeSameBotPM#3081ed9d
//  case *tg.InlineQueryPeerTypePM: // inlineQueryPeerTypePM#833c0fac
//  case *tg.InlineQueryPeerTypeChat: // inlineQueryPeerTypeChat#d766c50a
//  case *tg.InlineQueryPeerTypeMegagroup: // inlineQueryPeerTypeMegagroup#5ec4be43
//  case *tg.InlineQueryPeerTypeBroadcast: // inlineQueryPeerTypeBroadcast#6334ee9a
//  default: panic(v)
//  }
type InlineQueryPeerTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InlineQueryPeerTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInlineQueryPeerType implements binary de-serialization for InlineQueryPeerTypeClass.
func DecodeInlineQueryPeerType(buf *bin.Buffer) (InlineQueryPeerTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InlineQueryPeerTypeSameBotPMTypeID:
		// Decoding inlineQueryPeerTypeSameBotPM#3081ed9d.
		v := InlineQueryPeerTypeSameBotPM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypePMTypeID:
		// Decoding inlineQueryPeerTypePM#833c0fac.
		v := InlineQueryPeerTypePM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypeChatTypeID:
		// Decoding inlineQueryPeerTypeChat#d766c50a.
		v := InlineQueryPeerTypeChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypeMegagroupTypeID:
		// Decoding inlineQueryPeerTypeMegagroup#5ec4be43.
		v := InlineQueryPeerTypeMegagroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypeBroadcastTypeID:
		// Decoding inlineQueryPeerTypeBroadcast#6334ee9a.
		v := InlineQueryPeerTypeBroadcast{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InlineQueryPeerType boxes the InlineQueryPeerTypeClass providing a helper.
type InlineQueryPeerTypeBox struct {
	InlineQueryPeerType InlineQueryPeerTypeClass
}

// Decode implements bin.Decoder for InlineQueryPeerTypeBox.
func (b *InlineQueryPeerTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InlineQueryPeerTypeBox to nil")
	}
	v, err := DecodeInlineQueryPeerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InlineQueryPeerType = v
	return nil
}

// Encode implements bin.Encode for InlineQueryPeerTypeBox.
func (b *InlineQueryPeerTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InlineQueryPeerType == nil {
		return fmt.Errorf("unable to encode InlineQueryPeerTypeClass as nil")
	}
	return b.InlineQueryPeerType.Encode(buf)
}
