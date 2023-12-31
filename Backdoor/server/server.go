package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func main() {
	addr:=fmt.Sprintf("127.0.0.1:%d",443)
	http.HandleFunc("/upload",func(w http.ResponseWriter, r *http.Request ){
		if r.Method=="GET"{
			filedata,err:=os.ReadFile("static/upload.html")
			if err != nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,string(filedata))
		}else if r.Method=="POST" {
			commands :=r.FormValue("commands")
			cmdfile,err:=os.OpenFile("commands.txt",os.O_RDWR,0644)
			if err != nil {
				log.Fatal(err)
			}
			cmdfile.Truncate(0)
			_,err= cmdfile.Write([]byte(commands))
			if err != nil{
				log.Fatal(err)
			}
			fmt.Fprint(w,"Successfully updated commands file!")
		} 
	})
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		if r.Method=="GET"{
			commands,err:=os.ReadFile("commands.txt")
			if err !=nil{
				log.Fatal(err)
			}
			w.Header().Set("Content-Type","text/plain")
			fmt.Fprint(w,string(commands))
		}
	})
	fmt.Printf("Serving on http://%s\n",addr)
	http.ListenAndServe(addr,nil)
	
}
