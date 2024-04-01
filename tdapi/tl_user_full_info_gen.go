// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// UserFullInfo represents TL type `userFullInfo#809a3ac2`.
type UserFullInfo struct {
	// User profile photo set by the current user for the contact; may be null. If null and
	// user.profile_photo is null, then the photo is empty; otherwise, it is unknown.
	PersonalPhoto ChatPhoto
	// User profile photo; may be null. If null and user.profile_photo is null, then the
	// photo is empty; otherwise, it is unknown.
	Photo ChatPhoto
	// User profile photo visible if the main photo is hidden by privacy settings; may be
	// null. If null and user.profile_photo is null, then the photo is empty; otherwise, it
	// is unknown.
	PublicPhoto ChatPhoto
	// Block list to which the user is added; may be null if none
	BlockList BlockListClass
	// True, if the user can be called
	CanBeCalled bool
	// True, if a video call can be created with the user
	SupportsVideoCalls bool
	// True, if the user can't be called due to their privacy settings
	HasPrivateCalls bool
	// True, if the user can't be linked in forwarded messages due to their privacy settings
	HasPrivateForwards bool
	// True, if voice and video notes can't be sent or forwarded to the user
	HasRestrictedVoiceAndVideoNoteMessages bool
	// True, if the user has pinned stories
	HasPinnedStories bool
	// True, if the current user needs to explicitly allow to share their phone number with
	// the user when the method addContact is used
	NeedPhoneNumberPrivacyException bool
	// True, if the user set chat background for both chat users and it wasn't reverted yet
	SetChatBackground bool
	// A short user bio; may be null for bots
	Bio FormattedText
	// Birthdate of the user; may be null if unknown
	Birthdate Birthdate
	// Identifier of the personal chat of the user; 0 if none
	PersonalChatID int64
	// The list of available options for gifting Telegram Premium to the user
	PremiumGiftOptions []PremiumPaymentOption
	// Number of group chats where both the other user and the current user are a member; 0
	// for the current user
	GroupInCommonCount int32
	// Information about business settings for Telegram Business accounts; may be null if
	// none
	BusinessInfo BusinessInfo
	// For bots, information about the bot; may be null if the user isn't a bot
	BotInfo BotInfo
}

// UserFullInfoTypeID is TL type id of UserFullInfo.
const UserFullInfoTypeID = 0x809a3ac2

// Ensuring interfaces in compile-time for UserFullInfo.
var (
	_ bin.Encoder     = &UserFullInfo{}
	_ bin.Decoder     = &UserFullInfo{}
	_ bin.BareEncoder = &UserFullInfo{}
	_ bin.BareDecoder = &UserFullInfo{}
)

