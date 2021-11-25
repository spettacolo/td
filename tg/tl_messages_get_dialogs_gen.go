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

// MessagesGetDialogsRequest represents TL type `messages.getDialogs#a0f4cb4f`.
// Returns the current user dialog list.
//
// See https://core.telegram.org/method/messages.getDialogs for reference.
type MessagesGetDialogsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Exclude pinned dialogs
	ExcludePinned bool
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetDate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Offset peer for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetPeer InputPeerClass
	// Number of list elements to be returned
	Limit int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetDialogsRequestTypeID is TL type id of MessagesGetDialogsRequest.
const MessagesGetDialogsRequestTypeID = 0xa0f4cb4f

// Ensuring interfaces in compile-time for MessagesGetDialogsRequest.
var (
	_ bin.Encoder     = &MessagesGetDialogsRequest{}
	_ bin.Decoder     = &MessagesGetDialogsRequest{}
	_ bin.BareEncoder = &MessagesGetDialogsRequest{}
	_ bin.BareDecoder = &MessagesGetDialogsRequest{}
)

func (g *MessagesGetDialogsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.ExcludePinned == false) {
		return false
	}
	if !(g.FolderID == 0) {
		return false
	}
	if !(g.OffsetDate == 0) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.OffsetPeer == nil) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetDialogsRequest) String() string {
	if g == nil {
		return "MessagesGetDialogsRequest(nil)"
	}
	type Alias MessagesGetDialogsRequest
	return fmt.Sprintf("MessagesGetDialogsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetDialogsRequest from given interface.
func (g *MessagesGetDialogsRequest) FillFrom(from interface {
	GetExcludePinned() (value bool)
	GetFolderID() (value int, ok bool)
	GetOffsetDate() (value int)
	GetOffsetID() (value int)
	GetOffsetPeer() (value InputPeerClass)
	GetLimit() (value int)
	GetHash() (value int64)
}) {
	g.ExcludePinned = from.GetExcludePinned()
	if val, ok := from.GetFolderID(); ok {
		g.FolderID = val
	}

	g.OffsetDate = from.GetOffsetDate()
	g.OffsetID = from.GetOffsetID()
	g.OffsetPeer = from.GetOffsetPeer()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetDialogsRequest) TypeID() uint32 {
	return MessagesGetDialogsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetDialogsRequest) TypeName() string {
	return "messages.getDialogs"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetDialogsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getDialogs",
		ID:   MessagesGetDialogsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExcludePinned",
			SchemaName: "exclude_pinned",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "FolderID",
			SchemaName: "folder_id",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "OffsetID",
			SchemaName: "offset_id",
		},
		{
			Name:       "OffsetPeer",
			SchemaName: "offset_peer",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetDialogsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogs#a0f4cb4f as nil")
	}
	b.PutID(MessagesGetDialogsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetDialogsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getDialogs#a0f4cb4f as nil")
	}
	if !(g.ExcludePinned == false) {
		g.Flags.Set(0)
	}
	if !(g.FolderID == 0) {
		g.Flags.Set(1)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0f4cb4f: field flags: %w", err)
	}
	if g.Flags.Has(1) {
		b.PutInt(g.FolderID)
	}
	b.PutInt(g.OffsetDate)
	b.PutInt(g.OffsetID)
	if g.OffsetPeer == nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0f4cb4f: field offset_peer is nil")
	}
	if err := g.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getDialogs#a0f4cb4f: field offset_peer: %w", err)
	}
	b.PutInt(g.Limit)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetDialogsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogs#a0f4cb4f to nil")
	}
	if err := b.ConsumeID(MessagesGetDialogsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetDialogsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getDialogs#a0f4cb4f to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field flags: %w", err)
		}
	}
	g.ExcludePinned = g.Flags.Has(0)
	if g.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field folder_id: %w", err)
		}
		g.FolderID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field offset_date: %w", err)
		}
		g.OffsetDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field offset_peer: %w", err)
		}
		g.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getDialogs#a0f4cb4f: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// SetExcludePinned sets value of ExcludePinned conditional field.
func (g *MessagesGetDialogsRequest) SetExcludePinned(value bool) {
	if value {
		g.Flags.Set(0)
		g.ExcludePinned = true
	} else {
		g.Flags.Unset(0)
		g.ExcludePinned = false
	}
}

// GetExcludePinned returns value of ExcludePinned conditional field.
func (g *MessagesGetDialogsRequest) GetExcludePinned() (value bool) {
	return g.Flags.Has(0)
}

// SetFolderID sets value of FolderID conditional field.
func (g *MessagesGetDialogsRequest) SetFolderID(value int) {
	g.Flags.Set(1)
	g.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (g *MessagesGetDialogsRequest) GetFolderID() (value int, ok bool) {
	if !g.Flags.Has(1) {
		return value, false
	}
	return g.FolderID, true
}

// GetOffsetDate returns value of OffsetDate field.
func (g *MessagesGetDialogsRequest) GetOffsetDate() (value int) {
	return g.OffsetDate
}

// GetOffsetID returns value of OffsetID field.
func (g *MessagesGetDialogsRequest) GetOffsetID() (value int) {
	return g.OffsetID
}

// GetOffsetPeer returns value of OffsetPeer field.
func (g *MessagesGetDialogsRequest) GetOffsetPeer() (value InputPeerClass) {
	return g.OffsetPeer
}

// GetLimit returns value of Limit field.
func (g *MessagesGetDialogsRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *MessagesGetDialogsRequest) GetHash() (value int64) {
	return g.Hash
}

// MessagesGetDialogs invokes method messages.getDialogs#a0f4cb4f returning error if any.
// Returns the current user dialog list.
//
// Possible errors:
//  400 FOLDER_ID_INVALID: Invalid folder ID.
//  400 OFFSET_PEER_ID_INVALID: The provided offset peer is invalid.
//
// See https://core.telegram.org/method/messages.getDialogs for reference.
func (c *Client) MessagesGetDialogs(ctx context.Context, request *MessagesGetDialogsRequest) (MessagesDialogsClass, error) {
	var result MessagesDialogsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Dialogs, nil
}
