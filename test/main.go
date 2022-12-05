package main

import (
	"bet/core/game"
	"bet/model"
	"bet/net/tron/pkg/address"
	"bet/utils"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
)

//	func Retry() (error, error) {
//		ctx, cancel := context.WithTimeout(
//			context.Background(),
//			10*time.Second,
//		)
//		defer cancel()
//		for {
//			res, err := dosomething()
//			if err == nil {
//				return res, nil
//			}
//			select {
//			case <-ctx.Done():
//				return nil, fmt.Errorf("timeout while dosomething (last error: %w)", err)
//			case <-time.After(5 * time.Second):
//			}
//		}
//	}

func GenerateKey() (wif string, address1 string) {
	pri, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", ""
	}
	if len(pri.D.Bytes()) != 32 {
		for {
			pri, err = btcec.NewPrivateKey(btcec.S256())
			if err != nil {
				continue
			}
			if len(pri.D.Bytes()) == 32 {
				break
			}
		}
	}
	address1 = address.PubkeyToAddress(pri.ToECDSA().PublicKey).String()
	wif = hex.EncodeToString(pri.D.Bytes())
	return
}

func main() {

	sysGamePlayList := []model.SysGamePlay{}

	utils.DB.Find(&sysGamePlayList, "game_code=? and status=1", 1000)
	GenerateKey()
	a := 1136.1
	b := a * 100
	utils.Init()
	as := &model.SysBalanceCode{}
	issue := game.GetIssue(1000)
	fmt.Print(issue)
	err := as.SysBalanceCodeDB().First(as, 100).Error
	if err != nil {

	}
	as.GetByCodeCache(1)
	fmt.Println(as)
	fmt.Println(b)
}
