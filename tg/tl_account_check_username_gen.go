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

// AccountCheckUsernameRequest represents TL type `account.checkUsername#2714d86c`.
// Validates a username and checks availability.
//
// See https://core.telegram.org/method/account.checkUsername for reference.
type AccountCheckUsernameRequest struct {
	// usernameAccepted characters: A-z (case-insensitive), 0-9 and underscores.Length: 5-32 characters.
	Username string
}

// AccountCheckUsernameRequestTypeID is TL type id of AccountCheckUsernameRequest.
const AccountCheckUsernameRequestTypeID = 0x2714d86c

func (c *AccountCheckUsernameRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountCheckUsernameRequest) String() string {
	if c == nil {
		return "AccountCheckUsernameRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountCheckUsernameRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tUsername: ")
	sb.WriteString(fmt.Sprint(c.Username))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *AccountCheckUsernameRequest) TypeID() uint32 {
	return AccountCheckUsernameRequestTypeID
}

// Encode implements bin.Encoder.
func (c *AccountCheckUsernameRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.checkUsername#2714d86c as nil")
	}
	b.PutID(AccountCheckUsernameRequestTypeID)
	b.PutString(c.Username)
	return nil
}

// GetUsername returns value of Username field.
func (c *AccountCheckUsernameRequest) GetUsername() (value string) {
	return c.Username
}

// Decode implements bin.Decoder.
func (c *AccountCheckUsernameRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.checkUsername#2714d86c to nil")
	}
	if err := b.ConsumeID(AccountCheckUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.checkUsername#2714d86c: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.checkUsername#2714d86c: field username: %w", err)
		}
		c.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountCheckUsernameRequest.
var (
	_ bin.Encoder = &AccountCheckUsernameRequest{}
	_ bin.Decoder = &AccountCheckUsernameRequest{}
)

// AccountCheckUsername invokes method account.checkUsername#2714d86c returning error if any.
// Validates a username and checks availability.
//
// Possible errors:
//  400 USERNAME_INVALID: Unacceptable username
//
// See https://core.telegram.org/method/account.checkUsername for reference.
func (c *Client) AccountCheckUsername(ctx context.Context, username string) (bool, error) {
	var result BoolBox

	request := &AccountCheckUsernameRequest{
		Username: username,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
