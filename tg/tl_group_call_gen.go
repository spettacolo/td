// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// GroupCallDiscarded represents TL type `groupCallDiscarded#7780bcb4`.
// An ended group call
//
// See https://core.telegram.org/constructor/groupCallDiscarded for reference.
type GroupCallDiscarded struct {
	// Group call ID
	ID int64
	// Group call access hash
	AccessHash int64
	// Group call duration
	Duration int
}

// GroupCallDiscardedTypeID is TL type id of GroupCallDiscarded.
const GroupCallDiscardedTypeID = 0x7780bcb4

// construct implements constructor of GroupCallClass.
func (g GroupCallDiscarded) construct() GroupCallClass { return &g }

// Ensuring interfaces in compile-time for GroupCallDiscarded.
var (
	_ bin.Encoder     = &GroupCallDiscarded{}
	_ bin.Decoder     = &GroupCallDiscarded{}
	_ bin.BareEncoder = &GroupCallDiscarded{}
	_ bin.BareDecoder = &GroupCallDiscarded{}

	_ GroupCallClass = &GroupCallDiscarded{}
)

func (g *GroupCallDiscarded) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.AccessHash == 0) {
		return false
	}
	if !(g.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallDiscarded) String() string {
	if g == nil {
		return "GroupCallDiscarded(nil)"
	}
	type Alias GroupCallDiscarded
	return fmt.Sprintf("GroupCallDiscarded%+v", Alias(*g))
}

// FillFrom fills GroupCallDiscarded from given interface.
func (g *GroupCallDiscarded) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetDuration() (value int)
}) {
	g.ID = from.GetID()
	g.AccessHash = from.GetAccessHash()
	g.Duration = from.GetDuration()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCallDiscarded) TypeID() uint32 {
	return GroupCallDiscardedTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCallDiscarded) TypeName() string {
	return "groupCallDiscarded"
}

// TypeInfo returns info about TL type.
func (g *GroupCallDiscarded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCallDiscarded",
		ID:   GroupCallDiscardedTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GroupCallDiscarded) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallDiscarded#7780bcb4 as nil")
	}
	b.PutID(GroupCallDiscardedTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCallDiscarded) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallDiscarded#7780bcb4 as nil")
	}
	b.PutLong(g.ID)
	b.PutLong(g.AccessHash)
	b.PutInt(g.Duration)
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCallDiscarded) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallDiscarded#7780bcb4 to nil")
	}
	if err := b.ConsumeID(GroupCallDiscardedTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCallDiscarded) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallDiscarded#7780bcb4 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: field access_hash: %w", err)
		}
		g.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallDiscarded#7780bcb4: field duration: %w", err)
		}
		g.Duration = value
	}
	return nil
}

// GetID returns value of ID field.
func (g *GroupCallDiscarded) GetID() (value int64) {
	return g.ID
}

// GetAccessHash returns value of AccessHash field.
func (g *GroupCallDiscarded) GetAccessHash() (value int64) {
	return g.AccessHash
}

// GetDuration returns value of Duration field.
func (g *GroupCallDiscarded) GetDuration() (value int) {
	return g.Duration
}

// GroupCall represents TL type `groupCall#d597650c`.
// Info about a group call or livestream
//
// See https://core.telegram.org/constructor/groupCall for reference.
type GroupCall struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user should be muted upon joining the call
	JoinMuted bool
	// Whether the current user can change the value of the join_muted flag using phone
	// toggleGroupCallSettings¹
	//
	// Links:
	//  1) https://core.telegram.org/method/phone.toggleGroupCallSettings
	CanChangeJoinMuted bool
	// Specifies the ordering to use when locally sorting by date and displaying in the UI
	// group call participants.
	JoinDateAsc bool
	// Whether we subscribed to the scheduled call
	ScheduleStartSubscribed bool
	// Whether you can start streaming video into the call
	CanStartVideo bool
	// Whether the group call is currently being recorded
	RecordVideoActive bool
	// Group call ID
	ID int64
	// Group call access hash
	AccessHash int64
	// Participant count
	ParticipantsCount int
	// Group call title
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// DC ID to be used for livestream chunks
	//
	// Use SetStreamDCID and GetStreamDCID helpers.
	StreamDCID int
	// When was the recording started
	//
	// Use SetRecordStartDate and GetRecordStartDate helpers.
	RecordStartDate int
	// When is the call scheduled to start
	//
	// Use SetScheduleDate and GetScheduleDate helpers.
	ScheduleDate int
	// Number of people currently streaming video into the call
	//
	// Use SetUnmutedVideoCount and GetUnmutedVideoCount helpers.
	UnmutedVideoCount int
	// Maximum number of people allowed to stream video into the call
	UnmutedVideoLimit int
	// Version
	Version int
}

