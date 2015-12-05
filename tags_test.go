package apidemic

import (
	"testing"
)

func TestTags(t *testing.T) {
	src := "characters_n,max=30"
	sample := []struct {
		tag, value string
	}{
		{"type", "characters_n"},
		{"max", "30"},
	}

	tags := make(Tags)
	tags.Load(src)
	for _, v := range sample {
		k, ok := tags.Get(v.tag)
		if !ok {
			t.Errorf("expected %s to exist %#v", v.tag, tags)
		}
		if k != v.value {
			t.Errorf("expected %s got %s", v.value, k)
		}
	}

	max, err := tags.Int("max")
	if err != nil {
		t.Fatal(err)
	}
	if max != 30 {
		t.Errorf("expected %d got %d", 30, max)
	}
}
