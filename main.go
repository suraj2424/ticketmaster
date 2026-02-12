package main
import (
	"fmt"
	"net/http"
	"github.com/suraj2424/db"
	"github.com/joho/godotenv"
)

func main(){
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Hello  Suraj")
	})

	fmt.Println("Server starting on http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

