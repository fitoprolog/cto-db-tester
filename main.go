package main

import(
  "net/http"
  "database/sql"
  _ "github.com/lib/pq"
  "log"
)

func main(){
  //credentials hardcoded but don't worry there is no public access
  db,_:= sql.Open("postgres", "host=ctolearn.cluster-ro-cjincqaxcmb8.us-east-1.rds.amazonaws.com port=5432 user=postgres password=hdAXa4yVe7HWRXb dbname=postgres sslmode=disable")
  http.HandleFunc("/",func(res http.ResponseWriter, req *http.Request){
    row:=db.QueryRow("SELECT version(),NOW()")
    version :=""
    now     :=""
    err:=row.Scan(&version,&now)
    if err != nil{
      log.Printf(err.Error())
    }
    res.Write([]byte(version+" "+now))
  })
  http.ListenAndServe(":8081", nil)
}
