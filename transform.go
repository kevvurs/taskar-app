package main

import (
  "log"
	"time"
)

var utc, seattle *time.Location

func init() {
	utc, _ = time.LoadLocation("UTC")
	seattle, _ = time.LoadLocation("America/Los_Angeles")
}

func Event2Record(de *DeviceEvent) []*SensorRecord {
  t, err := time.ParseInLocation(Props.DateLayout, de.Timestamp, utc)
  if err != nil {
  	log.Printf("ERROR: timestamp not valid <%v>\n", err)
  	return nil
  }
  if t.Year() == 2000 {
    // timestamp initial value is 2000-2-1 00:00:00
  	log.Println("WARN: GPS signal is disconnected")
  	// return nil
  }
  t = t.In(seattle)  // the timestamp from the device (unstable)
  now := time.Now().In(seattle)  // post-generated
  rs := make([]*SensorRecord, len(de.Packets))
  for idx, sp := range de.Packets {
  	rs[idx] = &SensorRecord{
  		ReportTime:  now,
  		SensorName:  sp.Name,
  		SensorValue: sp.Value,
  	}
  }
  return rs
}

func Filter(param string) (time.Time, error) {
	return time.ParseInLocation(time.RFC3339, param, seattle)
}
