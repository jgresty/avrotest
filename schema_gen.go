// Code generated by avrogen. DO NOT EDIT.

package main

import (
	"github.com/heetch/avro/avrotypegen"
)

type Data struct {
	Data *[]*string `json:"data"`
}

// AvroRecord implements the avro.AvroRecord interface.
func (Data) AvroRecord() avrotypegen.RecordInfo {
	return avrotypegen.RecordInfo{
		Schema: `{"connect.name":"test.Data","connect.version":1,"default":null,"fields":[{"default":null,"name":"data","type":["null",{"items":["null","string"],"type":"array"}]}],"name":"test.Data","type":"record"}`,
	}
}
