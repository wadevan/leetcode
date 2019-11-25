package golang

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test742(t *testing.T) {
	ints := []int{1, 2, 3, 4, NULL, NULL, NULL, 5, NULL, 6}
	node := Ints2TreeNode(ints)
	leaf := findClosestLeaf(node, 2)
	assert.Equal(t, 3, leaf)
}

func Test386(t *testing.T) {
	order := lexicalOrder(13)
	assert.Equal(t, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}, order)
}

func Test93(t *testing.T) {
	s := "25525511135"
	addresses := restoreIpAddresses(s)
	assert.Equal(t, []string{"255.255.11.135", "255.255.111.35"}, addresses)
}

func Test92(t *testing.T){
	ints := []int{1,2,3,4,5}
	list := Ints2List(ints)
	res := reverseBetween(list, 2, 4)
	assertInts := []int{1,4,3,2,5}
	list2Ints := List2Ints(res)
	assert.Equal(t, assertInts,list2Ints )

	ints =[]int{5}
	list = Ints2List(ints)
	res = reverseBetween(list, 1, 1)
	assertInts = []int{5}
	list2Ints = List2Ints(res)
	assert.Equal(t, assertInts,list2Ints )

}