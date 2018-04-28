package domaintree16

// !!!DON'T EDIT!!! Generated with infobloxopen/go-trees/etc from domaintree{{.bits}} with etc -s uint16 -d dtuintX.yaml -t ./domaintree\{\{.bits\}\}

import (
	"fmt"
	"testing"

	"github.com/pmezard/go-difflib/difflib"

	"github.com/infobloxopen/go-trees/domain"
)

func TestLabelNewTree(t *testing.T) {
	r := newLabelTree()
	assertLabelTree(r, TestEmptyTree, "empty tree", t)
}

func TestLabelInsert(t *testing.T) {
	var r *labelTree

	n := new(Node)

	r = r.insert("k", n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node tree", t)

	r = newLabelTree()
	r = r.insert("k", n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node tree", t)

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("1", nil)
	r = r.insert("2", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "tree 012", t)

	r = newLabelTree()
	r = r.insert("1", nil)
	r = r.insert("2", nil)
	r = r.insert("0", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "tree 120", t)

	r = newLabelTree()
	r = r.insert("2", nil)
	r = r.insert("0", nil)
	r = r.insert("1", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "tree 201", t)

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("2", nil)
	r = r.insert("1", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "tree 021", t)

	r = newLabelTree()
	r = r.insert("1", nil)
	r = r.insert("0", nil)
	r = r.insert("2", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "tree 102", t)

	r = newLabelTree()
	r = r.insert("2", nil)
	r = r.insert("1", nil)
	r = r.insert("0", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "tree 210", t)

	r = newLabelTree()
	r = r.insert("1", nil)
	r = r.insert("0", nil)
	r = r.insert("4", nil)
	r = r.insert("2", nil)
	r = r.insert("3", nil)
	assertLabelTree(r, TestFiveNodeTree, "tree 10423", t)

	r = newLabelTree()
	r = r.insert("f", nil)
	r = r.insert("e", nil)
	r = r.insert("d", nil)
	r = r.insert("c", nil)
	r = r.insert("b", nil)
	r = r.insert("a", nil)
	r = r.insert("9", nil)
	r = r.insert("8", nil)
	r = r.insert("7", nil)
	r = r.insert("6", nil)
	r = r.insert("5", nil)
	r = r.insert("4", nil)
	r = r.insert("3", nil)
	r = r.insert("2", nil)
	r = r.insert("1", nil)
	r = r.insert("0", nil)
	assertLabelTree(r, Test16InversedNodeTree, "tree inversed 16 nodes", t)

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("1", nil)
	r = r.insert("2", nil)
	r = r.insert("3", nil)
	r = r.insert("4", nil)
	r = r.insert("5", nil)
	r = r.insert("6", nil)
	r = r.insert("7", nil)
	r = r.insert("8", nil)
	r = r.insert("9", nil)
	r = r.insert("a", nil)
	r = r.insert("b", nil)
	r = r.insert("c", nil)
	r = r.insert("d", nil)
	r = r.insert("e", nil)
	r = r.insert("f", nil)
	assertLabelTree(r, Test16DirectNodeTree, "tree direct 16 nodes", t)

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("2", nil)
	r = r.insert("4", nil)
	r = r.insert("6", nil)
	r = r.insert("8", nil)
	r = r.insert("a", nil)
	r = r.insert("c", nil)
	r = r.insert("e", nil)
	r = r.insert("1", nil)
	r = r.insert("3", nil)
	r = r.insert("5", nil)
	r = r.insert("7", nil)
	r = r.insert("9", nil)
	r = r.insert("b", nil)
	r = r.insert("d", nil)
	r = r.insert("f", nil)
	assertLabelTree(r, Test16AlternatingNodeTree, "tree alternating 16 nodes", t)

	r = newLabelTree()
	r = r.insert("f", nil)
	r = r.insert("d", nil)
	r = r.insert("b", nil)
	r = r.insert("9", nil)
	r = r.insert("7", nil)
	r = r.insert("5", nil)
	r = r.insert("3", nil)
	r = r.insert("1", nil)
	r = r.insert("e", nil)
	r = r.insert("c", nil)
	r = r.insert("a", nil)
	r = r.insert("8", nil)
	r = r.insert("6", nil)
	r = r.insert("4", nil)
	r = r.insert("2", nil)
	r = r.insert("0", nil)
	assertLabelTree(r, Test16AlternatingInversedNodeTree, "tree alternating inversed 16 nodes", t)

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("3", nil)
	r = r.insert("6", nil)
	r = r.insert("9", nil)
	r = r.insert("c", nil)
	r = r.insert("f", nil)
	r = r.insert("1", nil)
	r = r.insert("2", nil)
	r = r.insert("4", nil)
	r = r.insert("5", nil)
	r = r.insert("7", nil)
	r = r.insert("8", nil)
	r = r.insert("a", nil)
	r = r.insert("b", nil)
	r = r.insert("d", nil)
	r = r.insert("e", nil)
	assertLabelTree(r, Test16_3AltNodeTree, "tree alternating by 3 16 nodes", t)

	r = newLabelTree()
	r = r.insert("00", nil)
	r = r.insert("02", nil)
	r = r.insert("04", nil)
	r = r.insert("06", nil)
	r = r.insert("08", nil)
	r = r.insert("0a", nil)
	r = r.insert("0c", nil)
	r = r.insert("0e", nil)
	r = r.insert("10", nil)
	r = r.insert("12", nil)
	r = r.insert("14", nil)
	r = r.insert("16", nil)
	r = r.insert("18", nil)
	r = r.insert("1a", nil)
	r = r.insert("1c", nil)
	r = r.insert("1e", nil)
	r = r.insert("01", nil)
	r = r.insert("03", nil)
	r = r.insert("05", nil)
	r = r.insert("07", nil)
	r = r.insert("09", nil)
	r = r.insert("0b", nil)
	r = r.insert("0d", nil)
	r = r.insert("0f", nil)
	r = r.insert("11", nil)
	r = r.insert("13", nil)
	r = r.insert("15", nil)
	r = r.insert("17", nil)
	r = r.insert("19", nil)
	r = r.insert("1b", nil)
	r = r.insert("1d", nil)
	r = r.insert("1f", nil)
	assertLabelTree(r, Test32AlternatingNodeTree, "tree with alternating 32 nodes", t)

	n1 := new(Node)
	n2 := new(Node)
	r = nil
	r = r.insert("k", n1)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n1), "tree with same node first insertion", t)
	r = r.insert("k", n2)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n2), "tree with same node second insertion", t)
}

