package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	node := l.Head
	for node != nil {
		if cond(node) {
			f(node)
		}
		node = node.Next
	}
}

func printNodeData(node *NodeL) {
	// Print data without using packages
	println(node.Data)
}

func main() {
	// Create a sample list
	list := &List{
		Head: &NodeL{Data: 5, Next: &NodeL{Data: -3, Next: &NodeL{Data: "Hello", Next: nil}}},
		Tail: nil,
	}

	// Apply IsPositiveNode to nodes with positive values
	println("Positive Nodes:")
	ListForEachIf(list, printNodeData, IsPositiveNode)

	// Apply IsAlNode to nodes with non-numeric data
	println("Non-Numeric Nodes:")
	ListForEachIf(list, printNodeData, IsAlNode)
}
