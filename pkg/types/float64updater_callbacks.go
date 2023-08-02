// Code generated by "callbackgen -type Float64Updater"; DO NOT EDIT.

package types

import ()

func (F *Float64Updater) OnUpdate(cb func(v float64)) {
	F.updateCallbacks = append(F.updateCallbacks, cb)
}

func (F *Float64Updater) EmitUpdate(v float64) {
	for _, cb := range F.updateCallbacks {
		cb(v)
	}
}