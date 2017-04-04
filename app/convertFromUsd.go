package main
 
import (
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)
 
type CurrencyFromUSD struct {
    Name        string	`json:"name"`
    Prefix 		string	`json:"prefix"`
    Postfix  	string	`json:"postfix"`
    FromUSD   	float32	`json:"fromUSD"`
}
 
var currencies []CurrencyFromUSD

func GetCurrenciesFromUSDEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(currencies)
}
 
func main() {
	router := mux.NewRouter()
    currencies = append(currencies, CurrencyFromUSD{Name: "CAD", Prefix: "$", Postfix: " CAD", FromUSD: 1.34})
    currencies = append(currencies, CurrencyFromUSD{Name: "EURO", Prefix: "&euro;", Postfix: " EUR", FromUSD: 0.94})
	
	router.HandleFunc("/currenciesFromUSD", GetCurrenciesFromUSDEndpoint).Methods("GET")

    log.Fatal(http.ListenAndServe(":8089", router))
}