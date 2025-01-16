package main

import "reflect"

type Op interface {
	Name() string
	Arity() int
	SymbolTemplate() string
	TypeClass() TypeClass
	IsFunc() bool
}

type TypedOp interface {
	Op
	Kind() reflect.Kind
}

type BinOp interface {
	Op
	IsBinOp()
}

type UnaryOp interface {
	Op
	IsUnaryOp()
}

type RetSamer interface {
	RetSame() bool
}

type basicBinOp struct {
	symbol string
	name   string
	isFunc bool
	is     TypeClass
}

type arithOp struct {
	basicBinOp
	TypeClassName   string
	HasIdentity     bool
	Identity        int
	IsInv           bool
	Inv             string
	IsCommutative   bool
	IsInvolutionary bool
}

type cmpOp struct {
	basicBinOp
	TypeClassName string
	Inv string
	IsTransitive bool
	IsSymmetric bool
}

func (op basicBinOp) Name() string           { return op.name }
func (op basicBinOp) Arity() int             { return 2 }
func (op basicBinOp) SymbolTemplate() string { return op.symbol }
func (op basicBinOp) TypeClass() TypeClass   { return op.is }
func (op basicBinOp) IsFunc() bool           { return op.isFunc }
func (op basicBinOp) IsBinOp()               {}

type TypedBinOp struct {
	BinOp
	k reflect.Kind
}

// IsFunc contains special conditions
func (op TypedBinOp) IsFunc() bool {
	if op.Name() == "Mod" && isFloatCmplx(op.k) {
		return true
	}
	return op.BinOp.IsFunc()
}

func (op TypedBinOp) Kind() reflect.Kind { return op.k }

type unaryOp struct {
	symbol string
	name   string
	isfunc bool
	is     TypeClass

	TypeClassName string // mainly used in tests
	Inv string
}

func (op unaryOp) Name() string           { return op.name }
func (op unaryOp) Arity() int             { return 1 }
func (op unaryOp) SymbolTemplate() string { return op.symbol }
func (op unaryOp) TypeClass() TypeClass   { return op.is }
func (op unaryOp) IsFunc() bool           { return op.isfunc }
func (op unaryOp) IsUnaryOp()             {}

type TypedUnaryOp struct {
	UnaryOp
	k reflect.Kind
}

func (op TypedUnaryOp) Kind() reflect.Kind { return op.k }

type specialUnaryOp struct {
	unaryOp
	additionalParams []string
}

type Level int

const (
	Basic Level = iota
	InternalE
	StdEng
	Dense
	API
)
