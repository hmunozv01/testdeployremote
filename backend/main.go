package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!
ESTE ES MI AMBIENTE DEVELOP ahora si que si

Cambio en repo develop para despliegue local.

`)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
