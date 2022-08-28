package block

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transations  []string
}
