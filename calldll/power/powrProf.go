package main

import (
	"syscall"
)

/**
功能未实现
*/

// 第一层
type GLOBAL_POWER_POLICY struct {
	user  GLOBAL_USER_POWER_POLICY
	match GLOBAL_MACHINE_POWER_POLICY
}

// 第一层
type PPOWER_POLICY struct {
	user  USER_POWER_POLICY
	match MACHINE_POWER_POLICY
}

// 第二层
type USER_POWER_POLICY struct {
}

// 第二层
type MACHINE_POWER_POLICY struct {
}

// 第二层
type GLOBAL_USER_POWER_POLICY struct {
}

// 第二层
type GLOBAL_MACHINE_POWER_POLICY struct {
}

type pGlobalPowerPolicy GLOBAL_POWER_POLICY
type pPowerPolicy PPOWER_POLICY

func main() {
	mod := syscall.NewLazyDLL("PowrProf.dll")
	proc := mod.NewProc("GetCurrentPowerPolicies")

	proc.Call()

}
