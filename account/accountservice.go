package account

import (
	"sync"

	"github.com/aergoio/aergo/account/key"

	"github.com/aergoio/aergo-actor/actor"

	"github.com/aergoio/aergo-lib/log"
	cfg "github.com/aergoio/aergo/config"
	"github.com/aergoio/aergo/message"
	"github.com/aergoio/aergo/pkg/component"
	"github.com/aergoio/aergo/types"
)

type AccountService struct {
	*component.BaseComponent
	cfg         *cfg.Config
	ks          *key.Store
	accountLock sync.RWMutex
	accounts    []*types.Account
	testConfig  bool
}

//NewAccountService create account service
func NewAccountService(cfg *cfg.Config) *AccountService {
	actor := &AccountService{
		cfg: cfg,
	}
	actor.BaseComponent = component.NewBaseComponent(message.AccountsSvc, actor, log.NewLogger("account"))

	return actor
}

func (as *AccountService) BeforeStart() {
	as.ks = key.NewStore(as.cfg.DataDir)

	as.accounts = []*types.Account{}
	addresses, err := as.ks.GetAddresses()
	if err != nil {
		as.Logger.Error().Err(err).Msg("could not open addresses")
	}
	for _, v := range addresses {
		as.accounts = append(as.accounts, &types.Account{Address: v})
	}
}

func (as *AccountService) AfterStart() {}

func (as *AccountService) BeforeStop() {
	as.ks.CloseStore()
	as.accounts = nil
}

func (as *AccountService) Statistics() *map[string]interface{} {
	return nil
}

func (as *AccountService) Receive(context actor.Context) {

	switch msg := context.Message().(type) {
	case *message.GetAccounts:
		accountList := as.getAccounts()
		context.Respond(&message.GetAccountsRsp{Accounts: &types.AccountList{Accounts: accountList}})
	case *message.CreateAccount:
		account, _ := as.createAccount(msg.Passphrase)
		context.Respond(&message.CreateAccountRsp{Account: account})
	case *message.LockAccount:
		account, err := as.lockAccount(msg.Account.Address, msg.Passphrase)
		context.Respond(&message.AccountRsp{Account: account, Err: err})
	case *message.UnlockAccount:
		account, err := as.unlockAccount(msg.Account.Address, msg.Passphrase)
		context.Respond(&message.AccountRsp{Account: account, Err: err})
	case *message.ImportAccount:
		account, err := as.importAccount(msg.Wif, msg.OldPass, msg.NewPass)
		context.Respond(&message.ImportAccountRsp{Account: account, Err: err})
	case *message.ExportAccount:
		wif, err := as.exportAccount(msg.Account.Address, msg.Pass)
		context.Respond(&message.ExportAccountRsp{Wif: wif, Err: err})
	case *message.SignTx:
		err := as.signTx(context, msg.Tx)
		if err != nil {
			context.Respond(&message.SignTxRsp{Tx: nil, Err: err})
		}
	case *message.VerifyTx:
		err := as.verifyTx(msg.Tx)
		if err != nil {
			context.Respond(&message.VerifyTxRsp{Tx: nil, Err: err})
		} else {
			context.Respond(&message.VerifyTxRsp{Tx: msg.Tx, Err: nil})
		}
	}
}

func (as *AccountService) getAccounts() []*types.Account {
	as.accountLock.RLock()
	defer as.accountLock.RUnlock()
	return as.accounts
}

func (as *AccountService) createAccount(passphrase string) (*types.Account, error) {
	address, err := as.ks.CreateKey(passphrase)
	if err != nil {
		return nil, err
	}
	account := types.NewAccount(address)

	//append list
	as.accountLock.Lock()
	//TODO: performance turning here
	as.ks.SaveAddress(address)
	as.accounts = append(as.accounts, account)
	as.accountLock.Unlock()
	return account, nil
}

func (as *AccountService) importAccount(wif []byte, old string, new string) (*types.Account, error) {
	address, err := as.ks.ImportKey(wif, old, new)
	if err != nil {
		return nil, err
	}
	//append list
	account := &types.Account{Address: address}
	as.accountLock.Lock()
	//TODO: performance turning here
	as.ks.SaveAddress(address)
	as.accounts = append(as.accounts, account)
	as.accountLock.Unlock()
	return account, nil
}

func (as *AccountService) exportAccount(address []byte, pass string) ([]byte, error) {
	wif, err := as.ks.ExportKey(address, pass)
	if err != nil {
		return nil, err
	}
	return wif, nil
}

func (as *AccountService) unlockAccount(address []byte, passphrase string) (*types.Account, error) {
	addr, err := as.ks.Unlock(address, passphrase)
	if err != nil {
		as.Warn().Err(err).Msg("could not find the key")
		return nil, err
	}
	return &types.Account{Address: addr}, nil
}

func (as *AccountService) lockAccount(address []byte, passphrase string) (*types.Account, error) {
	addr, err := as.ks.Lock(address, passphrase)
	if err != nil {
		as.Warn().Err(err).Msg("could not load the key")
		return nil, err
	}
	return &types.Account{Address: addr}, nil
}

func (as *AccountService) signTx(c actor.Context, tx *types.Tx) error {
	//sign tx
	prop := actor.FromInstance(NewSigner(as.ks))
	signer := c.Spawn(prop)
	signer.Request(tx, c.Sender())
	return nil
}

func (as *AccountService) verifyTx(tx *types.Tx) error {
	return as.ks.VerifyTx(tx)
}
