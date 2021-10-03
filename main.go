package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/hiroki-kondo-git/call_alpha_vantage_api.git/parse"
)

func main() {
	// コマンドラインの受付エラー処理
	if len(os.Args) != 2 {
		log.Fatal("Usage: program symbol")
	}
	symbol := os.Args[1]

	apiKey, err := readAPIKey()
	if err != nil {
		log.Fatal(err)
	}

	json, err := getStock(symbol, apiKey)
	if err != nil {
		log.Fatal(err)
	}

	parse.ParseJSON(json)
}

func getStock(symbol, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", symbol, apiKey)

	// httpクライアント
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// レスポンスのボディを全て読み出す
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

	// resp.Body.Closeはここで実行
}

func readAPIKey() (string, error) {
	apiKey := os.Getenv("API_KEY")
	// API_KEYが環境変数に登録されていなかったらエラーを返す
	if apiKey == "" {
		return "", fmt.Errorf("API_KEY environment variable is not defined")
	}
	return apiKey, nil
}
