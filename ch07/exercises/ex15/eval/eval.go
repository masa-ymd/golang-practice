package eval

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

type ternary struct {
	op1, op2 rune
	x, y, z  Expr
}

func (t ternary) Eval(env Env) float64 {
	switch t.op1 {
	case '?':
		switch t.op2 {
		case ':':
			x := t.x.Eval(env)
			if math.IsInf(x, 0) || math.IsNaN(x) || x == 0 {
				return t.z.Eval(env)
			}
			return t.y.Eval(env)
		}
		panic(fmt.Sprintf("unsupported ternary operator: %q", t.op2))
	}
	panic(fmt.Sprintf("unsupported ternary operator: %q", t.op1))
}

func (t ternary) String() string {
	return fmt.Sprintf("(%s %c %s %c %s)", t.x, t.op1, t.y, t.op2, t.z)
}

func (t ternary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("?:", t.op1) {
		return fmt.Errorf("unexpected ternary op %q", t.op1)
	} else if !strings.ContainsRune("?:", t.op2) {
		return fmt.Errorf("unexpected ternary op %q", t.op2)
	}
	if err := t.x.Check(vars); err != nil {
		return err
	} else if err := t.y.Check(vars); err != nil {
		return err
	}
	return t.z.Check(vars)
}

type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unexpected binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return string(u.op) + string(u.x.String())
}

func (b binary) String() string {
	return fmt.Sprintf("%s %s %s", b.x.String(), string(b.op), b.y.String())
}

func (c call) String() string {
	return "hoge"
}

var numParams = map[string]int{"pow": 2, "sin": 1, "sqrt": 1}
