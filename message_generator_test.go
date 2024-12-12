package mewtocol

import (
	"fmt"
	"testing"
)

func Test_address(t *testing.T) {
	type args struct {
		ad uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				ad: 1,
			},
			want: "01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := address(tt.args.ad)
			fmt.Println("got", got)

			if got != tt.want {
				t.Errorf("address() = %v, want %v", got, tt.want)
			}
		})
	}
}
