package config

import (
      "io/ioutil"
      "encoding/json"
)

const (
      CONFIG_FILE = "/etc/galaxydb/conf.json"

      DEFAULT_KEYSPACE_SIZE = 65536
      DEFAULT_SYS_METRIC_INTERVAL = 300
      DEFAULT_APP_METRIC_INTERVAL = 300
      DEFAULT_TTL = -1
      DEFAULT_CRAWLER_INTERVAL = 300
      DEFAULT_SNAPSHIT_ENABLED = true
      DEFAULT_AOF_PERSISTENCE = false
      // 70mb
      DEFAULT_AOF_MAX_BYTES = 70000000
      DEFAULT_ENTRY_TIMESTAMP = true
      DEFAULT_ENABLE_ENCRYPTION = true
      DEFAULT_PASSPHRASE = "galaxydb"
)

type Configuration struct {

      KeyspaceSize int32

      SysMetricInterval int32

      AppMetricInterval int32

      DefaultTTL int32

      CrawlerInterval int32

      SnapshotInterval int32

      SnapshotEnabled bool

      PersistenceAOF bool

      AofMaxBytes int64

      EntryTimestamp bool

      EnableEncryption bool

      Passphrase string
}


