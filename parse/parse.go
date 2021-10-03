package parse

import (
	"encoding/json"
	"fmt"
	"sort"
)

// JSONのパース
type Schema struct {
	MetaData   map[string]string     `json:"Meta Data"`
	TimeSeries map[string]StockValue `json:"Time Series (5min)"`
}

type StockValue map[string]string

func ParseJSON(js []byte) error {
	var sch Schema
	if err := json.Unmarshal(js, &sch); err != nil {
		return err
	}
	// メタ情報の出力を順番通りにする
	var keys []string
	for key := range sch.MetaData {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s : %s\n", key, sch.MetaData[key])
	}
	// 株価の出力（最新時間のもののみ）
	for k, v := range sch.TimeSeries {
		fmt.Println(k, v)
		break
	}
	return nil

}