func (u *UserFullInfo) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.PersonalPhoto.Zero()) {
		return false
	}
	if !(u.Photo.Zero()) {
		return false
	}
	if !(u.PublicPhoto.Zero()) {
		return false
	}
	if !(u.BlockList == nil) {
		return false
	}
	if !(u.CanBeCalled == false) {
		return false
	}
	if !(u.SupportsVideoCalls == false) {
		return false
	}
	if !(u.HasPrivateCalls == false) {
		return false
	}
	if !(u.HasPrivateForwards == false) {
		return false
	}
	if !(u.HasRestrictedVoiceAndVideoNoteMessages == false) {
		return false
	}
	if !(u.HasPinnedStories == false) {
		return false
	}
	if !(u.NeedPhoneNumberPrivacyException == false) {
		return false
	}
	if !(u.SetChatBackground == false) {
		return false
	}
	if !(u.Bio.Zero()) {
		return false
	}
	if !(u.Birthdate.Zero()) {
		return false
	}
	if !(u.PersonalChatID == 0) {
		return false
	}
	if !(u.PremiumGiftOptions == nil) {
		return false
	}
	if !(u.GroupInCommonCount == 0) {
		return false
	}
	if !(u.BusinessInfo.Zero()) {
		return false
	}
	if !(u.BotInfo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserFullInfo) String() string {
	if u == nil {
		return "UserFullInfo(nil)"
	}
	type Alias UserFullInfo
	return fmt.Sprintf("UserFullInfo%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserFullInfo) TypeID() uint32 {
	return UserFullInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*UserFullInfo) TypeName() string {
	return "userFullInfo"
}

// TypeInfo returns info about TL type.
func (u *UserFullInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userFullInfo",
		ID:   UserFullInfoTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PersonalPhoto",
			SchemaName: "personal_photo",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "PublicPhoto",
			SchemaName: "public_photo",
		},
		{
			Name:       "BlockList",
			SchemaName: "block_list",
		},
		{
			Name:       "CanBeCalled",
			SchemaName: "can_be_called",
		},
		{
			Name:       "SupportsVideoCalls",
			SchemaName: "supports_video_calls",
		},
		{
			Name:       "HasPrivateCalls",
			SchemaName: "has_private_calls",
		},
		{
			Name:       "HasPrivateForwards",
			SchemaName: "has_private_forwards",
		},
		{
			Name:       "HasRestrictedVoiceAndVideoNoteMessages",
			SchemaName: "has_restricted_voice_and_video_note_messages",
		},
		{
			Name:       "HasPinnedStories",
			SchemaName: "has_pinned_stories",
		},
		{
			Name:       "NeedPhoneNumberPrivacyException",
			SchemaName: "need_phone_number_privacy_exception",
		},
		{
			Name:       "SetChatBackground",
			SchemaName: "set_chat_background",
		},
		{
			Name:       "Bio",
			SchemaName: "bio",
		},
		{
			Name:       "Birthdate",
			SchemaName: "birthdate",
		},
		{
			Name:       "PersonalChatID",
			SchemaName: "personal_chat_id",
		},
		{
			Name:       "PremiumGiftOptions",
			SchemaName: "premium_gift_options",
		},
		{
			Name:       "GroupInCommonCount",
			SchemaName: "group_in_common_count",
		},
		{
			Name:       "BusinessInfo",
			SchemaName: "business_info",
		},
		{
			Name:       "BotInfo",
			SchemaName: "bot_info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserFullInfo) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFullInfo#809a3ac2 as nil")
	}
	b.PutID(UserFullInfoTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserFullInfo) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFullInfo#809a3ac2 as nil")
	}
	if err := u.PersonalPhoto.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field personal_photo: %w", err)
	}
	if err := u.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field photo: %w", err)
	}
	if err := u.PublicPhoto.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field public_photo: %w", err)
	}
	if u.BlockList == nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field block_list is nil")
	}
	if err := u.BlockList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field block_list: %w", err)
	}
	b.PutBool(u.CanBeCalled)
	b.PutBool(u.SupportsVideoCalls)
	b.PutBool(u.HasPrivateCalls)
	b.PutBool(u.HasPrivateForwards)
	b.PutBool(u.HasRestrictedVoiceAndVideoNoteMessages)
	b.PutBool(u.HasPinnedStories)
	b.PutBool(u.NeedPhoneNumberPrivacyException)
	b.PutBool(u.SetChatBackground)
	if err := u.Bio.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field bio: %w", err)
	}
	if err := u.Birthdate.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field birthdate: %w", err)
	}
	b.PutInt53(u.PersonalChatID)
	b.PutInt(len(u.PremiumGiftOptions))
	for idx, v := range u.PremiumGiftOptions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare userFullInfo#809a3ac2: field premium_gift_options element with index %d: %w", idx, err)
		}
	}
	b.PutInt32(u.GroupInCommonCount)
	if err := u.BusinessInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field business_info: %w", err)
	}
	if err := u.BotInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field bot_info: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserFullInfo) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFullInfo#809a3ac2 to nil")
	}
	if err := b.ConsumeID(UserFullInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode userFullInfo#809a3ac2: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserFullInfo) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFullInfo#809a3ac2 to nil")
	}
	{
		if err := u.PersonalPhoto.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field personal_photo: %w", err)
		}
	}
	{
		if err := u.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field photo: %w", err)
		}
	}
	{
		if err := u.PublicPhoto.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field public_photo: %w", err)
		}
	}
	{
		value, err := DecodeBlockList(b)
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field block_list: %w", err)
		}
		u.BlockList = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field can_be_called: %w", err)
		}
		u.CanBeCalled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field supports_video_calls: %w", err)
		}
		u.SupportsVideoCalls = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_private_calls: %w", err)
		}
		u.HasPrivateCalls = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_private_forwards: %w", err)
		}
		u.HasPrivateForwards = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_restricted_voice_and_video_note_messages: %w", err)
		}
		u.HasRestrictedVoiceAndVideoNoteMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_pinned_stories: %w", err)
		}
		u.HasPinnedStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field need_phone_number_privacy_exception: %w", err)
		}
		u.NeedPhoneNumberPrivacyException = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field set_chat_background: %w", err)
		}
		u.SetChatBackground = value
	}
	{
		if err := u.Bio.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field bio: %w", err)
		}
	}
	{
		if err := u.Birthdate.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field birthdate: %w", err)
		}
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field personal_chat_id: %w", err)
		}
		u.PersonalChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field premium_gift_options: %w", err)
		}

		if headerLen > 0 {
			u.PremiumGiftOptions = make([]PremiumPaymentOption, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PremiumPaymentOption
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare userFullInfo#809a3ac2: field premium_gift_options: %w", err)
			}
			u.PremiumGiftOptions = append(u.PremiumGiftOptions, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field group_in_common_count: %w", err)
		}
		u.GroupInCommonCount = value
	}
	{
		if err := u.BusinessInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field business_info: %w", err)
		}
	}
	{
		if err := u.BotInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field bot_info: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UserFullInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode userFullInfo#809a3ac2 as nil")
	}
	b.ObjStart()
	b.PutID("userFullInfo")
	b.Comma()
	b.FieldStart("personal_photo")
	if err := u.PersonalPhoto.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field personal_photo: %w", err)
	}
	b.Comma()
	b.FieldStart("photo")
	if err := u.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field photo: %w", err)
	}
	b.Comma()
	b.FieldStart("public_photo")
	if err := u.PublicPhoto.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field public_photo: %w", err)
	}
	b.Comma()
	b.FieldStart("block_list")
	if u.BlockList == nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field block_list is nil")
	}
	if err := u.BlockList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field block_list: %w", err)
	}
	b.Comma()
	b.FieldStart("can_be_called")
	b.PutBool(u.CanBeCalled)
	b.Comma()
	b.FieldStart("supports_video_calls")
	b.PutBool(u.SupportsVideoCalls)
	b.Comma()
	b.FieldStart("has_private_calls")
	b.PutBool(u.HasPrivateCalls)
	b.Comma()
	b.FieldStart("has_private_forwards")
	b.PutBool(u.HasPrivateForwards)
	b.Comma()
	b.FieldStart("has_restricted_voice_and_video_note_messages")
	b.PutBool(u.HasRestrictedVoiceAndVideoNoteMessages)
	b.Comma()
	b.FieldStart("has_pinned_stories")
	b.PutBool(u.HasPinnedStories)
	b.Comma()
	b.FieldStart("need_phone_number_privacy_exception")
	b.PutBool(u.NeedPhoneNumberPrivacyException)
	b.Comma()
	b.FieldStart("set_chat_background")
	b.PutBool(u.SetChatBackground)
	b.Comma()
	b.FieldStart("bio")
	if err := u.Bio.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field bio: %w", err)
	}
	b.Comma()
	b.FieldStart("birthdate")
	if err := u.Birthdate.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field birthdate: %w", err)
	}
	b.Comma()
	b.FieldStart("personal_chat_id")
	b.PutInt53(u.PersonalChatID)
	b.Comma()
	b.FieldStart("premium_gift_options")
	b.ArrStart()
	for idx, v := range u.PremiumGiftOptions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field premium_gift_options element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("group_in_common_count")
	b.PutInt32(u.GroupInCommonCount)
	b.Comma()
	b.FieldStart("business_info")
	if err := u.BusinessInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field business_info: %w", err)
	}
	b.Comma()
	b.FieldStart("bot_info")
	if err := u.BotInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#809a3ac2: field bot_info: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UserFullInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode userFullInfo#809a3ac2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("userFullInfo"); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: %w", err)
			}
		case "personal_photo":
			if err := u.PersonalPhoto.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field personal_photo: %w", err)
			}
		case "photo":
			if err := u.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field photo: %w", err)
			}
		case "public_photo":
			if err := u.PublicPhoto.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field public_photo: %w", err)
			}
		case "block_list":
			value, err := DecodeTDLibJSONBlockList(b)
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field block_list: %w", err)
			}
			u.BlockList = value
		case "can_be_called":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field can_be_called: %w", err)
			}
			u.CanBeCalled = value
		case "supports_video_calls":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field supports_video_calls: %w", err)
			}
			u.SupportsVideoCalls = value
		case "has_private_calls":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_private_calls: %w", err)
			}
			u.HasPrivateCalls = value
		case "has_private_forwards":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_private_forwards: %w", err)
			}
			u.HasPrivateForwards = value
		case "has_restricted_voice_and_video_note_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_restricted_voice_and_video_note_messages: %w", err)
			}
			u.HasRestrictedVoiceAndVideoNoteMessages = value
		case "has_pinned_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field has_pinned_stories: %w", err)
			}
			u.HasPinnedStories = value
		case "need_phone_number_privacy_exception":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field need_phone_number_privacy_exception: %w", err)
			}
			u.NeedPhoneNumberPrivacyException = value
		case "set_chat_background":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field set_chat_background: %w", err)
			}
			u.SetChatBackground = value
		case "bio":
			if err := u.Bio.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field bio: %w", err)
			}
		case "birthdate":
			if err := u.Birthdate.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field birthdate: %w", err)
			}
		case "personal_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field personal_chat_id: %w", err)
			}
			u.PersonalChatID = value
		case "premium_gift_options":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value PremiumPaymentOption
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field premium_gift_options: %w", err)
				}
				u.PremiumGiftOptions = append(u.PremiumGiftOptions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field premium_gift_options: %w", err)
			}
		case "group_in_common_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field group_in_common_count: %w", err)
			}
			u.GroupInCommonCount = value
		case "business_info":
			if err := u.BusinessInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field business_info: %w", err)
			}
		case "bot_info":
			if err := u.BotInfo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode userFullInfo#809a3ac2: field bot_info: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPersonalPhoto returns value of PersonalPhoto field.
