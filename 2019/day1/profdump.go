// Copyright (c) 2019, Theodor-Iulian Ciobanu
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:

// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// +build profdump

package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"
)

func init() {
	var (
		cpuprofile   = flag.Bool("cpuprofile", false, "write cpu profile")
		goprofile    = flag.Int("goprofile", -1, "write goroutine profile")
		heapprofile  = flag.Int("heapprofile", -1, "write heap profile")
		allocprofile = flag.Int("allocprofile", -1, "write allocs profile")
		tcprofile    = flag.Int("tcprofile", -1, "write threadcreate profile")
		blockprofile = flag.Int("blockprofile", -1, "write block profile")
		mutexprofile = flag.Int("mutexprofile", -1, "write mutex profile")
		allprofile   = flag.Int("allprofile", -1, "write all profiles")
	)

	flag.Parse()

	if *allprofile >= 0 {
		*cpuprofile = true
		*goprofile = *allprofile
		*heapprofile = *allprofile
		*allocprofile = *allprofile
		*tcprofile = *allprofile
		*blockprofile = *allprofile
		*mutexprofile = *allprofile
	}

	getProfileFile := profileFileGetter("pprof")

	var cpufile *os.File
	startProfiling = func() {
		var err error
		if *cpuprofile {
			cpufile = getProfileFile("cpu", 0)
			if err = pprof.StartCPUProfile(cpufile); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
		}
	}

	stopProfiling = func() {
		var err error
		if *cpuprofile {
			pprof.StopCPUProfile()
			if err = cpufile.Close(); err != nil {
				log.Fatal("could not close CPU profile: ", err)
			}
		}

		runtime.GC()
		switch {
		case *goprofile >= 0:
			writeProfile("goroutine", *goprofile, getProfileFile)
			fallthrough
		case *heapprofile >= 0:
			writeProfile("heap", *heapprofile, getProfileFile)
			fallthrough
		case *allprofile >= 0:
			writeProfile("allocs", *allocprofile, getProfileFile)
			fallthrough
		case *tcprofile >= 0:
			writeProfile("threadcreate", *tcprofile, getProfileFile)
			fallthrough
		case *blockprofile >= 0:
			writeProfile("block", *blockprofile, getProfileFile)
			fallthrough
		case *mutexprofile >= 0:
			writeProfile("mutex", *mutexprofile, getProfileFile)
		}
	}
}

func writeProfile(name string, debug int,
	getProfileFile func(string, int) *os.File) {
	f := getProfileFile(name, debug)

	if profile := pprof.Lookup(name); profile != nil {
		if err := profile.WriteTo(f, debug); err != nil {
			log.Fatalf("could not write %s profile: %s", name, err.Error())
		}
	} else {
		log.Fatal("profile not found: ", name)
	}

	f.Close()
}

func profileFileGetter(prefix string) func(string, int) *os.File {
	if prefix != "" {
		if err := os.MkdirAll(prefix, 0755); err != nil {
			log.Fatal("could not create profiling folder: ", err)
		}
	}

	startTime := time.Now().UTC().Unix()
	baseName := filepath.Base(os.Args[0])

	return func(profile string, debug int) *os.File {
		ext := "gz"
		if debug > 0 {
			ext = "txt"
		}

		f, err := os.Create(fmt.Sprintf("%s/%s_%s_%d.prof.%s", prefix, baseName,
			profile, startTime, ext))
		if err != nil {
			log.Fatalf("could not create %s profile: ", profile, err)
		}

		return f
	}
}
