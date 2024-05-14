package main 
 
import (
	"fmt" 
	"net/http"
	"os"
) 

func FileServerHandler(w http.ResponseWriter, req *http.Request){
	request_file := fmt.Sprintf(".%v", req.URL.String())
	_, err := os.Stat(request_file)

	if os.IsNotExist(err) {
		fmt.Fprint(w, "File not found\n")
	} else {
		fs := http.FileServer(http.Dir(".")) 
		fs.ServeHTTP(w, req)
	}
}

func main(){
	PORT := "12013"
	fmt.Println("Launching server...")

	hh := http.HandlerFunc(FileServerHandler)
	http.Handle("/",hh)
	for {
		http.ListenAndServe(fmt.Sprintf(":%v", PORT), hh)
	}
}