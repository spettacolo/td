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

// StatsGroupTopPoster represents TL type `statsGroupTopPoster#18f3d0f7`.
// Information about an active user in a supergroup
//
// See https://core.telegram.org/constructor/statsGroupTopPoster for reference.
type StatsGroupTopPoster struct {
	// User ID
	UserID int
	// Number of messages for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Messages int
	// Average number of characters per message
	AvgChars int
}

// StatsGroupTopPosterTypeID is TL type id of StatsGroupTopPoster.
const StatsGroupTopPosterTypeID = 0x18f3d0f7

func (s *StatsGroupTopPoster) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Messages == 0) {
		return false
	}
	if !(s.AvgChars == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGroupTopPoster) String() string {
	if s == nil {
		return "StatsGroupTopPoster(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsGroupTopPoster")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(s.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tMessages: ")
	sb.WriteString(fmt.Sprint(s.Messages))
	sb.WriteString(",\n")
	sb.WriteString("\tAvgChars: ")
	sb.WriteString(fmt.Sprint(s.AvgChars))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *StatsGroupTopPoster) TypeID() uint32 {
	return StatsGroupTopPosterTypeID
}

// Encode implements bin.Encoder.
func (s *StatsGroupTopPoster) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopPoster#18f3d0f7 as nil")
	}
	b.PutID(StatsGroupTopPosterTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Messages)
	b.PutInt(s.AvgChars)
	return nil
}

// GetUserID returns value of UserID field.
func (s *StatsGroupTopPoster) GetUserID() (value int) {
	return s.UserID
}

// GetMessages returns value of Messages field.
func (s *StatsGroupTopPoster) GetMessages() (value int) {
	return s.Messages
}

// GetAvgChars returns value of AvgChars field.
func (s *StatsGroupTopPoster) GetAvgChars() (value int) {
	return s.AvgChars
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopPoster) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopPoster#18f3d0f7 to nil")
	}
	if err := b.ConsumeID(StatsGroupTopPosterTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field messages: %w", err)
		}
		s.Messages = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopPoster#18f3d0f7: field avg_chars: %w", err)
		}
		s.AvgChars = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopPoster.
var (
	_ bin.Encoder = &StatsGroupTopPoster{}
	_ bin.Decoder = &StatsGroupTopPoster{}
)
