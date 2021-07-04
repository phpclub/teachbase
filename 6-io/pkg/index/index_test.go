package index

import (
	"reflect"
	"teachbase/6-io/pkg/crawler"
	"testing"
)

var (
	docst = []crawler.Document{
		{
			ID:    10,
			Title: "Первый Слон",
		},
		{
			ID:    20,
			Title: "Второй Слон",
		},
		{
			ID:    30,
			Title: "Третий Слониха",
		},
	}
)

func TestReverse_Add(t *testing.T) {
	type args struct {
		docs []crawler.Document
	}
	i := make(Reverse)
	tests := []struct {
		name string
		idx  Reverse
		args args
		want int
	}{
		{
			name: "Test fill", // Получаем 5 уникальных токенов
			idx:  i,
			args: args{docs: docst},
			want: 5,
		},
		{
			name: "Test duplicate token", // Повторное добавление одинаковой фразы не создает новый токен
			idx:  i,
			args: args{docs: append(docst, crawler.Document{
				ID:    40,
				Title: "Третий Слониха",
			},
			)},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.idx.Add(tt.args.docs)
			got := len(tt.idx)
			if got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestReverse_Search(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		idx  Reverse
		args args
		want [][]int
	}{
		{
			name: "Simple search",
			idx:  make(Reverse),
			args: args{s: "Слон"},
			want: [][]int{{10, 20}},
		},
	}
	for _, tt := range tests {
		tt.idx.Add(docst)
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.idx.Search(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tokenize(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "split words",
			args: args{text: "string for token"},
			want: []string{"string", "for", "token"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tokenize(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
