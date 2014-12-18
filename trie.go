package mercury

import (
	//"errors"
	"strings"
)

type node struct {
	children     []*node
	path         string
	handler      Handle
	component    string
	isNamedParam bool
}

func (n *node) addNode(path string, handler Handle) {
	components := strings.Split(path, "/")[1:]
	aNode, component := n.traverse(components, nil)
	if aNode.path == path {
		panic("node has already been added to tree")
	}
	newNode := node{path: path, handler: handler, component: component, isNamedParam: false}
	if component[0] == ':' {
		newNode.isNamedParam = true
	}
	aNode.children = append(aNode.children, &newNode)
}

func (n *node) traverse(components []string, params map[string]string) (*node, string) {
	component := components[0]
	if len(n.children) > 0 {
		for _, child := range n.children {
			if component == child.component || child.isNamedParam {
				if child.isNamedParam && params != nil {
					params[child.component[1:]] = component
				}
				next := components[1:]
				if len(next) > 0 {
					return child.traverse(next, params)
				} else {
					return child, component
				}
			}
		}
	}
	return n, component
}
