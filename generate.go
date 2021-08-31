//go:generate go run github.com/mna/pigeon@latest -o grammars/sample_logic_pigeon/sample_logic.peg.go grammars/sample_logic_pigeon/sample_logic.peg
//go:generate go run github.com/pointlander/peg@latest -switch grammars/sample_logic_peg/sample_logic.peg
//go:generate go run github.com/pointlander/peg@latest -switch grammars/golang/golang.peg
//go:generate go test -v ./...
package main