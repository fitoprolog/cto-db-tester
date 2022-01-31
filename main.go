package main

import(
  "net/http"
  "net"
  "database/sql"
  _ "github.com/lib/pq"
  "log"
  //"strings"
)

func main(){
  //credentials hardcoded but don't worry there is no public access
  db,_:= sql.Open("postgres", "host=ctolearn.cluster-ro-cjincqaxcmb8.us-east-1.rds.amazonaws.com port=5432 user=postgres password=hdAXa4yVe7HWRXb dbname=postgres sslmode=disable")
  conn,_ := net.Dial("tcp", "golang.org:80")
  myIP:=conn.LocalAddr().String()
  //myIP=strings.Split(myIP,":")[0]
  http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
    stats:=""
    row,err:=db.Query("SELECT (usename || ' ' || application_name  || ' ' || client_addr)::text as results FROM pg_stat_activity")
    substats:=""
    if err !=nil{
      log.Printf(err.Error())
      return
    }
    for row.Next(){
      err=row.Scan(&substats)
      if err != nil{
        //log.Printf(err.Error())
        continue
      }
      stats+="\n"+substats
      log.Printf(substats)
    }
    res.Write([]byte(myIP+"\n"+stats))
  })
  http.ListenAndServe(":8080", nil)
}
