package parse_utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseSeed(t *testing.T) {
	type args struct {
		seedString string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "parse seed",
			args: args{
				seedString: "12345abcdef",
			},
			want:    1251004370415,
			wantErr: assert.NoError,
		},
		{
			name: "bad seed",
			args: args{
				seedString: "hello!!",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSeed(tt.args.seedString)
			if !tt.wantErr(t, err, fmt.Sprintf("ParseSeed(%v)", tt.args.seedString)) {
				return
			}
			assert.Equalf(t, tt.want, got, "ParseSeed(%v)", tt.args.seedString)
		})
	}

	t.Run("random seed", func(t *testing.T) {
		got, err := ParseSeed("")
		assert.NoError(t, err)
		assert.InDelta(t, time.Now().UnixNano(), got, 1_000_000)
	})
}
