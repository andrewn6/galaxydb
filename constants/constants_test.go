package constants

import (
				"testing"

				"github.com/anddddrew/galaxydb/tools"
)

func ConstantsTest(t *testing.T) {
			tools.AssertEqual(t, STORE_GET, "get", "")
			tools.AssertEqual(t, STORE_PUT, "put", "")
			tools.AssertEqual(t, STORE_ADD, "add", "")
			tools.AssertEqual(t, STORE_DELETE, "delete", "")
			tools.AssertEqual(t, STORE_FLUSH, "flush", "")
			tools.AssertEqual(t, STORE_NODE_SIZE, "nodeSize", "")
			tools.AssertEqual(t, STORE_APP_METRICS, "getAppMetrics", "")

			tools.AssertEqual(t, LRU_TYPE "LRU", "")
			tools.AssertEqual(t, LFU_TYPE, "LFU", "")
			tools.AssertEqual(t, MRU_TYPE, "MRU", "")
			tools.AssertEqual(t, ARC_TYPE, "ARC", "")
			tools.AssertEqual(t, TLRU_TYPE, "TLRU", "")

}
