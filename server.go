// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func main() {
// 	fileServer := http.FileServer(http.Dir("./static"))
// 	http.Handle("/", fileServer)
//     http.HandleFunc("/hello", helloHandler)


//     fmt.Printf("Starting server at port 8080\n")
//     if err := http.ListenAndServe(":8080", nil); err != nil {
//         log.Fatal(err)
//     }
// }

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/hello" {
// 		http.Error(w, "404 not found.", http.StatusNotFound)
// 		return
// 	}

// 	if r.Method != "GET" {
// 		http.Error(w, "Method is not supported.", http.StatusNotFound)
// 		return
// 	} 

// 	fmt.Fprint(w, "Hello!")
// }

// func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		r.ParseForm()
// 		// bodyContent = fmt.Sprintf("POST data received: %v", r.PostForm)
// 	}

// 	tmpl, err := index.ParseFiles("index.html")
// 	if err != nil {
// 		http.Error(w, "Unable to load template", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl.Execute(w, bodyContent)
// }

package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bodyContent := "No POST data received."

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		bodyContent = fmt.Sprintf("POST data received: %v", r.PostForm)
	}

	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html>
	<head>
		<title>GET and POST example</title>
	</head>
	<body>

	<h1>GET and POST example</h1>
	
	<form action="/" method="post">
		<label for="name">Name:</label>
		<input type="text" id="name" name="name"><br><br>
		
		<label for="email">Email:</label>
		<input type="text" id="email" name="email"><br><br>
		
		<input type="submit" value="Submit">
	</form>

	<h2>POST Data Received:</h2>
	<p>%s</p>
	
	</body>
	</html>
	`, bodyContent)

	// tmpl, err := template.ParseFiles("template.html")
	// if err != nil {
	// 	http.Error(w, "Unable to load template", http.StatusInternalServerError)
	// 	return
	// }

	// tmpl.Execute(w, bodyContent)

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
