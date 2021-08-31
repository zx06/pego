package sample_logic_peg

import (
	"fmt"
	"testing"
)

func TestSampleLogic(t *testing.T) {
	sl:=&SampleLogic{
		Buffer: "(eq1('1') and ne1('1')) or lt5('4')",
	}
	sl.Init()
	sl.Parse()
	fmt.Println(sl.AST().String())
}