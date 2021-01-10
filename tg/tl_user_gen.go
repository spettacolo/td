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

// UserEmpty represents TL type `userEmpty#200250ba`.
// Empty constructor, non-existent user.
//
// See https://core.telegram.org/constructor/userEmpty for reference.
type UserEmpty struct {
	// User identifier or 0
	ID int
}

// UserEmptyTypeID is TL type id of UserEmpty.
const UserEmptyTypeID = 0x200250ba

func (u *UserEmpty) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserEmpty) String() string {
	if u == nil {
		return "UserEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UserEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(u.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *UserEmpty) TypeID() uint32 {
	return UserEmptyTypeID
}

// Encode implements bin.Encoder.
func (u *UserEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userEmpty#200250ba as nil")
	}
	b.PutID(UserEmptyTypeID)
	b.PutInt(u.ID)
	return nil
}

// GetID returns value of ID field.
func (u *UserEmpty) GetID() (value int) {
	return u.ID
}

// Decode implements bin.Decoder.
func (u *UserEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userEmpty#200250ba to nil")
	}
	if err := b.ConsumeID(UserEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userEmpty#200250ba: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userEmpty#200250ba: field id: %w", err)
		}
		u.ID = value
	}
	return nil
}

// construct implements constructor of UserClass.
func (u UserEmpty) construct() UserClass { return &u }

// Ensuring interfaces in compile-time for UserEmpty.
var (
	_ bin.Encoder = &UserEmpty{}
	_ bin.Decoder = &UserEmpty{}

	_ UserClass = &UserEmpty{}
)

// User represents TL type `user#938458c1`.
// Indicates info about a certain user
//
// See https://core.telegram.org/constructor/user for reference.
type User struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this user indicates the currently logged in user
	Self bool
	// Whether this user is a contact
	Contact bool
	// Whether this user is a mutual contact
	MutualContact bool
	// Whether the account of this user was deleted
	Deleted bool
	// Is this user a bot?
	Bot bool
	// Can the bot see all messages in groups?
	BotChatHistory bool
	// Can the bot be added to groups?
	BotNochats bool
	// Whether this user is verified
	Verified bool
	// Access to this user must be restricted for the reason specified in restriction_reason
	Restricted bool
	// See min¹
	//
	// Links:
	//  1) https://core.telegram.org/api/min
	Min bool
	// Whether the bot can request our geolocation in inline mode
	BotInlineGeo bool
	// Whether this is an official support user
	Support bool
	// This may be a scam user
	Scam bool
	// If set, the profile picture for this user should be refetched
	ApplyMinPhoto bool
	// ID of the user
	ID int
	// Access hash of the user
	//
	// Use SetAccessHash and GetAccessHash helpers.
	AccessHash int64
	// First name
	//
	// Use SetFirstName and GetFirstName helpers.
	FirstName string
	// Last name
	//
	// Use SetLastName and GetLastName helpers.
	LastName string
	// Username
	//
	// Use SetUsername and GetUsername helpers.
	Username string
	// Phone number
	//
	// Use SetPhone and GetPhone helpers.
	Phone string
	// Profile picture of user
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo UserProfilePhotoClass
	// Online status of user
	//
	// Use SetStatus and GetStatus helpers.
	Status UserStatusClass
	// Version of the bot_info field in userFull¹, incremented every time it changes
	//
	// Links:
	//  1) https://core.telegram.org/constructor/userFull
	//
	// Use SetBotInfoVersion and GetBotInfoVersion helpers.
	BotInfoVersion int
	// Contains the reason why access to this user must be restricted.
	//
	// Use SetRestrictionReason and GetRestrictionReason helpers.
	RestrictionReason []RestrictionReason
	// Inline placeholder for this inline bot
	//
	// Use SetBotInlinePlaceholder and GetBotInlinePlaceholder helpers.
	BotInlinePlaceholder string
	// Language code of the user
	//
	// Use SetLangCode and GetLangCode helpers.
	LangCode string
}

// UserTypeID is TL type id of User.
const UserTypeID = 0x938458c1

