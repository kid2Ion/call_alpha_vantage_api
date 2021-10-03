# call_alpha_vantage_api
株価の取得API

https://www.alphavantage.co/
でAPIキーの取得、環境変数API_KEY=""に設定が必要

go run の後に株価取得したい企業のティッカーsymbolを入力

ティッカーシンボルは以下を参照
https://search.sbisec.co.jp/v2/popwin/info/stock/pop6040_usequity_list.html

ex) go run . GOOGL
