package tree

// Trie trie = new Trie();
// trie.insert("apple");
// trie.search("apple"); // true
// trie.search("app"); // false
// trie.startsWith("app"); // true
// trie.insert("app");
// trie.search("app"); // true

type trieNode struct {
	// rune vs. byte
	// rune: map[rune]*trieNode
	children map[rune]*trieNode
	// byte: [26]*trieNode
	//   26 english letters, there's no need to use a map.
	// children [26]*trieNode
	isVal bool
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{root: &trieNode{}}
}

func (t *Trie) Insert(str []rune) {
	cur := t.root
	for i := 0; i < len(str); i++ {
		child, _ := cur.children[str[i]]
		if child == nil {
			if cur.children == nil {
				cur.children = make(map[rune]*trieNode)
				child = &trieNode{}
				cur.children[str[i]] = child
			}
		}
		cur = child
	}
	cur.isVal = true
}

func (t *Trie) Search(str []rune) bool {
	cur := t.root
	for i := 0; i < len(str); i++ {
		child, _ := cur.children[str[i]]
		if child == nil {
			return false
		}
		cur = child
	}
	return cur.isVal == true
}

func (t *Trie) StartsWith(str []rune) bool {
	cur := t.root
	for i := 0; i < len(str); i++ {
		child, _ := cur.children[str[i]]
		if child == nil {
			return false
		}
		cur = child
	}
	return true
}
