package main

import (
	"reflect"
	"testing"
)

func TestDeck_deal(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name  string
		d     Deck
		args  args
		want  Hand
		want1 Deck
	}{
		{"Test 1 Card", newDeck(), args{1}, Hand(newDeck()[:1]), newDeck()[1:]},
		{"Test 4 Cards", newDeck(), args{4}, Hand(newDeck()[:4]), newDeck()[4:]},
		{"Test 0 Cards", newDeck(), args{0}, Hand{}, newDeck()},
		{"Test -1 Cards", newDeck(), args{-1}, Hand{}, newDeck()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.d.deal(tt.args.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deck.deal() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Deck.deal() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