func (u *User) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.Self == false) {
		return false
	}
	if !(u.Contact == false) {
		return false
	}
	if !(u.MutualContact == false) {
		return false
	}
	if !(u.Deleted == false) {
		return false
	}
	if !(u.Bot == false) {
		return false
	}
	if !(u.BotChatHistory == false) {
		return false
	}
	if !(u.BotNochats == false) {
		return false
	}
	if !(u.Verified == false) {
		return false
	}
	if !(u.Restricted == false) {
		return false
	}
	if !(u.Min == false) {
		return false
	}
	if !(u.BotInlineGeo == false) {
		return false
	}
	if !(u.Support == false) {
		return false
	}
	if !(u.Scam == false) {
		return false
	}
	if !(u.ApplyMinPhoto == false) {
		return false
	}
	if !(u.ID == 0) {
		return false
	}
	if !(u.AccessHash == 0) {
		return false
	}
	if !(u.FirstName == "") {
		return false
	}
	if !(u.LastName == "") {
		return false
	}
	if !(u.Username == "") {
		return false
	}
	if !(u.Phone == "") {
		return false
	}
	if !(u.Photo == nil) {
		return false
	}
	if !(u.Status == nil) {
		return false
	}
	if !(u.BotInfoVersion == 0) {
		return false
	}
	if !(u.RestrictionReason == nil) {
		return false
	}
	if !(u.BotInlinePlaceholder == "") {
		return false
	}
	if !(u.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *User) String() string {
	if u == nil {
		return "User(nil)"
	}
	var sb strings.Builder
	sb.WriteString("User")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(u.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(u.ID))
	sb.WriteString(",\n")
	if u.Flags.Has(0) {
		sb.WriteString("\tAccessHash: ")
		sb.WriteString(fmt.Sprint(u.AccessHash))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(1) {
		sb.WriteString("\tFirstName: ")
		sb.WriteString(fmt.Sprint(u.FirstName))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(2) {
		sb.WriteString("\tLastName: ")
		sb.WriteString(fmt.Sprint(u.LastName))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(3) {
		sb.WriteString("\tUsername: ")
		sb.WriteString(fmt.Sprint(u.Username))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(4) {
		sb.WriteString("\tPhone: ")
		sb.WriteString(fmt.Sprint(u.Phone))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(5) {
		sb.WriteString("\tPhoto: ")
		sb.WriteString(fmt.Sprint(u.Photo))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(6) {
		sb.WriteString("\tStatus: ")
		sb.WriteString(fmt.Sprint(u.Status))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(14) {
		sb.WriteString("\tBotInfoVersion: ")
		sb.WriteString(fmt.Sprint(u.BotInfoVersion))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(18) {
		sb.WriteByte('[')
		for _, v := range u.RestrictionReason {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	if u.Flags.Has(19) {
		sb.WriteString("\tBotInlinePlaceholder: ")
		sb.WriteString(fmt.Sprint(u.BotInlinePlaceholder))
		sb.WriteString(",\n")
	}
	if u.Flags.Has(22) {
		sb.WriteString("\tLangCode: ")
		sb.WriteString(fmt.Sprint(u.LangCode))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (u *User) TypeID() uint32 {
	return UserTypeID
}

// Encode implements bin.Encoder.
func (u *User) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode user#938458c1 as nil")
	}
	b.PutID(UserTypeID)
	if !(u.Self == false) {
		u.Flags.Set(10)
	}
	if !(u.Contact == false) {
		u.Flags.Set(11)
	}
	if !(u.MutualContact == false) {
		u.Flags.Set(12)
	}
	if !(u.Deleted == false) {
		u.Flags.Set(13)
	}
	if !(u.Bot == false) {
		u.Flags.Set(14)
	}
	if !(u.BotChatHistory == false) {
		u.Flags.Set(15)
	}
	if !(u.BotNochats == false) {
		u.Flags.Set(16)
	}
	if !(u.Verified == false) {
		u.Flags.Set(17)
	}
	if !(u.Restricted == false) {
		u.Flags.Set(18)
	}
	if !(u.Min == false) {
		u.Flags.Set(20)
	}
	if !(u.BotInlineGeo == false) {
		u.Flags.Set(21)
	}
	if !(u.Support == false) {
		u.Flags.Set(23)
	}
	if !(u.Scam == false) {
		u.Flags.Set(24)
	}
	if !(u.ApplyMinPhoto == false) {
		u.Flags.Set(25)
	}
	if !(u.AccessHash == 0) {
		u.Flags.Set(0)
	}
	if !(u.FirstName == "") {
		u.Flags.Set(1)
	}
	if !(u.LastName == "") {
		u.Flags.Set(2)
	}
	if !(u.Username == "") {
		u.Flags.Set(3)
	}
	if !(u.Phone == "") {
		u.Flags.Set(4)
	}
	if !(u.Photo == nil) {
		u.Flags.Set(5)
	}
	if !(u.Status == nil) {
		u.Flags.Set(6)
	}
	if !(u.BotInfoVersion == 0) {
		u.Flags.Set(14)
	}
	if !(u.RestrictionReason == nil) {
		u.Flags.Set(18)
	}
	if !(u.BotInlinePlaceholder == "") {
		u.Flags.Set(19)
	}
	if !(u.LangCode == "") {
		u.Flags.Set(22)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode user#938458c1: field flags: %w", err)
	}
	b.PutInt(u.ID)
	if u.Flags.Has(0) {
		b.PutLong(u.AccessHash)
	}
	if u.Flags.Has(1) {
		b.PutString(u.FirstName)
	}
	if u.Flags.Has(2) {
		b.PutString(u.LastName)
	}
	if u.Flags.Has(3) {
		b.PutString(u.Username)
	}
	if u.Flags.Has(4) {
		b.PutString(u.Phone)
	}
	if u.Flags.Has(5) {
		if u.Photo == nil {
			return fmt.Errorf("unable to encode user#938458c1: field photo is nil")
		}
		if err := u.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode user#938458c1: field photo: %w", err)
		}
	}
	if u.Flags.Has(6) {
		if u.Status == nil {
			return fmt.Errorf("unable to encode user#938458c1: field status is nil")
		}
		if err := u.Status.Encode(b); err != nil {
			return fmt.Errorf("unable to encode user#938458c1: field status: %w", err)
		}
	}
	if u.Flags.Has(14) {
		b.PutInt(u.BotInfoVersion)
	}
	if u.Flags.Has(18) {
		b.PutVectorHeader(len(u.RestrictionReason))
		for idx, v := range u.RestrictionReason {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode user#938458c1: field restriction_reason element with index %d: %w", idx, err)
			}
		}
	}
	if u.Flags.Has(19) {
		b.PutString(u.BotInlinePlaceholder)
	}
	if u.Flags.Has(22) {
		b.PutString(u.LangCode)
	}
	return nil
}

// SetSelf sets value of Self conditional field.
func (u *User) SetSelf(value bool) {
	if value {
		u.Flags.Set(10)
		u.Self = true
	} else {
		u.Flags.Unset(10)
		u.Self = false
	}
}

// GetSelf returns value of Self conditional field.
func (u *User) GetSelf() (value bool) {
	return u.Flags.Has(10)
}

// SetContact sets value of Contact conditional field.
func (u *User) SetContact(value bool) {
	if value {
		u.Flags.Set(11)
		u.Contact = true
	} else {
		u.Flags.Unset(11)
		u.Contact = false
	}
}

// GetContact returns value of Contact conditional field.
func (u *User) GetContact() (value bool) {
	return u.Flags.Has(11)
}

// SetMutualContact sets value of MutualContact conditional field.
func (u *User) SetMutualContact(value bool) {
	if value {
		u.Flags.Set(12)
		u.MutualContact = true
	} else {
		u.Flags.Unset(12)
		u.MutualContact = false
	}
}

// GetMutualContact returns value of MutualContact conditional field.
func (u *User) GetMutualContact() (value bool) {
	return u.Flags.Has(12)
}

// SetDeleted sets value of Deleted conditional field.
func (u *User) SetDeleted(value bool) {
	if value {
		u.Flags.Set(13)
		u.Deleted = true
	} else {
		u.Flags.Unset(13)
		u.Deleted = false
	}
}

// GetDeleted returns value of Deleted conditional field.
func (u *User) GetDeleted() (value bool) {
	return u.Flags.Has(13)
}

// SetBot sets value of Bot conditional field.
func (u *User) SetBot(value bool) {
	if value {
		u.Flags.Set(14)
		u.Bot = true
	} else {
		u.Flags.Unset(14)
		u.Bot = false
	}
}

// GetBot returns value of Bot conditional field.
func (u *User) GetBot() (value bool) {
	return u.Flags.Has(14)
}

// SetBotChatHistory sets value of BotChatHistory conditional field.
func (u *User) SetBotChatHistory(value bool) {
	if value {
		u.Flags.Set(15)
		u.BotChatHistory = true
	} else {
		u.Flags.Unset(15)
		u.BotChatHistory = false
	}
}

// GetBotChatHistory returns value of BotChatHistory conditional field.
func (u *User) GetBotChatHistory() (value bool) {
	return u.Flags.Has(15)
}

// SetBotNochats sets value of BotNochats conditional field.
func (u *User) SetBotNochats(value bool) {
	if value {
		u.Flags.Set(16)
		u.BotNochats = true
	} else {
		u.Flags.Unset(16)
		u.BotNochats = false
	}
}

// GetBotNochats returns value of BotNochats conditional field.
func (u *User) GetBotNochats() (value bool) {
	return u.Flags.Has(16)
}

// SetVerified sets value of Verified conditional field.
func (u *User) SetVerified(value bool) {
	if value {
		u.Flags.Set(17)
		u.Verified = true
	} else {
		u.Flags.Unset(17)
		u.Verified = false
	}
}

// GetVerified returns value of Verified conditional field.
func (u *User) GetVerified() (value bool) {
	return u.Flags.Has(17)
}

// SetRestricted sets value of Restricted conditional field.
func (u *User) SetRestricted(value bool) {
	if value {
		u.Flags.Set(18)
		u.Restricted = true
	} else {
		u.Flags.Unset(18)
		u.Restricted = false
	}
}

// GetRestricted returns value of Restricted conditional field.
func (u *User) GetRestricted() (value bool) {
	return u.Flags.Has(18)
}

// SetMin sets value of Min conditional field.
func (u *User) SetMin(value bool) {
	if value {
		u.Flags.Set(20)
		u.Min = true
	} else {
		u.Flags.Unset(20)
		u.Min = false
	}
}

// GetMin returns value of Min conditional field.
func (u *User) GetMin() (value bool) {
	return u.Flags.Has(20)
}

// SetBotInlineGeo sets value of BotInlineGeo conditional field.
func (u *User) SetBotInlineGeo(value bool) {
	if value {
		u.Flags.Set(21)
		u.BotInlineGeo = true
	} else {
		u.Flags.Unset(21)
		u.BotInlineGeo = false
	}
}

// GetBotInlineGeo returns value of BotInlineGeo conditional field.
func (u *User) GetBotInlineGeo() (value bool) {
	return u.Flags.Has(21)
}

// SetSupport sets value of Support conditional field.
func (u *User) SetSupport(value bool) {
	if value {
		u.Flags.Set(23)
		u.Support = true
	} else {
		u.Flags.Unset(23)
		u.Support = false
	}
}

// GetSupport returns value of Support conditional field.
func (u *User) GetSupport() (value bool) {
	return u.Flags.Has(23)
}

// SetScam sets value of Scam conditional field.
func (u *User) SetScam(value bool) {
	if value {
		u.Flags.Set(24)
		u.Scam = true
	} else {
		u.Flags.Unset(24)
		u.Scam = false
	}
}

// GetScam returns value of Scam conditional field.
func (u *User) GetScam() (value bool) {
	return u.Flags.Has(24)
}

// SetApplyMinPhoto sets value of ApplyMinPhoto conditional field.
func (u *User) SetApplyMinPhoto(value bool) {
	if value {
		u.Flags.Set(25)
		u.ApplyMinPhoto = true
	} else {
		u.Flags.Unset(25)
		u.ApplyMinPhoto = false
	}
}

// GetApplyMinPhoto returns value of ApplyMinPhoto conditional field.
func (u *User) GetApplyMinPhoto() (value bool) {
	return u.Flags.Has(25)
}

// GetID returns value of ID field.
func (u *User) GetID() (value int) {
	return u.ID
}

// SetAccessHash sets value of AccessHash conditional field.
func (u *User) SetAccessHash(value int64) {
	u.Flags.Set(0)
	u.AccessHash = value
}

// GetAccessHash returns value of AccessHash conditional field and
// boolean which is true if field was set.
func (u *User) GetAccessHash() (value int64, ok bool) {
	if !u.Flags.Has(0) {
		return value, false
	}
	return u.AccessHash, true
}

// SetFirstName sets value of FirstName conditional field.
func (u *User) SetFirstName(value string) {
	u.Flags.Set(1)
	u.FirstName = value
}

// GetFirstName returns value of FirstName conditional field and
// boolean which is true if field was set.
func (u *User) GetFirstName() (value string, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.FirstName, true
}

// SetLastName sets value of LastName conditional field.
func (u *User) SetLastName(value string) {
	u.Flags.Set(2)
	u.LastName = value
}

// GetLastName returns value of LastName conditional field and
// boolean which is true if field was set.
func (u *User) GetLastName() (value string, ok bool) {
	if !u.Flags.Has(2) {
		return value, false
	}
	return u.LastName, true
}

// SetUsername sets value of Username conditional field.
func (u *User) SetUsername(value string) {
	u.Flags.Set(3)
	u.Username = value
}

// GetUsername returns value of Username conditional field and
// boolean which is true if field was set.
func (u *User) GetUsername() (value string, ok bool) {
	if !u.Flags.Has(3) {
		return value, false
	}
	return u.Username, true
}

// SetPhone sets value of Phone conditional field.
func (u *User) SetPhone(value string) {
	u.Flags.Set(4)
	u.Phone = value
}

// GetPhone returns value of Phone conditional field and
// boolean which is true if field was set.
func (u *User) GetPhone() (value string, ok bool) {
	if !u.Flags.Has(4) {
		return value, false
	}
	return u.Phone, true
}

// SetPhoto sets value of Photo conditional field.
func (u *User) SetPhoto(value UserProfilePhotoClass) {
	u.Flags.Set(5)
	u.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (u *User) GetPhoto() (value UserProfilePhotoClass, ok bool) {
	if !u.Flags.Has(5) {
		return value, false
	}
	return u.Photo, true
}

// SetStatus sets value of Status conditional field.
func (u *User) SetStatus(value UserStatusClass) {
	u.Flags.Set(6)
	u.Status = value
}

// GetStatus returns value of Status conditional field and
// boolean which is true if field was set.
func (u *User) GetStatus() (value UserStatusClass, ok bool) {
	if !u.Flags.Has(6) {
		return value, false
	}
	return u.Status, true
}

// SetBotInfoVersion sets value of BotInfoVersion conditional field.
func (u *User) SetBotInfoVersion(value int) {
	u.Flags.Set(14)
	u.BotInfoVersion = value
}

// GetBotInfoVersion returns value of BotInfoVersion conditional field and
// boolean which is true if field was set.
func (u *User) GetBotInfoVersion() (value int, ok bool) {
	if !u.Flags.Has(14) {
		return value, false
	}
	return u.BotInfoVersion, true
}

// SetRestrictionReason sets value of RestrictionReason conditional field.
func (u *User) SetRestrictionReason(value []RestrictionReason) {
	u.Flags.Set(18)
	u.RestrictionReason = value
}

// GetRestrictionReason returns value of RestrictionReason conditional field and
// boolean which is true if field was set.
func (u *User) GetRestrictionReason() (value []RestrictionReason, ok bool) {
	if !u.Flags.Has(18) {
		return value, false
	}
	return u.RestrictionReason, true
}

// SetBotInlinePlaceholder sets value of BotInlinePlaceholder conditional field.
func (u *User) SetBotInlinePlaceholder(value string) {
	u.Flags.Set(19)
	u.BotInlinePlaceholder = value
}

// GetBotInlinePlaceholder returns value of BotInlinePlaceholder conditional field and
// boolean which is true if field was set.
func (u *User) GetBotInlinePlaceholder() (value string, ok bool) {
	if !u.Flags.Has(19) {
		return value, false
	}
	return u.BotInlinePlaceholder, true
}

// SetLangCode sets value of LangCode conditional field.
func (u *User) SetLangCode(value string) {
	u.Flags.Set(22)
	u.LangCode = value
}

// GetLangCode returns value of LangCode conditional field and
// boolean which is true if field was set.
func (u *User) GetLangCode() (value string, ok bool) {
	if !u.Flags.Has(22) {
		return value, false
	}
	return u.LangCode, true
}

// Decode implements bin.Decoder.
func (u *User) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode user#938458c1 to nil")
	}
	if err := b.ConsumeID(UserTypeID); err != nil {
		return fmt.Errorf("unable to decode user#938458c1: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field flags: %w", err)
		}
	}
	u.Self = u.Flags.Has(10)
	u.Contact = u.Flags.Has(11)
	u.MutualContact = u.Flags.Has(12)
	u.Deleted = u.Flags.Has(13)
	u.Bot = u.Flags.Has(14)
	u.BotChatHistory = u.Flags.Has(15)
	u.BotNochats = u.Flags.Has(16)
	u.Verified = u.Flags.Has(17)
	u.Restricted = u.Flags.Has(18)
	u.Min = u.Flags.Has(20)
	u.BotInlineGeo = u.Flags.Has(21)
	u.Support = u.Flags.Has(23)
	u.Scam = u.Flags.Has(24)
	u.ApplyMinPhoto = u.Flags.Has(25)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field id: %w", err)
		}
		u.ID = value
	}
	if u.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field access_hash: %w", err)
		}
		u.AccessHash = value
	}
	if u.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field first_name: %w", err)
		}
		u.FirstName = value
	}
	if u.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field last_name: %w", err)
		}
		u.LastName = value
	}
	if u.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field username: %w", err)
		}
		u.Username = value
	}
	if u.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field phone: %w", err)
		}
		u.Phone = value
	}
	if u.Flags.Has(5) {
		value, err := DecodeUserProfilePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field photo: %w", err)
		}
		u.Photo = value
	}
	if u.Flags.Has(6) {
		value, err := DecodeUserStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field status: %w", err)
		}
		u.Status = value
	}
	if u.Flags.Has(14) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field bot_info_version: %w", err)
		}
		u.BotInfoVersion = value
	}
	if u.Flags.Has(18) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field restriction_reason: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value RestrictionReason
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode user#938458c1: field restriction_reason: %w", err)
			}
			u.RestrictionReason = append(u.RestrictionReason, value)
		}
	}
	if u.Flags.Has(19) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field bot_inline_placeholder: %w", err)
		}
		u.BotInlinePlaceholder = value
	}
	if u.Flags.Has(22) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user#938458c1: field lang_code: %w", err)
		}
		u.LangCode = value
	}
	return nil
}

