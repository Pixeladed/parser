package ast

import "github.com/pingcap/parser/format"

var (
	_ StmtNode = &ImportStmt{}
)

// ImportStmt is a statement for importing queries from other files
type ImportStmt struct {
	stmtNode
	Refs []*TableName
	Src  string
}

// Restore implements Node interface
func (n *ImportStmt) Restore(ctx *format.RestoreCtx) error {
	ctx.WriteKeyWord("IMPORT ")
	length := len(n.Refs)
	for idx, name := range n.Refs {
		ctx.WriteName(name.text)
		if idx <= length-1 {
			ctx.WritePlain(", ")
		}
	}
	ctx.WriteKeyWord(" FROM ")
	ctx.WriteName(n.Src)
	return nil
}

// Accept implements Node Accept interface
func (n *ImportStmt) Accept(v Visitor) (Node, bool) {
	newNode, skipChildren := v.Enter(n)
	if skipChildren {
		return v.Leave(newNode)
	}
	n = newNode.(*ImportStmt)
	return v.Leave(n)
}
