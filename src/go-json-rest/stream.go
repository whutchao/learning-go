package main

import (
    "fmt"
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
    "time"
)

func main() {
  api := api.NewApi()
  api.Use(&rest.AccessLogApacheMiddleware{})
  api.Use(rest.DefaultDveStack...)
  router, err := api.MakeRouter(
    rest.Get("/stream", StreamThings),
  )
  if err != nil{
    log.Fatal(err)
  }
  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Thing struct{
  Name string
}

func StreamThings(w rest.ResponseWriter, r *rest.Request){
  cpt := 0
  for{
    cpt++
    w.WriteJson(
      &Thing{
        Name: fmt.Sprintf("thing #%d", cpt),
      },
    )
    w.(http.ResponseWriter).Write([]byte("\n"))
    w.(http.Flusher).Flush()
    time.Sleep(time.Duration(3) * time.Second)
  }
}
