package types

import "encoding/json"

var RpcTradeTypeTransList = []RpcTypeInfo{
	{
		"GetTokenSellOrderByStatus",
		TradeQueryTokenSellOrder{},
	},
	{
		"GetOnesSellOrderWithStatus",
		TradeQueryOnesSellOrder{},
	},
	{
		"GetOnesSellOrder",
		TradeQueryOnesSellOrder{},
	},

	{
		"GetTokenBuyOrderByStatus",
		TradeQueryTokenBuyOrder{},
	},
	{
		"GetOnesBuyOrderWithStatus",
		TradeQueryOnesBuyOrder{},
	},
	{
		"GetOnesBuyOrder",
		TradeQueryOnesBuyOrder{},
	},
}

// rpc query trade sell order part

type RpcReplySellOrder struct {
	TokenSymbol       string `json:"tokenSymbol"`
	Owner             string `json:"owner"`
	AmountPerBoardlot int64  `json:"amountPerBoardlot"`
	MinBoardlot       int64  `json:"minBoardlot"`
	PricePerBoardlot  int64  `json:"pricePerBoardlot"`
	TotalBoardlot     int64  `json:"totalBoardlot"`
	SoldBoardlot      int64  `json:"soldBoardlot"`
	BuyID             string `json:"buyID"`
	Status            int32  `json:"status"`
	SellID            string `json:"sellID"`
	TxHash            string `json:"txHash"`
	Height            int64  `json:"height"`
	Key               string `json:"key"`
}
type RpcReplySellOrders struct {
	SellOrders []*RpcReplySellOrder `json:"sellOrders"`
}

type TradeQueryTokenSellOrder struct {
}

func (t *TradeQueryTokenSellOrder) Input(message json.RawMessage) ([]byte, error) {
	var req ReqTokenSellOrder
	err := json.Unmarshal(message, &req)
	if err != nil {
		return nil, err
	}
	return Encode(&req), nil
}

func (t *TradeQueryTokenSellOrder) Output(reply *ReplySellOrders) (*RpcReplySellOrders, error) {
	str, err := json.Marshal(*reply)
	if err != nil {
		return nil, err
	}
	var rpcReply RpcReplySellOrders
	json.Unmarshal(str, &rpcReply)
	return &rpcReply, nil
}

type TradeQueryOnesSellOrder struct {
}

func (t *TradeQueryOnesSellOrder) Input(message json.RawMessage) ([]byte, error) {
	var req ReqAddrTokens
	err := json.Unmarshal(message, &req)
	if err != nil {
		return nil, err
	}
	return Encode(&req), nil
}

func (t *TradeQueryOnesSellOrder) Output(reply *ReplySellOrders) (*RpcReplySellOrders, error) {
	str, err := json.Marshal(*reply)
	if err != nil {
		return nil, err
	}
	var rpcReply RpcReplySellOrders
	json.Unmarshal(str, &rpcReply)
	return &rpcReply, nil
}

// rpc query trade buy order

type RpcReplyBuyOrder struct {
	TokenSymbol       string `json:"tokenSymbol"`
	Owner             string `json:"owner"`
	AmountPerBoardlot int64  `json:"amountPerBoardlot"`
	MinBoardlot       int64  `json:"minBoardlot"`
	PricePerBoardlot  int64  `json:"pricePerBoardlot"`
	TotalBoardlot     int64  `json:"totalBoardlot"`
	BoughtBoardlot    int64  `json:"boughtBoardlot"`
	BuyID             string `json:"buyID"`
	Status            int32  `json:"status"`
	SellID            string `json:"sellID"`
	TxHash            string `json:"txHash"`
	Height            int64  `json:"height"`
	Key               string `json:"key"`
}

type RpcReplyBuyOrders struct {
	BuyOrders []*RpcReplyBuyOrder `json:"buyOrders"`
}

type TradeQueryTokenBuyOrder struct {
}

func (t *TradeQueryTokenBuyOrder) Input(message json.RawMessage) ([]byte, error) {
	var req ReqTokenBuyOrder
	err := json.Unmarshal(message, &req)
	if err != nil {
		return nil, err
	}
	return Encode(&req), nil
}

func (t *TradeQueryTokenBuyOrder) Output(reply *ReplyBuyOrders) (*RpcReplyBuyOrders, error) {
	str, err := json.Marshal(*reply)
	if err != nil {
		return nil, err
	}
	var rpcReply RpcReplyBuyOrders
	json.Unmarshal(str, &rpcReply)
	return &rpcReply, nil
}

type TradeQueryOnesBuyOrder struct {
}

func (t *TradeQueryOnesBuyOrder) Input(message json.RawMessage) ([]byte, error) {
	var req ReqAddrTokens
	err := json.Unmarshal(message, &req)
	if err != nil {
		return nil, err
	}
	return Encode(&req), nil
}

func (t *TradeQueryOnesBuyOrder) Output(reply *ReplyBuyOrders) (*RpcReplyBuyOrders, error) {
	str, err := json.Marshal(*reply)
	if err != nil {
		return nil, err
	}
	var rpcReply RpcReplyBuyOrders
	json.Unmarshal(str, &rpcReply)
	return &rpcReply, nil
}