package huffman_tree

type HuffmanTree struct {
	encodingLength int
	encodeMap map[byte]string
	root *node
}

type node struct {
	left *node
	right *node
}
