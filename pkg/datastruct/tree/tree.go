package tree

import (
	"fmt"
	"strconv"
	"strings"
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

func (list *Proc) newProc(comm string, pid int, uid int) *Proc {
	return &Proc{comm, pid, uid, 0, nil, nil, nil}
}

func (list *Proc) addChild(parent *Proc, child *Proc) {
	var cnew *Child
	var walk **Child
	cnew.child = child
	for walk = &parent.children; *walk != nil; walk = &(*walk).next {
		if (*walk).child.pid > child.pid {
			break
		}
	}
	cnew.next = *walk
	*walk = cnew
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

func (list *Proc) addProc(comm string, pid int, uid int, ppid int) {
	// Create a new process instance
	// either it exists or new
	this := list.findProc(pid)
	if this == nil {
		this = list.newProc(comm, pid, uid)
	}
	// find the parent
	parent := list.findProc(ppid)
	if parent == nil {
		parent = list.newProc("?", ppid, 0)
	}
	list.addChild(parent, this)
	this.parent = parent
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
		list.addProc(comm, pid, uid, ppid)
	}
	return fmt.Sprint(list)
}
