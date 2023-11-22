package tree

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	sym_empty_2  = "  "
	sym_branch_2 = "|-"
	sym_verti_2  = "| "
	sym_last_2   = "`-"
	sym_single_3 = "---"
	sym_first_3  = "-+-"
)

// Using a concrete type to sketch out a few general functions
type Proc struct {
	comm     string
	pid      int
	uid      int
	ppid     int
	children *Child
	parent   *Proc
	next     *Proc
}

type Child struct {
	child *Proc
	next  *Child
}

func (tree *Proc) String() string {
	if tree == nil {
		return "(empty)"
	}

	result := ""

	for tree != nil {
		if len(result) > 0 {
			result += " > "
		}
		result += tree.comm
		tree = tree.next
	}
	return result
}

func (tree *Proc) Dump(current *Proc, level int, rep int, leaf bool, last bool, prev_uid int, closing int) {

	/* Build up the tree here */
	var o strings.Builder
	if tree == nil {
		fmt.Fprint(&o, "(empty)")
	}
	if current == nil {
		/* Base case, print the tree */
		fmt.Println(o.String())
	}
	if !leaf {
		for lvl := 0; lvl < level; lvl++ {
			// the original uses width[] here to determine the width of columns
			if (lvl == level - 1) {
				if last {
					o.writeString(sym_last_2)           // last
				} else {
					o.writeString(sym_branch_2)         // branch
				} else {
					if { /* if more[lvl+1] */
					o.writeString(sym_versym_verti_2)   // vert
				} else {
					/* print empty */
					o.writeString(sym_empty_2)
				}
		}
	}

}

func (list *Proc) newProc(comm string, pid int, uid int) *Proc {
	return &Proc{comm, pid, uid, 0, nil, nil, nil}
}

func (list *Proc) Insert(node *Proc) *Proc {
	if list == nil { // new list
		return node
	}
	if list.next == nil {
		list.next = node
	} else {
		node.next = list.next
		list.next = node
	}
	return list
}

func (parent *Proc) addChild(child *Proc) *Proc {
	cnew := &Child{
		child: &Proc{
			comm:     child.comm,
			pid:      child.pid,
			uid:      child.uid,
			ppid:     child.ppid,
			children: child.children,
			parent:   child.parent,
			next:     child.next,
		},
		next: nil,
	}

	if parent.children == nil {
		parent.children = cnew
		return parent
	} else {
		// if the parent has children
		// walk the list of children reading pids
		// insert the child at when pid > pid children
		walk := parent.children
		for walk != nil {
			if walk.child.pid > child.pid {
				break
			}
			walk = walk.next
		}
		cnew.next = walk
		walk = cnew
	}
	return parent
}

func (list *Proc) findProc(pid int) *Proc {
	for list != nil {
		if list.pid == pid {
			return list
		}
		list = list.next
	}
	return nil
}

func (list *Proc) addProc(comm string, pid int, uid int, ppid int) *Proc {
	this := list.findProc(pid)
	if this == nil {
		this = list.newProc(comm, pid, uid)
	}
	parent := list.findProc(ppid)
	if parent == nil {
		parent = list.newProc("?", ppid, 0)
	}
	//list.addChild(parent, this)
	parent.addChild(this)
	this.parent = parent
	list = list.Insert(this)
	return list
}

// Construct a tree by reading the proc pseudo-filesystem
// assumes that miniproc is a string containing records, separator=","
func readProc(miniproc string) string {

	var list *Proc
	records := strings.Split(miniproc, ",")
	for _, rr := range records {
		r := strings.TrimSpace(rr)
		items := strings.Split(r, " ")
		comm := items[0]
		pid, _ := strconv.Atoi(items[1])
		uid, _ := strconv.Atoi(items[2])
		ppid, _ := strconv.Atoi(items[3])
		list = list.addProc(comm, pid, uid, ppid)
	}
	return fmt.Sprint(list)
}
