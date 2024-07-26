package stringutil

import "testing"

func TestDelimit(t *testing.T) {
	tests := []struct {
		name string
		sl   []string
		sep  string
		last []string
		want string
	}{
		{"nil slice", nil, "", nil, ""},
		{"nil slice with sep", nil, ", ", nil, ""},
		{"nil slice with lsep", nil, "", []string{" and "}, ""},
		{"nil slice with both", nil, ", ", []string{" and "}, ""},
		{"len == 0", []string{}, "", nil, ""},
		{"len == 0 with sep", []string{}, ", ", nil, ""},
		{"len == 0 with lsep", []string{}, "", []string{" and "}, ""},
		{"len == 0 with both", []string{}, ", ", []string{" and "}, ""},
		{"len == 1", []string{"a"}, "", nil, "a"},
		{"len == 1 with sep", []string{"a"}, ", ", nil, "a"},
		{"len == 1 with lsep", []string{"a"}, "", []string{" and "}, "a"},
		{"len == 1 with both", []string{"a"}, ", ", []string{" and "}, "a"},
		{"len == 2", []string{"a", "b"}, "", nil, "ab"},
		{"len == 2 with sep", []string{"a", "b"}, ", ", nil, "a, b"},
		{"len == 2 with lsep", []string{"a", "b"}, "", []string{" and "}, "a and b"},
		{"len == 2 with both", []string{"a", "b"}, ", ", []string{" and "}, "a and b"},
		{"len == 3", []string{"a", "b", "c"}, "", nil, "abc"},
		{"len == 3 with sep", []string{"a", "b", "c"}, ", ", nil, "a, b, c"},
		{"len == 3 with lsep", []string{"a", "b", "c"}, "", []string{" and "}, "ab and c"},
		{"len == 3 with both", []string{"a", "b", "c"}, ", ", []string{" and "}, "a, b and c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.last) == 0 {
				if got := Delimit(tt.sl, tt.sep); got != tt.want {
					t.Errorf("Delimit() = %v, want %v", got, tt.want)
				}
			} else {
				if got := Delimit(tt.sl, tt.sep, tt.last[0]); got != tt.want {
					t.Errorf("Delimit() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
