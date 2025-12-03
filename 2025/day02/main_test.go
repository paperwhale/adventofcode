package main

import "testing"

var example = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224," +
	"1698522-1698528,446443-446449,38593856-38593862,565653-565659," +
	"824824821-824824827,2121212118-2121212124"

func TestPart1(t *testing.T) {
	if got, want := Part1(example), 1227775554; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := Part2(example), 4174379265; got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
