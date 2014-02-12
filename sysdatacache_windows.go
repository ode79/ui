// 11 february 2014
package main

import (
	"fmt"
	"sync"
)

// I need a way to get a sysData for a given HWND or a given HWND/control ID. So, this.

var (
	sysDatas = map[_HWND]*sdcEntry{}
	sysDatasLock sync.Mutex
	sysDataIDs = map[_HMENU]*sdcEntry{}
	sysDataIDsLock sync.Mutex
)

func addSysData(hwnd _HWND, s *sysData) {
	sysDatasLock.Lock()
	defer sysDatasLock.Unlock()
	sysDatas[hwnd] = s
}

func addIDSysData(id _HMENU, s *sysData) {
	sysDataIDsLock.Lock()
	defer sysDataIDsLock.Unlock()
	sysDataIDs[id] = s
}

func getSysData(hwnd _HWND) *sysData {
	sysDatasLock.Lock()
	defer sysDatasLock.Unlock()
	if ss, ok := sysDatas[hwnd]; ok {
		return ss
	}
	return nil
}

func getIDSysData(id _HMENU) *sysData {
	sysDataIDsLock.Lock()
	defer sysDataIDsLock.Unlock()
	if ss, ok := sysDataIDs[id]; ok {
		return ss
	}
	panic(fmt.Sprintf("getting nonexistent ID %d for HWND %d\n", id, hwnd))
}