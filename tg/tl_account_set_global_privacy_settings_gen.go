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

// AccountSetGlobalPrivacySettingsRequest represents TL type `account.setGlobalPrivacySettings#1edaaac2`.
// Set global privacy settings
//
// See https://core.telegram.org/method/account.setGlobalPrivacySettings for reference.
type AccountSetGlobalPrivacySettingsRequest struct {
	// Global privacy settings
	Settings GlobalPrivacySettings
}

// AccountSetGlobalPrivacySettingsRequestTypeID is TL type id of AccountSetGlobalPrivacySettingsRequest.
const AccountSetGlobalPrivacySettingsRequestTypeID = 0x1edaaac2

func (s *AccountSetGlobalPrivacySettingsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetGlobalPrivacySettingsRequest) String() string {
	if s == nil {
		return "AccountSetGlobalPrivacySettingsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountSetGlobalPrivacySettingsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tSettings: ")
	sb.WriteString(fmt.Sprint(s.Settings))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *AccountSetGlobalPrivacySettingsRequest) TypeID() uint32 {
	return AccountSetGlobalPrivacySettingsRequestTypeID
}

// Encode implements bin.Encoder.
func (s *AccountSetGlobalPrivacySettingsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setGlobalPrivacySettings#1edaaac2 as nil")
	}
	b.PutID(AccountSetGlobalPrivacySettingsRequestTypeID)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setGlobalPrivacySettings#1edaaac2: field settings: %w", err)
	}
	return nil
}

// GetSettings returns value of Settings field.
func (s *AccountSetGlobalPrivacySettingsRequest) GetSettings() (value GlobalPrivacySettings) {
	return s.Settings
}

// Decode implements bin.Decoder.
func (s *AccountSetGlobalPrivacySettingsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setGlobalPrivacySettings#1edaaac2 to nil")
	}
	if err := b.ConsumeID(AccountSetGlobalPrivacySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setGlobalPrivacySettings#1edaaac2: %w", err)
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.setGlobalPrivacySettings#1edaaac2: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetGlobalPrivacySettingsRequest.
var (
	_ bin.Encoder = &AccountSetGlobalPrivacySettingsRequest{}
	_ bin.Decoder = &AccountSetGlobalPrivacySettingsRequest{}
)

// AccountSetGlobalPrivacySettings invokes method account.setGlobalPrivacySettings#1edaaac2 returning error if any.
// Set global privacy settings
//
// See https://core.telegram.org/method/account.setGlobalPrivacySettings for reference.
// Can be used by bots.
func (c *Client) AccountSetGlobalPrivacySettings(ctx context.Context, settings GlobalPrivacySettings) (*GlobalPrivacySettings, error) {
	var result GlobalPrivacySettings

	request := &AccountSetGlobalPrivacySettingsRequest{
		Settings: settings,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
