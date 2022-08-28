package pool

import (
	"powblockchain/transaction"
)

type Pool struct {
	transactions []*transaction.Transaction
}
