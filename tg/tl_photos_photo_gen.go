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

// PhotosPhoto represents TL type `photos.photo#20212ca8`.
// Photo with auxiliary data.
//
// See https://core.telegram.org/constructor/photos.photo for reference.
type PhotosPhoto struct {
	// Photo
	Photo PhotoClass
	// Users
	Users []UserClass
}

// PhotosPhotoTypeID is TL type id of PhotosPhoto.
const PhotosPhotoTypeID = 0x20212ca8

func (p *PhotosPhoto) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Photo == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotosPhoto) String() string {
	if p == nil {
		return "PhotosPhoto(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhotosPhoto")
	sb.WriteString("{\n")
	sb.WriteString("\tPhoto: ")
	sb.WriteString(fmt.Sprint(p.Photo))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range p.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (p *PhotosPhoto) TypeID() uint32 {
	return PhotosPhotoTypeID
}

// Encode implements bin.Encoder.
func (p *PhotosPhoto) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photo#20212ca8 as nil")
	}
	b.PutID(PhotosPhotoTypeID)
	if p.Photo == nil {
		return fmt.Errorf("unable to encode photos.photo#20212ca8: field photo is nil")
	}
	if err := p.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.photo#20212ca8: field photo: %w", err)
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photo#20212ca8: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photo#20212ca8: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPhoto returns value of Photo field.
func (p *PhotosPhoto) GetPhoto() (value PhotoClass) {
	return p.Photo
}

// GetUsers returns value of Users field.
func (p *PhotosPhoto) GetUsers() (value []UserClass) {
	return p.Users
}

// Decode implements bin.Decoder.
func (p *PhotosPhoto) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photo#20212ca8 to nil")
	}
	if err := b.ConsumeID(PhotosPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.photo#20212ca8: %w", err)
	}
	{
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.photo#20212ca8: field photo: %w", err)
		}
		p.Photo = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photo#20212ca8: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photo#20212ca8: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhotosPhoto.
var (
	_ bin.Encoder = &PhotosPhoto{}
	_ bin.Decoder = &PhotosPhoto{}
)
