package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertHexToBase64(t *testing.T)  {
	tests := [] struct {
		name string
		s    string
		want string
	}{
		{
			name: "Test01",
			s:    "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			want: "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		},
	}

	for _,  tt :=  range tests {
		got := ConvertHexToBase64(tt.s)
		require.Equal(t, tt.want, got)
	}
}