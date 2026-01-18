package additional

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type myHandlerRu struct{}

// func (h *myHandlerRu) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Привет мир"))
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }
// func aboutHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "This is about page.")
// }

// func main() {
// 	http.HandleFunc("/", helloHandler)
// 	http.HandleFunc("/about", aboutHandler)
// 	http.Handle("/ru/", &myHandlerRu{})
// 	err := http.ListenAndServe(":8080", nil)
// 	if err != nil {
// 		fmt.Println("ListenAndServe error", err)
// 		return
// 	}
// }
