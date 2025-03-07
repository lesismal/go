// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package pprof

import (
	"fmt"
	"io"
	"runtime"
	"syscall"
)

// Adds MaxRSS to platforms that are supported.
func addMaxRSS(w io.Writer) {
	var rssToBytes uintptr
	switch runtime.GOOS {
	case "aix", "android", "dragonfly", "freebsd", "linux", "netbsd", "openbsd":
		rssToBytes = 1024
	case "darwin", "ios":
		rssToBytes = 1
	case "illumos", "solaris":
		rssToBytes = uintptr(syscall.Getpagesize())
	default:
		panic("unsupported OS")
	}

	var rusage syscall.Rusage
	syscall.Getrusage(0, &rusage)
	fmt.Fprintf(w, "# MaxRSS = %d\n", uintptr(rusage.Maxrss)*rssToBytes)
}
