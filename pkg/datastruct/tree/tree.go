package tree

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
func readProc() {}
