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

// AccountGetThemeRequest represents TL type `account.getTheme#8d9d742b`.
// Get theme information
//
// See https://core.telegram.org/method/account.getTheme for reference.
type AccountGetThemeRequest struct {
	// Theme format, a string that identifies the theming engines supported by the client
	Format string
	// Theme
	Theme InputThemeClass
	// Document ID
	DocumentID int64
}

// AccountGetThemeRequestTypeID is TL type id of AccountGetThemeRequest.
const AccountGetThemeRequestTypeID = 0x8d9d742b

// Ensuring interfaces in compile-time for AccountGetThemeRequest.
var (
	_ bin.Encoder     = &AccountGetThemeRequest{}
	_ bin.Decoder     = &AccountGetThemeRequest{}
	_ bin.BareEncoder = &AccountGetThemeRequest{}
	_ bin.BareDecoder = &AccountGetThemeRequest{}
)

func (g *AccountGetThemeRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Format == "") {
		return false
	}
	if !(g.Theme == nil) {
		return false
	}
	if !(g.DocumentID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetThemeRequest) String() string {
	if g == nil {
		return "AccountGetThemeRequest(nil)"
	}
	type Alias AccountGetThemeRequest
	return fmt.Sprintf("AccountGetThemeRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetThemeRequest from given interface.
func (g *AccountGetThemeRequest) FillFrom(from interface {
	GetFormat() (value string)
	GetTheme() (value InputThemeClass)
	GetDocumentID() (value int64)
}) {
	g.Format = from.GetFormat()
	g.Theme = from.GetTheme()
	g.DocumentID = from.GetDocumentID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetThemeRequest) TypeID() uint32 {
	return AccountGetThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetThemeRequest) TypeName() string {
	return "account.getTheme"
}

// TypeInfo returns info about TL type.
func (g *AccountGetThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getTheme",
		ID:   AccountGetThemeRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Format",
			SchemaName: "format",
		},
		{
			Name:       "Theme",
			SchemaName: "theme",
		},
		{
			Name:       "DocumentID",
			SchemaName: "document_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetThemeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getTheme#8d9d742b as nil")
	}
	b.PutID(AccountGetThemeRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetThemeRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getTheme#8d9d742b as nil")
	}
	b.PutString(g.Format)
	if g.Theme == nil {
		return fmt.Errorf("unable to encode account.getTheme#8d9d742b: field theme is nil")
	}
	if err := g.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getTheme#8d9d742b: field theme: %w", err)
	}
	b.PutLong(g.DocumentID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetThemeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getTheme#8d9d742b to nil")
	}
	if err := b.ConsumeID(AccountGetThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getTheme#8d9d742b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetThemeRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getTheme#8d9d742b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getTheme#8d9d742b: field format: %w", err)
		}
		g.Format = value
	}
	{
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getTheme#8d9d742b: field theme: %w", err)
		}
		g.Theme = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.getTheme#8d9d742b: field document_id: %w", err)
		}
		g.DocumentID = value
	}
	return nil
}

// GetFormat returns value of Format field.
func (g *AccountGetThemeRequest) GetFormat() (value string) {
	return g.Format
}

// GetTheme returns value of Theme field.
func (g *AccountGetThemeRequest) GetTheme() (value InputThemeClass) {
	return g.Theme
}

// GetDocumentID returns value of DocumentID field.
func (g *AccountGetThemeRequest) GetDocumentID() (value int64) {
	return g.DocumentID
}

// AccountGetTheme invokes method account.getTheme#8d9d742b returning error if any.
// Get theme information
//
// Possible errors:
//  400 THEME_FORMAT_INVALID: Invalid theme format provided.
//  400 THEME_INVALID: Invalid theme provided.
//
// See https://core.telegram.org/method/account.getTheme for reference.
func (c *Client) AccountGetTheme(ctx context.Context, request *AccountGetThemeRequest) (*Theme, error) {
	var result Theme

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
