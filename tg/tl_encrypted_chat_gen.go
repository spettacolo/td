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

// EncryptedChatEmpty represents TL type `encryptedChatEmpty#ab7ec0a0`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/encryptedChatEmpty for reference.
type EncryptedChatEmpty struct {
	// Chat ID
	ID int
}

// EncryptedChatEmptyTypeID is TL type id of EncryptedChatEmpty.
const EncryptedChatEmptyTypeID = 0xab7ec0a0

func (e *EncryptedChatEmpty) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatEmpty) String() string {
	if e == nil {
		return "EncryptedChatEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EncryptedChatEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *EncryptedChatEmpty) TypeID() uint32 {
	return EncryptedChatEmptyTypeID
}

// Encode implements bin.Encoder.
func (e *EncryptedChatEmpty) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatEmpty#ab7ec0a0 as nil")
	}
	b.PutID(EncryptedChatEmptyTypeID)
	b.PutInt(e.ID)
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChatEmpty) GetID() (value int) {
	return e.ID
}

// Decode implements bin.Decoder.
func (e *EncryptedChatEmpty) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatEmpty#ab7ec0a0 to nil")
	}
	if err := b.ConsumeID(EncryptedChatEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatEmpty#ab7ec0a0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatEmpty#ab7ec0a0: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatEmpty) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatEmpty.
var (
	_ bin.Encoder = &EncryptedChatEmpty{}
	_ bin.Decoder = &EncryptedChatEmpty{}

	_ EncryptedChatClass = &EncryptedChatEmpty{}
)

// EncryptedChatWaiting represents TL type `encryptedChatWaiting#3bf703dc`.
// Chat waiting for approval of second participant.
//
// See https://core.telegram.org/constructor/encryptedChatWaiting for reference.
type EncryptedChatWaiting struct {
	// Chat ID
	ID int
	// Checking sum depending on user ID
	AccessHash int64
	// Date of chat creation
	Date int
	// Chat creator ID
	AdminID int
	// ID of second chat participant
	ParticipantID int
}

// EncryptedChatWaitingTypeID is TL type id of EncryptedChatWaiting.
const EncryptedChatWaitingTypeID = 0x3bf703dc

func (e *EncryptedChatWaiting) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.AdminID == 0) {
		return false
	}
	if !(e.ParticipantID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatWaiting) String() string {
	if e == nil {
		return "EncryptedChatWaiting(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EncryptedChatWaiting")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(e.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(e.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tAdminID: ")
	sb.WriteString(fmt.Sprint(e.AdminID))
	sb.WriteString(",\n")
	sb.WriteString("\tParticipantID: ")
	sb.WriteString(fmt.Sprint(e.ParticipantID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *EncryptedChatWaiting) TypeID() uint32 {
	return EncryptedChatWaitingTypeID
}

// Encode implements bin.Encoder.
func (e *EncryptedChatWaiting) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatWaiting#3bf703dc as nil")
	}
	b.PutID(EncryptedChatWaitingTypeID)
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutInt(e.AdminID)
	b.PutInt(e.ParticipantID)
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChatWaiting) GetID() (value int) {
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedChatWaiting) GetAccessHash() (value int64) {
	return e.AccessHash
}

// GetDate returns value of Date field.
func (e *EncryptedChatWaiting) GetDate() (value int) {
	return e.Date
}

// GetAdminID returns value of AdminID field.
func (e *EncryptedChatWaiting) GetAdminID() (value int) {
	return e.AdminID
}

// GetParticipantID returns value of ParticipantID field.
func (e *EncryptedChatWaiting) GetParticipantID() (value int) {
	return e.ParticipantID
}

// Decode implements bin.Decoder.
func (e *EncryptedChatWaiting) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatWaiting#3bf703dc to nil")
	}
	if err := b.ConsumeID(EncryptedChatWaitingTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatWaiting#3bf703dc: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatWaiting) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatWaiting.
var (
	_ bin.Encoder = &EncryptedChatWaiting{}
	_ bin.Decoder = &EncryptedChatWaiting{}

	_ EncryptedChatClass = &EncryptedChatWaiting{}
)

// EncryptedChatRequested represents TL type `encryptedChatRequested#62718a82`.
// Request to create an encrypted chat.
//
// See https://core.telegram.org/constructor/encryptedChatRequested for reference.
type EncryptedChatRequested struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	//
	// Use SetFolderID and GetFolderID helpers.
	FolderID int
	// Chat ID
	ID int
	// Check sum depending on user ID
	AccessHash int64
	// Chat creation date
	Date int
	// Chat creator ID
	AdminID int
	// ID of second chat participant
	ParticipantID int
	// A = g ^ a mod p, see Wikipedia¹
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
	GA []byte
}

// EncryptedChatRequestedTypeID is TL type id of EncryptedChatRequested.
const EncryptedChatRequestedTypeID = 0x62718a82

func (e *EncryptedChatRequested) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.FolderID == 0) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.AdminID == 0) {
		return false
	}
	if !(e.ParticipantID == 0) {
		return false
	}
	if !(e.GA == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatRequested) String() string {
	if e == nil {
		return "EncryptedChatRequested(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EncryptedChatRequested")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(e.Flags))
	sb.WriteString(",\n")
	if e.Flags.Has(0) {
		sb.WriteString("\tFolderID: ")
		sb.WriteString(fmt.Sprint(e.FolderID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(e.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(e.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tAdminID: ")
	sb.WriteString(fmt.Sprint(e.AdminID))
	sb.WriteString(",\n")
	sb.WriteString("\tParticipantID: ")
	sb.WriteString(fmt.Sprint(e.ParticipantID))
	sb.WriteString(",\n")
	sb.WriteString("\tGA: ")
	sb.WriteString(fmt.Sprint(e.GA))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *EncryptedChatRequested) TypeID() uint32 {
	return EncryptedChatRequestedTypeID
}

// Encode implements bin.Encoder.
func (e *EncryptedChatRequested) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatRequested#62718a82 as nil")
	}
	b.PutID(EncryptedChatRequestedTypeID)
	if !(e.FolderID == 0) {
		e.Flags.Set(0)
	}
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode encryptedChatRequested#62718a82: field flags: %w", err)
	}
	if e.Flags.Has(0) {
		b.PutInt(e.FolderID)
	}
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutInt(e.AdminID)
	b.PutInt(e.ParticipantID)
	b.PutBytes(e.GA)
	return nil
}

// SetFolderID sets value of FolderID conditional field.
func (e *EncryptedChatRequested) SetFolderID(value int) {
	e.Flags.Set(0)
	e.FolderID = value
}

// GetFolderID returns value of FolderID conditional field and
// boolean which is true if field was set.
func (e *EncryptedChatRequested) GetFolderID() (value int, ok bool) {
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.FolderID, true
}

// GetID returns value of ID field.
func (e *EncryptedChatRequested) GetID() (value int) {
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedChatRequested) GetAccessHash() (value int64) {
	return e.AccessHash
}

// GetDate returns value of Date field.
func (e *EncryptedChatRequested) GetDate() (value int) {
	return e.Date
}

// GetAdminID returns value of AdminID field.
func (e *EncryptedChatRequested) GetAdminID() (value int) {
	return e.AdminID
}

// GetParticipantID returns value of ParticipantID field.
func (e *EncryptedChatRequested) GetParticipantID() (value int) {
	return e.ParticipantID
}

// GetGA returns value of GA field.
func (e *EncryptedChatRequested) GetGA() (value []byte) {
	return e.GA
}

// Decode implements bin.Decoder.
func (e *EncryptedChatRequested) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatRequested#62718a82 to nil")
	}
	if err := b.ConsumeID(EncryptedChatRequestedTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field flags: %w", err)
		}
	}
	if e.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field folder_id: %w", err)
		}
		e.FolderID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatRequested#62718a82: field g_a: %w", err)
		}
		e.GA = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatRequested) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatRequested.
