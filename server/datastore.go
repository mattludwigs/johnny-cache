package server

import (
  "encoding/json"
)

type DataStore struct {
  Method string `json:"method"`
  NameSpace string `json:"namespace,omitempty"`
  Data map[string]interface{} `json:"data,omitempty"`
  Query map[string]interface{} `json:"query,omitempty"`
}

type StatusResponse struct {
  Status string `json:"status"`
}

func FromJSON(b []byte) DataStore {
  d := DataStore{}
  err := json.Unmarshal(b, &d)
  HandleError(err, "FROM JSON")
  return d
}

func ToJSON(d interface{}) []byte {
  json, err := json.Marshal(d)
  HandleError(err, "TO JSON")
  return json
}

