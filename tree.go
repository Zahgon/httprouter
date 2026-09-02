package httprouter

func min(a, b int) int { _ = "STUB: not implemented"; return 0 }

func longestCommonPrefix(a, b string) int { _ = "STUB: not implemented"; return 0 }

func findWildcard(path string) (wilcard string, i int, valid bool) {
	_ = "STUB: not implemented"
	return "", 0, false
}

func countParams(path string) uint16 { _ = "STUB: not implemented"; return 0 }

type nodeType uint8

const (
	static nodeType = iota
	root
	param
	catchAll
)

type node struct {
	path      string
	indices   string
	wildChild bool
	nType     nodeType
	priority  uint32
	children  []*node
	handle    Handle
}

func (n *node) incrementChildPrio(pos int) int { _ = "STUB: not implemented"; return 0 }

func (n *node) addRoute(path string, handle Handle) { _ = "STUB: not implemented"; return }

func (n *node) insertChild(path, fullPath string, handle Handle) { _ = "STUB: not implemented"; return }

func (n *node) getValue(path string, params func() *Params) (handle Handle, ps *Params, tsr bool) {
	_ = "STUB: not implemented"
	return *new(Handle), nil, false
}

func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (fixedPath string, found bool) {
	_ = "STUB: not implemented"
	return "", false
}

func shiftNRuneBytes(rb [4]byte, n int) [4]byte { _ = "STUB: not implemented"; return [4]byte{} }

func (n *node) findCaseInsensitivePathRec(path string, ciPath []byte, rb [4]byte, fixTrailingSlash bool) []byte {
	_ = "STUB: not implemented"
	return nil
}
