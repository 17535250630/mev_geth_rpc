package ethclient

type (
	EstimateGasBundleRes struct {
		IsSuccess  bool
		GasUsed    uint64
		ErrMessage string
		Return     []byte
		SpendTime  int64
	}
	EstimateGasRes struct {
		GasUsed   uint64
		Return    []byte
		SpendTime int64
	}
)