func TestLabelInplaceInsert(t *testing.T) {
	n := new(Node)

	r := newLabelTree()

	r.inplaceInsert("k", n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node inplace tree", t)

	r = newLabelTree()
	r.inplaceInsert("k", n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node inplace tree", t)

	r = newLabelTree()
	r.inplaceInsert("0", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("2", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "inplace tree 012", t)

	r = newLabelTree()
	r.inplaceInsert("1", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("0", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "inplace tree 120", t)

	r = newLabelTree()
	r.inplaceInsert("2", nil)
	r.inplaceInsert("0", nil)
	r.inplaceInsert("1", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "inplace tree 201", t)

	r = newLabelTree()
	r.inplaceInsert("0", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("1", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "inplace tree 021", t)

	r = newLabelTree()
	r.inplaceInsert("1", nil)
	r.inplaceInsert("0", nil)
	r.inplaceInsert("2", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "inplace tree 102", t)

	r = newLabelTree()
	r.inplaceInsert("2", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("0", nil)
	assertLabelTree(r, TestThreeNodeTreeRed, "inplace tree 210", t)

	r = newLabelTree()
	r.inplaceInsert("1", nil)
	r.inplaceInsert("0", nil)
	r.inplaceInsert("4", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("3", nil)
	assertLabelTree(r, TestFiveNodeTree, "inplace tree 10423", t)

	r = newLabelTree()
	r.inplaceInsert("f", nil)
	r.inplaceInsert("e", nil)
	r.inplaceInsert("d", nil)
	r.inplaceInsert("c", nil)
	r.inplaceInsert("b", nil)
	r.inplaceInsert("a", nil)
	r.inplaceInsert("9", nil)
	r.inplaceInsert("8", nil)
	r.inplaceInsert("7", nil)
	r.inplaceInsert("6", nil)
	r.inplaceInsert("5", nil)
	r.inplaceInsert("4", nil)
	r.inplaceInsert("3", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("0", nil)
	assertLabelTree(r, Test16InversedNodeTree, "inplace tree inversed 16 nodes", t)

	r = newLabelTree()
	r.inplaceInsert("0", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("3", nil)
	r.inplaceInsert("4", nil)
	r.inplaceInsert("5", nil)
	r.inplaceInsert("6", nil)
	r.inplaceInsert("7", nil)
	r.inplaceInsert("8", nil)
	r.inplaceInsert("9", nil)
	r.inplaceInsert("a", nil)
	r.inplaceInsert("b", nil)
	r.inplaceInsert("c", nil)
	r.inplaceInsert("d", nil)
	r.inplaceInsert("e", nil)
	r.inplaceInsert("f", nil)
	assertLabelTree(r, Test16DirectNodeTree, "inplace tree direct 16 nodes", t)

	r = newLabelTree()
	r.inplaceInsert("0", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("4", nil)
	r.inplaceInsert("6", nil)
	r.inplaceInsert("8", nil)
	r.inplaceInsert("a", nil)
	r.inplaceInsert("c", nil)
	r.inplaceInsert("e", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("3", nil)
	r.inplaceInsert("5", nil)
	r.inplaceInsert("7", nil)
	r.inplaceInsert("9", nil)
	r.inplaceInsert("b", nil)
	r.inplaceInsert("d", nil)
	r.inplaceInsert("f", nil)
	assertLabelTree(r, Test16AlternatingNodeTree, "inplace tree alternating 16 nodes", t)

	r = newLabelTree()
	r.inplaceInsert("f", nil)
	r.inplaceInsert("d", nil)
	r.inplaceInsert("b", nil)
	r.inplaceInsert("9", nil)
	r.inplaceInsert("7", nil)
	r.inplaceInsert("5", nil)
	r.inplaceInsert("3", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("e", nil)
	r.inplaceInsert("c", nil)
	r.inplaceInsert("a", nil)
	r.inplaceInsert("8", nil)
	r.inplaceInsert("6", nil)
	r.inplaceInsert("4", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("0", nil)
	assertLabelTree(r, Test16AlternatingInversedNodeTree, "inplace tree alternating inversed 16 nodes", t)

	r = newLabelTree()
	r.inplaceInsert("0", nil)
	r.inplaceInsert("3", nil)
	r.inplaceInsert("6", nil)
	r.inplaceInsert("9", nil)
	r.inplaceInsert("c", nil)
	r.inplaceInsert("f", nil)
	r.inplaceInsert("1", nil)
	r.inplaceInsert("2", nil)
	r.inplaceInsert("4", nil)
	r.inplaceInsert("5", nil)
	r.inplaceInsert("7", nil)
	r.inplaceInsert("8", nil)
	r.inplaceInsert("a", nil)
	r.inplaceInsert("b", nil)
	r.inplaceInsert("d", nil)
	r.inplaceInsert("e", nil)
	assertLabelTree(r, Test16_3AltNodeTree, "inplace tree alternating by 3 16 nodes", t)

	r = newLabelTree()
	r.inplaceInsert("00", nil)
	r.inplaceInsert("02", nil)
	r.inplaceInsert("04", nil)
	r.inplaceInsert("06", nil)
	r.inplaceInsert("08", nil)
	r.inplaceInsert("0a", nil)
	r.inplaceInsert("0c", nil)
	r.inplaceInsert("0e", nil)
	r.inplaceInsert("10", nil)
	r.inplaceInsert("12", nil)
	r.inplaceInsert("14", nil)
	r.inplaceInsert("16", nil)
	r.inplaceInsert("18", nil)
	r.inplaceInsert("1a", nil)
	r.inplaceInsert("1c", nil)
	r.inplaceInsert("1e", nil)
	r.inplaceInsert("01", nil)
	r.inplaceInsert("03", nil)
	r.inplaceInsert("05", nil)
	r.inplaceInsert("07", nil)
	r.inplaceInsert("09", nil)
	r.inplaceInsert("0b", nil)
	r.inplaceInsert("0d", nil)
	r.inplaceInsert("0f", nil)
	r.inplaceInsert("11", nil)
	r.inplaceInsert("13", nil)
	r.inplaceInsert("15", nil)
	r.inplaceInsert("17", nil)
	r.inplaceInsert("19", nil)
	r.inplaceInsert("1b", nil)
	r.inplaceInsert("1d", nil)
	r.inplaceInsert("1f", nil)
	assertLabelTree(r, Test32AlternatingNodeTree, "inplace tree with alternating 32 nodes", t)

	r = nil
	assertPanic(func() { r.inplaceInsert("00", n) }, "nil tree inplace insertion", t)
}

func TestLabelGet(t *testing.T) {
	var r *labelTree

	v, ok := r.get("0")
	if ok {
		t.Errorf("Expected nothing but got %T (%#v)", v, v)
	}

	n0 := new(Node)
	n1 := new(Node)
	n2 := new(Node)
	n3 := new(Node)
	n4 := new(Node)

	r = newLabelTree()
	r = r.insert("1", n1)
	r = r.insert("0", n0)
	r = r.insert("4", n4)
	r = r.insert("2", n2)
	r = r.insert("3", n3)

	v, ok = r.get("3")
	if !ok {
		t.Errorf("Expected %p but got nothing", n3)
	} else if v != n3 {
		t.Errorf("Expected %p but got %p", n3, v)
	}

	v, ok = r.get("f")
	if ok {
		t.Errorf("Expected nothing but got %p", v)
	}
}

func TestLabelEnumerate(t *testing.T) {
	var r *labelTree

	assertEnumerate(r.enumerate(), "empty tree", t)

	n0 := new(Node)
	n1 := new(Node)
	n2 := new(Node)
	n3 := new(Node)
	n4 := new(Node)

	r = newLabelTree()
	r = r.insert("1", n1)
	r = r.insert("0", n0)
	r = r.insert("4", n4)
	r = r.insert("2", n2)
	r = r.insert("3", n3)
	assertEnumerate(r.enumerate(), "enumeration of tree 10423", t,
		fmt.Sprintf("\"0\": \"%p\"\n", n0),
		fmt.Sprintf("\"1\": \"%p\"\n", n1),
		fmt.Sprintf("\"2\": \"%p\"\n", n2),
		fmt.Sprintf("\"3\": \"%p\"\n", n3),
		fmt.Sprintf("\"4\": \"%p\"\n", n4),
	)
}

func TestLabelDelete(t *testing.T) {
	var r *labelTree

	r, ok := r.del("test")
	if ok {
		t.Errorf("Expected nothing to be deleted from empty tree but something has been deleted:\n%s", r.dot())
	}

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("3", nil)
	r = r.insert("6", nil)
	r = r.insert("9", nil)
	r = r.insert("c", nil)
	r = r.insert("f", nil)
	r = r.insert("1", nil)
	r = r.insert("2", nil)
	r = r.insert("4", nil)
	r = r.insert("5", nil)
	r = r.insert("7", nil)
	r = r.insert("8", nil)
	r = r.insert("a", nil)
	r = r.insert("b", nil)
	r = r.insert("d", nil)
	r = r.insert("e", nil)

	r, ok = r.del("81")
	if ok {
		t.Errorf("Expected nothing to be deleted by key \"81\" but something has been deleted")
	}
	assertLabelTree(r, TestTreeAfterNonExistingNodeDel, "tree after non-existing node deletion", t)

	r, ok = r.del("6")
	if !ok {
		t.Errorf("Expected node \"6\" to be deleted but got nothing")
	}
	assertLabelTree(r, TestTreeAfterNode6Deletion, "tree after node 6 deletion", t)

	r, ok = r.del("7")
	if !ok {
		t.Errorf("Expected node \"7\" to be deleted but got nothing")
	}
	r, ok = r.del("8")
	if !ok {
		t.Errorf("Expected node \"8\" to be deleted but got nothing")
	}
	r, ok = r.del("5")
	if !ok {
		t.Errorf("Expected node \"5\" to be deleted but got nothing")
	}
	r, ok = r.del("9")
	if !ok {
		t.Errorf("Expected node \"9\" to be deleted but got nothing")
	}
	assertLabelTree(r, TestTreeAfterNodes7859Deletion, "tree after nodes 7, 8, 5 and 9 deletion", t)

	r, ok = r.del("c")
	if !ok {
		t.Errorf("Expected node \"C\" to be deleted but got nothing")
	}
	r, ok = r.del("e")
	if !ok {
		t.Errorf("Expected node \"E\" to be deleted but got nothing")
	}
	r, ok = r.del("d")
	if !ok {
		t.Errorf("Expected node \"D\" to be deleted but got nothing")
	}
	r, ok = r.del("a")
	if !ok {
		t.Errorf("Expected node \"A\" to be deleted but got nothing")
	}
	r, ok = r.del("b")
	if !ok {
		t.Errorf("Expected node \"B\" to be deleted but got nothing")
	}
	r, ok = r.del("4")
	if !ok {
		t.Errorf("Expected node \"4\" to be deleted but got nothing")
	}
	r, ok = r.del("f")
	if !ok {
		t.Errorf("Expected node \"F\" to be deleted but got nothing")
	}
	r, ok = r.del("0")
	if !ok {
		t.Errorf("Expected node \"0\" to be deleted but got nothing")
	}
	r, ok = r.del("3")
	if !ok {
		t.Errorf("Expected node \"3\" to be deleted but got nothing")
	}
	r, ok = r.del("1")
	if !ok {
		t.Errorf("Expected node \"1\" to be deleted but got nothing")
	}
	r, ok = r.del("2")
	if !ok {
		t.Errorf("Expected node \"2\" to be deleted but got nothing")
	}
	assertLabelTree(r, TestEmptyTree, "tree after rest nodes deletion", t)
}

func TestLabelIsEmpty(t *testing.T) {
	var r *labelTree

	if !r.isEmpty() {
		t.Errorf("Expected nil tree to be empty")
	}

	r = newLabelTree()
	r = r.insert("0", nil)
	r = r.insert("3", nil)
	r = r.insert("6", nil)
	if r.isEmpty() {
		t.Errorf("Expected three nodes tree to be not empty")
	}

	r, ok := r.del("3")
	if !ok {
		t.Errorf("Expected element \"3\" to be deleted")
	}

	if r.isEmpty() {
		t.Errorf("Expected two nodes tree to be not empty")
	}

	r, ok = r.del("0")
	r, ok = r.del("6")

	if !r.isEmpty() {
		t.Errorf("Expected empty non-nil tree to be empty")
	}
}

func TestLabelRawMethods(t *testing.T) {
	var r *labelTree

	n := new(Node)

	r = r.rawInsert([]byte("k"), n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node tree", t)

	r = newLabelTree()
	r = r.rawInsert([]byte("k"), n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node tree", t)

	r = newLabelTree()
	r.rawInplaceInsert([]byte("k"), n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node inplace tree", t)

	r = newLabelTree()
	r.rawInplaceInsert([]byte("k"), n)
	assertLabelTree(r, fmt.Sprintf(TestSingleNodeTree, n), "single node inplace tree", t)

	r = nil
	v, ok := r.rawGet([]byte("0"))
	if ok {
		t.Errorf("Expected nothing but got %T (%#v)", v, v)
	}

	n0 := new(Node)
	n1 := new(Node)
	n2 := new(Node)
	n3 := new(Node)
	n4 := new(Node)

	r = newLabelTree()
	r = r.insert("1", n1)
	r = r.insert("0", n0)
	r = r.insert("4", n4)
	r = r.insert("2", n2)
	r = r.insert("3", n3)

	v, ok = r.rawGet([]byte("3"))
	if !ok {
		t.Errorf("Expected %p but got nothing", n3)
	} else if v != n3 {
		t.Errorf("Expected %p but got %p", n3, v)
	}

	var e *labelTree
	assertRawEnumerate(e.rawEnumerate(), "empty tree", t)

	assertRawEnumerate(r.rawEnumerate(), "raw enumeration of tree 10423", t,
		fmt.Sprintf("\"0\": \"%p\"\n", n0),
		fmt.Sprintf("\"1\": \"%p\"\n", n1),
		fmt.Sprintf("\"2\": \"%p\"\n", n2),
		fmt.Sprintf("\"3\": \"%p\"\n", n3),
		fmt.Sprintf("\"4\": \"%p\"\n", n4),
	)

	e, ok = e.rawDel([]byte("0"))
	if ok {
		t.Errorf("Expected nothing to be deleted from empty tree but something has been deleted:\n%s", e.dot())
	}

	_, ok = r.rawDel([]byte("0"))
	if !ok {
		t.Errorf("Expected node \"0\" to be deleted but got nothing")
	}
}

const (
	TestEmptyTree = `digraph d {
N0 [label="nil" style=filled fontcolor=white fillcolor=black]
}
`

	TestSingleNodeTree = `digraph d {
N0 [label="k: \"k\" v: %p" style=filled fontcolor=white fillcolor=black]
}
`

	TestThreeNodeTreeRed = `digraph d {
N0 [label="1" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="0" style=filled fillcolor=red]
N2 [label="2" style=filled fillcolor=red]
}
`

	TestFiveNodeTree = `digraph d {
N0 [label="1" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="0" style=filled fontcolor=white fillcolor=black]
N2 [label="3" style=filled fontcolor=white fillcolor=black]
N2 -> { N3 N4 }
N3 [label="2" style=filled fillcolor=red]
N4 [label="4" style=filled fillcolor=red]
}
`

	Test16InversedNodeTree = `digraph d {
N0 [label="c" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="8" style=filled fillcolor=red]
N1 -> { N3 N4 }
N2 [label="e" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="4" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="a" style=filled fontcolor=white fillcolor=black]
N4 -> { N9 N10 }
N5 [label="d" style=filled fontcolor=white fillcolor=black]
N6 [label="f" style=filled fontcolor=white fillcolor=black]
N7 [label="2" style=filled fillcolor=red]
N7 -> { N11 N12 }
N8 [label="6" style=filled fillcolor=red]
N8 -> { N13 N14 }
N9 [label="9" style=filled fontcolor=white fillcolor=black]
N10 [label="b" style=filled fontcolor=white fillcolor=black]
N11 [label="1" style=filled fontcolor=white fillcolor=black]
N11 -> { N15 N16 }
N12 [label="3" style=filled fontcolor=white fillcolor=black]
N13 [label="5" style=filled fontcolor=white fillcolor=black]
N14 [label="7" style=filled fontcolor=white fillcolor=black]
N15 [label="0" style=filled fillcolor=red]
N16 [label="nil" style=filled fontcolor=white fillcolor=black]
}
`

	Test16DirectNodeTree = `digraph d {
N0 [label="3" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="1" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="7" style=filled fillcolor=red]
N2 -> { N5 N6 }
N3 [label="0" style=filled fontcolor=white fillcolor=black]
N4 [label="2" style=filled fontcolor=white fillcolor=black]
N5 [label="5" style=filled fontcolor=white fillcolor=black]
N5 -> { N7 N8 }
N6 [label="b" style=filled fontcolor=white fillcolor=black]
N6 -> { N9 N10 }
N7 [label="4" style=filled fontcolor=white fillcolor=black]
N8 [label="6" style=filled fontcolor=white fillcolor=black]
N9 [label="9" style=filled fillcolor=red]
N9 -> { N11 N12 }
N10 [label="d" style=filled fillcolor=red]
N10 -> { N13 N14 }
N11 [label="8" style=filled fontcolor=white fillcolor=black]
N12 [label="a" style=filled fontcolor=white fillcolor=black]
N13 [label="c" style=filled fontcolor=white fillcolor=black]
N14 [label="e" style=filled fontcolor=white fillcolor=black]
N14 -> { N15 N16 }
N15 [label="nil" style=filled fontcolor=white fillcolor=black]
N16 [label="f" style=filled fillcolor=red]
}
`

	Test16AlternatingNodeTree = `digraph d {
N0 [label="6" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="2" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="a" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="0" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="4" style=filled fontcolor=white fillcolor=black]
N4 -> { N9 N10 }
N5 [label="8" style=filled fontcolor=white fillcolor=black]
N5 -> { N11 N12 }
N6 [label="c" style=filled fillcolor=red]
N6 -> { N13 N14 }
N7 [label="nil" style=filled fontcolor=white fillcolor=black]
N8 [label="1" style=filled fillcolor=red]
N9 [label="3" style=filled fillcolor=red]
N10 [label="5" style=filled fillcolor=red]
N11 [label="7" style=filled fillcolor=red]
N12 [label="9" style=filled fillcolor=red]
N13 [label="b" style=filled fontcolor=white fillcolor=black]
N14 [label="e" style=filled fontcolor=white fillcolor=black]
N14 -> { N15 N16 }
N15 [label="d" style=filled fillcolor=red]
N16 [label="f" style=filled fillcolor=red]
}
`

	Test16AlternatingInversedNodeTree = `digraph d {
N0 [label="9" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="5" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="d" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="3" style=filled fillcolor=red]
N3 -> { N7 N8 }
N4 [label="7" style=filled fontcolor=white fillcolor=black]
N4 -> { N9 N10 }
N5 [label="b" style=filled fontcolor=white fillcolor=black]
N5 -> { N11 N12 }
N6 [label="f" style=filled fontcolor=white fillcolor=black]
N6 -> { N13 N14 }
N7 [label="1" style=filled fontcolor=white fillcolor=black]
N7 -> { N15 N16 }
N8 [label="4" style=filled fontcolor=white fillcolor=black]
N9 [label="6" style=filled fillcolor=red]
N10 [label="8" style=filled fillcolor=red]
N11 [label="a" style=filled fillcolor=red]
N12 [label="c" style=filled fillcolor=red]
N13 [label="e" style=filled fillcolor=red]
N14 [label="nil" style=filled fontcolor=white fillcolor=black]
N15 [label="0" style=filled fillcolor=red]
N16 [label="2" style=filled fillcolor=red]
}
`

	Test16_3AltNodeTree = `digraph d {
N0 [label="5" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="3" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="9" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="1" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="4" style=filled fontcolor=white fillcolor=black]
N5 [label="7" style=filled fontcolor=white fillcolor=black]
N5 -> { N9 N10 }
N6 [label="c" style=filled fillcolor=red]
N6 -> { N11 N12 }
N7 [label="0" style=filled fillcolor=red]
N8 [label="2" style=filled fillcolor=red]
N9 [label="6" style=filled fillcolor=red]
N10 [label="8" style=filled fillcolor=red]
N11 [label="a" style=filled fontcolor=white fillcolor=black]
N11 -> { N13 N14 }
N12 [label="e" style=filled fontcolor=white fillcolor=black]
N12 -> { N15 N16 }
N13 [label="nil" style=filled fontcolor=white fillcolor=black]
N14 [label="b" style=filled fillcolor=red]
N15 [label="d" style=filled fillcolor=red]
N16 [label="f" style=filled fillcolor=red]
}
`

	Test32AlternatingNodeTree = `digraph d {
N0 [label="0e" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="06" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="16" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="02" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="0a" style=filled fontcolor=white fillcolor=black]
N4 -> { N9 N10 }
N5 [label="12" style=filled fontcolor=white fillcolor=black]
N5 -> { N11 N12 }
N6 [label="1a" style=filled fontcolor=white fillcolor=black]
N6 -> { N13 N14 }
N7 [label="00" style=filled fontcolor=white fillcolor=black]
N7 -> { N15 N16 }
N8 [label="04" style=filled fontcolor=white fillcolor=black]
N8 -> { N17 N18 }
N9 [label="08" style=filled fontcolor=white fillcolor=black]
N9 -> { N19 N20 }
N10 [label="0c" style=filled fontcolor=white fillcolor=black]
N10 -> { N21 N22 }
N11 [label="10" style=filled fontcolor=white fillcolor=black]
N11 -> { N23 N24 }
N12 [label="14" style=filled fontcolor=white fillcolor=black]
N12 -> { N25 N26 }
N13 [label="18" style=filled fontcolor=white fillcolor=black]
N13 -> { N27 N28 }
N14 [label="1c" style=filled fillcolor=red]
N14 -> { N29 N30 }
N15 [label="nil" style=filled fontcolor=white fillcolor=black]
N16 [label="01" style=filled fillcolor=red]
N17 [label="03" style=filled fillcolor=red]
N18 [label="05" style=filled fillcolor=red]
N19 [label="07" style=filled fillcolor=red]
N20 [label="09" style=filled fillcolor=red]
N21 [label="0b" style=filled fillcolor=red]
N22 [label="0d" style=filled fillcolor=red]
N23 [label="0f" style=filled fillcolor=red]
N24 [label="11" style=filled fillcolor=red]
N25 [label="13" style=filled fillcolor=red]
N26 [label="15" style=filled fillcolor=red]
N27 [label="17" style=filled fillcolor=red]
N28 [label="19" style=filled fillcolor=red]
N29 [label="1b" style=filled fontcolor=white fillcolor=black]
N30 [label="1e" style=filled fontcolor=white fillcolor=black]
N30 -> { N31 N32 }
N31 [label="1d" style=filled fillcolor=red]
N32 [label="1f" style=filled fillcolor=red]
}
`

	TestTreeAfterNonExistingNodeDel = `digraph d {
N0 [label="5" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="3" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="9" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="1" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="4" style=filled fontcolor=white fillcolor=black]
N5 [label="7" style=filled fontcolor=white fillcolor=black]
N5 -> { N9 N10 }
N6 [label="c" style=filled fillcolor=red]
N6 -> { N11 N12 }
N7 [label="0" style=filled fillcolor=red]
N8 [label="2" style=filled fillcolor=red]
N9 [label="6" style=filled fillcolor=red]
N10 [label="8" style=filled fillcolor=red]
N11 [label="a" style=filled fontcolor=white fillcolor=black]
N11 -> { N13 N14 }
N12 [label="e" style=filled fontcolor=white fillcolor=black]
N12 -> { N15 N16 }
N13 [label="nil" style=filled fontcolor=white fillcolor=black]
N14 [label="b" style=filled fillcolor=red]
N15 [label="d" style=filled fillcolor=red]
N16 [label="f" style=filled fillcolor=red]
}
`

	TestTreeAfterNode6Deletion = `digraph d {
N0 [label="5" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="3" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="c" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="1" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="4" style=filled fontcolor=white fillcolor=black]
N5 [label="9" style=filled fillcolor=red]
N5 -> { N9 N10 }
N6 [label="e" style=filled fontcolor=white fillcolor=black]
N6 -> { N11 N12 }
N7 [label="0" style=filled fillcolor=red]
N8 [label="2" style=filled fillcolor=red]
N9 [label="7" style=filled fontcolor=white fillcolor=black]
N9 -> { N13 N14 }
N10 [label="a" style=filled fontcolor=white fillcolor=black]
N10 -> { N15 N16 }
N11 [label="d" style=filled fillcolor=red]
N12 [label="f" style=filled fillcolor=red]
N13 [label="nil" style=filled fontcolor=white fillcolor=black]
N14 [label="8" style=filled fillcolor=red]
N15 [label="nil" style=filled fontcolor=white fillcolor=black]
N16 [label="b" style=filled fillcolor=red]
}
`

	TestTreeAfterNodes7859Deletion = `digraph d {
N0 [label="a" style=filled fontcolor=white fillcolor=black]
N0 -> { N1 N2 }
N1 [label="2" style=filled fontcolor=white fillcolor=black]
N1 -> { N3 N4 }
N2 [label="c" style=filled fontcolor=white fillcolor=black]
N2 -> { N5 N6 }
N3 [label="1" style=filled fontcolor=white fillcolor=black]
N3 -> { N7 N8 }
N4 [label="4" style=filled fontcolor=white fillcolor=black]
N4 -> { N9 N10 }
N5 [label="b" style=filled fontcolor=white fillcolor=black]
N6 [label="e" style=filled fontcolor=white fillcolor=black]
N6 -> { N11 N12 }
N7 [label="0" style=filled fillcolor=red]
N8 [label="nil" style=filled fontcolor=white fillcolor=black]
N9 [label="3" style=filled fillcolor=red]
N10 [label="nil" style=filled fontcolor=white fillcolor=black]
N11 [label="d" style=filled fillcolor=red]
N12 [label="f" style=filled fillcolor=red]
}
`
)

func assertLabelTree(r *labelTree, e, desc string, t *testing.T) {
	assertStringLists(difflib.SplitLines(r.dot()), difflib.SplitLines(e), desc, t)
}

func assertEnumerate(ch chan labelPair, desc string, t *testing.T, e ...string) {
	pairs := []string{}
	for p := range ch {
		pairs = append(pairs, fmt.Sprintf("%q: \"%p\"\n", p.Key, p.Value))
	}

	assertStringLists(pairs, e, desc, t)
}

func assertRawEnumerate(ch chan labelRawPair, desc string, t *testing.T, e ...string) {
	pairs := []string{}
	for p := range ch {
		pairs = append(pairs, fmt.Sprintf("%q: \"%p\"\n", domain.Label(p.Key), p.Value))
	}

	assertStringLists(pairs, e, desc, t)
}

func assertStringLists(v, e []string, desc string, t *testing.T) {
	ctx := difflib.ContextDiff{
		A:        e,
		B:        v,
		FromFile: "Expected",
		ToFile:   "Got"}

	diff, err := difflib.GetContextDiffString(ctx)
	if err != nil {
		panic(fmt.Errorf("can't compare \"%s\": %s", desc, err))
	}

	if len(diff) > 0 {
		t.Errorf("\"%s\" doesn't match:\n%s", desc, diff)
	}
}

func assertPanic(f func(), desc string, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic from %s but got nothing", desc)
		}
	}()

	f()
}