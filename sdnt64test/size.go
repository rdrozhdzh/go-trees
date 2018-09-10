package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func printAlloc(desc string) {
	runtime.GC()

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)

	log.Printf("Alloc (%s): %s, Sys: %s", desc, makeSize(m.Alloc), makeSize(m.Sys))
}

func makeSize(s uint64) string {
	if s < 2*1024 {
		return fmt.Sprintf("%d", s)
	}

	f := float64(s) / 1024
	if f < 2*1024 {
		return fmt.Sprintf("%.02f kB", f)
	}

	f /= 1024
	if f < 2*1024 {
		return fmt.Sprintf("%.02f MB", f)
	}

	f /= 1024
	if f < 2*1024 {
		return fmt.Sprintf("%.02f GB", f)
	}

	f /= 1024
	return fmt.Sprintf("%.02f TB", f)
}

func dumpMemProf() error {
	f, err := os.Create(conf.memProf)
	if err != nil {
		return err
	}
	defer f.Close()

	runtime.GC()

	return pprof.WriteHeapProfile(f)
}
