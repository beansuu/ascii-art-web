package main

import (
	"fmt"
	"net/http"
	"ascii-art-web/ascii-art/ascii_art"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	bodyContent := "No POST data received."

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		// bodyContent = fmt.Sprintf("POST data received: %v", r.PostForm)
		text := r.Form.Get("text")
		style := r.Form.Get("banner")
		bodyContent = ascii_art.PrintArt(text,style)
		
	}

	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>ASCII Art Web</title>
	</head>
	<body style="background-color: #888888">
		<form action="/ascii-art" method="POST">
			<fieldset>
				<legend>ASCII-art-web</legend>

				<div>
					<label>Banner</label>
				</div>
				<div>
					<!-- <input name="banner" list="banner-list" type="text">
					<datalist id="banner-list">
						<option value="standard">
						<option value="shadow">
						<option value="thinkertoy">
					</datalist> -->
					<input type="checkbox" name="banner" value="standard">Standard
					<input type="checkbox" name="banner" value="shadow">Shadow
					<input type="checkbox" name="banner" value="thinkertoy">Thinkertoy
				</div>

				<div>
					<label>Input text</label>
				</div>
				<div>
					<input id="add-text" name="text" type="text">
				</div>
			</fieldset>
			<input type="submit" value="Generate">
		</form>
		<p>%s</p>
	</body>
	</html>
	`, bodyContent)


}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}
