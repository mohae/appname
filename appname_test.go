package appname

import (
	"path/filepath"
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct{
		name string
		expected string
	}{
		{"test", "test"},
		{"test.exe", "test"},
		{filepath.Join("foo", "test"), "test"},
		{filepath.Join("foo", "test.exe"), "test"},
		{filepath.Join("foo", "bar", "test"), "test"},
		{filepath.Join("foo", "bar", "test.exe"), "test"},
		{filepath.Join("foo.bar"), "foo.bar"},
		{filepath.Join("foo.bar.exe"), "foo.bar"},
		{filepath.Join("biz", "foo.bar"), "foo.bar"},
		{filepath.Join("biz", "foo.bar.exe"), "foo.bar"},
	}

	for _, test := range tests {
		v := get(test.name)
		if v != test.expected {
			t.Errorf("got %s, want %s", v, test.expected)
		}
	}

	// now test Get
	s := Get()
	if s != "appname.test" {
		t.Errorf("got %s, want appname.test", s)
	}
}
