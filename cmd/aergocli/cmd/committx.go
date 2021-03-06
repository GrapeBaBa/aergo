/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package cmd

import (
	"context"

	sha256 "github.com/minio/sha256-simd"

	"io/ioutil"

	"github.com/aergoio/aergo/cmd/aergocli/util"
	"github.com/aergoio/aergo/types"
	"github.com/mr-tron/base58/base58"
	"github.com/spf13/cobra"
)

var committxCmd = &cobra.Command{
	Use:   "committx",
	Short: "Send transaction",
	Args:  cobra.MinimumNArgs(0),
	Run:   execCommitTX,
}

var recipient string

var jsonTx string
var jsonPath string

func init() {
	rootCmd.AddCommand(committxCmd)
	committxCmd.Flags().StringVar(&jsonTx, "jsontx", "", "Transaction list json")
	committxCmd.Flags().StringVar(&jsonPath, "jsontxpath", "", "Transaction list json file path")
}

func execCommitTX(cmd *cobra.Command, args []string) {
	var msg *types.CommitResultList
	if jsonPath != "" {
		b, readerr := ioutil.ReadFile(jsonPath)
		if readerr != nil {
			cmd.Printf("Failed: %s\n", readerr.Error())
			return
		}
		jsonTx = string(b)
	}

	if jsonTx != "" {
		txlist, err := util.ParseBase58Tx([]byte(jsonTx))
		if err != nil {
			cmd.Printf("Failed: %s\n", err.Error())
			return
		}
		msg, err = client.CommitTX(context.Background(), &types.TxList{Txs: txlist})
		if err != nil {
			cmd.Printf("Failed: %s\n", err.Error())
			return
		}
	}
	for i, r := range msg.Results {
		cmd.Println(i+1, ":", base58.Encode(r.Hash), r.Error, r.Detail)
	}
}

func caculateHash(body *types.TxBody) []byte {
	input := append(body.Recipient, body.Payload...)
	sum1 := sha256.Sum256(input)
	sum2 := sha256.Sum256(sum1[:])
	return sum2[:]
}
