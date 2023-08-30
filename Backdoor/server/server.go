package backdoor

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)
func main() {
	addr:="127.0.0.1"+string(GeneratePortNumber())
	http.HandleFunc("/upload",func(wr http.ResponseWriter, r *http.Request ){
		filedata,err:=os.ReadFile("static/upload.html")
		if err != nil{
			log.Fatal(err)
		}
		fmt.Fprintf(wr,string(filedata))
	})
	http.ListenAndServe(addr,nil)
}
func GeneratePortNumber() uint16{
	randsource:=rand.New(rand.NewSource(time.Now().UnixNano()))
	randport:=randsource.Intn(49151)+1024
	return uint16(randport)
}