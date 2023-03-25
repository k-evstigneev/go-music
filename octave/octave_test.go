package octave

import (
	"reflect"
	"testing"
)

func Test_iterate(t *testing.T) {
	type args struct {
		cursor *Note
		times  int
	}

	octaveLoop := CreateOctaveLoop()

	tests := []struct {
		name string
		args args
		want *Note
	}{
		{name: "1 iteration", args: args{octaveLoop.root, 1}, want: octaveLoop.root.Next},
		{name: "2 iterations", args: args{octaveLoop.root, 2}, want: octaveLoop.root.Next.Next},
		{name: "12 iterations", args: args{octaveLoop.root, 12}, want: octaveLoop.root},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iterate(tt.args.cursor, tt.args.times); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}
