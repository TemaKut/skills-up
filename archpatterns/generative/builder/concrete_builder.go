package main

import "strconv"

type RandomPositionStringBuilder struct {
	symbols map[string]struct{}
}

func NewRandomPositionStringBuilder() *RandomPositionStringBuilder {
	return &RandomPositionStringBuilder{
		symbols: make(map[string]struct{}),
	}
}

func (r *RandomPositionStringBuilder) SetStar() {
	r.symbols["*"] = struct{}{}
}

func (r *RandomPositionStringBuilder) SetNumber(uint65 uint64) {
	r.symbols[strconv.Itoa(int(uint65))] = struct{}{}
}

func (r *RandomPositionStringBuilder) SetString(str string) {
	r.symbols[str] = struct{}{}
}

func (r *RandomPositionStringBuilder) Build() string {
	var resultStr string

	for symbol, _ := range r.symbols {
		resultStr += symbol
	}

	return resultStr
}
