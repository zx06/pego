package sample_logic_pigeon

import (
	"strconv"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	Functions = map[string]func(string) bool{
		"eq1": func(s string) bool { return s == "1" },
		"ne1": func(s string) bool { return s != "1" },
		"gt5": func(s string) bool { r, _ := strconv.Atoi(s); return r > 5 },
		"lt5": func(s string) bool { r, _ := strconv.Atoi(s); return r < 5 },
	}
	type args struct {
		rule string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				rule: "(eq1('1') and ne1('1')) or lt5('4')",
			},
			want: true,
		},
		{
			args: args{
				rule: `(((eq1('1') and ne1('1') and gt5('123')) or lt5('4')) and gt5('6') ) or (gt5("12313") and lt5('100'))`,
			},
			want: true,
		},
		{
			args: args{
				rule: `((eq1('1') and ne1('1')) or lt5('4')) and gt5("6")`,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseReader("", strings.NewReader(tt.args.rule))
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Errorf("%s = %v, want %v", tt.args.rule, got, tt.want)
			}

		})
	}
}