// GroupCallTypeID is TL type id of GroupCall.
const GroupCallTypeID = 0xd597650c

// construct implements constructor of GroupCallClass.
func (g GroupCall) construct() GroupCallClass { return &g }

// Ensuring interfaces in compile-time for GroupCall.
var (
	_ bin.Encoder     = &GroupCall{}
	_ bin.Decoder     = &GroupCall{}
	_ bin.BareEncoder = &GroupCall{}
	_ bin.BareDecoder = &GroupCall{}

	_ GroupCallClass = &GroupCall{}
)

func (g *GroupCall) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.JoinMuted == false) {
		return false
	}
	if !(g.CanChangeJoinMuted == false) {
		return false
	}
	if !(g.JoinDateAsc == false) {
		return false
	}
	if !(g.ScheduleStartSubscribed == false) {
		return false
	}
	if !(g.CanStartVideo == false) {
		return false
	}
	if !(g.RecordVideoActive == false) {
		return false
	}
	if !(g.ID == 0) {
		return false
	}
	if !(g.AccessHash == 0) {
		return false
	}
	if !(g.ParticipantsCount == 0) {
		return false
	}
	if !(g.Title == "") {
		return false
	}
	if !(g.StreamDCID == 0) {
		return false
	}
	if !(g.RecordStartDate == 0) {
		return false
	}
	if !(g.ScheduleDate == 0) {
		return false
	}
	if !(g.UnmutedVideoCount == 0) {
		return false
	}
	if !(g.UnmutedVideoLimit == 0) {
		return false
	}
	if !(g.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCall) String() string {
	if g == nil {
		return "GroupCall(nil)"
	}
	type Alias GroupCall
	return fmt.Sprintf("GroupCall%+v", Alias(*g))
}

// FillFrom fills GroupCall from given interface.
func (g *GroupCall) FillFrom(from interface {
	GetJoinMuted() (value bool)
	GetCanChangeJoinMuted() (value bool)
	GetJoinDateAsc() (value bool)
	GetScheduleStartSubscribed() (value bool)
	GetCanStartVideo() (value bool)
	GetRecordVideoActive() (value bool)
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetParticipantsCount() (value int)
	GetTitle() (value string, ok bool)
	GetStreamDCID() (value int, ok bool)
	GetRecordStartDate() (value int, ok bool)
	GetScheduleDate() (value int, ok bool)
	GetUnmutedVideoCount() (value int, ok bool)
	GetUnmutedVideoLimit() (value int)
	GetVersion() (value int)
}) {
	g.JoinMuted = from.GetJoinMuted()
	g.CanChangeJoinMuted = from.GetCanChangeJoinMuted()
	g.JoinDateAsc = from.GetJoinDateAsc()
	g.ScheduleStartSubscribed = from.GetScheduleStartSubscribed()
	g.CanStartVideo = from.GetCanStartVideo()
	g.RecordVideoActive = from.GetRecordVideoActive()
	g.ID = from.GetID()
	g.AccessHash = from.GetAccessHash()
	g.ParticipantsCount = from.GetParticipantsCount()
	if val, ok := from.GetTitle(); ok {
		g.Title = val
	}

	if val, ok := from.GetStreamDCID(); ok {
		g.StreamDCID = val
	}

	if val, ok := from.GetRecordStartDate(); ok {
		g.RecordStartDate = val
	}

	if val, ok := from.GetScheduleDate(); ok {
		g.ScheduleDate = val
	}

	if val, ok := from.GetUnmutedVideoCount(); ok {
		g.UnmutedVideoCount = val
	}

	g.UnmutedVideoLimit = from.GetUnmutedVideoLimit()
	g.Version = from.GetVersion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCall) TypeID() uint32 {
	return GroupCallTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCall) TypeName() string {
	return "groupCall"
}

