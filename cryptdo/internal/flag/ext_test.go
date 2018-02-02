package flag

import "testing"

func TestExt(t *testing.T) {
	var tcs = []struct {
		input, output string
	}{
		{input: ".ext", output: ".ext"},
		{input: "ext", output: ".ext"},
	}

	for _, tc := range tcs {
		ext := new(Ext)

		if err := ext.UnmarshalFlag(tc.input); err != nil {
			t.Fatalf("unexpected unmarshalling error: %q", err)
		}

		if tc.output != string(*ext) {
			t.Errorf("expected extension %q when the user uses `--extension %q` as the flag, got: %q", tc.output, tc.input, *ext)
		}
	}
}

func TestEmptyExt(t *testing.T) {
	ext := new(Ext)

	if err := ext.UnmarshalFlag(""); err != ErrEmptyExt {
		t.Errorf("expected error %q, got %+v", ErrEmptyExt, err)
	}
}
