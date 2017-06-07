package clientApis

import (
  "gopkg.in/resty.v0"

  "github.com/danielhood/loco.clr/config"
)

func GetToken() (string, error) {
  tokenResp, err := resty.R().Get(config.LocoServer() + "/token")

  if err != nil {
    return "", err
  }

  token := tokenResp.String()

  return token, nil
}
