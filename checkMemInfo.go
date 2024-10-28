package main

import (
	"github.com/c9s/goprocinfo/linux"
	"github.com/infrasonar/go-libagent"
)

func readMemInfo(state map[string][]map[string]any) error {
	mem, err := linux.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return err
	}

	item := map[string]any{
		"name":              "memInfo",
		"memFree":           mem.MemFree,
		"memTotal":          mem.MemTotal,
		"memAvailable":      mem.MemAvailable,
		"buffers":           mem.Buffers,
		"cached":            mem.Cached,
		"swapCached":        mem.SwapCached,
		"active":            mem.Active,
		"inactive":          mem.Inactive,
		"activeAnon":        mem.ActiveAnon,
		"inactiveAnon":      mem.InactiveAnon,
		"activeFile":        mem.ActiveFile,
		"inactiveFile":      mem.InactiveFile,
		"unevictable":       mem.Unevictable,
		"mlocked":           mem.Mlocked,
		"swapTotal":         mem.SwapTotal,
		"swapFree":          mem.SwapFree,
		"dirty":             mem.Dirty,
		"writeback":         mem.Writeback,
		"anonPages":         mem.AnonPages,
		"mapped":            mem.Mapped,
		"shmem":             mem.Shmem,
		"slab":              mem.Slab,
		"sReclaimable":      mem.SReclaimable,
		"sUnreclaim":        mem.SUnreclaim,
		"kernelStack":       mem.KernelStack,
		"pageTables":        mem.PageTables,
		"nFSUnstable":       mem.NFS_Unstable,
		"bounce":            mem.Bounce,
		"writebackTmp":      mem.WritebackTmp,
		"commitLimit":       mem.CommitLimit,
		"committedAS":       mem.Committed_AS,
		"vmallocTotal":      mem.VmallocTotal,
		"vmallocUsed":       mem.VmallocUsed,
		"vmallocChunk":      mem.VmallocChunk,
		"hardwareCorrupted": mem.HardwareCorrupted,
		"anonHugePages":     mem.AnonHugePages,
		"hugePagesTotal":    mem.HugePages_Total,
		"hugePagesFree":     mem.HugePages_Free,
		"hugePagesRsvd":     mem.HugePages_Rsvd,
		"hugePagesSurp":     mem.HugePages_Surp,
		"hugepagesize":      mem.Hugepagesize,
		"directMap4k":       mem.DirectMap4k,
		"directMap2M":       mem.DirectMap2M,
		"directMap1G":       mem.DirectMap1G,
	}

	state["memInfo"] = []map[string]any{item}
	return nil
}

func CheckMemInfo(check *libagent.Check) (map[string][]map[string]any, error) {
	state := map[string][]map[string]any{}

	err := readMemInfo(state)
	if err != nil {
		return nil, err
	}

	return state, nil
}
