package aspartasm

type ast struct {
	token
	children []ast
}

func newAst() *ast {
	return &ast{token{unknown, "root"}, nil}
}

//Appends the token to the tree as a child node.
//Returns the new child node.
func (a *ast) append(t token) *ast {
	child := ast{t, nil}
	a.children = append(a.children, child)
	return &child
}
