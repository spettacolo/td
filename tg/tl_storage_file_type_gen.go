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

// StorageFileUnknown represents TL type `storage.fileUnknown#aa963b05`.
// Unknown type.
//
// See https://core.telegram.org/constructor/storage.fileUnknown for reference.
type StorageFileUnknown struct {
}

// StorageFileUnknownTypeID is TL type id of StorageFileUnknown.
const StorageFileUnknownTypeID = 0xaa963b05

func (f *StorageFileUnknown) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileUnknown) String() string {
	if f == nil {
		return "StorageFileUnknown(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileUnknown")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileUnknown) TypeID() uint32 {
	return StorageFileUnknownTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileUnknown) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileUnknown#aa963b05 as nil")
	}
	b.PutID(StorageFileUnknownTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileUnknown) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileUnknown#aa963b05 to nil")
	}
	if err := b.ConsumeID(StorageFileUnknownTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileUnknown#aa963b05: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileUnknown) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileUnknown.
var (
	_ bin.Encoder = &StorageFileUnknown{}
	_ bin.Decoder = &StorageFileUnknown{}

	_ StorageFileTypeClass = &StorageFileUnknown{}
)

// StorageFilePartial represents TL type `storage.filePartial#40bc6f52`.
// Part of a bigger file.
//
// See https://core.telegram.org/constructor/storage.filePartial for reference.
type StorageFilePartial struct {
}

// StorageFilePartialTypeID is TL type id of StorageFilePartial.
const StorageFilePartialTypeID = 0x40bc6f52

func (f *StorageFilePartial) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFilePartial) String() string {
	if f == nil {
		return "StorageFilePartial(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFilePartial")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFilePartial) TypeID() uint32 {
	return StorageFilePartialTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFilePartial) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePartial#40bc6f52 as nil")
	}
	b.PutID(StorageFilePartialTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFilePartial) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePartial#40bc6f52 to nil")
	}
	if err := b.ConsumeID(StorageFilePartialTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.filePartial#40bc6f52: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFilePartial) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFilePartial.
var (
	_ bin.Encoder = &StorageFilePartial{}
	_ bin.Decoder = &StorageFilePartial{}

	_ StorageFileTypeClass = &StorageFilePartial{}
)

// StorageFileJpeg represents TL type `storage.fileJpeg#7efe0e`.
// JPEG image. MIME type: image/jpeg.
//
// See https://core.telegram.org/constructor/storage.fileJpeg for reference.
type StorageFileJpeg struct {
}

// StorageFileJpegTypeID is TL type id of StorageFileJpeg.
const StorageFileJpegTypeID = 0x7efe0e

func (f *StorageFileJpeg) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileJpeg) String() string {
	if f == nil {
		return "StorageFileJpeg(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileJpeg")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileJpeg) TypeID() uint32 {
	return StorageFileJpegTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileJpeg) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileJpeg#7efe0e as nil")
	}
	b.PutID(StorageFileJpegTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileJpeg) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileJpeg#7efe0e to nil")
	}
	if err := b.ConsumeID(StorageFileJpegTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileJpeg#7efe0e: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileJpeg) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileJpeg.
var (
	_ bin.Encoder = &StorageFileJpeg{}
	_ bin.Decoder = &StorageFileJpeg{}

	_ StorageFileTypeClass = &StorageFileJpeg{}
)

// StorageFileGif represents TL type `storage.fileGif#cae1aadf`.
// GIF image. MIME type: image/gif.
//
// See https://core.telegram.org/constructor/storage.fileGif for reference.
type StorageFileGif struct {
}

// StorageFileGifTypeID is TL type id of StorageFileGif.
const StorageFileGifTypeID = 0xcae1aadf

func (f *StorageFileGif) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileGif) String() string {
	if f == nil {
		return "StorageFileGif(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileGif")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileGif) TypeID() uint32 {
	return StorageFileGifTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileGif) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileGif#cae1aadf as nil")
	}
	b.PutID(StorageFileGifTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileGif) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileGif#cae1aadf to nil")
	}
	if err := b.ConsumeID(StorageFileGifTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileGif#cae1aadf: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileGif) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileGif.
var (
	_ bin.Encoder = &StorageFileGif{}
	_ bin.Decoder = &StorageFileGif{}

	_ StorageFileTypeClass = &StorageFileGif{}
)

// StorageFilePng represents TL type `storage.filePng#a4f63c0`.
// PNG image. MIME type: image/png.
//
// See https://core.telegram.org/constructor/storage.filePng for reference.
type StorageFilePng struct {
}

// StorageFilePngTypeID is TL type id of StorageFilePng.
const StorageFilePngTypeID = 0xa4f63c0

func (f *StorageFilePng) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFilePng) String() string {
	if f == nil {
		return "StorageFilePng(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFilePng")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFilePng) TypeID() uint32 {
	return StorageFilePngTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFilePng) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePng#a4f63c0 as nil")
	}
	b.PutID(StorageFilePngTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFilePng) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePng#a4f63c0 to nil")
	}
	if err := b.ConsumeID(StorageFilePngTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.filePng#a4f63c0: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFilePng) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFilePng.
var (
	_ bin.Encoder = &StorageFilePng{}
	_ bin.Decoder = &StorageFilePng{}

	_ StorageFileTypeClass = &StorageFilePng{}
)

// StorageFilePdf represents TL type `storage.filePdf#ae1e508d`.
// PDF document image. MIME type: application/pdf.
//
// See https://core.telegram.org/constructor/storage.filePdf for reference.
type StorageFilePdf struct {
}

// StorageFilePdfTypeID is TL type id of StorageFilePdf.
const StorageFilePdfTypeID = 0xae1e508d

func (f *StorageFilePdf) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFilePdf) String() string {
	if f == nil {
		return "StorageFilePdf(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFilePdf")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFilePdf) TypeID() uint32 {
	return StorageFilePdfTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFilePdf) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.filePdf#ae1e508d as nil")
	}
	b.PutID(StorageFilePdfTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFilePdf) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.filePdf#ae1e508d to nil")
	}
	if err := b.ConsumeID(StorageFilePdfTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.filePdf#ae1e508d: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFilePdf) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFilePdf.
var (
	_ bin.Encoder = &StorageFilePdf{}
	_ bin.Decoder = &StorageFilePdf{}

	_ StorageFileTypeClass = &StorageFilePdf{}
)

// StorageFileMp3 represents TL type `storage.fileMp3#528a0677`.
// Mp3 audio. MIME type: audio/mpeg.
//
// See https://core.telegram.org/constructor/storage.fileMp3 for reference.
type StorageFileMp3 struct {
}

// StorageFileMp3TypeID is TL type id of StorageFileMp3.
const StorageFileMp3TypeID = 0x528a0677

func (f *StorageFileMp3) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileMp3) String() string {
	if f == nil {
		return "StorageFileMp3(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileMp3")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileMp3) TypeID() uint32 {
	return StorageFileMp3TypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileMp3) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMp3#528a0677 as nil")
	}
	b.PutID(StorageFileMp3TypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileMp3) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMp3#528a0677 to nil")
	}
	if err := b.ConsumeID(StorageFileMp3TypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileMp3#528a0677: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileMp3) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileMp3.
var (
	_ bin.Encoder = &StorageFileMp3{}
	_ bin.Decoder = &StorageFileMp3{}

	_ StorageFileTypeClass = &StorageFileMp3{}
)

// StorageFileMov represents TL type `storage.fileMov#4b09ebbc`.
// Quicktime video. MIME type: video/quicktime.
//
// See https://core.telegram.org/constructor/storage.fileMov for reference.
type StorageFileMov struct {
}

// StorageFileMovTypeID is TL type id of StorageFileMov.
const StorageFileMovTypeID = 0x4b09ebbc

func (f *StorageFileMov) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileMov) String() string {
	if f == nil {
		return "StorageFileMov(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileMov")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileMov) TypeID() uint32 {
	return StorageFileMovTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileMov) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMov#4b09ebbc as nil")
	}
	b.PutID(StorageFileMovTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileMov) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMov#4b09ebbc to nil")
	}
	if err := b.ConsumeID(StorageFileMovTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileMov#4b09ebbc: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileMov) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileMov.
var (
	_ bin.Encoder = &StorageFileMov{}
	_ bin.Decoder = &StorageFileMov{}

	_ StorageFileTypeClass = &StorageFileMov{}
)

// StorageFileMp4 represents TL type `storage.fileMp4#b3cea0e4`.
// MPEG-4 video. MIME type: video/mp4.
//
// See https://core.telegram.org/constructor/storage.fileMp4 for reference.
type StorageFileMp4 struct {
}

// StorageFileMp4TypeID is TL type id of StorageFileMp4.
const StorageFileMp4TypeID = 0xb3cea0e4

func (f *StorageFileMp4) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileMp4) String() string {
	if f == nil {
		return "StorageFileMp4(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileMp4")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileMp4) TypeID() uint32 {
	return StorageFileMp4TypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileMp4) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileMp4#b3cea0e4 as nil")
	}
	b.PutID(StorageFileMp4TypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileMp4) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileMp4#b3cea0e4 to nil")
	}
	if err := b.ConsumeID(StorageFileMp4TypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileMp4#b3cea0e4: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileMp4) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileMp4.
var (
	_ bin.Encoder = &StorageFileMp4{}
	_ bin.Decoder = &StorageFileMp4{}

	_ StorageFileTypeClass = &StorageFileMp4{}
)

// StorageFileWebp represents TL type `storage.fileWebp#1081464c`.
// WEBP image. MIME type: image/webp.
//
// See https://core.telegram.org/constructor/storage.fileWebp for reference.
type StorageFileWebp struct {
}

// StorageFileWebpTypeID is TL type id of StorageFileWebp.
const StorageFileWebpTypeID = 0x1081464c

func (f *StorageFileWebp) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *StorageFileWebp) String() string {
	if f == nil {
		return "StorageFileWebp(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StorageFileWebp")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (f *StorageFileWebp) TypeID() uint32 {
	return StorageFileWebpTypeID
}

// Encode implements bin.Encoder.
func (f *StorageFileWebp) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode storage.fileWebp#1081464c as nil")
	}
	b.PutID(StorageFileWebpTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (f *StorageFileWebp) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode storage.fileWebp#1081464c to nil")
	}
	if err := b.ConsumeID(StorageFileWebpTypeID); err != nil {
		return fmt.Errorf("unable to decode storage.fileWebp#1081464c: %w", err)
	}
	return nil
}

