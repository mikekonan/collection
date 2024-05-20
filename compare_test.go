package collection_test

import (
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestMin(t *testing.T) {
	cases := []struct {
		name string
		l    int
		r    int
		want int
	}{
		{name: "positive numbers", l: 5, r: 10, want: 5},
		{name: "negative numbers", l: -5, r: -10, want: -10},
		{name: "same numbers", l: 10, r: 10, want: 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Min(tc.l, tc.r)
			if got != tc.want {
				t.Errorf("Min(%d, %d) = %d; want %d", tc.l, tc.r, got, tc.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		name string
		l    int
		r    int
		want int
	}{
		{name: "positive numbers", l: 5, r: 10, want: 10},
		{name: "negative numbers", l: -5, r: -10, want: -5},
		{name: "same numbers", l: 10, r: 10, want: 10},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.Max(tc.l, tc.r)
			if got != tc.want {
				t.Errorf("Max(%d, %d) = %d; want %d", tc.l, tc.r, got, tc.want)
			}
		})
	}
}

func TestMinOf(t *testing.T) {
	cases := []struct {
		name     string
		elements []int
		want     int
	}{
		{name: "positive numbers", elements: []int{5, 10, 3}, want: 3},
		{name: "negative numbers", elements: []int{-5, -10, -3}, want: -10},
		{name: "same numbers", elements: []int{10, 10, 10}, want: 10},
		{name: "empty slice", elements: []int{}, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MinOf(tc.elements...)
			if got != tc.want {
				t.Errorf("MinOf(%v) = %d; want %d", tc.elements, got, tc.want)
			}
		})
	}
}

func TestMaxOf(t *testing.T) {
	cases := []struct {
		name     string
		elements []int
		want     int
	}{
		{name: "positive numbers", elements: []int{5, 10, 3}, want: 10},
		{name: "negative numbers", elements: []int{-5, -10, -3}, want: -3},
		{name: "same numbers", elements: []int{10, 10, 10}, want: 10},
		{name: "empty slice", elements: []int{}, want: 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := collection.MaxOf(tc.elements...)
			if got != tc.want {
				t.Errorf("MaxOf(%v) = %d; want %d", tc.elements, got, tc.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	cases := []struct {
		name string
		l    []int
		r    []int
		want bool
	}{
		{name: "equal", l: []int{5, 10, 3}, r: []int{5, 10, 3}, want: true},
		{name: "wrong order", l: []int{5, 10, 3}, r: []int{3, 10, 5}, want: false},
		{name: "empty", l: []int{}, r: nil, want: true},
		{name: "not equal", l: []int{1, 2, 3, 5, 10, 3}, r: []int{5, 10, 3}, want: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			got := collection.Equal(tc.l, tc.r)
			if got != tc.want {
				t.Errorf("Equal = %v; want %v", got, tc.want)
			}
		})
	}
}

func TestMapEqual(t *testing.T) {
	cases := []struct {
		name string
		l    map[string]int
		r    map[string]int
		want bool
	}{
		{name: "equal", l: map[string]int{"5": 5, "10": 10, "3": 3}, r: map[string]int{"5": 5, "10": 10, "3": 3}, want: true},
		{name: "different order", l: map[string]int{"5": 5, "10": 10, "3": 3}, r: map[string]int{"3": 3, "5": 5, "10": 10}, want: true},
		{name: "empty", l: map[string]int{}, r: nil, want: true},
		{name: "not equal", l: map[string]int{"5": 5, "10": 10, "3": 3}, r: map[string]int{"5": 5, "12": 12}, want: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			got := collection.MapEqual(tc.l, tc.r)
			if got != tc.want {
				t.Errorf("MapEqual = %v; want %v", got, tc.want)
			}
		})
	}
}
