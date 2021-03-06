// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package serialization

import "strconv"

type OrderType int8

const (
	OrderTypeUNKNOWN OrderType = 0
	OrderTypeLIMIT   OrderType = 1
	OrderTypeMARKET  OrderType = 2
)

var EnumNamesOrderType = map[OrderType]string{
	OrderTypeUNKNOWN: "UNKNOWN",
	OrderTypeLIMIT:   "LIMIT",
	OrderTypeMARKET:  "MARKET",
}

var EnumValuesOrderType = map[string]OrderType{
	"UNKNOWN": OrderTypeUNKNOWN,
	"LIMIT":   OrderTypeLIMIT,
	"MARKET":  OrderTypeMARKET,
}

func (v OrderType) String() string {
	if s, ok := EnumNamesOrderType[v]; ok {
		return s
	}
	return "OrderType(" + strconv.FormatInt(int64(v), 10) + ")"
}
