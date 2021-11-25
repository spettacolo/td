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

// AccountResetPasswordFailedWait represents TL type `account.resetPasswordFailedWait#e3779861`.
// You recently requested a password reset that was canceled, please wait until the
// specified date before requesting another reset.
//
// See https://core.telegram.org/constructor/account.resetPasswordFailedWait for reference.
type AccountResetPasswordFailedWait struct {
	// Wait until this date before requesting another reset.
	RetryDate int
}

// AccountResetPasswordFailedWaitTypeID is TL type id of AccountResetPasswordFailedWait.
const AccountResetPasswordFailedWaitTypeID = 0xe3779861

// construct implements constructor of AccountResetPasswordResultClass.
func (r AccountResetPasswordFailedWait) construct() AccountResetPasswordResultClass { return &r }

// Ensuring interfaces in compile-time for AccountResetPasswordFailedWait.
var (
	_ bin.Encoder     = &AccountResetPasswordFailedWait{}
	_ bin.Decoder     = &AccountResetPasswordFailedWait{}
	_ bin.BareEncoder = &AccountResetPasswordFailedWait{}
	_ bin.BareDecoder = &AccountResetPasswordFailedWait{}

	_ AccountResetPasswordResultClass = &AccountResetPasswordFailedWait{}
)

func (r *AccountResetPasswordFailedWait) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.RetryDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetPasswordFailedWait) String() string {
	if r == nil {
		return "AccountResetPasswordFailedWait(nil)"
	}
	type Alias AccountResetPasswordFailedWait
	return fmt.Sprintf("AccountResetPasswordFailedWait%+v", Alias(*r))
}

// FillFrom fills AccountResetPasswordFailedWait from given interface.
func (r *AccountResetPasswordFailedWait) FillFrom(from interface {
	GetRetryDate() (value int)
}) {
	r.RetryDate = from.GetRetryDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetPasswordFailedWait) TypeID() uint32 {
	return AccountResetPasswordFailedWaitTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetPasswordFailedWait) TypeName() string {
	return "account.resetPasswordFailedWait"
}

// TypeInfo returns info about TL type.
func (r *AccountResetPasswordFailedWait) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetPasswordFailedWait",
		ID:   AccountResetPasswordFailedWaitTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RetryDate",
			SchemaName: "retry_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetPasswordFailedWait) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPasswordFailedWait#e3779861 as nil")
	}
	b.PutID(AccountResetPasswordFailedWaitTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetPasswordFailedWait) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPasswordFailedWait#e3779861 as nil")
	}
	b.PutInt(r.RetryDate)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetPasswordFailedWait) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPasswordFailedWait#e3779861 to nil")
	}
	if err := b.ConsumeID(AccountResetPasswordFailedWaitTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetPasswordFailedWait#e3779861: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetPasswordFailedWait) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPasswordFailedWait#e3779861 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.resetPasswordFailedWait#e3779861: field retry_date: %w", err)
		}
		r.RetryDate = value
	}
	return nil
}

// GetRetryDate returns value of RetryDate field.
func (r *AccountResetPasswordFailedWait) GetRetryDate() (value int) {
	return r.RetryDate
}

// AccountResetPasswordRequestedWait represents TL type `account.resetPasswordRequestedWait#e9effc7d`.
// You successfully requested a password reset, please wait until the specified date
// before finalizing the reset.
//
// See https://core.telegram.org/constructor/account.resetPasswordRequestedWait for reference.
type AccountResetPasswordRequestedWait struct {
	// Wait until this date before finalizing the reset.
	UntilDate int
}

// AccountResetPasswordRequestedWaitTypeID is TL type id of AccountResetPasswordRequestedWait.
const AccountResetPasswordRequestedWaitTypeID = 0xe9effc7d

// construct implements constructor of AccountResetPasswordResultClass.
func (r AccountResetPasswordRequestedWait) construct() AccountResetPasswordResultClass { return &r }

// Ensuring interfaces in compile-time for AccountResetPasswordRequestedWait.
var (
	_ bin.Encoder     = &AccountResetPasswordRequestedWait{}
	_ bin.Decoder     = &AccountResetPasswordRequestedWait{}
	_ bin.BareEncoder = &AccountResetPasswordRequestedWait{}
	_ bin.BareDecoder = &AccountResetPasswordRequestedWait{}

	_ AccountResetPasswordResultClass = &AccountResetPasswordRequestedWait{}
)

func (r *AccountResetPasswordRequestedWait) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.UntilDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetPasswordRequestedWait) String() string {
	if r == nil {
		return "AccountResetPasswordRequestedWait(nil)"
	}
	type Alias AccountResetPasswordRequestedWait
	return fmt.Sprintf("AccountResetPasswordRequestedWait%+v", Alias(*r))
}

// FillFrom fills AccountResetPasswordRequestedWait from given interface.
func (r *AccountResetPasswordRequestedWait) FillFrom(from interface {
	GetUntilDate() (value int)
}) {
	r.UntilDate = from.GetUntilDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetPasswordRequestedWait) TypeID() uint32 {
	return AccountResetPasswordRequestedWaitTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetPasswordRequestedWait) TypeName() string {
	return "account.resetPasswordRequestedWait"
}

