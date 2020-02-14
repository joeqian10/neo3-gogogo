package examples

import (
	"fmt"

	"github.com/joeqian10/neo3-gogogo/rpc"
)

func main() {
	var rpc = rpc.NewClient("http://127.0.0.1:10332")

	response := rpc.GetBestBlockHash()
	r := response.Result
	fmt.Println(r)

	response1 := rpc.GetRawTransaction("0xf694b1d630de0480dc2bb47a4f4edc37877565b4be1226ffbe03e91e9f2a25cf")
	fmt.Println(response1.Result)

	result := rpc.OpenWallet("D:\\temp\\wallet.json", "1111")
	fmt.Println(result.Result)
}
