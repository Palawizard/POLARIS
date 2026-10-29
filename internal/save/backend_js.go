//go:build js

package save

import "syscall/js"

// The page exposes `polarisSave` (get/set/remove on localStorage, wrapped
// in try/catch because storage can be disabled or full).

func bridge() js.Value {
	return js.Global().Get("polarisSave")
}

func read() (s string) {
	defer func() {
		if recover() != nil {
			s = ""
		}
	}()
	b := bridge()
	if b.IsUndefined() || b.IsNull() {
		return ""
	}
	v := b.Call("get")
	if v.Type() != js.TypeString {
		return ""
	}
	return v.String()
}

func write(s string) {
	defer func() { _ = recover() }()
	if b := bridge(); !b.IsUndefined() && !b.IsNull() {
		b.Call("set", s)
	}
}

func remove() {
	defer func() { _ = recover() }()
	if b := bridge(); !b.IsUndefined() && !b.IsNull() {
		b.Call("remove")
	}
}