var (
	_ bin.Encoder = &EncryptedChatRequested{}
	_ bin.Decoder = &EncryptedChatRequested{}

	_ EncryptedChatClass = &EncryptedChatRequested{}
)

// EncryptedChat represents TL type `encryptedChat#fa56ce36`.
// Encrypted chat
//
// See https://core.telegram.org/constructor/encryptedChat for reference.
type EncryptedChat struct {
	// Chat ID
	ID int
	// Check sum dependant on the user ID
	AccessHash int64
	// Date chat was created
	Date int
	// Chat creator ID
	AdminID int
	// ID of the second chat participant
	ParticipantID int
	// B = g ^ b mod p, if the currently authorized user is the chat's creator,or A = g ^ a mod p otherwiseSee Wikipedia¹ for more info
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
	GAOrB []byte
	// 64-bit fingerprint of received key
	KeyFingerprint int64
}

// EncryptedChatTypeID is TL type id of EncryptedChat.
const EncryptedChatTypeID = 0xfa56ce36

func (e *EncryptedChat) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.AccessHash == 0) {
		return false
	}
	if !(e.Date == 0) {
		return false
	}
	if !(e.AdminID == 0) {
		return false
	}
	if !(e.ParticipantID == 0) {
		return false
	}
	if !(e.GAOrB == nil) {
		return false
	}
	if !(e.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChat) String() string {
	if e == nil {
		return "EncryptedChat(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EncryptedChat")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(e.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(e.Date))
	sb.WriteString(",\n")
	sb.WriteString("\tAdminID: ")
	sb.WriteString(fmt.Sprint(e.AdminID))
	sb.WriteString(",\n")
	sb.WriteString("\tParticipantID: ")
	sb.WriteString(fmt.Sprint(e.ParticipantID))
	sb.WriteString(",\n")
	sb.WriteString("\tGAOrB: ")
	sb.WriteString(fmt.Sprint(e.GAOrB))
	sb.WriteString(",\n")
	sb.WriteString("\tKeyFingerprint: ")
	sb.WriteString(fmt.Sprint(e.KeyFingerprint))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *EncryptedChat) TypeID() uint32 {
	return EncryptedChatTypeID
}

