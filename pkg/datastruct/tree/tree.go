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
	parent   *Proc
	children *Proc
	next     *Child
}

type Child struct {
	child *Proc
	next  *Child
}

func (list *Proc) newProc(comm string, pid int, uid int, ppid int) {
}

func (list *Proc) addChild(child *Proc) {}

func (list *Proc) addProc(comm string, pid int, uid int, ppid int) {
	// Create a new process instance
	// find the position of it's parent in list
	// if the parent has existing children, add this to the list
}

func findProc() {}

// Construct a tree by reading the proc pseudo-filesystem
// assumes that miniproc is a string containing records, separator=","
func readProc(miniproc string) string {

	var list *Proc
	records := strings.Split(miniproc, ",")
	for _, r := range records {
		items := strings.Split(r, " ")
		comm := items[0]
		pid, _ := strconv.Atoi(items[1])
		uid, _ := strconv.Atoi(items[2])
		ppid, _ := strconv.Atoi(items[3])
		list.addProc(comm, pid, uid, ppid)
	}
	return fmt.Sprint(list)
}


//		{name: "test construction logic for tree",
//			args: args{"(systemd) 1 0 0, (ModemManager) 2 1 0, (abrd-dbug) 3 1 0"}},
