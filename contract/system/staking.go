/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */
package system

import (
	"encoding/binary"

	"github.com/aergoio/aergo/state"
	"github.com/aergoio/aergo/types"
)

var stakingkey = []byte("staking")

const StakingDelay = 10

func staking(txBody *types.TxBody, senderState *types.State,
	scs *state.ContractState, blockNo types.BlockNo) error {

	err := validateForStaking(txBody, scs, blockNo)
	if err != nil {
		return err
	}

	staked, _, err := getStaking(scs, txBody.Account)
	if err != nil {
		return err
	}

	err = setStaking(scs, txBody.Account, staked+txBody.Amount, blockNo)
	if err != nil {
		return err
	}

	senderState.Balance = senderState.Balance - txBody.Amount
	return nil
}

func unstaking(txBody *types.TxBody, senderState *types.State, scs *state.ContractState, blockNo types.BlockNo) error {
	staked, _, err := validateForUnstaking(txBody, scs, blockNo)
	if err != nil {
		return err
	}
	amount := txBody.Amount
	var backToBalance uint64
	if staked < amount {
		amount = 0
		backToBalance = staked
	} else {
		amount = staked - txBody.Amount
		backToBalance = txBody.Amount
	}
	err = setStaking(scs, txBody.Account, amount, 0)
	//blockNo will be updated in voting
	if err != nil {
		return err
	}
	err = voting(txBody, scs, blockNo)
	if err != nil {
		return err
	}

	senderState.Balance = senderState.Balance + backToBalance
	return nil
}

func setStaking(scs *state.ContractState, who []byte, balance uint64, blockNo uint64) error {
	key := append(stakingkey, who...)
	v := make([]byte, 16)
	binary.LittleEndian.PutUint64(v, balance)
	binary.LittleEndian.PutUint64(v[8:], blockNo) //TODO:change to block no
	//logger.Info().Str("key", util.enc.ToString(key)).Msg("VOTE setStaking")
	//logger.Info().Uint64("balance", balance).Uint64("blockNo", blockNo).Msg("VOTE setStaking")
	return scs.SetData(key, v)
}

func getStaking(scs *state.ContractState, who []byte) (uint64, uint64, error) {
	key := append(stakingkey, who...)
	data, err := scs.GetData(key)
	if err != nil {
		return 0, 0, err
	}
	var staked uint64
	var blockNo uint64
	if cap(data) == 0 {
		staked = 0
		blockNo = 0
	} else if cap(data) >= 8 {
		staked = binary.LittleEndian.Uint64(data[:8])
		blockNo = 0
		if cap(data) >= 16 {
			blockNo = binary.LittleEndian.Uint64(data[8:16])
		}
	}
	//logger.Info().Str("key", util.enc.ToString(key)).Msg("VOTE getStaking")
	//logger.Info().Uint64("staked", staked).Uint64("blockNo", blockNo).Msg("VOTE getStaking")
	return staked, blockNo, nil
}
