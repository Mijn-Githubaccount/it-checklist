package main
import (
	"net/http"
	"Challenge/afdeling"
	"Challenge/inloggen"
	"Challenge/gegevens_weergeven"
)
func main(){
	http.HandleFunc("/", inloggen.Inloggen)
	http.HandleFunc("/accountaanmaken", inloggen.Accountaanmaken)
	http.HandleFunc("/gebruiker_inlog", inloggen.Gebruiker_inlog)
	http.HandleFunc("/gebruiker_aanmelden", inloggen.Gebruiker_aanmelden)
	http.HandleFunc("/hr", afdeling.Hr)
	http.HandleFunc("/sales", afdeling.Sales)
	http.HandleFunc("/marketing", afdeling.Marketing)
	http.HandleFunc("/browse", gegevens_weergeven.Gegevens_weergeven)
	http.ListenAndServe(":8000", nil)
}