package main

import "testing"

func TestFullName(t *testing.T) {
	tests := []struct {
		name   string
		values User
		want   string
	}{
		{
			name: "First values",
			values: User{
				FirstName: "Миша",
				LastName:  "Попов",
			},
			want: "Миша Попов",
		},
		{
			name: "Second values",
			values: User{
				FirstName: "Миша",
				LastName:  "",
			},
			want: "Миша ",
		},
		{
			name: "Third values",
			values: User{
				FirstName: "",
				LastName:  "Попов",
			},
			want: " Попов",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := (tt.values.FullName()); res != tt.want {
				t.Errorf("Error in test: %s. Waited: %s. Got: %s", tt.name, tt.want, res)
			}
		})
	}
}
