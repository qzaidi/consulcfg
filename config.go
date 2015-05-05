package config

import (
  "reflect"
  "fmt"
  "strconv"
  "log"
  consul "github.com/hashicorp/consul/api"
)


// recursively read a structure from consul via reflection
func readStruct (getter func(string) string, prefix string, st reflect.Type, v reflect.Value) {
  for i := 0; i < st.NumField(); i++ {
    field := st.Field(i)
    val := v.Field(i)
    switch field.Type.Kind() {
    case reflect.Struct:
      readStruct(getter,prefix + "/" + field.Name,field.Type,val)
    case reflect.String:
      key := fmt.Sprintf("%s/%s",prefix,field.Name)
      if  cval := getter(key); cval != "" {
        val.SetString(cval)
      }
    case reflect.Int:
      key := fmt.Sprintf("%s/%s",prefix,field.Name)
      if  cval := getter(key); cval != "" {
        i,err := strconv.ParseInt(cval,10,64)
        if err == nil {
          val.SetInt(i)
        }
      }
    }
  }
}


func ReadConfig(namespace string, cfg interface{}) bool {
  client,_ := consul.NewClient(consul.DefaultConfig())
  kv := client.KV()
  status := false // if at least one field was read successfully, status will be true

  readConsulKey := func(name string) string {
    pair, _, err := kv.Get(name,nil)
    if err != nil {
      log.Println("err",err)
      return ""
    }
    status = true
    return string(pair.Value[:])
  }

  readStruct(readConsulKey,"wallet",reflect.TypeOf(cfg).Elem(),reflect.ValueOf(cfg).Elem())
  return status
}

// marshals struct into consul key value pairs
func SaveConfig(namespace string, cfg interface{}) bool {
  return false
}
