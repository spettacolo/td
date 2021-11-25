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

// UsersSetSecureValueErrorsRequest represents TL type `users.setSecureValueErrors#90c894b5`.
// Notify the user that the sent passport¹ data contains some errors The user will not
// be able to re-submit their Passport data to you until the errors are fixed (the
// contents of the field for which you returned the error must change).
// Use this if the data submitted by the user doesn't satisfy the standards your service
// requires for any reason. For example, if a birthday date seems invalid, a submitted
// document is blurry, a scan shows evidence of tampering, etc. Supply some details in
// the error message to make sure the user knows how to correct the issues.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/users.setSecureValueErrors for reference.
type UsersSetSecureValueErrorsRequest struct {
	// The user
	ID InputUserClass
	// Errors
	Errors []SecureValueErrorClass
}

// UsersSetSecureValueErrorsRequestTypeID is TL type id of UsersSetSecureValueErrorsRequest.
const UsersSetSecureValueErrorsRequestTypeID = 0x90c894b5

// Ensuring interfaces in compile-time for UsersSetSecureValueErrorsRequest.
var (
	_ bin.Encoder     = &UsersSetSecureValueErrorsRequest{}
	_ bin.Decoder     = &UsersSetSecureValueErrorsRequest{}
	_ bin.BareEncoder = &UsersSetSecureValueErrorsRequest{}
	_ bin.BareDecoder = &UsersSetSecureValueErrorsRequest{}
)

func (s *UsersSetSecureValueErrorsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == nil) {
		return false
	}
	if !(s.Errors == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *UsersSetSecureValueErrorsRequest) String() string {
	if s == nil {
		return "UsersSetSecureValueErrorsRequest(nil)"
	}
	type Alias UsersSetSecureValueErrorsRequest
	return fmt.Sprintf("UsersSetSecureValueErrorsRequest%+v", Alias(*s))
}

// FillFrom fills UsersSetSecureValueErrorsRequest from given interface.
func (s *UsersSetSecureValueErrorsRequest) FillFrom(from interface {
	GetID() (value InputUserClass)
	GetErrors() (value []SecureValueErrorClass)
}) {
	s.ID = from.GetID()
	s.Errors = from.GetErrors()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UsersSetSecureValueErrorsRequest) TypeID() uint32 {
	return UsersSetSecureValueErrorsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UsersSetSecureValueErrorsRequest) TypeName() string {
	return "users.setSecureValueErrors"
}

// TypeInfo returns info about TL type.
func (s *UsersSetSecureValueErrorsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "users.setSecureValueErrors",
		ID:   UsersSetSecureValueErrorsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Errors",
			SchemaName: "errors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *UsersSetSecureValueErrorsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode users.setSecureValueErrors#90c894b5 as nil")
	}
	b.PutID(UsersSetSecureValueErrorsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *UsersSetSecureValueErrorsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode users.setSecureValueErrors#90c894b5 as nil")
	}
	if s.ID == nil {
		return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field id is nil")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field id: %w", err)
	}
	b.PutVectorHeader(len(s.Errors))
	for idx, v := range s.Errors {
		if v == nil {
			return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field errors element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.setSecureValueErrors#90c894b5: field errors element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *UsersSetSecureValueErrorsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode users.setSecureValueErrors#90c894b5 to nil")
	}
	if err := b.ConsumeID(UsersSetSecureValueErrorsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *UsersSetSecureValueErrorsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode users.setSecureValueErrors#90c894b5 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: field id: %w", err)
		}
		s.ID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: field errors: %w", err)
		}

		if headerLen > 0 {
			s.Errors = make([]SecureValueErrorClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureValueError(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.setSecureValueErrors#90c894b5: field errors: %w", err)
			}
			s.Errors = append(s.Errors, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (s *UsersSetSecureValueErrorsRequest) GetID() (value InputUserClass) {
	return s.ID
}

// GetErrors returns value of Errors field.
func (s *UsersSetSecureValueErrorsRequest) GetErrors() (value []SecureValueErrorClass) {
	return s.Errors
}

// MapErrors returns field Errors wrapped in SecureValueErrorClassArray helper.
func (s *UsersSetSecureValueErrorsRequest) MapErrors() (value SecureValueErrorClassArray) {
	return SecureValueErrorClassArray(s.Errors)
}

// UsersSetSecureValueErrors invokes method users.setSecureValueErrors#90c894b5 returning error if any.
// Notify the user that the sent passport¹ data contains some errors The user will not
// be able to re-submit their Passport data to you until the errors are fixed (the
// contents of the field for which you returned the error must change).
// Use this if the data submitted by the user doesn't satisfy the standards your service
// requires for any reason. For example, if a birthday date seems invalid, a submitted
// document is blurry, a scan shows evidence of tampering, etc. Supply some details in
// the error message to make sure the user knows how to correct the issues.
//
// Links:
//  1) https://core.telegram.org/passport
//
// Possible errors:
//  400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/users.setSecureValueErrors for reference.
// Can be used by bots.
func (c *Client) UsersSetSecureValueErrors(ctx context.Context, request *UsersSetSecureValueErrorsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
