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

// HighScore represents TL type `highScore#58fffcd0`.
// Game highscore
//
// See https://core.telegram.org/constructor/highScore for reference.
type HighScore struct {
	// Position in highscore list
	Pos int
	// User ID
	UserID int
	// Score
	Score int
}

// HighScoreTypeID is TL type id of HighScore.
const HighScoreTypeID = 0x58fffcd0

func (h *HighScore) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Pos == 0) {
		return false
	}
	if !(h.UserID == 0) {
		return false
	}
	if !(h.Score == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *HighScore) String() string {
	if h == nil {
		return "HighScore(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HighScore")
	sb.WriteString("{\n")
	sb.WriteString("\tPos: ")
	sb.WriteString(fmt.Sprint(h.Pos))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(h.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tScore: ")
	sb.WriteString(fmt.Sprint(h.Score))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (h *HighScore) TypeID() uint32 {
	return HighScoreTypeID
}

// Encode implements bin.Encoder.
func (h *HighScore) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode highScore#58fffcd0 as nil")
	}
	b.PutID(HighScoreTypeID)
	b.PutInt(h.Pos)
	b.PutInt(h.UserID)
	b.PutInt(h.Score)
	return nil
}

// GetPos returns value of Pos field.
func (h *HighScore) GetPos() (value int) {
	return h.Pos
}

// GetUserID returns value of UserID field.
func (h *HighScore) GetUserID() (value int) {
	return h.UserID
}

// GetScore returns value of Score field.
func (h *HighScore) GetScore() (value int) {
	return h.Score
}

// Decode implements bin.Decoder.
func (h *HighScore) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode highScore#58fffcd0 to nil")
	}
	if err := b.ConsumeID(HighScoreTypeID); err != nil {
		return fmt.Errorf("unable to decode highScore#58fffcd0: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode highScore#58fffcd0: field pos: %w", err)
		}
		h.Pos = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode highScore#58fffcd0: field user_id: %w", err)
		}
		h.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode highScore#58fffcd0: field score: %w", err)
		}
		h.Score = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HighScore.
var (
	_ bin.Encoder = &HighScore{}
	_ bin.Decoder = &HighScore{}
)
