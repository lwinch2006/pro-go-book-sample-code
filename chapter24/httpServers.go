package main

import (
	"chapter24/data"
	"chapter24/httpServers"
	"chapter24/utils"
	"html/template"
	"net/http"
	"strconv"
)

func HttpServers1() {
	utils.Printfln("HttpServers1()")

	for _, p := range data.Products {
		utils.Printfln("Product: %v, Category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}
}

func HttpServers2() {
	utils.Printfln("HttpServers2()")

	err := http.ListenAndServe(":5002", httpServers.StringHandler{Message: "Hello World"})
	if err != nil {
		utils.Printfln("Http Server creation error: %v", err.Error())
	}
}

func HttpServers3() {
	utils.Printfln("HttpServers3()")

	err := http.ListenAndServe(":5003", httpServers.StringHandler3{Message: "Hello World"})
	if err != nil {
		utils.Printfln("Http Server creation error: %v", err.Error())
	}
}

func HttpServers4() {
	utils.Printfln("HttpServers4()")

	err := http.ListenAndServe(":5004", httpServers.StringHandler4{Message: "Hello World"})
	if err != nil {
		utils.Printfln("Http Server creation error: %v", err.Error())
	}
}

func HttpServers5() {
	utils.Printfln("HttpServers5()")

	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	http.Handle("/message", httpServers.StringHandler{Message: "Hello World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())

	err := http.ListenAndServe(":5005", nil)
	if err != nil {
		utils.Printfln("Http Server creation error: %v", err.Error())
	}
}

func HttpServers6() {
	utils.Printfln("HttpServers6()")

	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	http.Handle("/message", httpServers.StringHandler{Message: "Hello World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())

	go func() {
		err := http.ListenAndServeTLS(":5556", "httpServers/cert/localhost-2023-06-06-073805.cer", "httpServers/cert/localhost-2023-06-06-073805.pkey", nil)
		if err != nil {
			utils.Printfln("HTTPS Server creation error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5006", nil)
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers7() {
	utils.Printfln("HttpServers7()")

	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	http.Handle("/message", httpServers.StringHandler7{Message: "Hello World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())

	go func() {
		err := http.ListenAndServeTLS(":5557", "httpServers/cert/localhost-2023-06-06-073805.cer", "httpServers/cert/localhost-2023-06-06-073805.pkey", nil)
		if err != nil {
			utils.Printfln("HTTPS Server creation error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5007", http.HandlerFunc(httpServers.RedirectToHttps))
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers8() {
	utils.Printfln("HttpServers8()")

	http.Handle("/", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	http.Handle("/message", httpServers.StringHandler8{Message: "Hello World"})
	http.Handle("/favicon.ico", http.NotFoundHandler())

	fsHandler := http.FileServer(http.Dir("./httpServers/wwwroot"))
	http.Handle("/files/", http.StripPrefix("/files", fsHandler))

	go func() {
		err := http.ListenAndServeTLS(":5558", "httpServers/cert/localhost-2023-06-06-073805.cer", "httpServers/cert/localhost-2023-06-06-073805.pkey", nil)
		if err != nil {
			utils.Printfln("HTTPS Server creation error: %v", err.Error())
		}
	}()

	err := http.ListenAndServe(":5008", http.HandlerFunc(httpServers.RedirectToHttps8))
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers9() {
	utils.Printfln("HttpServers9()")

	var err error
	httpServers.HtmlTemplates = template.New("all")
	httpServers.HtmlTemplates.Funcs(map[string]interface{}{
		"intVal": strconv.Atoi,
	})

	httpServers.HtmlTemplates, err = httpServers.HtmlTemplates.ParseGlob("httpServers/templates/*.html")
	if err == nil {
		http.Handle("/templates/", http.StripPrefix("/templates/", http.HandlerFunc(httpServers.HandleHtmlTemplate)))
	} else {
		panic(err)
	}

	fsHandler := http.FileServer(http.Dir("./httpServers/wwwroot"))
	http.Handle("/", fsHandler)

	err = http.ListenAndServe(":5009", nil)
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers10() {
	utils.Printfln("HttpServers10()")
	http.Handle("/json", http.HandlerFunc(httpServers.JsonHandler))

	err := http.ListenAndServe(":5010", nil)
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers11() {
	utils.Printfln("HttpServers11()")

	var err error
	httpServers.HtmlTemplates = template.New("all")
	httpServers.HtmlTemplates.Funcs(map[string]interface{}{
		"intVal": strconv.Atoi,
	})

	httpServers.HtmlTemplates, err = httpServers.HtmlTemplates.ParseGlob("httpServers/templates/*.html")
	if err == nil {
		http.Handle("/templates/", http.StripPrefix("/templates/", http.HandlerFunc(httpServers.HandleHtmlTemplate)))
		http.Handle("/forms/edit", http.HandlerFunc(httpServers.ProcessFormData))
	} else {
		panic(err)
	}

	fsHandler := http.FileServer(http.Dir("./httpServers/wwwroot"))
	http.Handle("/", fsHandler)

	err = http.ListenAndServe(":5011", nil)
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers12() {
	utils.Printfln("HttpServers12()")

	var err error
	httpServers.HtmlTemplates = template.New("all")
	httpServers.HtmlTemplates.Funcs(map[string]interface{}{
		"intVal": strconv.Atoi,
	})

	httpServers.HtmlTemplates, err = httpServers.HtmlTemplates.ParseGlob("httpServers/templates/*.html")
	if err == nil {
		http.Handle("/templates/", http.StripPrefix("/templates/", http.HandlerFunc(httpServers.HandleHtmlTemplate)))
		http.Handle("/forms/edit", http.HandlerFunc(httpServers.ProcessFormData))
		http.Handle("/forms/singleupload", http.HandlerFunc(httpServers.UploadSingleFileHandler))
		http.Handle("/forms/multiupload", http.HandlerFunc(httpServers.UploadMultipleFileHandler))
	} else {
		panic(err)
	}

	fsHandler := http.FileServer(http.Dir("./httpServers/wwwroot"))
	http.Handle("/", fsHandler)

	err = http.ListenAndServe(":5012", nil)
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}

func HttpServers13() {
	utils.Printfln("HttpServers13()")
	http.Handle("/cookies", http.HandlerFunc(httpServers.GetAndSetCookies))

	err := http.ListenAndServe(":5013", nil)
	if err != nil {
		utils.Printfln("HTTP Server creation error: %v", err.Error())
	}
}
