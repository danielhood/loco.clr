package main

import (
  "crypto/tls"
  "log"

  "gopkg.in/resty.v0"

)

const (
  server = "https://loco.local:8080"
  Version = "0.1"
)

func main() {
  log.Print("loco.clr v", Version, "starting...")

  log.Print("getting token from server: ", server)

  resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })
  tokenResp, err := resty.R().Get(server + "/token")

  if err != nil {
    log.Print("Error: ", err)
    return
  }

  token := tokenResp.String()
  log.Print("got token: ", token)

  log.Print("getting object list")
  resty.SetAuthToken(token)
  objectsResp, err := resty.R().Get(server + "/object")

  log.Print("Objects: ", objectsResp.String())

  log.Print("Done")
}
