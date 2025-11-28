package __obf_4f13ac5aa043b5ce

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_a0c6719477333085 uint32) *Heap {
	__obf_b1fecdf21015160d := Nodes{}
	heap.Init(&__obf_b1fecdf21015160d)
	return &Heap{Nodes: __obf_b1fecdf21015160d, K: __obf_a0c6719477333085}
}

func (__obf_b1fecdf21015160d *Heap) Add(__obf_09f91c0944512e18 *Node) *Node {
	if __obf_b1fecdf21015160d.K > uint32(len(__obf_b1fecdf21015160d.Nodes)) {
		heap.Push(&__obf_b1fecdf21015160d.Nodes, __obf_09f91c0944512e18)
	} else if __obf_09f91c0944512e18.Count > __obf_b1fecdf21015160d.Nodes[0].Count {
		__obf_a0c6b5ffefc09b86 := heap.Pop(&__obf_b1fecdf21015160d.Nodes)
		heap.Push(&__obf_b1fecdf21015160d.Nodes, __obf_09f91c0944512e18)
		__obf_2f7077ce39e8d038 := __obf_a0c6b5ffefc09b86.(*Node)
		return __obf_2f7077ce39e8d038
	}
	return nil
}

func (__obf_b1fecdf21015160d *Heap) Pop() *Node {
	__obf_a0c6b5ffefc09b86 := heap.Pop(&__obf_b1fecdf21015160d.Nodes)
	return __obf_a0c6b5ffefc09b86.(*Node)
}

func (__obf_b1fecdf21015160d *Heap) Fix(__obf_9c5668fe593f5615 int, __obf_7a971de362a6bf24 uint32) {
	__obf_b1fecdf21015160d.Nodes[__obf_9c5668fe593f5615].Count = __obf_7a971de362a6bf24
	heap.Fix(&__obf_b1fecdf21015160d.Nodes, __obf_9c5668fe593f5615)
}

func (__obf_b1fecdf21015160d *Heap) Min() uint32 {
	if len(__obf_b1fecdf21015160d.Nodes) == 0 {
		return 0
	}
	return __obf_b1fecdf21015160d.Nodes[0].Count
}

func (__obf_b1fecdf21015160d *Heap) Find(__obf_797707a3dac0ebb7 string) (int, bool) {
	for __obf_e3e988c1360a57a4 := range __obf_b1fecdf21015160d.Nodes {
		if __obf_b1fecdf21015160d.Nodes[__obf_e3e988c1360a57a4].Key == __obf_797707a3dac0ebb7 {
			return __obf_e3e988c1360a57a4, true
		}
	}
	return 0, false
}

func (__obf_b1fecdf21015160d *Heap) Sorted() Nodes {
	__obf_181e19e2594435f5 := append([]*Node(nil), __obf_b1fecdf21015160d.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_181e19e2594435f5)))
	return __obf_181e19e2594435f5
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_9178770692fd13ed Nodes) Len() int {
	return len(__obf_9178770692fd13ed)
}

func (__obf_9178770692fd13ed Nodes) Less(__obf_e3e988c1360a57a4, __obf_b45ef0e024b2a74c int) bool {
	return (__obf_9178770692fd13ed[__obf_e3e988c1360a57a4].Count < __obf_9178770692fd13ed[__obf_b45ef0e024b2a74c].Count) || (__obf_9178770692fd13ed[__obf_e3e988c1360a57a4].Count == __obf_9178770692fd13ed[__obf_b45ef0e024b2a74c].Count && __obf_9178770692fd13ed[__obf_e3e988c1360a57a4].Key > __obf_9178770692fd13ed[__obf_b45ef0e024b2a74c].Key)
}

func (__obf_9178770692fd13ed Nodes) Swap(__obf_e3e988c1360a57a4, __obf_b45ef0e024b2a74c int) {
	__obf_9178770692fd13ed[__obf_e3e988c1360a57a4], __obf_9178770692fd13ed[__obf_b45ef0e024b2a74c] = __obf_9178770692fd13ed[__obf_b45ef0e024b2a74c], __obf_9178770692fd13ed[__obf_e3e988c1360a57a4]
}

func (__obf_9178770692fd13ed *Nodes) Push(__obf_09f91c0944512e18 any) {
	*__obf_9178770692fd13ed = append(*__obf_9178770692fd13ed, __obf_09f91c0944512e18.(*Node))
}

func (__obf_9178770692fd13ed *Nodes) Pop() any {
	var __obf_09f91c0944512e18 *Node
	__obf_09f91c0944512e18, *__obf_9178770692fd13ed = (*__obf_9178770692fd13ed)[len((*__obf_9178770692fd13ed))-1], (*__obf_9178770692fd13ed)[:len((*__obf_9178770692fd13ed))-1]
	return __obf_09f91c0944512e18
}