func (u *UserFullInfo) GetPersonalPhoto() (value ChatPhoto) {
	if u == nil {
		return
	}
	return u.PersonalPhoto
}

// GetPhoto returns value of Photo field.
func (u *UserFullInfo) GetPhoto() (value ChatPhoto) {
	if u == nil {
		return
	}
	return u.Photo
}

// GetPublicPhoto returns value of PublicPhoto field.
func (u *UserFullInfo) GetPublicPhoto() (value ChatPhoto) {
	if u == nil {
		return
	}
	return u.PublicPhoto
}

// GetBlockList returns value of BlockList field.
func (u *UserFullInfo) GetBlockList() (value BlockListClass) {
	if u == nil {
		return
	}
	return u.BlockList
}

// GetCanBeCalled returns value of CanBeCalled field.
func (u *UserFullInfo) GetCanBeCalled() (value bool) {
	if u == nil {
		return
	}
	return u.CanBeCalled
}

// GetSupportsVideoCalls returns value of SupportsVideoCalls field.
func (u *UserFullInfo) GetSupportsVideoCalls() (value bool) {
	if u == nil {
		return
	}
	return u.SupportsVideoCalls
}

// GetHasPrivateCalls returns value of HasPrivateCalls field.
func (u *UserFullInfo) GetHasPrivateCalls() (value bool) {
	if u == nil {
		return
	}
	return u.HasPrivateCalls
}

