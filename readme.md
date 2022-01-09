golang web udemy


jika tidak menggunakan create template cache, maka seperti ini 


```go
parsedTemplate, _ := template.ParseFiles("./template/" + tmpl)
	err = parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template :", err)
		return
	}
```