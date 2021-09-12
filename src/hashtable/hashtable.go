package hashtable

type hashTable struct {
	data [][][]interface{}
}

func NewHashTable(size int) hashTable {
	return hashTable{data: make([][][]interface{}, size)}
}

func (h hashTable) hashMethod(key string) (hash int) {
	var str []rune = []rune(key)
	for i := 0; i < len(key); i++ {
		hash = (hash + int(str[i])*i) % len(h.data)
	}
	return
}

func (h *hashTable) Set(key string, value int) {
	var address int = h.hashMethod(key)
	if h.data[address] == nil {
		h.data[address] = make([][]interface{}, 0, 2)
	}
	h.data[address] = append(h.data[address], []interface{}{key, value})

}

func (h hashTable) Get(key string) interface{} {
	var address = h.hashMethod(key)
	var currentBucket = h.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key {
				return currentBucket[i][1]
			}
		}
	}
	return nil
}

func (h hashTable) Delete(key string) interface{} {
	var address = h.hashMethod(key)
	var currentBucket = h.data[address]
	if currentBucket != nil {
		for i := 0; i < len(currentBucket); i++ {
			if currentBucket[i][0] == key {
				h.data[address] = append(h.data[address][:i], h.data[address][i+1:]...)
				return currentBucket[i][1]
			}
		}
	}
	return nil
}

func (h *hashTable) GetHashes() interface{} {
	var hashes []interface{}
	for i := 0; i < len(h.data); i++ {
		if h.data[i] != nil {
			for j := 0; j < len(h.data[i]); j++ {
				hashes = append(hashes, h.data[i][j][0])
			}
		}
	}
	return hashes
}
