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

// ContactsGetTopPeersRequest represents TL type `contacts.getTopPeers#973478b6`.
// Get most used peers
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
type ContactsGetTopPeersRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Users we've chatted most frequently with
	Correspondents bool
	// Most used bots
	BotsPm bool
	// Most used inline bots
	BotsInline bool
	// Most frequently called users
	PhoneCalls bool
	// Users to which the users often forwards messages to
	ForwardUsers bool
	// Chats to which the users often forwards messages to
	ForwardChats bool
	// Often-opened groups and supergroups
	Groups bool
	// Most frequently visited channels
	Channels bool
	// Offset for pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// ContactsGetTopPeersRequestTypeID is TL type id of ContactsGetTopPeersRequest.
const ContactsGetTopPeersRequestTypeID = 0x973478b6

// Ensuring interfaces in compile-time for ContactsGetTopPeersRequest.
var (
	_ bin.Encoder     = &ContactsGetTopPeersRequest{}
	_ bin.Decoder     = &ContactsGetTopPeersRequest{}
	_ bin.BareEncoder = &ContactsGetTopPeersRequest{}
	_ bin.BareDecoder = &ContactsGetTopPeersRequest{}
)

func (g *ContactsGetTopPeersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Correspondents == false) {
		return false
	}
	if !(g.BotsPm == false) {
		return false
	}
	if !(g.BotsInline == false) {
		return false
	}
	if !(g.PhoneCalls == false) {
		return false
	}
	if !(g.ForwardUsers == false) {
		return false
	}
	if !(g.ForwardChats == false) {
		return false
	}
	if !(g.Groups == false) {
		return false
	}
	if !(g.Channels == false) {
		return false
	}
	if !(g.Offset == 0) {
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
func (g *ContactsGetTopPeersRequest) String() string {
	if g == nil {
		return "ContactsGetTopPeersRequest(nil)"
	}
	type Alias ContactsGetTopPeersRequest
	return fmt.Sprintf("ContactsGetTopPeersRequest%+v", Alias(*g))
}

// FillFrom fills ContactsGetTopPeersRequest from given interface.
func (g *ContactsGetTopPeersRequest) FillFrom(from interface {
	GetCorrespondents() (value bool)
	GetBotsPm() (value bool)
	GetBotsInline() (value bool)
	GetPhoneCalls() (value bool)
	GetForwardUsers() (value bool)
	GetForwardChats() (value bool)
	GetGroups() (value bool)
	GetChannels() (value bool)
	GetOffset() (value int)
	GetLimit() (value int)
	GetHash() (value int64)
}) {
	g.Correspondents = from.GetCorrespondents()
	g.BotsPm = from.GetBotsPm()
	g.BotsInline = from.GetBotsInline()
	g.PhoneCalls = from.GetPhoneCalls()
	g.ForwardUsers = from.GetForwardUsers()
	g.ForwardChats = from.GetForwardChats()
	g.Groups = from.GetGroups()
	g.Channels = from.GetChannels()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsGetTopPeersRequest) TypeID() uint32 {
	return ContactsGetTopPeersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsGetTopPeersRequest) TypeName() string {
	return "contacts.getTopPeers"
}

// TypeInfo returns info about TL type.
func (g *ContactsGetTopPeersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.getTopPeers",
		ID:   ContactsGetTopPeersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Correspondents",
			SchemaName: "correspondents",
			Null:       !g.Flags.Has(0),
		},
		{
			Name:       "BotsPm",
			SchemaName: "bots_pm",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "BotsInline",
			SchemaName: "bots_inline",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "PhoneCalls",
			SchemaName: "phone_calls",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "ForwardUsers",
			SchemaName: "forward_users",
			Null:       !g.Flags.Has(4),
		},
		{
			Name:       "ForwardChats",
			SchemaName: "forward_chats",
			Null:       !g.Flags.Has(5),
		},
		{
			Name:       "Groups",
			SchemaName: "groups",
			Null:       !g.Flags.Has(10),
		},
		{
			Name:       "Channels",
			SchemaName: "channels",
			Null:       !g.Flags.Has(15),
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
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
func (g *ContactsGetTopPeersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getTopPeers#973478b6 as nil")
	}
	b.PutID(ContactsGetTopPeersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ContactsGetTopPeersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getTopPeers#973478b6 as nil")
	}
	if !(g.Correspondents == false) {
		g.Flags.Set(0)
	}
	if !(g.BotsPm == false) {
		g.Flags.Set(1)
	}
	if !(g.BotsInline == false) {
		g.Flags.Set(2)
	}
	if !(g.PhoneCalls == false) {
		g.Flags.Set(3)
	}
	if !(g.ForwardUsers == false) {
		g.Flags.Set(4)
	}
	if !(g.ForwardChats == false) {
		g.Flags.Set(5)
	}
	if !(g.Groups == false) {
		g.Flags.Set(10)
	}
	if !(g.Channels == false) {
		g.Flags.Set(15)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getTopPeers#973478b6: field flags: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetTopPeersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getTopPeers#973478b6 to nil")
	}
	if err := b.ConsumeID(ContactsGetTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getTopPeers#973478b6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ContactsGetTopPeersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getTopPeers#973478b6 to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#973478b6: field flags: %w", err)
		}
	}
	g.Correspondents = g.Flags.Has(0)
	g.BotsPm = g.Flags.Has(1)
	g.BotsInline = g.Flags.Has(2)
	g.PhoneCalls = g.Flags.Has(3)
	g.ForwardUsers = g.Flags.Has(4)
	g.ForwardChats = g.Flags.Has(5)
	g.Groups = g.Flags.Has(10)
	g.Channels = g.Flags.Has(15)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#973478b6: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#973478b6: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getTopPeers#973478b6: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// SetCorrespondents sets value of Correspondents conditional field.
func (g *ContactsGetTopPeersRequest) SetCorrespondents(value bool) {
	if value {
		g.Flags.Set(0)
		g.Correspondents = true
	} else {
		g.Flags.Unset(0)
		g.Correspondents = false
	}
}

// GetCorrespondents returns value of Correspondents conditional field.
func (g *ContactsGetTopPeersRequest) GetCorrespondents() (value bool) {
	return g.Flags.Has(0)
}

// SetBotsPm sets value of BotsPm conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsPm(value bool) {
	if value {
		g.Flags.Set(1)
		g.BotsPm = true
	} else {
		g.Flags.Unset(1)
		g.BotsPm = false
	}
}

