package define

type AesResult struct {
	Key  string
	Data string
}

const (
	AesCbcKeySize   = 32
	AesCbcIvVec     = 16
	AecCbcBlockSize = 32
)