// GetHasPrivateForwards returns value of HasPrivateForwards field.
func (u *UserFullInfo) GetHasPrivateForwards() (value bool) {
	if u == nil {
		return
	}
	return u.HasPrivateForwards
}

// GetHasRestrictedVoiceAndVideoNoteMessages returns value of HasRestrictedVoiceAndVideoNoteMessages field.
func (u *UserFullInfo) GetHasRestrictedVoiceAndVideoNoteMessages() (value bool) {
	if u == nil {
		return
	}
	return u.HasRestrictedVoiceAndVideoNoteMessages
}

// GetHasPinnedStories returns value of HasPinnedStories field.
func (u *UserFullInfo) GetHasPinnedStories() (value bool) {
	if u == nil {
		return
	}
	return u.HasPinnedStories
}

// GetNeedPhoneNumberPrivacyException returns value of NeedPhoneNumberPrivacyException field.
func (u *UserFullInfo) GetNeedPhoneNumberPrivacyException() (value bool) {
	if u == nil {
		return
	}
	return u.NeedPhoneNumberPrivacyException
}

// GetSetChatBackground returns value of SetChatBackground field.
func (u *UserFullInfo) GetSetChatBackground() (value bool) {
	if u == nil {
		return
	}
	return u.SetChatBackground
}

// GetBio returns value of Bio field.
func (u *UserFullInfo) GetBio() (value FormattedText) {
	if u == nil {
		return
	}
	return u.Bio
}

// GetBirthdate returns value of Birthdate field.
func (u *UserFullInfo) GetBirthdate() (value Birthdate) {
	if u == nil {
		return
	}
	return u.Birthdate
}

// GetPersonalChatID returns value of PersonalChatID field.
func (u *UserFullInfo) GetPersonalChatID() (value int64) {
	if u == nil {
		return
	}
	return u.PersonalChatID
}

// GetPremiumGiftOptions returns value of PremiumGiftOptions field.
func (u *UserFullInfo) GetPremiumGiftOptions() (value []PremiumPaymentOption) {
	if u == nil {
		return
	}
	return u.PremiumGiftOptions
}

// GetGroupInCommonCount returns value of GroupInCommonCount field.
func (u *UserFullInfo) GetGroupInCommonCount() (value int32) {
	if u == nil {
		return
	}
	return u.GroupInCommonCount
}

// GetBusinessInfo returns value of BusinessInfo field.
func (u *UserFullInfo) GetBusinessInfo() (value BusinessInfo) {
	if u == nil {
		return
	}
	return u.BusinessInfo
}

// GetBotInfo returns value of BotInfo field.
func (u *UserFullInfo) GetBotInfo() (value BotInfo) {
	if u == nil {
		return
	}
	return u.BotInfo
}
