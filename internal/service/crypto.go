package service

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/JorgeAdd/go-final-skills-cto/cryptoAPI/internal/database"
)

// type Bitso struct {
// 	Success string        `json:"success"`
// 	Payload Payload `json:"payload"`
// }

type Payload struct {
	Id         string `json:"id"`
	High       string `json:"high"`
	Last_mxn   string `json:"last_mxn"`
	Last_usd   string `json:"last_usd,omitempty"`
	Last_hkd   string `json:"last_hkd,omitempty"`
	Created_at string `json:"created_at"`
	Volume     string `json:"volume"`
	Vwap       string `json:"vwap"`
	Low        string `json:"low"`
	Ask        string `json:"ask"`
	Bid        string `json:"bid"`
	Change_24  string `json:"change_24"`
}

type Filters struct {
	Start_date string `json:"start_date"`
	End_date   string `json:"end_date"`
	Usd_book   bool   `json:"usd_book"`
	Hkd_book   bool   `json:"hkd_book"`
}

func getConnection() *sql.DB {
	return database.GetDB()
}

func GetCryptoInfo() []Payload {
	db := getConnection()
	defer db.Close()
	var payload Payload
	var payloads []Payload

	payloadRes, errPayload := db.Query("SELECT * FROM crypto_values.crypto_values")
	if errPayload != nil {
		panic(errPayload)
	}
	defer payloadRes.Close()
	for payloadRes.Next() {
		errPayload = payloadRes.Scan(&payload.Id, &payload.High, &payload.Last_mxn, &payload.Last_usd, &payload.Last_hkd,
			&payload.Created_at, &payload.Volume, &payload.Vwap, &payload.Low, &payload.Ask, &payload.Bid, &payload.Change_24)
		if errPayload != nil {
			panic(errPayload)
		}
		payloads = append(payloads, payload)
	}

	return payloads
}

func GetCryptoFilterDate(body string) []Payload {
	db := getConnection()
	defer db.Close()
	var payload Payload
	var payloads []Payload
	var filters Filters

	json.Unmarshal([]byte(body), &filters)

	fmt.Println(filters)
	var payloadRes *sql.Rows
	var errPayload error

	payloadRes, errPayload = db.Query("SELECT * FROM crypto_values.crypto_values WHERE created_at BETWEEN '" + filters.Start_date + "' AND '" + filters.End_date + "';")
	if errPayload != nil {
		panic(errPayload)
	}

	defer payloadRes.Close()
	for payloadRes.Next() {
		errPayload = payloadRes.Scan(&payload.Id, &payload.High, &payload.Last_mxn, &payload.Last_usd, &payload.Last_hkd,
			&payload.Created_at, &payload.Volume, &payload.Vwap, &payload.Low, &payload.Ask, &payload.Bid, &payload.Change_24)
		if errPayload != nil {
			panic(errPayload)
		}
		payloads = append(payloads, payload)
	}

	return payloads
}

func GetCryptoFilterBook(body string) []Payload {
	db := getConnection()
	defer db.Close()
	var payload Payload
	var payloads []Payload
	var filters Filters

	filters.Usd_book = false
	filters.Hkd_book = false
	json.Unmarshal([]byte(body), &filters)

	payloadRes, errPayload := db.Query("SELECT * FROM crypto_values.crypto_values")
	if errPayload != nil {
		panic(errPayload)
	}
	defer payloadRes.Close()
	for payloadRes.Next() {
		errPayload = payloadRes.Scan(&payload.Id, &payload.High, &payload.Last_mxn, &payload.Last_usd, &payload.Last_hkd,
			&payload.Created_at, &payload.Volume, &payload.Vwap, &payload.Low, &payload.Ask, &payload.Bid, &payload.Change_24)
		if errPayload != nil {
			panic(errPayload)
		}
		if !filters.Usd_book {
			payload.Last_usd = ""
		}

		if !filters.Hkd_book {
			payload.Last_hkd = ""
		}

		payloads = append(payloads, payload)
	}

	return payloads
}