// TypeInfo returns info about TL type.
func (r *AccountResetPasswordRequestedWait) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetPasswordRequestedWait",
		ID:   AccountResetPasswordRequestedWaitTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UntilDate",
			SchemaName: "until_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetPasswordRequestedWait) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPasswordRequestedWait#e9effc7d as nil")
	}
	b.PutID(AccountResetPasswordRequestedWaitTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetPasswordRequestedWait) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPasswordRequestedWait#e9effc7d as nil")
	}
	b.PutInt(r.UntilDate)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetPasswordRequestedWait) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPasswordRequestedWait#e9effc7d to nil")
	}
	if err := b.ConsumeID(AccountResetPasswordRequestedWaitTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetPasswordRequestedWait#e9effc7d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetPasswordRequestedWait) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPasswordRequestedWait#e9effc7d to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.resetPasswordRequestedWait#e9effc7d: field until_date: %w", err)
		}
		r.UntilDate = value
	}
	return nil
}

// GetUntilDate returns value of UntilDate field.
func (r *AccountResetPasswordRequestedWait) GetUntilDate() (value int) {
	return r.UntilDate
}

// AccountResetPasswordOk represents TL type `account.resetPasswordOk#e926d63e`.
// The 2FA password was reset successfully.
//
// See https://core.telegram.org/constructor/account.resetPasswordOk for reference.
type AccountResetPasswordOk struct {
}

// AccountResetPasswordOkTypeID is TL type id of AccountResetPasswordOk.
const AccountResetPasswordOkTypeID = 0xe926d63e

// construct implements constructor of AccountResetPasswordResultClass.
func (r AccountResetPasswordOk) construct() AccountResetPasswordResultClass { return &r }

// Ensuring interfaces in compile-time for AccountResetPasswordOk.
var (
	_ bin.Encoder     = &AccountResetPasswordOk{}
	_ bin.Decoder     = &AccountResetPasswordOk{}
	_ bin.BareEncoder = &AccountResetPasswordOk{}
	_ bin.BareDecoder = &AccountResetPasswordOk{}

	_ AccountResetPasswordResultClass = &AccountResetPasswordOk{}
)

func (r *AccountResetPasswordOk) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetPasswordOk) String() string {
	if r == nil {
		return "AccountResetPasswordOk(nil)"
	}
	type Alias AccountResetPasswordOk
	return fmt.Sprintf("AccountResetPasswordOk%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetPasswordOk) TypeID() uint32 {
	return AccountResetPasswordOkTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetPasswordOk) TypeName() string {
	return "account.resetPasswordOk"
}

// TypeInfo returns info about TL type.
func (r *AccountResetPasswordOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetPasswordOk",
		ID:   AccountResetPasswordOkTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetPasswordOk) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPasswordOk#e926d63e as nil")
	}
	b.PutID(AccountResetPasswordOkTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetPasswordOk) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetPasswordOk#e926d63e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetPasswordOk) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPasswordOk#e926d63e to nil")
	}
	if err := b.ConsumeID(AccountResetPasswordOkTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetPasswordOk#e926d63e: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetPasswordOk) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetPasswordOk#e926d63e to nil")
	}
	return nil
}

// AccountResetPasswordResultClass represents account.ResetPasswordResult generic type.
//
// See https://core.telegram.org/type/account.ResetPasswordResult for reference.
//
// Example:
//  g, err := tg.DecodeAccountResetPasswordResult(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AccountResetPasswordFailedWait: // account.resetPasswordFailedWait#e3779861
//  case *tg.AccountResetPasswordRequestedWait: // account.resetPasswordRequestedWait#e9effc7d
//  case *tg.AccountResetPasswordOk: // account.resetPasswordOk#e926d63e
//  default: panic(v)
//  }
type AccountResetPasswordResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AccountResetPasswordResultClass

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

// DecodeAccountResetPasswordResult implements binary de-serialization for AccountResetPasswordResultClass.
func DecodeAccountResetPasswordResult(buf *bin.Buffer) (AccountResetPasswordResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AccountResetPasswordFailedWaitTypeID:
		// Decoding account.resetPasswordFailedWait#e3779861.
		v := AccountResetPasswordFailedWait{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountResetPasswordResultClass: %w", err)
		}
		return &v, nil
	case AccountResetPasswordRequestedWaitTypeID:
		// Decoding account.resetPasswordRequestedWait#e9effc7d.
		v := AccountResetPasswordRequestedWait{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountResetPasswordResultClass: %w", err)
		}
		return &v, nil
	case AccountResetPasswordOkTypeID:
		// Decoding account.resetPasswordOk#e926d63e.
		v := AccountResetPasswordOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AccountResetPasswordResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AccountResetPasswordResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// AccountResetPasswordResult boxes the AccountResetPasswordResultClass providing a helper.
type AccountResetPasswordResultBox struct {
	ResetPasswordResult AccountResetPasswordResultClass
}

// Decode implements bin.Decoder for AccountResetPasswordResultBox.
func (b *AccountResetPasswordResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AccountResetPasswordResultBox to nil")
	}
	v, err := DecodeAccountResetPasswordResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ResetPasswordResult = v
	return nil
}

// Encode implements bin.Encode for AccountResetPasswordResultBox.
func (b *AccountResetPasswordResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ResetPasswordResult == nil {
		return fmt.Errorf("unable to encode AccountResetPasswordResultClass as nil")
	}
	return b.ResetPasswordResult.Encode(buf)
}
