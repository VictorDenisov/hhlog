package main

import (
	"fmt"
	"strings"
)

func ParseFilter(s string) (Filter, error) {
	if len(strings.TrimSpace(s)) == 0 {
		return &True{}, nil
	}
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

var (
	flipSign = map[string]string{
		"==": "==",
		"<=": ">=",
		">=": "<=",
	}
)

func ParseExpr(s string) (Filter, error) {
	pos := strings.Index(s, "==")
	sign := "=="
	if pos >= 0 {
		goto found
	}
	pos = strings.Index(s, "<=")
	if pos >= 0 {
		sign = "<="
		goto found
	}
	pos = strings.Index(s, ">=")
	if pos >= 0 {
		sign = ">="
		goto found
	}
	const errorMessage = "Only ==, <=, >= operators are supported"
	return nil, fmt.Errorf(errorMessage)
found:
	ops := make([]string, 2)
	ops[0] = strings.TrimSpace(s[0:pos])
	ops[1] = strings.TrimSpace(s[pos+2 : len(s)])
	if ops[1][0] == '%' {
		ops[0], ops[1] = ops[1], ops[0]
		sign = flipSign[sign]
	}
	templateHandler, ok := templateHandlers[ops[0]]
	if !ok {
		return nil, fmt.Errorf("Unknown flag in the filter: %v", ops[0])
	}
	switch sign {
	case "==":
		return &Eq{templateHandler.getter(), ops[1]}, nil
	case "<=":
		return &EqLess{templateHandler.getter(), ops[1]}, nil
	case ">=":
		return &EqMore{templateHandler.getter(), ops[1]}, nil
	}
	return nil, fmt.Errorf(errorMessage)

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

type EqLess struct {
	fieldGetter FieldGetter
	value       string
}

var _ Filter = &EqLess{}

func (e *EqLess) run(c *Contact) bool {
	e.fieldGetter.get(c)
	vv := &ValueVisitor{}
	e.fieldGetter.accept(vv)
	return vv.val <= e.value
}

type EqMore struct {
	fieldGetter FieldGetter
	value       string
}

var _ Filter = &EqMore{}

func (e *EqMore) run(c *Contact) bool {
	e.fieldGetter.get(c)
	vv := &ValueVisitor{}
	e.fieldGetter.accept(vv)
	return vv.val >= e.value
}

type True struct {
}

var _ Filter = &True{}

func (t *True) run(c *Contact) bool {
	return true
}
