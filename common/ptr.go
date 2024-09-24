package common

import "github.com/gogf/gf/v2/container/gvar"

func String(v string) *string {
	return &v
}

func Int(v int) *int {
	return &v
}

func Int64(v int64) *int64 {
	return &v
}

func Float64(v float64) *float64 {
	return &v
}

func VarString(v *gvar.Var) *string {
	if v.IsNil() {
		return nil
	}
	return String(v.String())
}

func VarInt(v *gvar.Var) *int {
	if v.IsNil() {
		return nil
	}
	return Int(v.Int())
}

func VarInt64(v *gvar.Var) *int64 {
	if v.IsNil() {
		return nil
	}
	return Int64(v.Int64())
}

func VarFloat64(v *gvar.Var) *float64 {
	if v.IsNil() {
		return nil
	}
	return Float64(v.Float64())
}
