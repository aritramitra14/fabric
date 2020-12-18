package timedkeys


import (
	//"bytes"
	//"errors"
	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/rwsetutil"
	"github.com/hyperledger/fabric/protos/common"
	"strconv"
	//"sync"
	"time"

	//"time"
	putils "github.com/hyperledger/fabric/protos/utils"
)


type TimedKeys struct {

	level1 map[string][]byte
	//level2 map[string][]byte
	//level3 map[string][]byte
	//lock sync.RWMutex
	//lastShiftTime time.Time

}

func (tk *TimedKeys) Commit(block *common.Block) error{
	newBlockTime, err4 := ptypes.Timestamp(block.Metadata.BlockTime)
	if err4 !=nil {
		return err4

	}
	var stamp string
	keytime := newBlockTime.UnixNano()/(int64(time.Millisecond))
	stamp = strconv.FormatInt(keytime,10)
	if tk.level1 == nil {
		tk.level1 = make(map[string][]byte)
	}
	for _, envBytes := range block.Data.Data {
		env, err := putils.GetEnvelopeFromBlock(envBytes)
		if err !=nil{
			return err
		}
		payload, err := putils.GetPayload(env)
		if err !=nil{
			return err
		}
		chdr, err := putils.UnmarshalChannelHeader(payload.Header.ChannelHeader)
		if err !=nil{
			return err
		}

		if common.HeaderType(chdr.Type) == common.HeaderType_ENDORSER_TRANSACTION {
			respPayload, err := putils.GetActionFromEnvelope(envBytes)
			if err !=nil{
				return err
			}
			txRWSet := &rwsetutil.TxRwSet{}
			if err = txRWSet.FromProtoBytes(respPayload.Results); err != nil {
				return err
			}
			for _, nsRWSet := range txRWSet.NsRwSets {
				for _, kvWrite := range nsRWSet.KvRwSet.Writes {
					writeKey := kvWrite.Key+stamp
					writeValue := kvWrite.Value
					tk.level1[writeKey] = writeValue
				}
			}


		}



	}

	return nil
}


func (tk *TimedKeys) GetHistory(sk string, ek string) [][]byte {

	res :=  make([][]byte,0)

	for k:= range tk.level1 {

		if (k> sk) && (k<ek) {
			res = append(res, tk.level1[k])
		}
	}

	return res
}





func (tk *TimedKeys) GetCount() int {
	return len(tk.level1)
}

