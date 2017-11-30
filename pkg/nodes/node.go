package nodes

// HTMLNode represents a single HTML element.
type HTMLNode struct {
	Type       string
	Parent     *HTMLNode
	Children   []*HTMLNode
	Attributes NodeAttributes
}

// NodeAttributes object stores attributes for a HTMLNode object.
type NodeAttributes struct {
	ID                string
	Classes           map[string]v
	BooleanAttributes map[string]v
	CustomAttributes  map[string]string
}

type v struct{}

// NewNode returns a new HTMLNode with no Parent.
func NewNode() *HTMLNode {

	n := new(HTMLNode)
	n.Children = make([]*HTMLNode, 0)
	n.Attributes.Classes = make(map[string]v)
	n.Attributes.BooleanAttributes = make(map[string]v)
	n.Attributes.CustomAttributes = make(map[string]string)

	return n
}

// NewChild creates a new HTMLNode object and attaches it to the parent node.
func (n *HTMLNode) NewChild() *HTMLNode {

	child := NewNode()
	child.Parent = n
	n.Children = append(n.Children, child)

	return child
}

// SetID sets the ID of the node.
func (n *HTMLNode) SetID(s string) {
	n.Attributes.ID = s
}

// AddClass adds a class to to the object.
func (n *HTMLNode) AddClass(s string) {
	n.Attributes.Classes[s] = v{}
}

// SetAttribute sets attribute s1 to the value of s2.
// Overwrites existing value.
func (n *HTMLNode) SetAttribute(s1, s2 string) {
	n.Attributes.CustomAttributes[s1] = s2
}

// SetBooleanAttribute sets a boolean attribute (key with no value).
func (n *HTMLNode) SetBooleanAttribute(s string) {
	n.Attributes.BooleanAttributes[s] = v{}
}
