package server

func HandleGet(dstore DataStore) []map[string]interface{} {
  query := dstore.Query
  keys := GetKeys(query)
  store := cache.Stores[dstore.NameSpace]

  res := make([]map[string]interface{}, 0)

  // @TODO: Probably should do some better stuff here
  for _, key := range keys {
    for _, value := range store {
      if _, okay := value[key.(string)]; okay {
        if query[key.(string)] == value[key.(string)] {
          res = append(res, value)
        }
      }
    }
  }

  return res
}

func GetKeys(s map[string]interface{}) []interface{} {
  keys := make([]interface{}, 0, len(s))
  for k := range s {
    keys = append(keys, k)
  }
  return keys
}
