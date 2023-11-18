package main

import (
	"os"
	"os/exec"
	"syscall"
	"time"
	"unsafe"

	"github.com/lxn/walk"
)

var ClashStatus string
var TunStatus string

var MainWindow *walk.MainWindow
var Icon *walk.Icon
var NotifyIcon *walk.NotifyIcon
var IcoData []byte = []byte{
	0x00, 0x00, 0x01, 0x00, 0x01, 0x00, 0x20, 0x20, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0xa8, 0x10, 0x00, 0x00, 0x16, 0x00, 0x00, 0x00, 0x28, 0x00,
	0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x40, 0x00, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x04, 0x00, 0x98, 0xf3, 0x02, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x1a, 0x29, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1a,
	0x29, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x02, 0x00, 0x98, 0xf3, 0x04, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x56, 0x00, 0x98, 0xf3, 0x9b, 0x00, 0x98, 0xf3, 0x57, 0x00, 0x98,
	0xf3, 0x1d, 0x00, 0x98, 0xf3, 0x01, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x01, 0x00, 0x98,
	0xf3, 0x1d, 0x00, 0x98, 0xf3, 0x57, 0x00, 0x98, 0xf3, 0x9b, 0x00, 0x98,
	0xf3, 0x56, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x1a, 0x00, 0x98,
	0xf3, 0x62, 0x00, 0x98, 0xf3, 0xa3, 0x00, 0x98, 0xf3, 0xba, 0x00, 0x98,
	0xf3, 0x96, 0x00, 0x98, 0xf3, 0x52, 0x00, 0x98, 0xf3, 0x19, 0x00, 0x98,
	0xf3, 0x01, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x01, 0x00, 0x98, 0xf3, 0x19, 0x00, 0x98,
	0xf3, 0x52, 0x00, 0x98, 0xf3, 0x96, 0x00, 0x98, 0xf3, 0xba, 0x00, 0x98,
	0xf3, 0xa3, 0x00, 0x98, 0xf3, 0x62, 0x00, 0x98, 0xf3, 0x1a, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x04, 0x00, 0x98, 0xf3, 0x28, 0x00, 0x98, 0xf3, 0x67, 0x00, 0x98,
	0xf3, 0xa6, 0x00, 0x98, 0xf3, 0xba, 0x00, 0x98, 0xf3, 0x70, 0x00, 0x98,
	0xf3, 0x03, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x04, 0x02,
	0x01, 0x00, 0x8c, 0x3d, 0x14, 0x00, 0x8c, 0x3d, 0x14, 0x00, 0x8c, 0x3d,
	0x14, 0x00, 0x8c, 0x3d, 0x14, 0x00, 0x90, 0x3d, 0x15, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x03, 0x00, 0x98,
	0xf3, 0x70, 0x00, 0x98, 0xf3, 0xba, 0x00, 0x98, 0xf3, 0xa6, 0x00, 0x98,
	0xf3, 0x67, 0x00, 0x98, 0xf3, 0x28, 0x00, 0x98, 0xf3, 0x04, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x01, 0x00, 0x98,
	0xf3, 0x29, 0x00, 0x98, 0xf3, 0x38, 0x00, 0x98, 0xf3, 0x02, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x99, 0x4c, 0x1e, 0x00, 0x93, 0x45,
	0x1b, 0x00, 0x95, 0x47, 0x1b, 0x00, 0x93, 0x45, 0x1a, 0x00, 0x8b, 0x3a,
	0x12, 0x00, 0x92, 0x44, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x02, 0x00, 0x98, 0xf3, 0x37, 0x00, 0x98,
	0xf3, 0x29, 0x00, 0x98, 0xf3, 0x01, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x3d, 0x00, 0x98,
	0xf3, 0x6b, 0x00, 0x98, 0xf3, 0x69, 0x00, 0x98, 0xf3, 0x69, 0x00, 0x98,
	0xf3, 0x69, 0x00, 0x98, 0xf3, 0x69, 0x00, 0x98, 0xf3, 0x68, 0x00, 0x98,
	0xf3, 0x47, 0x00, 0x98, 0xf3, 0x03, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xe2, 0xab, 0x56, 0x00, 0x90, 0x42, 0x18, 0x00, 0x92, 0x43,
	0x18, 0x37, 0x90, 0x42, 0x18, 0x44, 0x92, 0x44, 0x19, 0x00, 0x8d, 0x3e,
	0x15, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x02, 0x00, 0x98, 0xf3, 0x46, 0x00, 0x98, 0xf3, 0x68, 0x00, 0x98,
	0xf3, 0x69, 0x00, 0x98, 0xf3, 0x69, 0x00, 0x98, 0xf3, 0x69, 0x00, 0x98,
	0xf3, 0x69, 0x00, 0x98, 0xf3, 0x6b, 0x00, 0x98, 0xf3, 0x3d, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf4, 0x00, 0x00, 0x98, 0xf3, 0x50, 0x00, 0x98, 0xf4, 0x8b, 0x00, 0x99,
	0xf6, 0x88, 0x00, 0x9a, 0xf8, 0x86, 0x00, 0x9a, 0xf9, 0x86, 0x00, 0x99,
	0xf5, 0x89, 0x00, 0x98, 0xf3, 0x8c, 0x00, 0x98, 0xf3, 0x61, 0x00, 0x98,
	0xf3, 0x04, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9f, 0x54,
	0x22, 0x00, 0xa2, 0x58, 0x24, 0x05, 0x9a, 0x4e, 0x1f, 0xa9, 0x97, 0x4a,
	0x1d, 0xc1, 0x96, 0x48, 0x1c, 0x0e, 0x97, 0x49, 0x1c, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x03, 0x00, 0x98,
	0xf3, 0x60, 0x00, 0x98, 0xf3, 0x8c, 0x00, 0x99, 0xf5, 0x88, 0x00, 0x9a,
	0xf9, 0x86, 0x00, 0x9a, 0xf9, 0x86, 0x00, 0x99, 0xf6, 0x88, 0x00, 0x98,
	0xf4, 0x8b, 0x00, 0x98, 0xf3, 0x50, 0x00, 0x98, 0xf4, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x14, 0x92, 0xdc, 0x00, 0x57, 0x7b, 0x8b, 0x00, 0x31, 0x8a,
	0xbb, 0x00, 0xff, 0x3b, 0x00, 0x05, 0xb5, 0x56, 0x14, 0x1b, 0xaa, 0x55,
	0x1d, 0x3a, 0xa7, 0x52, 0x1c, 0x42, 0xb3, 0x4a, 0x08, 0x14, 0x30, 0x83,
	0xb4, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa0, 0x55, 0x23, 0x00, 0xa1, 0x56,
	0x24, 0x03, 0x9f, 0x55, 0x23, 0x35, 0x9c, 0x51, 0x21, 0x3b, 0x98, 0x4b,
	0x1d, 0x06, 0x99, 0x4c, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x2e, 0x7f,
	0xb0, 0x00, 0xa6, 0x3d, 0x00, 0x17, 0x9b, 0x43, 0x11, 0x43, 0x9c, 0x43,
	0x0f, 0x38, 0xa4, 0x3d, 0x03, 0x19, 0xf9, 0x0a, 0x00, 0x04, 0x2c, 0x81,
	0xb4, 0x00, 0x34, 0x78, 0xa4, 0x00, 0x06, 0x95, 0xea, 0x00, 0xab, 0x63,
	0x2c, 0x00, 0xaa, 0x61, 0x2a, 0x00, 0xaa, 0x62, 0x2b, 0x51, 0xa9, 0x60,
	0x2a, 0xb8, 0xa7, 0x5e, 0x29, 0xdc, 0xa5, 0x5b, 0x27, 0xf3, 0xa2, 0x57,
	0x25, 0xf7, 0xa1, 0x55, 0x24, 0x59, 0xa1, 0x56, 0x25, 0x00, 0x00, 0x9a,
	0xf8, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xa0, 0x56, 0x23, 0x00, 0xa1, 0x57, 0x24, 0x00, 0x9f, 0x54,
	0x23, 0x00, 0x9c, 0x50, 0x20, 0x00, 0x98, 0x4b, 0x1e, 0x00, 0x99, 0x4c,
	0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x98, 0xf3, 0x00, 0x00, 0x98,
	0xf3, 0x00, 0x00, 0x9e, 0xff, 0x00, 0x97, 0x49, 0x1c, 0x00, 0x97, 0x49,
	0x1b, 0x63, 0x97, 0x49, 0x1b, 0xf8, 0x97, 0x49, 0x1b, 0xf3, 0x96, 0x48,
	0x1a, 0xdb, 0x95, 0x46, 0x1a, 0xb6, 0x95, 0x46, 0x1a, 0x4a, 0x95, 0x46,
	0x1a, 0x00, 0x95, 0x45, 0x1a, 0x00, 0xac, 0x65, 0x2d, 0x00, 0xaa, 0x62,
	0x2b, 0x00, 0xac, 0x64, 0x2c, 0x91, 0xab, 0x62, 0x2b, 0xff, 0xa9, 0x60,
	0x2a, 0xff, 0xa7, 0x5e, 0x29, 0xff, 0xa6, 0x5c, 0x28, 0xff, 0xa5, 0x5a,
	0x27, 0x5d, 0xa5, 0x5a, 0x27, 0x00, 0xa4, 0x59, 0x26, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x9b, 0x4e,
	0x20, 0x00, 0x9b, 0x4f, 0x20, 0x00, 0x9b, 0x4e, 0x20, 0x67, 0x9b, 0x4e,
	0x1f, 0xff, 0x9a, 0x4d, 0x1f, 0xff, 0x99, 0x4c, 0x1e, 0xff, 0x98, 0x4a,
	0x1d, 0xff, 0x97, 0x49, 0x1c, 0x85, 0x97, 0x4a, 0x1c, 0x00, 0x96, 0x49,
	0x1c, 0x00, 0xae, 0x68, 0x2e, 0x00, 0xad, 0x65, 0x2e, 0x00, 0xae, 0x67,
	0x2e, 0x8e, 0xad, 0x65, 0x2d, 0xff, 0xab, 0x62, 0x2c, 0xff, 0xa9, 0x60,
	0x2b, 0xff, 0xa8, 0x5e, 0x29, 0xfe, 0xa7, 0x5d, 0x29, 0x5d, 0xa7, 0x5d,
	0x29, 0x00, 0xa7, 0x5d, 0x29, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xa0, 0x54, 0x24, 0x00, 0xa0, 0x54,
	0x24, 0x00, 0xa0, 0x54, 0x24, 0x67, 0x9f, 0x54, 0x23, 0xff, 0x9e, 0x52,
	0x22, 0xff, 0x9d, 0x51, 0x21, 0xff, 0x9b, 0x4f, 0x20, 0xff, 0x9b, 0x4e,
	0x20, 0x83, 0x9b, 0x4f, 0x20, 0x00, 0x9a, 0x4e, 0x1f, 0x00, 0xb1, 0x6a,
	0x2f, 0x00, 0xae, 0x67, 0x2e, 0x00, 0xb0, 0x69, 0x2f, 0x8e, 0xaf, 0x67,
	0x2e, 0xff, 0xac, 0x64, 0x2c, 0xff, 0xaa, 0x61, 0x2b, 0xff, 0xa9, 0x60,
	0x2a, 0xfe, 0xa8, 0x5f, 0x2a, 0x5d, 0xa8, 0x5f, 0x2a, 0x00, 0xa8, 0x5f,
	0x2a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0xa3, 0x59, 0x27, 0x00, 0xa3, 0x59, 0x27, 0x00, 0xa3, 0x59,
	0x27, 0x67, 0xa3, 0x58, 0x26, 0xff, 0xa2, 0x57, 0x26, 0xff, 0xa1, 0x56,
	0x25, 0xff, 0x9f, 0x54, 0x24, 0xff, 0x9e, 0x53, 0x23, 0x83, 0x9f, 0x54,
	0x23, 0x00, 0x9e, 0x52, 0x22, 0x00, 0xb3, 0x6d, 0x30, 0x00, 0xb0, 0x69,
	0x2f, 0x00, 0xb2, 0x6c, 0x30, 0x8e, 0xb0, 0x69, 0x2e, 0xff, 0xad, 0x65,
	0x2d, 0xff, 0xaa, 0x62, 0x2b, 0xff, 0xa9, 0x60, 0x2b, 0xfe, 0xa9, 0x61,
	0x2b, 0x5d, 0xa9, 0x61, 0x2b, 0x00, 0xa9, 0x61, 0x2b, 0x00, 0xae, 0x67,
	0x2e, 0x00, 0xff, 0x70, 0xbe, 0x00, 0xaf, 0x67, 0x2f, 0x00, 0xaf, 0x67,
	0x2f, 0x00, 0xaf, 0x67, 0x2f, 0x00, 0xaf, 0x67, 0x2f, 0x00, 0xae, 0x66,
	0x2e, 0x00, 0xad, 0x65, 0x2e, 0x00, 0xab, 0x63, 0x2d, 0x00, 0xab, 0x63,
	0x2d, 0x00, 0xaa, 0x62, 0x2c, 0x00, 0xac, 0x63, 0x2d, 0x00, 0xa6, 0x5d,
	0x29, 0x00, 0xa6, 0x5c, 0x29, 0x00, 0xa6, 0x5d, 0x29, 0x67, 0xa6, 0x5c,
	0x29, 0xff, 0xa5, 0x5b, 0x28, 0xff, 0xa5, 0x5b, 0x28, 0xff, 0xa3, 0x59,
	0x27, 0xff, 0xa2, 0x57, 0x26, 0x83, 0xa3, 0x58, 0x27, 0x00, 0xa2, 0x57,
	0x25, 0x00, 0xb5, 0x6f, 0x31, 0x00, 0xb1, 0x6a, 0x2f, 0x00, 0xb4, 0x6d,
	0x30, 0x8e, 0xb1, 0x6a, 0x2f, 0xff, 0xae, 0x66, 0x2d, 0xff, 0xab, 0x62,
	0x2c, 0xff, 0xaa, 0x62, 0x2b, 0xfe, 0xaa, 0x62, 0x2b, 0x5d, 0xaa, 0x62,
	0x2b, 0x00, 0xaf, 0x67, 0x2e, 0x00, 0xaf, 0x67, 0x2e, 0x00, 0xaf, 0x67,
	0x2e, 0x00, 0xaf, 0x67, 0x2f, 0x18, 0xaf, 0x67, 0x2f, 0x1d, 0xaf, 0x67,
	0x2f, 0x08, 0xad, 0x66, 0x2f, 0x02, 0xac, 0x64, 0x2e, 0x02, 0xad, 0x64,
	0x2d, 0x09, 0xab, 0x63, 0x2d, 0x1e, 0xab, 0x62, 0x2d, 0x16, 0xac, 0x65,
	0x2e, 0x00, 0xad, 0x63, 0x2d, 0x00, 0xad, 0x64, 0x2e, 0x00, 0xa9, 0x5f,
	0x2b, 0x00, 0xa9, 0x60, 0x2b, 0x67, 0xa8, 0x5f, 0x2b, 0xff, 0xa8, 0x5f,
	0x2a, 0xff, 0xa8, 0x5e, 0x2a, 0xff, 0xa6, 0x5c, 0x29, 0xff, 0xa5, 0x5b,
	0x28, 0x83, 0xa6, 0x5c, 0x29, 0x00, 0xa4, 0x5a, 0x27, 0x00, 0xb6, 0x70,
	0x31, 0x00, 0xb3, 0x6c, 0x30, 0x00, 0xb5, 0x6f, 0x31, 0x8e, 0xb3, 0x6c,
	0x30, 0xff, 0xaf, 0x67, 0x2e, 0xff, 0xac, 0x64, 0x2c, 0xff, 0xac, 0x63,
	0x2c, 0xfe, 0xac, 0x63, 0x2c, 0x5d, 0xac, 0x63, 0x2c, 0x00, 0xb0, 0x68,
	0x2f, 0x00, 0xb0, 0x68, 0x2e, 0x00, 0xaf, 0x67, 0x2e, 0x37, 0xaf, 0x67,
	0x2e, 0xc7, 0xaf, 0x67, 0x2f, 0xdc, 0xaf, 0x67, 0x2f, 0xba, 0xaf, 0x68,
	0x2f, 0xa7, 0xae, 0x67, 0x2e, 0xa8, 0xad, 0x66, 0x2e, 0xbc, 0xac, 0x64,
	0x2d, 0xde, 0xac, 0x64, 0x2d, 0xc2, 0xac, 0x64, 0x2e, 0x31, 0xae, 0x67,
	0x2f, 0x00, 0xaf, 0x67, 0x2f, 0x00, 0xaa, 0x62, 0x2c, 0x00, 0xaa, 0x62,
	0x2c, 0x67, 0xaa, 0x62, 0x2c, 0xff, 0xaa, 0x62, 0x2c, 0xff, 0xaa, 0x61,
	0x2c, 0xff, 0xa8, 0x5f, 0x2a, 0xff, 0xa6, 0x5d, 0x29, 0x83, 0xa7, 0x5e,
	0x2a, 0x00, 0xa6, 0x5c, 0x29, 0x00, 0xb8, 0x73, 0x33, 0x00, 0xb5, 0x6f,
	0x31, 0x00, 0xb7, 0x72, 0x32, 0x8e, 0xb5, 0x6f, 0x31, 0xff, 0xb1, 0x6a,
	0x2f, 0xff, 0xae, 0x66, 0x2d, 0xff, 0xad, 0x65, 0x2d, 0xfe, 0xae, 0x65,
	0x2d, 0x5d, 0xae, 0x66, 0x2d, 0x00, 0xb2, 0x6b, 0x30, 0x00, 0xb0, 0x68,
	0x2e, 0x44, 0xaf, 0x67, 0x2e, 0xd9, 0xaf, 0x67, 0x2e, 0xff, 0xaf, 0x67,
	0x2e, 0xff, 0xb0, 0x68, 0x2f, 0xff, 0xb0, 0x69, 0x2f, 0xff, 0xaf, 0x68,
	0x2f, 0xff, 0xaf, 0x67, 0x2e, 0xff, 0xae, 0x67, 0x2e, 0xff, 0xae, 0x66,
	0x2e, 0xff, 0xae, 0x66, 0x2e, 0xd3, 0xaf, 0x67, 0x2f, 0x3e, 0xb1, 0x6a,
	0x30, 0x00, 0xac, 0x64, 0x2e, 0x00, 0xac, 0x63, 0x2e, 0x67, 0xac, 0x63,
	0x2e, 0xff, 0xac, 0x63, 0x2d, 0xff, 0xab, 0x62, 0x2d, 0xff, 0xa9, 0x60,
	0x2c, 0xff, 0xa8, 0x5e, 0x2b, 0x83, 0xa8, 0x5f, 0x2b, 0x00, 0xa7, 0x5d,
	0x2a, 0x00, 0xba, 0x76, 0x34, 0x00, 0xb8, 0x73, 0x33, 0x00, 0xb9, 0x75,
	0x34, 0x8e, 0xb8, 0x73, 0x33, 0xff, 0xb4, 0x6e, 0x31, 0xff, 0xb1, 0x6a,
	0x2f, 0xff, 0xb0, 0x69, 0x2e, 0xfe, 0xb0, 0x68, 0x2e, 0x5b, 0xb2, 0x6b,
	0x30, 0x00, 0xb2, 0x6a, 0x30, 0x53, 0xb1, 0x69, 0x2f, 0xe3, 0xb0, 0x69,
	0x2f, 0xff, 0xb0, 0x68, 0x2e, 0xff, 0xb1, 0x69, 0x2f, 0xff, 0xb2, 0x6b,
	0x30, 0xff, 0xb3, 0x6c, 0x30, 0xff, 0xb2, 0x6c, 0x30, 0xff, 0xb2, 0x6b,
	0x30, 0xff, 0xb1, 0x6a, 0x30, 0xff, 0xb1, 0x6a, 0x30, 0xff, 0xb1, 0x6a,
	0x30, 0xff, 0xb1, 0x6a, 0x30, 0xde, 0xb1, 0x6a, 0x30, 0x4c, 0xaf, 0x67,
	0x2f, 0x00, 0xad, 0x65, 0x2e, 0x65, 0xad, 0x65, 0x2e, 0xff, 0xac, 0x64,
	0x2e, 0xff, 0xac, 0x63, 0x2e, 0xff, 0xaa, 0x61, 0x2c, 0xff, 0xa9, 0x5f,
	0x2c, 0x83, 0xa9, 0x60, 0x2c, 0x00, 0xa8, 0x5f, 0x2b, 0x00, 0xbd, 0x7a,
	0x36, 0x00, 0xbc, 0x79, 0x36, 0x00, 0xbd, 0x7a, 0x36, 0x8e, 0xbb, 0x78,
	0x35, 0xff, 0xb8, 0x73, 0x33, 0xff, 0xb5, 0x70, 0x32, 0xff, 0xb4, 0x6e,
	0x31, 0xff, 0xb5, 0x6f, 0x31, 0x82, 0xb5, 0x70, 0x32, 0x66, 0xb4, 0x6e,
	0x31, 0xeb, 0xb4, 0x6d, 0x31, 0xff, 0xb4, 0x6d, 0x30, 0xff, 0xb4, 0x6d,
	0x30, 0xff, 0xb5, 0x6e, 0x31, 0xff, 0xb6, 0x70, 0x32, 0xff, 0xb7, 0x71,
	0x32, 0xff, 0xb6, 0x70, 0x32, 0xff, 0xb5, 0x6e, 0x31, 0xff, 0xb4, 0x6d,
	0x31, 0xff, 0xb4, 0x6d, 0x31, 0xff, 0xb4, 0x6e, 0x31, 0xff, 0xb4, 0x6d,
	0x31, 0xff, 0xb3, 0x6d, 0x31, 0xe7, 0xb3, 0x6c, 0x31, 0x60, 0xb0, 0x68,
	0x30, 0x8a, 0xae, 0x66, 0x2f, 0xff, 0xad, 0x65, 0x2f, 0xff, 0xad, 0x64,
	0x2e, 0xff, 0xab, 0x62, 0x2d, 0xff, 0xa9, 0x61, 0x2c, 0x83, 0xaa, 0x61,
	0x2d, 0x00, 0xa9, 0x60, 0x2c, 0x00, 0xc1, 0x80, 0x39, 0x00, 0xbf, 0x7e,
	0x38, 0x00, 0xc0, 0x7f, 0x39, 0x8e, 0xbe, 0x7d, 0x38, 0xff, 0xbc, 0x79,
	0x36, 0xff, 0xba, 0x76, 0x35, 0xff, 0xb9, 0x75, 0x34, 0xff, 0xb9, 0x75,
	0x34, 0xf6, 0xb9, 0x75, 0x34, 0xf9, 0xb9, 0x75, 0x34, 0xff, 0xb9, 0x75,
	0x34, 0xff, 0xb9, 0x75, 0x34, 0xff, 0xba, 0x75, 0x34, 0xff, 0xbb, 0x76,
	0x34, 0xff, 0xbb, 0x77, 0x35, 0xf6, 0xbb, 0x77, 0x35, 0xf8, 0xba, 0x75,
	0x34, 0xf8, 0xb8, 0x72, 0x32, 0xf5, 0xb7, 0x70, 0x32, 0xff, 0xb6, 0x70,
	0x31, 0xff, 0xb6, 0x70, 0x32, 0xff, 0xb6, 0x70, 0x32, 0xff, 0xb6, 0x6f,
	0x32, 0xff, 0xb4, 0x6e, 0x32, 0xf8, 0xb2, 0x6b, 0x31, 0xf7, 0xb1, 0x69,
	0x30, 0xff, 0xb0, 0x67, 0x30, 0xff, 0xae, 0x65, 0x2f, 0xff, 0xac, 0x63,
	0x2e, 0xff, 0xab, 0x62, 0x2d, 0x83, 0xab, 0x63, 0x2d, 0x00, 0xaa, 0x61,
	0x2d, 0x00, 0xc3, 0x84, 0x3c, 0x00, 0xc0, 0x81, 0x3a, 0x00, 0xc2, 0x83,
	0x3c, 0x8e, 0xc0, 0x81, 0x3a, 0xff, 0xbe, 0x7d, 0x39, 0xff, 0xbd, 0x7c,
	0x38, 0xff, 0xbe, 0x7c, 0x38, 0xff, 0xbe, 0x7d, 0x38, 0xff, 0xbf, 0x7d,
	0x38, 0xff, 0xbf, 0x7e, 0x38, 0xff, 0xc0, 0x7e, 0x38, 0xff, 0xc0, 0x7f,
	0x38, 0xff, 0xc1, 0x80, 0x38, 0xff, 0xc0, 0x7f, 0x38, 0xae, 0xbf, 0x7c,
	0x37, 0x43, 0xbe, 0x7b, 0x36, 0x47, 0xbc, 0x78, 0x35, 0x47, 0xba, 0x75,
	0x34, 0x43, 0xba, 0x75, 0x34, 0xb4, 0xba, 0x74, 0x34, 0xff, 0xb9, 0x73,
	0x34, 0xff, 0xb8, 0x73, 0x33, 0xff, 0xb8, 0x73, 0x34, 0xff, 0xb7, 0x72,
	0x34, 0xff, 0xb6, 0x70, 0x33, 0xff, 0xb4, 0x6e, 0x33, 0xff, 0xb2, 0x6b,
	0x32, 0xff, 0xb0, 0x67, 0x30, 0xff, 0xad, 0x65, 0x2e, 0xff, 0xac, 0x63,
	0x2e, 0x83, 0xac, 0x64, 0x2e, 0x00, 0xab, 0x63, 0x2d, 0x00, 0xc3, 0x85,
	0x3d, 0x00, 0xc1, 0x82, 0x3b, 0x00, 0xc3, 0x84, 0x3d, 0x8e, 0xc1, 0x82,
	0x3c, 0xff, 0xc1, 0x81, 0x3b, 0xff, 0xc1, 0x81, 0x3b, 0xff, 0xc2, 0x83,
	0x3c, 0xff, 0xc3, 0x85, 0x3d, 0xff, 0xc4, 0x87, 0x3d, 0xff, 0xc5, 0x87,
	0x3d, 0xff, 0xc5, 0x88, 0x3e, 0xff, 0xc6, 0x89, 0x3e, 0xfe, 0xc6, 0x89,
	0x3d, 0x9e, 0xc5, 0x87, 0x3c, 0x15, 0xc0, 0x7e, 0x38, 0x00, 0xbe, 0x7a,
	0x36, 0x00, 0xbc, 0x78, 0x35, 0x00, 0xbb, 0x76, 0x34, 0x00, 0xbe, 0x7a,
	0x36, 0x19, 0xbf, 0x7c, 0x37, 0xa6, 0xbf, 0x7c, 0x38, 0xff, 0xbd, 0x7a,
	0x37, 0xff, 0xbd, 0x79, 0x37, 0xff, 0xbc, 0x78, 0x36, 0xff, 0xbb, 0x77,
	0x36, 0xff, 0xb9, 0x75, 0x35, 0xff, 0xb7, 0x71, 0x34, 0xff, 0xb3, 0x6c,
	0x32, 0xff, 0xb0, 0x68, 0x30, 0xff, 0xae, 0x66, 0x2f, 0x83, 0xaf, 0x67,
	0x2f, 0x00, 0xae, 0x65, 0x2e, 0x00, 0xc3, 0x85, 0x3d, 0x00, 0xc2, 0x84,
	0x3d, 0x00, 0xc3, 0x84, 0x3d, 0x8e, 0xc2, 0x84, 0x3d, 0xff, 0xc2, 0x84,
	0x3d, 0xff, 0xc4, 0x87, 0x3e, 0xff, 0xc6, 0x8a, 0x40, 0xff, 0xc7, 0x8c,
	0x40, 0xff, 0xc8, 0x8d, 0x41, 0xff, 0xc8, 0x8e, 0x41, 0xff, 0xc9, 0x8f,
	0x42, 0xfb, 0xca, 0x90, 0x42, 0x8f, 0xca, 0x90, 0x42, 0x0e, 0xc6, 0x89,
	0x3e, 0x00, 0xc4, 0x86, 0x3b, 0x00, 0xc2, 0x81, 0x39, 0x00, 0xbc, 0x77,
	0x35, 0x00, 0xbd, 0x79, 0x35, 0x00, 0xbf, 0x7d, 0x38, 0x00, 0xc4, 0x85,
	0x3c, 0x11, 0xc4, 0x85, 0x3d, 0x96, 0xc4, 0x85, 0x3d, 0xfc, 0xc3, 0x84,
	0x3c, 0xff, 0xc2, 0x82, 0x3b, 0xff, 0xc1, 0x80, 0x3b, 0xff, 0xbf, 0x7d,
	0x39, 0xff, 0xbc, 0x79, 0x37, 0xff, 0xb8, 0x73, 0x35, 0xff, 0xb4, 0x6e,
	0x32, 0xff, 0xb2, 0x6c, 0x31, 0x83, 0xb3, 0x6d, 0x32, 0x00, 0xb1, 0x6b,
	0x31, 0x00, 0xc4, 0x86, 0x3e, 0x00, 0xc3, 0x85, 0x3e, 0x00, 0xc4, 0x86,
	0x3e, 0x8e, 0xc4, 0x86, 0x3e, 0xff, 0xc5, 0x88, 0x3f, 0xff, 0xc7, 0x8b,
	0x41, 0xff, 0xc8, 0x8e, 0x42, 0xff, 0xc9, 0x8f, 0x42, 0xff, 0xca, 0x90,
	0x43, 0xff, 0xca, 0x91, 0x44, 0xf6, 0xcb, 0x92, 0x44, 0x7f, 0xcd, 0x94,
	0x45, 0x08, 0xca, 0x90, 0x42, 0x00, 0xc9, 0x8e, 0x41, 0x00, 0xc8, 0x8c,
	0x3f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc1, 0x81,
	0x39, 0x00, 0xc3, 0x83, 0x3b, 0x00, 0xc5, 0x87, 0x3d, 0x00, 0xc9, 0x8e,
	0x41, 0x0b, 0xc9, 0x8d, 0x41, 0x86, 0xc8, 0x8d, 0x41, 0xf8, 0xc7, 0x8b,
	0x41, 0xff, 0xc6, 0x89, 0x40, 0xff, 0xc4, 0x86, 0x3e, 0xff, 0xc1, 0x81,
	0x3c, 0xff, 0xbd, 0x7b, 0x39, 0xff, 0xba, 0x76, 0x37, 0xff, 0xb9, 0x74,
	0x36, 0x83, 0xba, 0x75, 0x36, 0x00, 0xb8, 0x73, 0x35, 0x00, 0xc5, 0x87,
	0x3f, 0x00, 0xc5, 0x89, 0x40, 0x00, 0xc5, 0x88, 0x3f, 0x8e, 0xc6, 0x89,
	0x40, 0xff, 0xc7, 0x8b, 0x41, 0xff, 0xc9, 0x8e, 0x42, 0xff, 0xca, 0x90,
	0x43, 0xff, 0xca, 0x90, 0x44, 0xff, 0xca, 0x91, 0x44, 0xf1, 0xcb, 0x91,
	0x44, 0x6f, 0xcc, 0x95, 0x47, 0x04, 0xcb, 0x92, 0x44, 0x00, 0xcb, 0x92,
	0x44, 0x00, 0xcc, 0x91, 0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc6, 0x8a,
	0x3f, 0x00, 0xc8, 0x8b, 0x40, 0x00, 0xc9, 0x8e, 0x41, 0x00, 0xcd, 0x95,
	0x45, 0x06, 0xcb, 0x92, 0x44, 0x76, 0xcb, 0x91, 0x44, 0xf3, 0xca, 0x8f,
	0x43, 0xff, 0xc8, 0x8d, 0x42, 0xff, 0xc5, 0x88, 0x40, 0xff, 0xc2, 0x83,
	0x3d, 0xff, 0xbf, 0x7e, 0x3b, 0xff, 0xbe, 0x7c, 0x3a, 0x83, 0xbf, 0x7d,
	0x3a, 0x00, 0xbe, 0x7c, 0x39, 0x00, 0xc6, 0x8a, 0x40, 0x00, 0xc8, 0x8c,
	0x42, 0x00, 0xc7, 0x8a, 0x41, 0x8e, 0xc8, 0x8c, 0x42, 0xff, 0xc9, 0x8f,
	0x43, 0xff, 0xcb, 0x91, 0x44, 0xff, 0xcb, 0x92, 0x44, 0xff, 0xcb, 0x91,
	0x44, 0xea, 0xca, 0x91, 0x44, 0x5f, 0xcc, 0x90, 0x43, 0x01, 0xcb, 0x92,
	0x45, 0x00, 0xcb, 0x93, 0x45, 0x00, 0xcb, 0x93, 0x44, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xca, 0x90,
	0x42, 0x00, 0xcb, 0x91, 0x43, 0x00, 0xcb, 0x93, 0x44, 0x00, 0xcf, 0x99,
	0x47, 0x02, 0xcc, 0x94, 0x45, 0x66, 0xcb, 0x92, 0x45, 0xed, 0xca, 0x90,
	0x44, 0xff, 0xc8, 0x8e, 0x42, 0xff, 0xc6, 0x89, 0x40, 0xff, 0xc3, 0x84,
	0x3d, 0xff, 0xc1, 0x81, 0x3c, 0x83, 0xc2, 0x82, 0x3c, 0x00, 0xc1, 0x80,
	0x3b, 0x00, 0xc8, 0x8c, 0x42, 0x00, 0xc9, 0x8f, 0x43, 0x00, 0xc8, 0x8d,
	0x42, 0x90, 0xc9, 0x8f, 0x43, 0xff, 0xcb, 0x91, 0x44, 0xff, 0xcc, 0x93,
	0x45, 0xff, 0xcc, 0x93, 0x45, 0xe1, 0xcb, 0x92, 0x44, 0x50, 0xc9, 0x94,
	0x46, 0x00, 0xcb, 0x91, 0x44, 0x00, 0xcb, 0x91, 0x44, 0x00, 0xcb, 0x92,
	0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x93,
	0x44, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xbc, 0x75,
	0x41, 0x00, 0xcc, 0x94, 0x45, 0x57, 0xcb, 0x93, 0x45, 0xe5, 0xcb, 0x91,
	0x44, 0xff, 0xc9, 0x8e, 0x42, 0xff, 0xc5, 0x88, 0x3f, 0xff, 0xc3, 0x85,
	0x3d, 0x85, 0xc4, 0x86, 0x3e, 0x00, 0xc2, 0x83, 0x3d, 0x00, 0xc8, 0x8d,
	0x42, 0x00, 0xca, 0x90, 0x44, 0x00, 0xc9, 0x8e, 0x43, 0x49, 0xca, 0x90,
	0x43, 0x9a, 0xcb, 0x92, 0x45, 0xaf, 0xcc, 0x94, 0x45, 0xb3, 0xcc, 0x94,
	0x45, 0x42, 0xcb, 0x92, 0x44, 0x00, 0xcb, 0x92, 0x44, 0x00, 0xcb, 0x91,
	0x44, 0x00, 0xca, 0x91, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x95,
	0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcb, 0x92,
	0x44, 0x00, 0xcc, 0x94, 0x45, 0x48, 0xcc, 0x93, 0x45, 0xb5, 0xca, 0x91,
	0x43, 0xae, 0xc7, 0x8b, 0x40, 0x98, 0xc4, 0x87, 0x3e, 0x42, 0xc5, 0x88,
	0x3f, 0x00, 0xc4, 0x85, 0x3d, 0x00, 0xc8, 0x8c, 0x42, 0x00, 0xd1, 0x9d,
	0x4d, 0x00, 0xc8, 0x8d, 0x42, 0x00, 0xc8, 0x8d, 0x43, 0x00, 0xcb, 0x92,
	0x44, 0x04, 0xcc, 0x93, 0x45, 0x06, 0xcc, 0x94, 0x45, 0x00, 0xcc, 0x93,
	0x45, 0x00, 0xcb, 0x93, 0x45, 0x00, 0xcb, 0x92, 0x44, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xcc, 0x95,
	0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcb, 0x92,
	0x44, 0x00, 0xcc, 0x95, 0x45, 0x07, 0xcd, 0x95, 0x45, 0x04, 0xc1, 0x80,
	0x3b, 0x00, 0xc4, 0x85, 0x3d, 0x00, 0xc4, 0x86, 0x3f, 0x00, 0xc3, 0x84,
	0x3d, 0x00, 0xc8, 0x8d, 0x42, 0x00, 0xc9, 0x8f, 0x43, 0x00, 0xc9, 0x8e,
	0x42, 0x00, 0xca, 0x90, 0x44, 0x00, 0xcb, 0x92, 0x45, 0x00, 0xcc, 0x93,
	0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcc, 0x94, 0x45, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x36, 0x27,
	0x12, 0x00, 0xcc, 0x94, 0x45, 0x00, 0xcc, 0x95, 0x45, 0x00, 0xcc, 0x94,
	0x45, 0x00, 0xcb, 0x93, 0x44, 0x00, 0xc9, 0x8d, 0x41, 0x00, 0xc3, 0x85,
	0x3e, 0x00, 0xc6, 0x89, 0x3f, 0x00, 0xc5, 0x86, 0x3e, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x01, 0xff,
	0xff, 0x80, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x0f, 0xf0, 0x00, 0x00, 0x0f,
	0xf0, 0x00, 0x00, 0x08, 0x10, 0x00, 0x00, 0x08, 0x10, 0x00, 0x00, 0x08,
	0x10, 0x00, 0x00, 0x08, 0x10, 0x00, 0x00, 0x08, 0x10, 0x00, 0x00, 0x08,
	0x10, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0x3f,
	0xfc, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x80, 0x00, 0x00, 0x03,
	0xc0, 0x00, 0x00, 0x07, 0xe0, 0x00, 0x00, 0x0f, 0xf0, 0x00, 0x00, 0x1f,
	0xf8, 0x00, 0x00, 0x3f, 0xfc, 0x00, 0x00, 0xff, 0xfe, 0x00, 0xff, 0xff,
	0xff, 0xff,
}

