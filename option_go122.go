//go:build go1.22
// +build go1.22

package mo

func (o *Option[T]) scanConvertValue(src any) error {
	_ = "STUB: not implemented"
	// we try to convertAssign values that we can't directly assign because ConvertValue
	// will return immediately for v that is already a Value, even if it is a different
	// Value type than the one we expect here.
	return nil
}
