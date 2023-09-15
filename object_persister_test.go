package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_save_object(t *testing.T) {
	type args struct {
		filename string
		s        Serializer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Save deck test", args{"./test_files/deckText.txt", newDeck()}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Remove(tt.args.filename)
			if err := save_object(tt.args.filename, tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("save_object() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_load_object(t *testing.T) {
	type args struct {
		filename string
		s        Serializer
	}
	tests := []struct {
		name string
		args args
		want Serializer
	}{
		// TODO: Add test cases.
		{"Success test", args{"./test_files/full_deck.txt", Deck{}}, newDeck()},
		{"Fail test", args{"./test_files/fail.txt", Deck{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := load_object(tt.args.filename, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("load_object() = %v, want %v", got, tt.want)
			}
		})
	}
}
