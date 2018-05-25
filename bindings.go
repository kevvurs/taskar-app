package main

import "time"

type SensorPacket struct {
	Name   string   `json:"name"`
	Value  float64  `json:"value"`
}

type DeviceEvent struct {
	Packets    []SensorPacket `json:"data"`
	Timestamp  string         `json:"time"`
}

type SensorRecord struct {
	ReportTime   time.Time
	SensorName   string
	SensorValue  float64
}

type props struct {
	DataSource  string     `yaml:"data-source"`
	DateLayout  string     `yaml:"date-layout"`
	Query       queryProp  `yaml:"query"`
}

type queryProp struct {
	InsertRecord  string  `yaml:"insert-record"`
}
