//go:build !go1.22
// +build !go1.22

package mo

//
// sql.Null[T] has been introduce in go1.22
// This file is a copy of stdlib and ensure retro-compatibility.
// See https://github.com/samber/mo/pull/49
//

import (
	"errors"
	"reflect"
)

var errNilPtr = errors.New("destination pointer is nil")

func cloneBytes(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func asString(src any) string { _ = "STUB: not implemented"; return "" }

func asBytes(buf []byte, rv reflect.Value) (b []byte, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func strconvErr(err error) error { _ = "STUB: not implemented"; return nil }

// convertAssignRows copies to dest the value in src, converting it if possible.
// An error is returned if the copy would result in loss of information.
// dest should be a pointer type. If rows is passed in, the rows will
// be used as the parent for any cursor values converted from a
// driver.Rows to a *Rows.
func convertAssign(dest, src any) error {
	_ = "STUB: not implemented"
	// Common cases, without reflect.
	return nil
}

// The following conversions use a string value as an intermediate representation
// to convert between various numeric types.
//
// This also allows scanning into user defined types such as "type Int int64".
// For symmetry, also check for string destination types.

func (o *Option[T]) scanConvertValue(src any) error {
	_ = "STUB: not implemented"
	// we try to convertAssign values that we can't directly assign because ConvertValue
	// will return immediately for v that is already a Value, even if it is a different
	// Value type than the one we expect here.
	return nil
}
