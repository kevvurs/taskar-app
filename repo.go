package main

import (
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func upload(rs []*SensorRecord) error {
  db, err := sql.Open("mysql", Props.DataSource)
  if err != nil {
  	return err
  }
  defer db.Close()
  tx, err := db.Begin()
  if err != nil {
  	return err
  }
  stmt, err := tx.Prepare(Props.Query.InsertRecord)
  if err != nil {
  	tx.Rollback()
  	return err
  }
  defer stmt.Close()
  for _, r := range rs {
//  log.Printf("INFO: record- %v", *r)
  	_, err := stmt.Exec(r.ReportTime, r.SensorName, r.SensorValue)
 	  if err != nil {
 	    tx.Rollback()
 	   	return err
 	  }
  }
  tx.Commit()
	return nil
}