// construct implements constructor of StorageFileTypeClass.
func (f StorageFileWebp) construct() StorageFileTypeClass { return &f }

// Ensuring interfaces in compile-time for StorageFileWebp.
var (
	_ bin.Encoder = &StorageFileWebp{}
	_ bin.Decoder = &StorageFileWebp{}

	_ StorageFileTypeClass = &StorageFileWebp{}
)

// StorageFileTypeClass represents storage.FileType generic type.
//
// See https://core.telegram.org/type/storage.FileType for reference.
//
// Example:
//  g, err := DecodeStorageFileType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *StorageFileUnknown: // storage.fileUnknown#aa963b05
//  case *StorageFilePartial: // storage.filePartial#40bc6f52
//  case *StorageFileJpeg: // storage.fileJpeg#7efe0e
//  case *StorageFileGif: // storage.fileGif#cae1aadf
//  case *StorageFilePng: // storage.filePng#a4f63c0
//  case *StorageFilePdf: // storage.filePdf#ae1e508d
//  case *StorageFileMp3: // storage.fileMp3#528a0677
//  case *StorageFileMov: // storage.fileMov#4b09ebbc
//  case *StorageFileMp4: // storage.fileMp4#b3cea0e4
//  case *StorageFileWebp: // storage.fileWebp#1081464c
//  default: panic(v)
//  }
type StorageFileTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() StorageFileTypeClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeStorageFileType implements binary de-serialization for StorageFileTypeClass.
func DecodeStorageFileType(buf *bin.Buffer) (StorageFileTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StorageFileUnknownTypeID:
		// Decoding storage.fileUnknown#aa963b05.
		v := StorageFileUnknown{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFilePartialTypeID:
		// Decoding storage.filePartial#40bc6f52.
		v := StorageFilePartial{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileJpegTypeID:
		// Decoding storage.fileJpeg#7efe0e.
		v := StorageFileJpeg{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileGifTypeID:
		// Decoding storage.fileGif#cae1aadf.
		v := StorageFileGif{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFilePngTypeID:
		// Decoding storage.filePng#a4f63c0.
		v := StorageFilePng{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFilePdfTypeID:
		// Decoding storage.filePdf#ae1e508d.
		v := StorageFilePdf{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileMp3TypeID:
		// Decoding storage.fileMp3#528a0677.
		v := StorageFileMp3{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileMovTypeID:
		// Decoding storage.fileMov#4b09ebbc.
		v := StorageFileMov{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileMp4TypeID:
		// Decoding storage.fileMp4#b3cea0e4.
		v := StorageFileMp4{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	case StorageFileWebpTypeID:
		// Decoding storage.fileWebp#1081464c.
		v := StorageFileWebp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StorageFileTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// StorageFileType boxes the StorageFileTypeClass providing a helper.
type StorageFileTypeBox struct {
	FileType StorageFileTypeClass
}

// Decode implements bin.Decoder for StorageFileTypeBox.
func (b *StorageFileTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StorageFileTypeBox to nil")
	}
	v, err := DecodeStorageFileType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FileType = v
	return nil
}

// Encode implements bin.Encode for StorageFileTypeBox.
func (b *StorageFileTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FileType == nil {
		return fmt.Errorf("unable to encode StorageFileTypeClass as nil")
	}
	return b.FileType.Encode(buf)
}
