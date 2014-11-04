package aspartasm

type AST struct {
	Token    `json:"Token"`
	Children []AST
}

func newAst() *AST {
	return &AST{Token{unknown, "root"}, nil}
}

//Appends the Token to the tree as a child node.
//Returns the new child node.
func (a *AST) append(t Token) *AST {
	child := AST{t, nil}
	a.Children = append(a.Children, child)
	return &a.Children[len(a.Children)-1]
}

func (a *AST) appendAll(ts ...Token) {
	for _, t := range ts {
		a.append(t)
	}
}