// TypeInfo returns info about TL type.
func (g *GroupCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCall",
		ID:   GroupCallTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JoinMuted",
			SchemaName: "join_muted",
			Null:       !g.Flags.Has(1),
		},
		{
			Name:       "CanChangeJoinMuted",
			SchemaName: "can_change_join_muted",
			Null:       !g.Flags.Has(2),
		},
		{
			Name:       "JoinDateAsc",
			SchemaName: "join_date_asc",
			Null:       !g.Flags.Has(6),
		},
		{
			Name:       "ScheduleStartSubscribed",
			SchemaName: "schedule_start_subscribed",
			Null:       !g.Flags.Has(8),
		},
		{
			Name:       "CanStartVideo",
			SchemaName: "can_start_video",
			Null:       !g.Flags.Has(9),
		},
		{
			Name:       "RecordVideoActive",
			SchemaName: "record_video_active",
			Null:       !g.Flags.Has(11),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "ParticipantsCount",
			SchemaName: "participants_count",
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !g.Flags.Has(3),
		},
		{
			Name:       "StreamDCID",
			SchemaName: "stream_dc_id",
			Null:       !g.Flags.Has(4),
		},
		{
			Name:       "RecordStartDate",
			SchemaName: "record_start_date",
			Null:       !g.Flags.Has(5),
		},
		{
			Name:       "ScheduleDate",
			SchemaName: "schedule_date",
			Null:       !g.Flags.Has(7),
		},
		{
			Name:       "UnmutedVideoCount",
			SchemaName: "unmuted_video_count",
			Null:       !g.Flags.Has(10),
		},
		{
			Name:       "UnmutedVideoLimit",
			SchemaName: "unmuted_video_limit",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GroupCall) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCall#d597650c as nil")
	}
	b.PutID(GroupCallTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCall) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCall#d597650c as nil")
	}
	if !(g.JoinMuted == false) {
		g.Flags.Set(1)
	}
	if !(g.CanChangeJoinMuted == false) {
		g.Flags.Set(2)
	}
	if !(g.JoinDateAsc == false) {
		g.Flags.Set(6)
	}
	if !(g.ScheduleStartSubscribed == false) {
		g.Flags.Set(8)
	}
	if !(g.CanStartVideo == false) {
		g.Flags.Set(9)
	}
	if !(g.RecordVideoActive == false) {
		g.Flags.Set(11)
	}
	if !(g.Title == "") {
		g.Flags.Set(3)
	}
	if !(g.StreamDCID == 0) {
		g.Flags.Set(4)
	}
	if !(g.RecordStartDate == 0) {
		g.Flags.Set(5)
	}
	if !(g.ScheduleDate == 0) {
		g.Flags.Set(7)
	}
	if !(g.UnmutedVideoCount == 0) {
		g.Flags.Set(10)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode groupCall#d597650c: field flags: %w", err)
	}
	b.PutLong(g.ID)
	b.PutLong(g.AccessHash)
	b.PutInt(g.ParticipantsCount)
	if g.Flags.Has(3) {
		b.PutString(g.Title)
	}
	if g.Flags.Has(4) {
		b.PutInt(g.StreamDCID)
	}
	if g.Flags.Has(5) {
		b.PutInt(g.RecordStartDate)
	}
	if g.Flags.Has(7) {
		b.PutInt(g.ScheduleDate)
	}
	if g.Flags.Has(10) {
		b.PutInt(g.UnmutedVideoCount)
	}
	b.PutInt(g.UnmutedVideoLimit)
	b.PutInt(g.Version)
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCall) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCall#d597650c to nil")
	}
	if err := b.ConsumeID(GroupCallTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCall#d597650c: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCall) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCall#d597650c to nil")
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field flags: %w", err)
		}
	}
	g.JoinMuted = g.Flags.Has(1)
	g.CanChangeJoinMuted = g.Flags.Has(2)
	g.JoinDateAsc = g.Flags.Has(6)
	g.ScheduleStartSubscribed = g.Flags.Has(8)
	g.CanStartVideo = g.Flags.Has(9)
	g.RecordVideoActive = g.Flags.Has(11)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field access_hash: %w", err)
		}
		g.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field participants_count: %w", err)
		}
		g.ParticipantsCount = value
	}
	if g.Flags.Has(3) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field title: %w", err)
		}
		g.Title = value
	}
	if g.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field stream_dc_id: %w", err)
		}
		g.StreamDCID = value
	}
	if g.Flags.Has(5) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field record_start_date: %w", err)
		}
		g.RecordStartDate = value
	}
	if g.Flags.Has(7) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field schedule_date: %w", err)
		}
		g.ScheduleDate = value
	}
	if g.Flags.Has(10) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field unmuted_video_count: %w", err)
		}
		g.UnmutedVideoCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field unmuted_video_limit: %w", err)
		}
		g.UnmutedVideoLimit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode groupCall#d597650c: field version: %w", err)
		}
		g.Version = value
	}
	return nil
}

