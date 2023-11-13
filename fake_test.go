package gox

import "testing"

func TestFakeAnimalName(t *testing.T) {
	tests := []struct {
		name       string
		want       string
		doNotCheck bool
	}{
		{
			name:       "success 1",
			want:       "",
			doNotCheck: true,
		},
		{
			name:       "success 2",
			want:       "",
			doNotCheck: true,
		},
		{
			name:       "success 3",
			want:       "",
			doNotCheck: true,
		},
		{
			name:       "success 4",
			want:       "",
			doNotCheck: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FakeAnimalName(); tt.doNotCheck {
				t.Logf("FakeAnimalName() = %v", got)
			} else if got != tt.want {
				t.Errorf("FakeAnimalName() = %v, want %v", got, tt.want)
			}
		})
	}
}
