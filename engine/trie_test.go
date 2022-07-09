package engine

import (
	"reflect"
	"testing"
)

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/hello/:name"), []string{"hello", ":name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestInsert(t *testing.T) {
	pattern := "/hello/:name"
	parts := parsePattern(pattern)
	root := &Node{}
	root.insert(pattern, parts, 0)
}

func TestSearch(t *testing.T) {
	// 先插入
	pattern := "/hello/:name"
	parts := parsePattern(pattern)
	root := &Node{}
	root.insert(pattern, parts, 0)

	// 再搜索
	path := "/hello/geek"
	parts = parsePattern(path)
	n := root.search(parts, 0)
	if n == nil {
		t.Fatal("test search failed")
	}
}
