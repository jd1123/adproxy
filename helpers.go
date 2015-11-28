package main

import "github.com/jd1123/adproxy/modules"

func RegisterModule(m modules.Module) {
	mods = append(mods, m)
}