// SetJoinMuted sets value of JoinMuted conditional field.
func (g *GroupCall) SetJoinMuted(value bool) {
	if value {
		g.Flags.Set(1)
		g.JoinMuted = true
	} else {
		g.Flags.Unset(1)
		g.JoinMuted = false
	}
}

// GetJoinMuted returns value of JoinMuted conditional field.
func (g *GroupCall) GetJoinMuted() (value bool) {
	return g.Flags.Has(1)
}

// SetCanChangeJoinMuted sets value of CanChangeJoinMuted conditional field.
func (g *GroupCall) SetCanChangeJoinMuted(value bool) {
	if value {
		g.Flags.Set(2)
		g.CanChangeJoinMuted = true
	} else {
		g.Flags.Unset(2)
		g.CanChangeJoinMuted = false
	}
}

// GetCanChangeJoinMuted returns value of CanChangeJoinMuted conditional field.
func (g *GroupCall) GetCanChangeJoinMuted() (value bool) {
	return g.Flags.Has(2)
}

// SetJoinDateAsc sets value of JoinDateAsc conditional field.
func (g *GroupCall) SetJoinDateAsc(value bool) {
	if value {
		g.Flags.Set(6)
		g.JoinDateAsc = true
	} else {
		g.Flags.Unset(6)
		g.JoinDateAsc = false
	}
}

// GetJoinDateAsc returns value of JoinDateAsc conditional field.
func (g *GroupCall) GetJoinDateAsc() (value bool) {
	return g.Flags.Has(6)
}

// SetScheduleStartSubscribed sets value of ScheduleStartSubscribed conditional field.
func (g *GroupCall) SetScheduleStartSubscribed(value bool) {
	if value {
		g.Flags.Set(8)
		g.ScheduleStartSubscribed = true
	} else {
		g.Flags.Unset(8)
		g.ScheduleStartSubscribed = false
	}
}

// GetScheduleStartSubscribed returns value of ScheduleStartSubscribed conditional field.
func (g *GroupCall) GetScheduleStartSubscribed() (value bool) {
	return g.Flags.Has(8)
}

// SetCanStartVideo sets value of CanStartVideo conditional field.
func (g *GroupCall) SetCanStartVideo(value bool) {
	if value {
		g.Flags.Set(9)
		g.CanStartVideo = true
	} else {
		g.Flags.Unset(9)
		g.CanStartVideo = false
	}
}

// GetCanStartVideo returns value of CanStartVideo conditional field.
func (g *GroupCall) GetCanStartVideo() (value bool) {
	return g.Flags.Has(9)
}

// SetRecordVideoActive sets value of RecordVideoActive conditional field.
func (g *GroupCall) SetRecordVideoActive(value bool) {
	if value {
		g.Flags.Set(11)
		g.RecordVideoActive = true
	} else {
		g.Flags.Unset(11)
		g.RecordVideoActive = false
	}
}

// GetRecordVideoActive returns value of RecordVideoActive conditional field.
func (g *GroupCall) GetRecordVideoActive() (value bool) {
	return g.Flags.Has(11)
}

// GetID returns value of ID field.
func (g *GroupCall) GetID() (value int64) {
	return g.ID
}

// GetAccessHash returns value of AccessHash field.
func (g *GroupCall) GetAccessHash() (value int64) {
	return g.AccessHash
}

// GetParticipantsCount returns value of ParticipantsCount field.
func (g *GroupCall) GetParticipantsCount() (value int) {
	return g.ParticipantsCount
}

