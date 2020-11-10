package main // this is the part we want run when we start our application

// import packages (contain prebuilt functionss)
import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	// writing html like this is not a good idea as it
	// has security risks
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc) // acts as a basic router for us
	http.ListenAndServe(":3000", nil) // start the server and listen
	// on port :3000. passing nil
	// uses defualt servmucks, but
	// would accept different one here
}
