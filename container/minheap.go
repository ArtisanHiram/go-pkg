package __obf_1fda7fbdeda52f1e

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_9058efdba9a0277b uint32) *Heap {
	__obf_79bb9dc2d03a5e29 := Nodes{}
	heap.Init(&__obf_79bb9dc2d03a5e29)
	return &Heap{Nodes: __obf_79bb9dc2d03a5e29, K: __obf_9058efdba9a0277b}
}

func (__obf_79bb9dc2d03a5e29 *Heap) Add(__obf_35a7ec03898cd9ce *Node) *Node {
	if __obf_79bb9dc2d03a5e29.K > uint32(len(__obf_79bb9dc2d03a5e29.Nodes)) {
		heap.Push(&__obf_79bb9dc2d03a5e29.Nodes, __obf_35a7ec03898cd9ce)
	} else if __obf_35a7ec03898cd9ce.Count > __obf_79bb9dc2d03a5e29.Nodes[0].Count {
		__obf_6a85a79659f1cd3e := heap.Pop(&__obf_79bb9dc2d03a5e29.Nodes)
		heap.Push(&__obf_79bb9dc2d03a5e29.Nodes, __obf_35a7ec03898cd9ce)
		__obf_1e412470104e608f := __obf_6a85a79659f1cd3e.(*Node)
		return __obf_1e412470104e608f
	}
	return nil
}

func (__obf_79bb9dc2d03a5e29 *Heap) Pop() *Node {
	__obf_6a85a79659f1cd3e := heap.Pop(&__obf_79bb9dc2d03a5e29.Nodes)
	return __obf_6a85a79659f1cd3e.(*Node)
}

func (__obf_79bb9dc2d03a5e29 *Heap) Fix(__obf_8fa6c1a4ba942a8a int, __obf_c29f343498177aa5 uint32) {
	__obf_79bb9dc2d03a5e29.Nodes[__obf_8fa6c1a4ba942a8a].Count = __obf_c29f343498177aa5
	heap.Fix(&__obf_79bb9dc2d03a5e29.Nodes, __obf_8fa6c1a4ba942a8a)
}

func (__obf_79bb9dc2d03a5e29 *Heap) Min() uint32 {
	if len(__obf_79bb9dc2d03a5e29.Nodes) == 0 {
		return 0
	}
	return __obf_79bb9dc2d03a5e29.Nodes[0].Count
}

func (__obf_79bb9dc2d03a5e29 *Heap) Find(__obf_95c4c09a56c1a070 string) (int, bool) {
	for __obf_457766c6d5154891 := range __obf_79bb9dc2d03a5e29.Nodes {
		if __obf_79bb9dc2d03a5e29.Nodes[__obf_457766c6d5154891].Key == __obf_95c4c09a56c1a070 {
			return __obf_457766c6d5154891, true
		}
	}
	return 0, false
}

func (__obf_79bb9dc2d03a5e29 *Heap) Sorted() Nodes {
	__obf_95f79fb5a643f49b := append([]*Node(nil), __obf_79bb9dc2d03a5e29.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_95f79fb5a643f49b)))
	return __obf_95f79fb5a643f49b
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_6d56bf10a4deb62c Nodes) Len() int {
	return len(__obf_6d56bf10a4deb62c)
}

func (__obf_6d56bf10a4deb62c Nodes) Less(__obf_457766c6d5154891, __obf_d1976d7634a537e7 int) bool {
	return (__obf_6d56bf10a4deb62c[__obf_457766c6d5154891].Count < __obf_6d56bf10a4deb62c[__obf_d1976d7634a537e7].Count) || (__obf_6d56bf10a4deb62c[__obf_457766c6d5154891].Count == __obf_6d56bf10a4deb62c[__obf_d1976d7634a537e7].Count && __obf_6d56bf10a4deb62c[__obf_457766c6d5154891].Key > __obf_6d56bf10a4deb62c[__obf_d1976d7634a537e7].Key)
}

func (__obf_6d56bf10a4deb62c Nodes) Swap(__obf_457766c6d5154891, __obf_d1976d7634a537e7 int) {
	__obf_6d56bf10a4deb62c[__obf_457766c6d5154891], __obf_6d56bf10a4deb62c[__obf_d1976d7634a537e7] = __obf_6d56bf10a4deb62c[__obf_d1976d7634a537e7], __obf_6d56bf10a4deb62c[__obf_457766c6d5154891]
}

func (__obf_6d56bf10a4deb62c *Nodes) Push(__obf_35a7ec03898cd9ce any) {
	*__obf_6d56bf10a4deb62c = append(*__obf_6d56bf10a4deb62c, __obf_35a7ec03898cd9ce.(*Node))
}

func (__obf_6d56bf10a4deb62c *Nodes) Pop() any {
	var __obf_35a7ec03898cd9ce *Node
	__obf_35a7ec03898cd9ce, *__obf_6d56bf10a4deb62c = (*__obf_6d56bf10a4deb62c)[len((*__obf_6d56bf10a4deb62c))-1], (*__obf_6d56bf10a4deb62c)[:len((*__obf_6d56bf10a4deb62c))-1]
	return __obf_35a7ec03898cd9ce
}
