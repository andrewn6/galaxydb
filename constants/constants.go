package constants

const (
	STORE_GET       = "get"
	STORE_PUT       = "put"
	STORE_ADD       = "add"
	STORE_DELETE    = "delete"
	STORE_FLUSH     = "flush"
	STORE_NODE_SIZE = "nodeSize"
	// grafana probably? or custom
	STORE_APP_METRICS = "getAppMetrics"

	LRU_TYPE  = "LRU"
	LFU_TYPE  = "LFU"
	MRU_TYPE  = "MRU"
	ARC_TYPE  = "ARC"
	TLRU_TYPE = "TLRU"
)
