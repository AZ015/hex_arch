package arithmetic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdapter_Addition(t *testing.T) {
	tests := []struct {
		name    string
		a       int32
		b       int32
		want    int32
		wantErr bool
	}{
		{
			"addition success",
			1,
			1,
			2,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := NewAdapter()
			got, err := arith.Addition(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestAdapter_Subtraction(t *testing.T) {
	tests := []struct {
		name    string
		a       int32
		b       int32
		want    int32
		wantErr bool
	}{
		{
			"subtraction success",
			10,
			1,
			9,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := NewAdapter()
			got, err := arith.Subtraction(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestAdapter_Multiplication(t *testing.T) {
	tests := []struct {
		name    string
		a       int32
		b       int32
		want    int32
		wantErr bool
	}{
		{
			"multiplication success",
			10,
			2,
			20,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := NewAdapter()
			got, err := arith.Multiplication(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}

func TestAdapter_Division(t *testing.T) {
	tests := []struct {
		name    string
		a       int32
		b       int32
		want    int32
		wantErr bool
	}{
		{
			"division success",
			10,
			2,
			5,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arith := NewAdapter()
			got, err := arith.Division(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