func StartClash() {
	// 启动Clash
	Command := exec.Command("cmd", "/c", "clash.meta-windows-amd64 -d .")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Start()
}
func StopClash() {
	// 卸载Clash Tunnel网卡
	Command := exec.Command("powershell", `pnputil /remove-device (Get-PnpDevice | Where-Object{$_.Name -eq "Meta Tunnel"}).InstanceId`)
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
	// 停止Clash
	Command = exec.Command("cmd", "/c", "taskkill /f /im clash.meta-windows-amd64.exe")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Run()
}
func ShowMessage() {
	// 延时
	time.Sleep(time.Duration(2) * time.Second)
	// 读取Clash运行状态
	Command := exec.Command("powershell", "Get-Process clash.meta-windows-amd64")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err := Command.Output()
	if err == nil {
		ClashStatus = "Running"
	} else {
		ClashStatus = "Stopped"
	}
	// 读取Tun运行状态
	Command = exec.Command("powershell", "Get-NetAdapter Meta")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_, err = Command.Output()
	if err == nil {
		TunStatus = "Running"
	} else {
		TunStatus = "Stopped"
	}
	// 显示通知
	NotifyIcon.ShowMessage("", "Clash Meta ---- "+ClashStatus+"\r\n"+"Tun Adapter --- "+TunStatus)
}
func init() {
	// 阻止多次启动
	Mutex, _ := syscall.UTF16PtrFromString("ClashMetaLite")
	_, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("CreateMutexW").Call(0, 0, uintptr(unsafe.Pointer(Mutex)))
	if int(err.(syscall.Errno)) != 0 {
		os.Exit(1)
	}
	// 创建ICO
	os.WriteFile("icon.ico", IcoData, 0644)
}
func main() {
	// 定义托盘图标文字
	MainWindow, _ = walk.NewMainWindow()
	Icon, _ = walk.Resources.Icon("icon.ico")
	// 删除图标
	Command := exec.Command("cmd", "/c", "del /f /q icon.ico")
	Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	Command.Start()
	NotifyIcon, _ = walk.NewNotifyIcon(MainWindow)
	defer NotifyIcon.Dispose()
	NotifyIcon.SetIcon(Icon)
	NotifyIcon.SetToolTip("ClashMetaLite")
	NotifyIcon.SetVisible(true)
	// 定义左键显示
	NotifyIcon.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}
		// 显示通知
		ShowMessage()
	})
	// 定义右键菜单元素
	Sign := walk.NewAction()
	Sign.SetText("Designed by Amaury")
	blank1 := walk.NewAction()
	blank1.SetText("-")
	blank2 := walk.NewAction()
	blank2.SetText("-")
	blank3 := walk.NewAction()
	blank3.SetText("-")
	blank4 := walk.NewAction()
	blank4.SetText("-")
	Start := walk.NewAction()
	Start.SetText("Start")
	Stop := walk.NewAction()
	Stop.SetText("Stop")
	Dashboard := walk.NewAction()
	Dashboard.SetText("Dashboard")
	Exit := walk.NewAction()
	Exit.SetText("Exit")
	// 定义右键菜单
	NotifyIcon.ContextMenu().Actions().Add(Sign)
	NotifyIcon.ContextMenu().Actions().Add(blank1)
	NotifyIcon.ContextMenu().Actions().Add(Start)
	NotifyIcon.ContextMenu().Actions().Add(blank2)
	NotifyIcon.ContextMenu().Actions().Add(Stop)
	NotifyIcon.ContextMenu().Actions().Add(blank3)
	NotifyIcon.ContextMenu().Actions().Add(Dashboard)
	NotifyIcon.ContextMenu().Actions().Add(blank4)
	NotifyIcon.ContextMenu().Actions().Add(Exit)
	// 启动Clash
	Start.Triggered().Attach(func() {
		// 停止Clash
		StopClash()
		// 启动Clash
		StartClash()
		// 显示通知
		ShowMessage()
	})
	// 停止Clash
	Stop.Triggered().Attach(func() {
		// 停止Clash
		StopClash()
		// 显示通知
		ShowMessage()
	})
	// 打开Dashboard
	Dashboard.Triggered().Attach(func() {
		Command := exec.Command("cmd", "/c", "start", "http://127.0.0.1:9090/ui/metacubexd-gh-pages/")
		Command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		Command.Start()
	})
	// Exit
	Exit.Triggered().Attach(func() {
		// 停止Clash
		StopClash()
		// 退出主程序
		walk.App().Exit(0)
	})
	// 停止Clash
	StopClash()
	// 启动Clash
	StartClash()
	// 显示通知
	ShowMessage()
	// 主程序运行
	MainWindow.Run()
}
