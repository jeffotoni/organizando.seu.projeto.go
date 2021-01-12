package fmts

import "testing"

func TestConcat(t *testing.T) {

	var many1Int []interface{}
	var many1String []interface{}

	many1Int = append(many1Int, []int{100, 98, 23})
	many1String = append(many1String, []string{"22", "jeff", "beta1"})

	type args struct {
		strs []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test_concat_", args{strs: many1Int}, "1009823"},
		{"test_concat_", args{strs: many1String}, "22jeffbeta1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Concat(tt.args.strs...); got != tt.want {
				t.Errorf("Concat() = %v, want %v", got, tt.want)
			}
		})
	}
}
