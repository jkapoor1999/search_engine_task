package models

import "testing"

func TestPagesByScore_Len(t *testing.T) {
	tests := []struct {
		name string
		u    PagesByScore
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Len(); got != tt.want {
				t.Errorf("PagesByScore.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPagesByScore_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		u    PagesByScore
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.u.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestPagesByScore_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		u    PagesByScore
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("PagesByScore.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}
