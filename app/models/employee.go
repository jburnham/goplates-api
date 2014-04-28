package models

import "encoding/xml"

type Employee struct {
	XMLName xml.Name `json:"-" db:"-" xml:"employee"`
	Id      int      `json:"id" xml:"id"`
	Name    string   `json:"name" xml:"name"`
	Email   string   `json:"-" xml:"-"`
	Phone   int      `json:"-" xml:"-"`
}
