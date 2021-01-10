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

// ChatAdminRights represents TL type `chatAdminRights#5fb224d5`.
// Represents the rights of an admin in a channel/supergroup¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/chatAdminRights for reference.
type ChatAdminRights struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, allows the admin to modify the description of the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	ChangeInfo bool
	// If set, allows the admin to post messages in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PostMessages bool
	// If set, allows the admin to also edit messages from other admins in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	EditMessages bool
	// If set, allows the admin to also delete messages from other admins in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	DeleteMessages bool
	// If set, allows the admin to ban users from the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	BanUsers bool
	// If set, allows the admin to invite users in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	InviteUsers bool
	// If set, allows the admin to pin messages in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PinMessages bool
	// If set, allows the admin to add other admins with the same (or more limited) permissions in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	AddAdmins bool
	// Whether this admin is anonymous
	Anonymous bool
	// ManageCall field of ChatAdminRights.
	ManageCall bool
}

// ChatAdminRightsTypeID is TL type id of ChatAdminRights.
const ChatAdminRightsTypeID = 0x5fb224d5

func (c *ChatAdminRights) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ChangeInfo == false) {
		return false
	}
	if !(c.PostMessages == false) {
		return false
	}
	if !(c.EditMessages == false) {
		return false
	}
	if !(c.DeleteMessages == false) {
		return false
	}
	if !(c.BanUsers == false) {
		return false
	}
	if !(c.InviteUsers == false) {
		return false
	}
	if !(c.PinMessages == false) {
		return false
	}
	if !(c.AddAdmins == false) {
		return false
	}
	if !(c.Anonymous == false) {
		return false
	}
	if !(c.ManageCall == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAdminRights) String() string {
	if c == nil {
		return "ChatAdminRights(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatAdminRights")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChatAdminRights) TypeID() uint32 {
	return ChatAdminRightsTypeID
}

// Encode implements bin.Encoder.
func (c *ChatAdminRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminRights#5fb224d5 as nil")
	}
	b.PutID(ChatAdminRightsTypeID)
	if !(c.ChangeInfo == false) {
		c.Flags.Set(0)
	}
	if !(c.PostMessages == false) {
		c.Flags.Set(1)
	}
	if !(c.EditMessages == false) {
		c.Flags.Set(2)
	}
	if !(c.DeleteMessages == false) {
		c.Flags.Set(3)
	}
	if !(c.BanUsers == false) {
		c.Flags.Set(4)
	}
	if !(c.InviteUsers == false) {
		c.Flags.Set(5)
	}
	if !(c.PinMessages == false) {
		c.Flags.Set(7)
	}
	if !(c.AddAdmins == false) {
		c.Flags.Set(9)
	}
	if !(c.Anonymous == false) {
		c.Flags.Set(10)
	}
	if !(c.ManageCall == false) {
		c.Flags.Set(11)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatAdminRights#5fb224d5: field flags: %w", err)
	}
	return nil
}

// SetChangeInfo sets value of ChangeInfo conditional field.
func (c *ChatAdminRights) SetChangeInfo(value bool) {
	if value {
		c.Flags.Set(0)
		c.ChangeInfo = true
	} else {
		c.Flags.Unset(0)
		c.ChangeInfo = false
	}
}

// GetChangeInfo returns value of ChangeInfo conditional field.
func (c *ChatAdminRights) GetChangeInfo() (value bool) {
	return c.Flags.Has(0)
}

// SetPostMessages sets value of PostMessages conditional field.
func (c *ChatAdminRights) SetPostMessages(value bool) {
	if value {
		c.Flags.Set(1)
		c.PostMessages = true
	} else {
		c.Flags.Unset(1)
		c.PostMessages = false
	}
}

// GetPostMessages returns value of PostMessages conditional field.
func (c *ChatAdminRights) GetPostMessages() (value bool) {
	return c.Flags.Has(1)
}

// SetEditMessages sets value of EditMessages conditional field.
func (c *ChatAdminRights) SetEditMessages(value bool) {
	if value {
		c.Flags.Set(2)
		c.EditMessages = true
	} else {
		c.Flags.Unset(2)
		c.EditMessages = false
	}
}

// GetEditMessages returns value of EditMessages conditional field.
func (c *ChatAdminRights) GetEditMessages() (value bool) {
	return c.Flags.Has(2)
}

