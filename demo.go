package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello! Wellcome to e-ticketing for Movie Show")
	fmt.Fprintf(w, "For more details visit 	  : http://127.0.0.1:2619/details/")
}

func details_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Copy the required link and paste it to browser!")
	fmt.Fprintln(w, "Index       : http://127.0.0.1:2619/")
	fmt.Fprintln(w, "About       : http://127.0.0.1:2619/about/")
	fmt.Fprintln(w, "Details       : http://127.0.0.1:2619/details/")
	fmt.Fprintln(w, "Theatres       : http://127.0.0.1:2619/theatre/")
	fmt.Fprintln(w, "Zamuna movie lists       : http://127.0.0.1:2619/zamuna/")
	fmt.Fprintln(w, "Bashundhra movie lists   : http://127.0.0.1:2619/bashundhara/")
	fmt.Fprintln(w, "Padma movie lists       : http://127.0.0.1:2619/padma/")
	fmt.Fprintln(w, "Meghna movie lists       : http://127.0.0.1:2619/meghna/")
	fmt.Fprintln(w, "Zamuna Show times       : http://127.0.0.1:2619/zamunashowtime/")
	fmt.Fprintln(w, "Bashundhara Show times   : http://127.0.0.1:2619/bashundharashowtime/")
	fmt.Fprintln(w, "Padma Show times         : http://127.0.0.1:2619/padmashowtime/")
	fmt.Fprintln(w, "Meghna Show times        : http://127.0.0.1:2619/meghnashowtime/")

}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "@Sadiquzzaman Shovon")
}

func theatre_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Which theatre do you want to choose?")
	fmt.Fprintln(w, "1. Zamuna Theatre,Badda.")
	fmt.Fprintln(w, "2. Bashundhara Theatre,Panthapath.")
	fmt.Fprintln(w, "3. Padma Theatre, Uttara.")
	fmt.Fprintln(w, "4. Jamuna Theatre, Dhanmondi.")
}

func zamuna_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Which movie do you want to choose?")
	fmt.Fprintln(w, "1. The Shawshank Redemption.")
	fmt.Fprintln(w, "2. The Godfather.")
	fmt.Fprintln(w, "3. The Dark Knight.")
	fmt.Fprintln(w, "4. 12 Angry Men.")
}

func zamunashowtime_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1. The Shawshank Redemption :      7.00 AM  && 9.00 AM")
	fmt.Fprintln(w, "2. The Godfather	:    	10.00 AM  &&  12.00 PM")
	fmt.Fprintln(w, "3. The Dark Knight 	:    	1.00 PM  &&  3.00 PM")
	fmt.Fprintln(w, "4. 12 Angry Men             :      4.00 AM  &&  6.00 PM")
}

func bashundhara_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Which movie do you want to choose?")
	fmt.Fprintln(w, "1. TeNet.")
	fmt.Fprintln(w, "2. Dark.")
	fmt.Fprintln(w, "3. The Lord of Rings.")
	fmt.Fprintln(w, "4. Fight Club.")
}

func bashundharashowtime_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1. TeNet             :    7.00 AM  && 9.00 AM")
	fmt.Fprintln(w, "2. Dark              :    10.00 AM  &&  12.00 PM")
	fmt.Fprintln(w, "3. The Lord of Rings :    1.00 PM  &&  3.00 PM")
	fmt.Fprintln(w, "4. Fight Club        :    4.00 AM  &&  6.00 PM")
}

func padma_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Which movie do you want to choose?")
	fmt.Fprintln(w, "1. Joker.")
	fmt.Fprintln(w, "2. The Godfather.")
	fmt.Fprintln(w, "3. The Dark k Dies.")
	fmt.Fprintln(w, "4. 12 Angry Men.")
}

func padmashowtime_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1. Joker 		:    7.00 AM  && 9.00 AM")
	fmt.Fprintln(w, "2. The Godfather 			:    10.00 AM  &&  12.00 PM")
	fmt.Fprintln(w, "3. The Dark k Dies 			:    1.00 PM  &&  3.00 PM")
	fmt.Fprintln(w, "4. 12 Angry Men 			:    4.00 AM  &&  6.00 PM")
}

func meghna_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Which movie do you want to choose?")
	fmt.Fprintln(w, "1. Interstaller.")
	fmt.Fprintln(w, "2. La Casa de Papel.")
	fmt.Fprintln(w, "3. The Professor.")
	fmt.Fprintln(w, "4. Inception.")
}

func meghnashowtime_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "1. Interstaller 	 :    7.00 AM  && 9.00 AM")
	fmt.Fprintln(w, "2. La Casa de Papel 	 :    10.00 AM  &&  12.00 PM")
	fmt.Fprintln(w, "3. The Professor	 :    1.00 PM  &&  3.00 PM")
	fmt.Fprintln(w, "4. Inception 		 :    4.00 AM  &&  6.00 PM")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.HandleFunc("/theatre/", theatre_handler)
	http.HandleFunc("/details/", details_handler)
	http.HandleFunc("/zamuna/", zamuna_handler)
	http.HandleFunc("/bashundhara/", bashundhara_handler)
	http.HandleFunc("/padma/", padma_handler)
	http.HandleFunc("/meghna/", meghna_handler)
	http.HandleFunc("/zamunashowtime/", zamunashowtime_handler)
	http.HandleFunc("/bashundharashowtime/", bashundharashowtime_handler)
	http.HandleFunc("/padmashowtime/", padmashowtime_handler)
	http.HandleFunc("/meghnashowtime/", meghnashowtime_handler)
	http.ListenAndServe(":2619", nil)
}