// Encode implements bin.Encoder.
func (e *EncryptedChat) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChat#fa56ce36 as nil")
	}
	b.PutID(EncryptedChatTypeID)
	b.PutInt(e.ID)
	b.PutLong(e.AccessHash)
	b.PutInt(e.Date)
	b.PutInt(e.AdminID)
	b.PutInt(e.ParticipantID)
	b.PutBytes(e.GAOrB)
	b.PutLong(e.KeyFingerprint)
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChat) GetID() (value int) {
	return e.ID
}

// GetAccessHash returns value of AccessHash field.
func (e *EncryptedChat) GetAccessHash() (value int64) {
	return e.AccessHash
}

// GetDate returns value of Date field.
func (e *EncryptedChat) GetDate() (value int) {
	return e.Date
}

// GetAdminID returns value of AdminID field.
func (e *EncryptedChat) GetAdminID() (value int) {
	return e.AdminID
}

// GetParticipantID returns value of ParticipantID field.
func (e *EncryptedChat) GetParticipantID() (value int) {
	return e.ParticipantID
}

// GetGAOrB returns value of GAOrB field.
func (e *EncryptedChat) GetGAOrB() (value []byte) {
	return e.GAOrB
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (e *EncryptedChat) GetKeyFingerprint() (value int64) {
	return e.KeyFingerprint
}

// Decode implements bin.Decoder.
func (e *EncryptedChat) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChat#fa56ce36 to nil")
	}
	if err := b.ConsumeID(EncryptedChatTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChat#fa56ce36: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field id: %w", err)
		}
		e.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field access_hash: %w", err)
		}
		e.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field date: %w", err)
		}
		e.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field admin_id: %w", err)
		}
		e.AdminID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field participant_id: %w", err)
		}
		e.ParticipantID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field g_a_or_b: %w", err)
		}
		e.GAOrB = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChat#fa56ce36: field key_fingerprint: %w", err)
		}
		e.KeyFingerprint = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChat) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChat.
var (
	_ bin.Encoder = &EncryptedChat{}
	_ bin.Decoder = &EncryptedChat{}

	_ EncryptedChatClass = &EncryptedChat{}
)

// EncryptedChatDiscarded represents TL type `encryptedChatDiscarded#13d6dd27`.
// Discarded or deleted chat.
//
// See https://core.telegram.org/constructor/encryptedChatDiscarded for reference.
type EncryptedChatDiscarded struct {
	// Chat ID
	ID int
}

