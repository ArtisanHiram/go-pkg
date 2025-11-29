package __obf_90a4883a02d0b41c

import (
	"container/heap"
	"sort"
)

type Heap struct {
	Nodes Nodes
	K     uint32
}

func NewHeap(__obf_7d788ca4b6a826e5 uint32) *Heap {
	__obf_ce6856a901600875 := Nodes{}
	heap.Init(&__obf_ce6856a901600875)
	return &Heap{Nodes: __obf_ce6856a901600875, K: __obf_7d788ca4b6a826e5}
}

func (__obf_ce6856a901600875 *Heap) Add(__obf_821546dddaaa0593 *Node) *Node {
	if __obf_ce6856a901600875.K > uint32(len(__obf_ce6856a901600875.Nodes)) {
		heap.Push(&__obf_ce6856a901600875.Nodes, __obf_821546dddaaa0593)
	} else if __obf_821546dddaaa0593.Count > __obf_ce6856a901600875.Nodes[0].Count {
		__obf_144fa304550d3a01 := heap.Pop(&__obf_ce6856a901600875.Nodes)
		heap.Push(&__obf_ce6856a901600875.Nodes, __obf_821546dddaaa0593)
		__obf_d91037d05dd8a42c := __obf_144fa304550d3a01.(*Node)
		return __obf_d91037d05dd8a42c
	}
	return nil
}

func (__obf_ce6856a901600875 *Heap) Pop() *Node {
	__obf_144fa304550d3a01 := heap.Pop(&__obf_ce6856a901600875.Nodes)
	return __obf_144fa304550d3a01.(*Node)
}

func (__obf_ce6856a901600875 *Heap) Fix(__obf_af8123cf8588a00d int, __obf_ee99c442f9d25bf4 uint32) {
	__obf_ce6856a901600875.
		Nodes[__obf_af8123cf8588a00d].Count = __obf_ee99c442f9d25bf4
	heap.Fix(&__obf_ce6856a901600875.Nodes, __obf_af8123cf8588a00d)
}

func (__obf_ce6856a901600875 *Heap) Min() uint32 {
	if len(__obf_ce6856a901600875.Nodes) == 0 {
		return 0
	}
	return __obf_ce6856a901600875.Nodes[0].Count
}

func (__obf_ce6856a901600875 *Heap) Find(__obf_5bba24c1758bbf28 string) (int, bool) {
	for __obf_f315019af10902c2 := range __obf_ce6856a901600875.Nodes {
		if __obf_ce6856a901600875.Nodes[__obf_f315019af10902c2].Key == __obf_5bba24c1758bbf28 {
			return __obf_f315019af10902c2, true
		}
	}
	return 0, false
}

func (__obf_ce6856a901600875 *Heap) Sorted() Nodes {
	__obf_555474ed181d3bd6 := append([]*Node(nil), __obf_ce6856a901600875.Nodes...)
	sort.Sort(sort.Reverse(Nodes(__obf_555474ed181d3bd6)))
	return __obf_555474ed181d3bd6
}

type Nodes []*Node

type Node struct {
	Key   string
	Count uint32
}

func (__obf_12d9f59bac8fe254 Nodes) Len() int {
	return len(__obf_12d9f59bac8fe254)
}

func (__obf_12d9f59bac8fe254 Nodes) Less(__obf_f315019af10902c2, __obf_e5cf5598742db663 int) bool {
	return (__obf_12d9f59bac8fe254[__obf_f315019af10902c2].Count < __obf_12d9f59bac8fe254[__obf_e5cf5598742db663].Count) || (__obf_12d9f59bac8fe254[__obf_f315019af10902c2].Count == __obf_12d9f59bac8fe254[__obf_e5cf5598742db663].Count && __obf_12d9f59bac8fe254[__obf_f315019af10902c2].Key > __obf_12d9f59bac8fe254[__obf_e5cf5598742db663].Key)
}

func (__obf_12d9f59bac8fe254 Nodes) Swap(__obf_f315019af10902c2, __obf_e5cf5598742db663 int) {
	__obf_12d9f59bac8fe254[__obf_f315019af10902c2], __obf_12d9f59bac8fe254[__obf_e5cf5598742db663] = __obf_12d9f59bac8fe254[__obf_e5cf5598742db663], __obf_12d9f59bac8fe254[__obf_f315019af10902c2]
}

func (__obf_12d9f59bac8fe254 *Nodes) Push(__obf_821546dddaaa0593 any) {
	*__obf_12d9f59bac8fe254 = append(*__obf_12d9f59bac8fe254, __obf_821546dddaaa0593.(*Node))
}

func (__obf_12d9f59bac8fe254 *Nodes) Pop() any {
	var __obf_821546dddaaa0593 *Node
	__obf_821546dddaaa0593, *__obf_12d9f59bac8fe254 = (*__obf_12d9f59bac8fe254)[len((*__obf_12d9f59bac8fe254))-1], (*__obf_12d9f59bac8fe254)[:len((*__obf_12d9f59bac8fe254))-1]
	return __obf_821546dddaaa0593
}
