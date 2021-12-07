package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFixedXOR(t *testing.T)  {
	tests := [] struct {
		name string
		a    string
		b    string
		want string
	}{
		{
			name: "Test01",
			a:    "1c0111001f010100061a024b53535009181c",
			b:    "686974207468652062756c6c277320657965",
			want: "746865206b696420646f6e277420706c6179",
		},
	}

	for _,  tt :=  range tests {
		got,err := FixedXOR(tt.a, tt.b)
		require.NoError(t, err)
		require.Equal(t, tt.want, got)
	}
}