// construct implements constructor of UserClass.
func (u User) construct() UserClass { return &u }

// Ensuring interfaces in compile-time for User.
var (
	_ bin.Encoder = &User{}
	_ bin.Decoder = &User{}

	_ UserClass = &User{}
)

// UserClass represents User generic type.
//
// See https://core.telegram.org/type/User for reference.
//
// Example:
//  g, err := DecodeUser(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UserEmpty: // userEmpty#200250ba
//  case *User: // user#938458c1
//  default: panic(v)
//  }
type UserClass interface {
	bin.Encoder
	bin.Decoder
	construct() UserClass

	// User identifier or 0
	GetID() (value int)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeUser implements binary de-serialization for UserClass.
func DecodeUser(buf *bin.Buffer) (UserClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserEmptyTypeID:
		// Decoding userEmpty#200250ba.
		v := UserEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserClass: %w", err)
		}
		return &v, nil
	case UserTypeID:
		// Decoding user#938458c1.
		v := User{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserClass: %w", bin.NewUnexpectedID(id))
	}
}

// User boxes the UserClass providing a helper.
type UserBox struct {
	User UserClass
}

// Decode implements bin.Decoder for UserBox.
func (b *UserBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserBox to nil")
	}
	v, err := DecodeUser(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.User = v
	return nil
}

// Encode implements bin.Encode for UserBox.
func (b *UserBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.User == nil {
		return fmt.Errorf("unable to encode UserClass as nil")
	}
	return b.User.Encode(buf)
}
