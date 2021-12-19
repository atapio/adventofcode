package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BinaryNode struct {
	parent *BinaryNode
	left   *BinaryNode
	right  *BinaryNode
	data   int
}

func (n *BinaryNode) Parse(input []string) ([]string, error) {
	var err error
	//fmt.Printf("Parse start: %s\n", input)

	if len(input) == 0 {
		return input, fmt.Errorf("end of input")
	}

	c, rest := input[0], input[1:]
	//fmt.Printf("c: %s rest: %v\n", c, rest)

	switch c {
	case "[":
		n.left = &BinaryNode{parent: n}
		rest, err = n.left.Parse(rest)
		if err != nil {
			return rest, err
		}

		c, rest = rest[0], rest[1:]
		if c != "," {
			fmt.Printf("c: %s rest: %v\n", c, rest)
			return rest, fmt.Errorf("invalid char %s", c)
		}

		n.right = &BinaryNode{parent: n}
		rest, err = n.right.Parse(rest)
		if err != nil {
			return rest, err
		}

		c, rest = rest[0], rest[1:]
		if c != "]" {
			fmt.Printf("c: %s rest: %v\n", c, rest)
			return rest, fmt.Errorf("invalid char %s", c)
		}
	default:
		d, err := strconv.Atoi(c)
		if err != nil {
			return rest, err
		}
		n.data = d
	}

	//fmt.Printf("node: %v\n", n)
	return rest, nil
}

func (n *BinaryNode) String() string {
	if n.left == nil && n.right == nil {
		return fmt.Sprintf("%d", n.data)
	}

	s := ""
	if n.left != nil {
		s = s + "[" + n.left.String()
	}

	if n.right != nil {
		s = s + "," + n.right.String() + "]"
	}
	return s
}

func (n *BinaryNode) Magnitude() int {
	if n.left == nil && n.right == nil {
		return n.data
	}

	v := 0
	v += 3*n.left.Magnitude() + 2*n.right.Magnitude()

	return v
}

func (n *BinaryNode) Explode(depth int) bool {
	//fmt.Printf("explode level %d: %v\n", depth, n)
	depth++
	if depth > 4 {
		if n.left == nil && n.right == nil {
			return false
		}
		fmt.Printf("exploding %v\n", n)
		leftVal := n.left.data
		rightVal := n.right.data
		n.AddLeft(leftVal)
		n.AddRight(rightVal)
		n.left = nil
		n.right = nil
		n.data = 0
		return true
	}
	if n.left != nil && n.left.Explode(depth) {
		return true
	}
	if n.right != nil && n.right.Explode(depth) {
		return true
	}
	return false
}

func (n *BinaryNode) AddLeft(v int) bool {
	// go up left node, when unvisited found, go down right child nodes until leaf
	parent := n.parent
	child := n

	for parent != nil {
		fmt.Printf("AddLeft %d parent: %v child: %v\n", v, n, n.parent)
		if parent.left != child {
			node := parent.left
			for node.left != nil || node.right != nil {
				node = node.right
			}
			if node.left == nil && node.right == nil {
				node.data += v
				return true
			}
			return false
		}
		child = parent
		parent = parent.parent
	}
	return false
}

func (n *BinaryNode) AddRight(v int) bool {
	parent := n.parent
	child := n

	for parent != nil {
		fmt.Printf("AddRight %d parent: %v child: %v\n", v, n, n.parent)
		if parent.right != child {
			node := parent.right
			for node.left != nil || node.right != nil {
				node = node.left
			}
			if node.left == nil && node.right == nil {
				node.data += v
				return true
			}
			return false
		}
		child = parent
		parent = parent.parent
	}
	return false
}

func (n *BinaryNode) Split() bool {
	if n.left == nil && n.right == nil {
		if n.data > 9 {
			// round down
			n.left = &BinaryNode{parent: n, data: n.data / 2}
			// round up
			n.right = &BinaryNode{parent: n, data: (n.data + 1) / 2}
			n.data = 0
			return true
		}
		return false
	}
	if n.left != nil && n.left.Split() {
		return true
	}
	if n.right != nil && n.right.Split() {
		return true
	}
	return false

}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) String() string {
	if t.root == nil {
		return "<nil>"
	}
	return t.root.String()
}

func (t *BinaryTree) Parse(line string) error {
	input := strings.Split(line, "")

	t.root = &BinaryNode{}
	remaining, err := t.root.Parse(input)
	if err != nil {
		return err
	}
	if len(remaining) != 0 {
		return fmt.Errorf("not all input parsed: %v", remaining)
	}

	fmt.Printf("input: %s\ntree: %v\n", line, t)

	return nil
}

func (t *BinaryTree) Magnitude() int {
	return t.root.Magnitude()
}

func (t *BinaryTree) Add(second *BinaryTree) {
	left := t.root
	t.root = &BinaryNode{left: left, right: second.root}
	t.root.left.parent = t.root
	t.root.right.parent = t.root

	fmt.Printf("after addition: %v\n", t)

	t.Reduce()
}

func (t *BinaryTree) Reduce() {
	for {
		if t.Explode() {
			fmt.Printf("after explode: %v\n", t)
			continue
		}

		if t.Split() {
			fmt.Printf("after split: %v\n", t)
			continue
		}
		break
	}
	fmt.Printf("Reduced: %v\n", t)
}

func (t *BinaryTree) Explode() bool {
	return t.root.Explode(0)
}

func (t *BinaryTree) Split() bool {
	return t.root.Split()
}

func findAnswer(lines []string) (int, error) {
	answer := 0

	line, rest := lines[0], lines[1:]

	bt := &BinaryTree{}
	err := bt.Parse(line)
	if err != nil {
		return 0, err
	}
	bt.Reduce()

	for _, line := range rest {
		second := &BinaryTree{}
		err := second.Parse(line)
		if err != nil {
			return 0, err
		}
		bt.Add(second)
	}

	fmt.Printf("result: %v\n", bt)

	answer = bt.Magnitude()

	return answer, nil
}

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	input := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}
	return input, nil
}

func main() {
	input, err := parseFile("input")
	if err != nil {
		log.Fatal(err)
	}

	answer, err := findAnswer(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("answer: %d\n", answer)
}