// SetDeleteMessages sets value of DeleteMessages conditional field.
func (c *ChatAdminRights) SetDeleteMessages(value bool) {
	if value {
		c.Flags.Set(3)
		c.DeleteMessages = true
	} else {
		c.Flags.Unset(3)
		c.DeleteMessages = false
	}
}

// GetDeleteMessages returns value of DeleteMessages conditional field.
func (c *ChatAdminRights) GetDeleteMessages() (value bool) {
	return c.Flags.Has(3)
}

// SetBanUsers sets value of BanUsers conditional field.
func (c *ChatAdminRights) SetBanUsers(value bool) {
	if value {
		c.Flags.Set(4)
		c.BanUsers = true
	} else {
		c.Flags.Unset(4)
		c.BanUsers = false
	}
}

// GetBanUsers returns value of BanUsers conditional field.
func (c *ChatAdminRights) GetBanUsers() (value bool) {
	return c.Flags.Has(4)
}

// SetInviteUsers sets value of InviteUsers conditional field.
func (c *ChatAdminRights) SetInviteUsers(value bool) {
	if value {
		c.Flags.Set(5)
		c.InviteUsers = true
	} else {
		c.Flags.Unset(5)
		c.InviteUsers = false
	}
}

// GetInviteUsers returns value of InviteUsers conditional field.
func (c *ChatAdminRights) GetInviteUsers() (value bool) {
	return c.Flags.Has(5)
}

// SetPinMessages sets value of PinMessages conditional field.
func (c *ChatAdminRights) SetPinMessages(value bool) {
	if value {
		c.Flags.Set(7)
		c.PinMessages = true
	} else {
		c.Flags.Unset(7)
		c.PinMessages = false
	}
}

// GetPinMessages returns value of PinMessages conditional field.
func (c *ChatAdminRights) GetPinMessages() (value bool) {
	return c.Flags.Has(7)
}

// SetAddAdmins sets value of AddAdmins conditional field.
func (c *ChatAdminRights) SetAddAdmins(value bool) {
	if value {
		c.Flags.Set(9)
		c.AddAdmins = true
	} else {
		c.Flags.Unset(9)
		c.AddAdmins = false
	}
}

// GetAddAdmins returns value of AddAdmins conditional field.
func (c *ChatAdminRights) GetAddAdmins() (value bool) {
	return c.Flags.Has(9)
}

// SetAnonymous sets value of Anonymous conditional field.
func (c *ChatAdminRights) SetAnonymous(value bool) {
	if value {
		c.Flags.Set(10)
		c.Anonymous = true
	} else {
		c.Flags.Unset(10)
		c.Anonymous = false
	}
}

// GetAnonymous returns value of Anonymous conditional field.
func (c *ChatAdminRights) GetAnonymous() (value bool) {
	return c.Flags.Has(10)
}

// SetManageCall sets value of ManageCall conditional field.
func (c *ChatAdminRights) SetManageCall(value bool) {
	if value {
		c.Flags.Set(11)
		c.ManageCall = true
	} else {
		c.Flags.Unset(11)
		c.ManageCall = false
	}
}

// GetManageCall returns value of ManageCall conditional field.
func (c *ChatAdminRights) GetManageCall() (value bool) {
	return c.Flags.Has(11)
}

// Decode implements bin.Decoder.
func (c *ChatAdminRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminRights#5fb224d5 to nil")
	}
	if err := b.ConsumeID(ChatAdminRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: field flags: %w", err)
		}
	}
	c.ChangeInfo = c.Flags.Has(0)
	c.PostMessages = c.Flags.Has(1)
	c.EditMessages = c.Flags.Has(2)
	c.DeleteMessages = c.Flags.Has(3)
	c.BanUsers = c.Flags.Has(4)
	c.InviteUsers = c.Flags.Has(5)
	c.PinMessages = c.Flags.Has(7)
	c.AddAdmins = c.Flags.Has(9)
	c.Anonymous = c.Flags.Has(10)
	c.ManageCall = c.Flags.Has(11)
	return nil
}

// Ensuring interfaces in compile-time for ChatAdminRights.
var (
	_ bin.Encoder = &ChatAdminRights{}
	_ bin.Decoder = &ChatAdminRights{}
)