// EncryptedChatDiscardedTypeID is TL type id of EncryptedChatDiscarded.
const EncryptedChatDiscardedTypeID = 0x13d6dd27

func (e *EncryptedChatDiscarded) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EncryptedChatDiscarded) String() string {
	if e == nil {
		return "EncryptedChatDiscarded(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EncryptedChatDiscarded")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *EncryptedChatDiscarded) TypeID() uint32 {
	return EncryptedChatDiscardedTypeID
}

// Encode implements bin.Encoder.
func (e *EncryptedChatDiscarded) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode encryptedChatDiscarded#13d6dd27 as nil")
	}
	b.PutID(EncryptedChatDiscardedTypeID)
	b.PutInt(e.ID)
	return nil
}

// GetID returns value of ID field.
func (e *EncryptedChatDiscarded) GetID() (value int) {
	return e.ID
}

// Decode implements bin.Decoder.
func (e *EncryptedChatDiscarded) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode encryptedChatDiscarded#13d6dd27 to nil")
	}
	if err := b.ConsumeID(EncryptedChatDiscardedTypeID); err != nil {
		return fmt.Errorf("unable to decode encryptedChatDiscarded#13d6dd27: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode encryptedChatDiscarded#13d6dd27: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// construct implements constructor of EncryptedChatClass.
func (e EncryptedChatDiscarded) construct() EncryptedChatClass { return &e }

// Ensuring interfaces in compile-time for EncryptedChatDiscarded.
var (
	_ bin.Encoder = &EncryptedChatDiscarded{}
	_ bin.Decoder = &EncryptedChatDiscarded{}

	_ EncryptedChatClass = &EncryptedChatDiscarded{}
)

// EncryptedChatClass represents EncryptedChat generic type.
//
// See https://core.telegram.org/type/EncryptedChat for reference.
//
// Example:
//  g, err := DecodeEncryptedChat(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *EncryptedChatEmpty: // encryptedChatEmpty#ab7ec0a0
//  case *EncryptedChatWaiting: // encryptedChatWaiting#3bf703dc
//  case *EncryptedChatRequested: // encryptedChatRequested#62718a82
//  case *EncryptedChat: // encryptedChat#fa56ce36
//  case *EncryptedChatDiscarded: // encryptedChatDiscarded#13d6dd27
//  default: panic(v)
//  }
type EncryptedChatClass interface {
	bin.Encoder
	bin.Decoder
	construct() EncryptedChatClass

	// Chat ID
	GetID() (value int)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeEncryptedChat implements binary de-serialization for EncryptedChatClass.
func DecodeEncryptedChat(buf *bin.Buffer) (EncryptedChatClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EncryptedChatEmptyTypeID:
		// Decoding encryptedChatEmpty#ab7ec0a0.
		v := EncryptedChatEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatWaitingTypeID:
		// Decoding encryptedChatWaiting#3bf703dc.
		v := EncryptedChatWaiting{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatRequestedTypeID:
		// Decoding encryptedChatRequested#62718a82.
		v := EncryptedChatRequested{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatTypeID:
		// Decoding encryptedChat#fa56ce36.
		v := EncryptedChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	case EncryptedChatDiscardedTypeID:
		// Decoding encryptedChatDiscarded#13d6dd27.
		v := EncryptedChatDiscarded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EncryptedChatClass: %w", bin.NewUnexpectedID(id))
	}
}

// EncryptedChat boxes the EncryptedChatClass providing a helper.
type EncryptedChatBox struct {
	EncryptedChat EncryptedChatClass
}

// Decode implements bin.Decoder for EncryptedChatBox.
func (b *EncryptedChatBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EncryptedChatBox to nil")
	}
	v, err := DecodeEncryptedChat(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EncryptedChat = v
	return nil
}

// Encode implements bin.Encode for EncryptedChatBox.
func (b *EncryptedChatBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EncryptedChat == nil {
		return fmt.Errorf("unable to encode EncryptedChatClass as nil")
	}
	return b.EncryptedChat.Encode(buf)
}
