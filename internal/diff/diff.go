package diff

type NodeType string

const (
	NodeAdded     NodeType = "added"
	NodeRemoved   NodeType = "removed"
	NodeChanged   NodeType = "changed"
	NodeUnchanged NodeType = "unchanged"
	NodeNested    NodeType = "nested"
)

type DiffNode struct {
	Key      string
	Type     NodeType
	Value    interface{}
	OldValue interface{}
	NewValue interface{}
	Children []DiffNode
}
