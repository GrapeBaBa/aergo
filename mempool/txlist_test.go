/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */
package mempool

import (
	"math/rand"
	"testing"

	"github.com/aergoio/aergo/types"
)

func TestListPutBasic(t *testing.T) {
	initTest(t)
	defer deinitTest()
	mpl := NewTxList(nil, uint64(1))

	count := 1000
	nonce := make([]int, count)
	for i := 0; i < count; i++ {
		nonce[i] = i + 1
	}
	rand.Shuffle(count, func(i, j int) {
		nonce[i], nonce[j] = nonce[j], nonce[i]
	})
	for i := 0; i < count; i++ {
		mpl.Put(genTx(0, 0, uint64(nonce[i]), 0))
	}

	ret := mpl.Get()
	if len(ret) != count {
		t.Error("put failed", len(ret), count)
	}
}

func TestListPutErrors(t *testing.T) {
	initTest(t)
	defer deinitTest()
	mpl := NewTxList(nil, uint64(10))
	added, err := mpl.Put(genTx(0, 0, uint64(1), 0))
	if added != 0 || err != types.ErrTxNonceTooLow {
		t.Errorf("put should be failed with ErrTxNonceTooLow, but %s", err)
	}

	added, err = mpl.Put(genTx(0, 0, uint64(10), 0))
	if added != 0 || err != nil || len(mpl.list) != 1 {
		t.Errorf("put should be not failed, but (%d)%s", added, err)
	}

	added, err = mpl.Put(genTx(0, 0, uint64(10), 0))
	if added != 0 || err != types.ErrTxAlreadyInMempool {
		t.Errorf("put should be failed with ErrTxNonceTooLow, but %s", err)
	}

}
func TestListDel(t *testing.T) {
	initTest(t)
	defer deinitTest()
	mpl := NewTxList(nil, uint64(1))

	ret, txs := mpl.SetMinNonce(uint64(3))
	if ret != 0 || mpl.Len() != 0 || len(txs) != 0 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	ret, txs = mpl.SetMinNonce(uint64(1))
	if ret != 0 || mpl.Len() != 0 || len(txs) != 0 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	count := 100

	//put 1~10 excpet 4 6 8
	for i := 0; i < count; i++ {
		if i == 3 || i == 5 || i == 7 {
			continue
		}
		mpl.Put(genTx(0, 0, uint64(i+1), 0))
	}
	// 1, |2, 3, | x, 5, x, 7, | x, 9... 14, |15... 100
	ret, txs = mpl.SetMinNonce(uint64(1))
	if ret != 0 || mpl.Len() != 3 || len(txs) != 0 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	ret, txs = mpl.SetMinNonce(uint64(2))
	if ret != 0 || mpl.Len() != 2 || len(txs) != 1 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	ret, txs = mpl.SetMinNonce(uint64(4))
	if ret != 0 || mpl.Len() != 0 || len(txs) != 2 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	ret, txs = mpl.SetMinNonce(uint64(8))
	if ret != 2 || mpl.Len() != 0 || len(txs) != 2 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	ret, txs = mpl.SetMinNonce(uint64(15))
	if ret != 92 || mpl.Len() != count-14 || len(txs) != 6 {
		t.Error(ret, mpl.Len(), len(txs))
	}

	txs = mpl.Get()
	if txs[0].GetBody().GetNonce() != 15 {
		t.Error(txs[0].GetBody().GetNonce())
	}

}

func TestListDelMiddle(t *testing.T) {
	initTest(t)
	defer deinitTest()
	mpl := NewTxList(nil, uint64(4))

	mpl.Put(genTx(0, 0, uint64(4), 0))
	mpl.Put(genTx(0, 0, uint64(5), 0))
	mpl.Put(genTx(0, 0, uint64(6), 0))

	if mpl.Len() != 3 {
		t.Error("should be 3 not ", len(mpl.list))
	}
	ret, txs := mpl.SetMinNonce(uint64(3))
	if ret != -3 || mpl.Len() != 0 || len(txs) != 0 {
		t.Error(ret, mpl.Len(), len(txs))
	}
	ret, txs = mpl.SetMinNonce(uint64(5))
	if ret != 3 || mpl.Len() != 2 || len(txs) != 1 {
		t.Error(ret, mpl.Len(), len(txs))
	}

}

func TestListPutRandom(t *testing.T) {
	initTest(t)
	defer deinitTest()
	mpl := NewTxList(nil, uint64(1))

	count := 100
	txs := make([]*types.Tx, count)
	for i := 0; i < count; i++ {
		txs[i] = genTx(0, 0, uint64(i+1), 0)
	}
	rand.Shuffle(count, func(i, j int) {
		txs[i], txs[j] = txs[j], txs[i]
	})
	for i := 0; i < count; i++ {
		mpl.Put(txs[i])
		if !sameTxs(mpl.GetAll(), txs[:i+1]) {
			t.Error("GetAll returns unproperly")
		}
	}
	ret := mpl.Get()
	if len(ret) != count {
		t.Error("put failed", len(ret), count)
	}
}
