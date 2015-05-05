package config

import (
  "testing"
)

type ServerCfg struct {
  Port string
}

type PromocodeCfg struct {
  MerchantGuid string
  SalesWalletGuid string
}

type UpstreamCfg struct {
  URI       string
  Timeout   int         // timeout in seconds
  ClientKey string
  HMACKey   string
  Platform string
  Promocode PromocodeCfg
}


type Config struct {
  Server ServerCfg
  Upstream UpstreamCfg
}

func TestReadConfigConsul(t *testing.T) {
  var cfg Config
  if ReadConfig("wallet",&cfg) == false {
    t.Fail()
  }
  t.Log("upstream.URI=",cfg.Upstream.URI)
  t.Log("upstream.ClientKey=",cfg.Upstream.ClientKey)
  t.Log("upstream.Platform=",cfg.Upstream.Platform)
  t.Log("upstream.HMACKey=",cfg.Upstream.HMACKey)
  t.Log("server.Port=",cfg.Server.Port)
}