// GetBotsPm returns value of BotsPm conditional field.
func (g *ContactsGetTopPeersRequest) GetBotsPm() (value bool) {
	return g.Flags.Has(1)
}

// SetBotsInline sets value of BotsInline conditional field.
func (g *ContactsGetTopPeersRequest) SetBotsInline(value bool) {
	if value {
		g.Flags.Set(2)
		g.BotsInline = true
	} else {
		g.Flags.Unset(2)
		g.BotsInline = false
	}
}

// GetBotsInline returns value of BotsInline conditional field.
func (g *ContactsGetTopPeersRequest) GetBotsInline() (value bool) {
	return g.Flags.Has(2)
}

// SetPhoneCalls sets value of PhoneCalls conditional field.
func (g *ContactsGetTopPeersRequest) SetPhoneCalls(value bool) {
	if value {
		g.Flags.Set(3)
		g.PhoneCalls = true
	} else {
		g.Flags.Unset(3)
		g.PhoneCalls = false
	}
}

// GetPhoneCalls returns value of PhoneCalls conditional field.
func (g *ContactsGetTopPeersRequest) GetPhoneCalls() (value bool) {
	return g.Flags.Has(3)
}

// SetForwardUsers sets value of ForwardUsers conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardUsers(value bool) {
	if value {
		g.Flags.Set(4)
		g.ForwardUsers = true
	} else {
		g.Flags.Unset(4)
		g.ForwardUsers = false
	}
}

// GetForwardUsers returns value of ForwardUsers conditional field.
func (g *ContactsGetTopPeersRequest) GetForwardUsers() (value bool) {
	return g.Flags.Has(4)
}

// SetForwardChats sets value of ForwardChats conditional field.
func (g *ContactsGetTopPeersRequest) SetForwardChats(value bool) {
	if value {
		g.Flags.Set(5)
		g.ForwardChats = true
	} else {
		g.Flags.Unset(5)
		g.ForwardChats = false
	}
}

// GetForwardChats returns value of ForwardChats conditional field.
func (g *ContactsGetTopPeersRequest) GetForwardChats() (value bool) {
	return g.Flags.Has(5)
}

// SetGroups sets value of Groups conditional field.
func (g *ContactsGetTopPeersRequest) SetGroups(value bool) {
	if value {
		g.Flags.Set(10)
		g.Groups = true
	} else {
		g.Flags.Unset(10)
		g.Groups = false
	}
}

// GetGroups returns value of Groups conditional field.
func (g *ContactsGetTopPeersRequest) GetGroups() (value bool) {
	return g.Flags.Has(10)
}

// SetChannels sets value of Channels conditional field.
func (g *ContactsGetTopPeersRequest) SetChannels(value bool) {
	if value {
		g.Flags.Set(15)
		g.Channels = true
	} else {
		g.Flags.Unset(15)
		g.Channels = false
	}
}

// GetChannels returns value of Channels conditional field.
func (g *ContactsGetTopPeersRequest) GetChannels() (value bool) {
	return g.Flags.Has(15)
}

// GetOffset returns value of Offset field.
func (g *ContactsGetTopPeersRequest) GetOffset() (value int) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *ContactsGetTopPeersRequest) GetLimit() (value int) {
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *ContactsGetTopPeersRequest) GetHash() (value int64) {
	return g.Hash
}

// ContactsGetTopPeers invokes method contacts.getTopPeers#973478b6 returning error if any.
// Get most used peers
//
// Possible errors:
//  400 TYPES_EMPTY: No top peer type was provided.
//
// See https://core.telegram.org/method/contacts.getTopPeers for reference.
func (c *Client) ContactsGetTopPeers(ctx context.Context, request *ContactsGetTopPeersRequest) (ContactsTopPeersClass, error) {
	var result ContactsTopPeersBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.TopPeers, nil
}
