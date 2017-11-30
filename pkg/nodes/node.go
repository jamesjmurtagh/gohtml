package nodes

// HTMLNode represents a single HTML element.
type HTMLNode struct {
	Type       string         `json:"type"`
	Children   []*HTMLNode    `json:"children, omitempty"`
	Attributes NodeAttributes `json:"attributes"`
}

// NodeAttributes object stores attributes for a HTMLNode object.
type NodeAttributes struct {
	ID                string            `json:"id,omitempty"`
	Classes           map[string]v      `json:"classes,omitempty"`
	BooleanAttributes map[string]v      `json:"boolean_attributes,omitempty"`
	CustomAttributes  map[string]string `json:"custom_attributes,omitempty"`
}

type v struct{}

// NewNode returns a new HTMLNode with no Parent.
func NewNode(s string) *HTMLNode {

	n := new(HTMLNode)
	n.Type = s
	n.Children = make([]*HTMLNode, 0)
	n.Attributes.Classes = make(map[string]v)
	n.Attributes.BooleanAttributes = make(map[string]v)
	n.Attributes.CustomAttributes = make(map[string]string)

	return n
}

// NewChild creates a new HTMLNode object and attaches it to the parent node.
func (n *HTMLNode) NewChild(s string) *HTMLNode {

	child := NewNode(s)
	n.Children = append(n.Children, child)

	return child
}

// SetID sets the ID of the node.
func (n *HTMLNode) SetID(s string) {
	n.Attributes.ID = s
}

// SetType sets the element type (ie. html, head, body, etc.)
func (n *HTMLNode) SetType(s string) {
	n.Type = s
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