// SetTitle sets value of Title conditional field.
func (g *GroupCall) SetTitle(value string) {
	g.Flags.Set(3)
	g.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (g *GroupCall) GetTitle() (value string, ok bool) {
	if !g.Flags.Has(3) {
		return value, false
	}
	return g.Title, true
}

// SetStreamDCID sets value of StreamDCID conditional field.
func (g *GroupCall) SetStreamDCID(value int) {
	g.Flags.Set(4)
	g.StreamDCID = value
}

// GetStreamDCID returns value of StreamDCID conditional field and
// boolean which is true if field was set.
func (g *GroupCall) GetStreamDCID() (value int, ok bool) {
	if !g.Flags.Has(4) {
		return value, false
	}
	return g.StreamDCID, true
}

// SetRecordStartDate sets value of RecordStartDate conditional field.
func (g *GroupCall) SetRecordStartDate(value int) {
	g.Flags.Set(5)
	g.RecordStartDate = value
}

// GetRecordStartDate returns value of RecordStartDate conditional field and
// boolean which is true if field was set.
func (g *GroupCall) GetRecordStartDate() (value int, ok bool) {
	if !g.Flags.Has(5) {
		return value, false
	}
	return g.RecordStartDate, true
}

// SetScheduleDate sets value of ScheduleDate conditional field.
func (g *GroupCall) SetScheduleDate(value int) {
	g.Flags.Set(7)
	g.ScheduleDate = value
}

// GetScheduleDate returns value of ScheduleDate conditional field and
// boolean which is true if field was set.
func (g *GroupCall) GetScheduleDate() (value int, ok bool) {
	if !g.Flags.Has(7) {
		return value, false
	}
	return g.ScheduleDate, true
}

// SetUnmutedVideoCount sets value of UnmutedVideoCount conditional field.
func (g *GroupCall) SetUnmutedVideoCount(value int) {
	g.Flags.Set(10)
	g.UnmutedVideoCount = value
}

// GetUnmutedVideoCount returns value of UnmutedVideoCount conditional field and
// boolean which is true if field was set.
func (g *GroupCall) GetUnmutedVideoCount() (value int, ok bool) {
	if !g.Flags.Has(10) {
		return value, false
	}
	return g.UnmutedVideoCount, true
}

// GetUnmutedVideoLimit returns value of UnmutedVideoLimit field.
func (g *GroupCall) GetUnmutedVideoLimit() (value int) {
	return g.UnmutedVideoLimit
}

// GetVersion returns value of Version field.
func (g *GroupCall) GetVersion() (value int) {
	return g.Version
}

// GroupCallClass represents GroupCall generic type.
//
// See https://core.telegram.org/type/GroupCall for reference.
//
// Example:
//  g, err := tg.DecodeGroupCall(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.GroupCallDiscarded: // groupCallDiscarded#7780bcb4
//  case *tg.GroupCall: // groupCall#d597650c
//  default: panic(v)
//  }
type GroupCallClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() GroupCallClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Group call ID
	GetID() (value int64)

	// Group call access hash
	GetAccessHash() (value int64)
}

// AsInput tries to map GroupCall to InputGroupCall.
func (g *GroupCall) AsInput() *InputGroupCall {
	value := new(InputGroupCall)
	value.ID = g.GetID()
	value.AccessHash = g.GetAccessHash()

	return value
}

// DecodeGroupCall implements binary de-serialization for GroupCallClass.
func DecodeGroupCall(buf *bin.Buffer) (GroupCallClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case GroupCallDiscardedTypeID:
		// Decoding groupCallDiscarded#7780bcb4.
		v := GroupCallDiscarded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GroupCallClass: %w", err)
		}
		return &v, nil
	case GroupCallTypeID:
		// Decoding groupCall#d597650c.
		v := GroupCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode GroupCallClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode GroupCallClass: %w", bin.NewUnexpectedID(id))
	}
}

// GroupCall boxes the GroupCallClass providing a helper.
type GroupCallBox struct {
	GroupCall GroupCallClass
}

// Decode implements bin.Decoder for GroupCallBox.
func (b *GroupCallBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode GroupCallBox to nil")
	}
	v, err := DecodeGroupCall(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.GroupCall = v
	return nil
}

// Encode implements bin.Encode for GroupCallBox.
func (b *GroupCallBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.GroupCall == nil {
		return fmt.Errorf("unable to encode GroupCallClass as nil")
	}
	return b.GroupCall.Encode(buf)
}
