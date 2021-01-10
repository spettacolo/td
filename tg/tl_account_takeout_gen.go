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

// AccountTakeout represents TL type `account.takeout#4dba4501`.
// Takout info
//
// See https://core.telegram.org/constructor/account.takeout for reference.
type AccountTakeout struct {
	// Takeout ID
	ID int64
}

// AccountTakeoutTypeID is TL type id of AccountTakeout.
const AccountTakeoutTypeID = 0x4dba4501

func (t *AccountTakeout) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountTakeout) String() string {
	if t == nil {
		return "AccountTakeout(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountTakeout")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(t.ID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *AccountTakeout) TypeID() uint32 {
	return AccountTakeoutTypeID
}

// Encode implements bin.Encoder.
func (t *AccountTakeout) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.takeout#4dba4501 as nil")
	}
	b.PutID(AccountTakeoutTypeID)
	b.PutLong(t.ID)
	return nil
}

// GetID returns value of ID field.
func (t *AccountTakeout) GetID() (value int64) {
	return t.ID
}

// Decode implements bin.Decoder.
func (t *AccountTakeout) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.takeout#4dba4501 to nil")
	}
	if err := b.ConsumeID(AccountTakeoutTypeID); err != nil {
		return fmt.Errorf("unable to decode account.takeout#4dba4501: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.takeout#4dba4501: field id: %w", err)
		}
		t.ID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountTakeout.
var (
	_ bin.Encoder = &AccountTakeout{}
	_ bin.Decoder = &AccountTakeout{}
)
