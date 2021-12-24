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

func InitializeConfiguration() Configuration {
        var config, err = InitializeFromConfig()
        if err != nil {
                config.SetDefaultParams()
        }

        return config
}

func (conf *Configuration) SetDefaultParams() {
        conf.KeyspaceSize = DEFAULT_KEYSPACE_SIZE
        conf.SysMetricInterval = DEFAULT_SYS_METRIC_INTERVAL
        conf.AppMetricInterval = DEFAULT_APP_METRIC_INTERVAL
        conf.DefaultTTL = DEFAULT_TTL
        conf.CrawlerInterval = DEFAULT_CRAWLER_INTERVAL
        conf.SnapshotInterval = DEFAULT_SNAPSHOT_INTERVAL
        conf.PersistenceAOF = DEFAULT_AOF_PERSISTENCE
        conf.AofMaxBytes = DEFAULT_AOF_MAX_BYTES
        conf.EntryTimestamp = DEFAULT_ENTRY_TIMESTAMP
        conf.EnableEncryption = DEFAULT_ENABLE_ENCRYPTION
        conf.Passphrase = DEFAULT_PASSPHRASE
}

func InitializeFromConfig() (Configuration, error) {
        var config Configuration

        file, err := ioutil.ReadFile(CONFIG_FILE)

        if err != nil {
                return Configuration{}, err
        }

        err = json.Unmarshal(file, &config)

        if err != nil {
                return Configuration{}, err
        }

        return config, nil
}
