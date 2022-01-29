package main

import(
  "net/http"
  "net"
  "database/sql"
  _ "github.com/lib/pq"
  "log"
  "strings"
)

func main(){
  //credentials hardcoded but don't worry there is no public access
  db,_:= sql.Open("postgres", "host=ctolearn.cluster-ro-cjincqaxcmb8.us-east-1.rds.amazonaws.com port=5432 user=postgres password=hdAXa4yVe7HWRXb dbname=postgres sslmode=disable")
  conn,_ := net.Dial("udp", "8.8.8.8:80")
  myIP:=conn.LocalAddr().String()
  myIP=strings.Split(myIP,":")[0]
  http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
    row:=db.QueryRow("SELECT version(),NOW()")
    version :=""
    now     :=""
    err:=row.Scan(&version,&now)
    if err != nil{
      log.Printf(err.Error())
    }
    res.Write([]byte(myIP+"\n"+version+" "+now))
  })
  http.ListenAndServe(":8080", nil)
}
