package main

import (
	"fmt"
	"strings"
)

func ParseFilter(s string) (Filter, error) {
	// TODO implement proper syntax tree parsing
	var err error
	exprs := strings.Split(s, "&&")
	and := &And{make([]Filter, len(exprs))}
	for i, e := range exprs {
		and.exprs[i], err = ParseExpr(e)
		if err != nil {
			return nil, err
		}
	}
	return and, nil
}

func ParseExpr(s string) (Filter, error) {
	pos := strings.Index(s, "==")
	if pos < 0 {
		return nil, fmt.Errorf("Only == operator is supported")
	}
	ops := make([]string, 2)
	ops[0] = strings.TrimSpace(s[0:pos])
	ops[1] = strings.TrimSpace(s[pos+2 : len(s)])
	if ops[1][0] == '%' {
		ops[0], ops[1] = ops[1], ops[0]
	}
	templateHandler, ok := templateHandlers[ops[0]]
	if !ok {
		return nil, fmt.Errorf("Unknown flag in the filter: %v", ops[0])
	}
	return &Eq{templateHandler.getter(), ops[1]}, nil
}

type Filter interface {
	run(c *Contact) bool
}

type And struct {
	exprs []Filter
}

var _ Filter = &And{}

func (e *And) run(c *Contact) bool {
	for _, e := range e.exprs {
		if !e.run(c) {
			return false
		}
	}
	return true
}

type Eq struct {
	fieldGetter FieldGetter
	value       string
}

var _ Filter = &Eq{}

func (e *Eq) run(c *Contact) bool {
	e.fieldGetter.get(c)
	vv := &ValueVisitor{}
	e.fieldGetter.accept(vv)
	return vv.val == e.value
}
