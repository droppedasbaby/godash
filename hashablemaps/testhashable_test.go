package hashablemaps_test

type TestHashable struct {
	Value string
}

func (t TestHashable) Hash() string {
	return t.Value
}
