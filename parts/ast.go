//   Copyright (C) 2022 Ashwin Godbole
//
//   This file is part of simla.
//
//   simla is free software: you can redistribute it and/or modify
//   it under the terms of the GNU General Public License as published by
//   the Free Software Foundation, either version 3 of the License, or
//   (at your option) any later version.
//
//   simla is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
//   GNU General Public License for more details.
//
//   You should have received a copy of the GNU General Public License
//   along with simla. If not, see <https://www.gnu.org/licenses/>.

package parts

// import "fmt"

type Node interface {
	NodeName() string
}

type ExprNode interface {
    Node
    IsExpr()
}

type StmtNode interface {
    Node
    IsStmt()
}

type TypeKind uint8
type Kind uint8

const (
    TInt TypeKind = iota
    TReal
    TChar
    TString
    TBool
    TVoid
    TFunc
    TArray
)

const (
    SDecl Kind = iota
    SExpr
    SCond
    SLoop
    SPrint
    SReturn
)

const (
    EAdd Kind = iota
    ESub
    EMul
    EDiv
    ERem

    EAnd
    EOr
    ENot
    ELst
    ELse
    EGrt
    EGre
    EEql
    ENeq

    EInt
    EReal
    EBool
    EStr
    EChar

    EVar
    EFnCall
)

type ParamList struct {
    Name string
    Type *Type
    Next *ParamList
}

type Type struct {
    TypeKind TypeKind
    Subtype  *Type
    Params   *ParamList
}

type Decl struct {
    Name  string
    Type  *Type
    Value *Node
    Code  *Node
    Next  *Decl
}

type Stmt struct {
    Kind     Kind
    Decl     *Decl
    InitExpr *Node
    Expr     *Node
    NextExpr *Node
    Body     *Stmt
    ElseBody *Stmt
    Next     *Stmt
}

type Expr struct {
    Kind    Kind
    Left    *Expr
    Right   *Expr
    Var     string
    Ival    int32
    Rval    float32
    Cval    byte
    Bval    bool
    Sval    string
}

func (o* Decl) NodeName() string {
    return "Decl: " + o.Name
}

func (o* Stmt) NodeName() string {
    return "Statement"
}

func (o* Expr) NodeName() string {
    return "Expression"
}

func (o* Decl) IsStmt() {}
func (o* Stmt) IsStmt() {}
func (o* Expr) IsExpr() {}

// procedures to make it easy to create ast nodes

func CreateDecl (
    name string, 
    dtype *Type,
    value *Node,
    code *Node, 
    next *Decl,
) *Decl {
    d := &Decl { 
        Name: name, 
        Type: dtype,
        Value: value, 
        Code: code,
        Next: next,
    }
    return d
}

func CreateStmt (
    kind Kind,
    decl *Decl,
    initExpr *Node,
    expr *Node,
    nextExpr *Node,
    body *Stmt,
    elseBody *Stmt,
    next *Stmt,
) *Stmt {
    s := &Stmt {
        Kind: kind,
        Decl: decl,
        InitExpr: initExpr,
        Expr: expr,
        NextExpr: nextExpr,
        Body: body,
        ElseBody: elseBody,
        Next: next,
    }
    return s
}

func CreateVarExpr(name string) *Expr {
    e := &Expr{ Var: name }
    return e
}

func CreateILitExpr(ival int32) *Expr {
    e := &Expr{ Ival: ival }
    return e
}

func CreateRLitExpr(rval float32) *Expr {
    e := &Expr{ Rval: rval }
    return e
}

func CreateCLitExpr(cval byte) *Expr {
    e := &Expr{ Cval: cval }
    return e
}

func CreateSLitExpr(sval string) *Expr {
    e := &Expr{ Sval: sval }
    return e
}

func CreateBLitExpr(bval bool) *Expr {
    e := &Expr{ Bval: bval }
    return e
}

func CreateBinExpr(kind Kind, left *Expr, right *Expr) *Expr {
    e := &Expr{ Kind: kind, Left: left, Right: right }
    return e
}

func CreateUniExpr(kind Kind, left *Expr) *Expr {
    e := &Expr{ Kind: kind, Left: left }
    return e
}
