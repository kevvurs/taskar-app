package main

import (
  "log"
	"time"
)

func Event2Record(de *DeviceEvent) []*SensorRecord {
  t, err := time.Parse(Props.DateLayout, de.Timestamp)
  if err != nil {
  	log.Printf("ERROR: timestamp not valid <%v>\n", err)
  	return nil
  }
  rs := make([]*SensorRecord, len(de.Packets))
  for idx, sp := range de.Packets {
  	rs[idx] = &SensorRecord{
  		ReportTime:  t,//.Format(time.RFC3339),
  		SensorName:  sp.Name,
  		SensorValue: sp.Value,
  	}
  }
  return rs
}
