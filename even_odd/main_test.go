package main

import (
	"reflect"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []string
	}{
		{
			name: "Test 1",
			s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: []string{"odd", "even", "odd", "even", "odd", "even", "odd", "even", "odd", "even"},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOrOdd(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOrOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}