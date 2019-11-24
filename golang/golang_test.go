package golang

import ("testing"
	"github.com/stretchr/testify/assert"
)

func Test742(t *testing.T) {
	ints := []int{1,2,3,4,NULL,NULL,NULL,5,NULL,6}
	node := Ints2TreeNode(ints)
	leaf := findClosestLeaf(node, 2)
	assert.Equal(t, 3, leaf)
}

func Test386(t *testing.T) {
	order := lexicalOrder(13)
	assert.Equal(t, []int{1,10,11,12,13,2,3,4,5,6,7,8,9}, order)
}