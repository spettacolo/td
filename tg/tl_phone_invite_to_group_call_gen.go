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

// PhoneInviteToGroupCallRequest represents TL type `phone.inviteToGroupCall#7b393160`.
//
// See https://core.telegram.org/method/phone.inviteToGroupCall for reference.
type PhoneInviteToGroupCallRequest struct {
	// Call field of PhoneInviteToGroupCallRequest.
	Call InputGroupCall
	// Users field of PhoneInviteToGroupCallRequest.
	Users []InputUserClass
}

// PhoneInviteToGroupCallRequestTypeID is TL type id of PhoneInviteToGroupCallRequest.
const PhoneInviteToGroupCallRequestTypeID = 0x7b393160

func (i *PhoneInviteToGroupCallRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Call.Zero()) {
		return false
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *PhoneInviteToGroupCallRequest) String() string {
	if i == nil {
		return "PhoneInviteToGroupCallRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneInviteToGroupCallRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tCall: ")
	sb.WriteString(fmt.Sprint(i.Call))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range i.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *PhoneInviteToGroupCallRequest) TypeID() uint32 {
	return PhoneInviteToGroupCallRequestTypeID
}

// Encode implements bin.Encoder.
func (i *PhoneInviteToGroupCallRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode phone.inviteToGroupCall#7b393160 as nil")
	}
	b.PutID(PhoneInviteToGroupCallRequestTypeID)
	if err := i.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.inviteToGroupCall#7b393160: field call: %w", err)
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode phone.inviteToGroupCall#7b393160: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode phone.inviteToGroupCall#7b393160: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (i *PhoneInviteToGroupCallRequest) GetCall() (value InputGroupCall) {
	return i.Call
}

// GetUsers returns value of Users field.
func (i *PhoneInviteToGroupCallRequest) GetUsers() (value []InputUserClass) {
	return i.Users
}

// Decode implements bin.Decoder.
func (i *PhoneInviteToGroupCallRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode phone.inviteToGroupCall#7b393160 to nil")
	}
	if err := b.ConsumeID(PhoneInviteToGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: %w", err)
	}
	{
		if err := i.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: field call: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode phone.inviteToGroupCall#7b393160: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneInviteToGroupCallRequest.
var (
	_ bin.Encoder = &PhoneInviteToGroupCallRequest{}
	_ bin.Decoder = &PhoneInviteToGroupCallRequest{}
)

// PhoneInviteToGroupCall invokes method phone.inviteToGroupCall#7b393160 returning error if any.
//
// See https://core.telegram.org/method/phone.inviteToGroupCall for reference.
func (c *Client) PhoneInviteToGroupCall(ctx context.Context, request *PhoneInviteToGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
