package nodes

import (
	"fmt"
	"strings"
)

// Parse takes a HTMLNode and parses to HTML, returning a byte array.
func Parse(n *HTMLNode) []byte {

	var out []byte
	out = parseNodes(n, 0)
	return out
}

func parseNodes(n *HTMLNode, indentLevel int) []byte {

	var indent string
	for i := 0; i < indentLevel; i++ {
		indent += "\t"
	}
	indentLevel++

	// Handle typeInnerHTML elements
	if n.Type == typeInnerHTML {
		return []byte(fmt.Sprintf("%s%s\n", indent, n.Attributes.ID))
	}

	arr := make([]byte, 0)
	var classes string
	for c := range n.Attributes.Classes {
		classes += fmt.Sprintf(" %s", c)
	}
	classes = strings.TrimSpace(classes)

	var customAttributes string
	for key, value := range n.Attributes.CustomAttributes {
		customAttributes += fmt.Sprintf(" %s=\"%s\"", key, value)
	}
	customAttributes = strings.TrimSpace(customAttributes)

	var boolAttributes string
	for bA := range n.Attributes.BooleanAttributes {
		boolAttributes += fmt.Sprintf(" %s", bA)
	}
	boolAttributes = strings.TrimSpace(boolAttributes)

	elementStart := fmt.Sprintf("<%s id=\"%s\" class=\"%s\" %s %s", n.Type, n.Attributes.ID, classes, customAttributes, boolAttributes)
	elementStart = fmt.Sprintf("%s%s>\n", indent, strings.TrimSpace(elementStart))

	arr = append(arr, []byte(elementStart)...)
	for _, cN := range n.Children {
		arr = append(arr, parseNodes(cN, indentLevel)...)
	}

	arr = append(arr, []byte(fmt.Sprintf("%s</%s>\n", indent, n.Type))...)
	return arr
}
