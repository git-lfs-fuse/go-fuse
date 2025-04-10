//go:build !linux

// Copyright 2024 the Go-FUSE Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fuse

import (
	"syscall"
)

func (a *Attr) FromStat(s *syscall.Stat_t) {
	a.Ino = uint64(s.Ino)
	a.Size = uint64(s.Size)
	a.Blocks = uint64(s.Blocks)
	a.Atime = uint64(s.Atimespec.Sec)
	a.Atimensec = uint32(s.Atimespec.Nsec)
	a.Mtime = uint64(s.Mtimespec.Sec)
	a.Mtimensec = uint32(s.Mtimespec.Nsec)
	a.Ctime = uint64(s.Ctimespec.Sec)
	a.Ctimensec = uint32(s.Ctimespec.Nsec)
	a.Crtime_ = uint64(s.Birthtimespec.Sec)
	a.Crtimensec_ = uint32(s.Birthtimespec.Nsec)
	a.Mode = uint32(s.Mode)
	a.Nlink = uint32(s.Nlink)
	a.Uid = uint32(s.Uid)
	a.Gid = uint32(s.Gid)
	a.Rdev = uint32(s.Rdev)
	a.Blksize = uint32(s.Blksize)
	a.Flags_ = uint32(s.Flags)
}
