package main
import (
	"fmt"
	"net/http"
	"github.com/suraj2424/ticketmaster/db"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	db.ConnectDB()
	defer db.Pool.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w ,"Hello Suraj")
	})

	fmt.Println("Backend running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

