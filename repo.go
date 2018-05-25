package main

import (
  "database/sql"
  "time"
	_ "github.com/go-sql-driver/mysql"
)

func save(rs []*SensorRecord) error {
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
  	_, err := stmt.Exec(r.ReportTime, r.SensorName, r.SensorValue)
 	  if err != nil {
 	    tx.Rollback()
 	   	return err
 	  }
  }
  tx.Commit()
	return nil
}

func fetch(filter *time.Time) ([]*SensorRecord, error) {
  var rs []*SensorRecord
	db, err := sql.Open("mysql", Props.DataSource)
  if err != nil {
  	return rs, err
  }
  defer db.Close()
  tx, err := db.Begin()
  if err != nil {
  	return rs, err
  }
  rows, err := tx.Query(Props.Query.SelectByTime, *filter)
  if err != nil {
  	tx.Rollback()
  	return rs, err
  }
  defer rows.Close()
  for rows.Next() {
		var (
			rt time.Time
			sn string
			sv float64
		)
		if err := rows.Scan(&rt, &sn, &sv); err != nil {
			tx.Rollback()
			return rs, err
		}
		r := &SensorRecord{
			ReportTime:  rt.In(seattle),
			SensorName:  sn,
			SensorValue: sv,
		}
		rs = append(rs, r)
  }
  tx.Commit()
	return rs, nil
}
