package main
import(
	"fmt"
	"log"
	"net/http"
	"time"
)
func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8000",nil))
}
func handler(w http.ResponseWriter, r *http.Request){
	t:=time.Now().UTC()
	fmt.Fprintf(w, "Current time is: %s\n", t)
}