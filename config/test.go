package config

import (
    "testing"
    "github.com/anddddrew/galaxydb/tools"
)

func TestInitializeConfiguration(t *testing.T) {
      conf := InitializeConfiguration()
      tools.AssertEqual(t, conf.KeyspaceSize, int32(65536), "")
}

