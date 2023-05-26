package main

import (
	"chapter23/constants"
	"chapter23/data"
	"chapter23/models"
	"chapter23/utils"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	textTemplate "text/template"
)

func WorkWithHtml1() {
	utils.Printfln("WorkWithHtml1()")

	for _, p := range data.Products {
		utils.Printfln("Product: %v, price: %.2f", p.Name, p.Price)
	}
}

func WorkWithHtml2() {
	utils.Printfln("WorkWithHtml2()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template1.html")); err == nil {
		utils.Printfln("Template name: %v", t.Name())
		t.Execute(os.Stdout, *data.Milk)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml3() {
	utils.Printfln("WorkWithHtml3()")

	if t, err := template.ParseGlob(filepath.Join(constants.HtmlTemplatesRootDir, "*")); err == nil {
		utils.Printfln("Template name: %v", t.Name())
		t.Execute(os.Stdout, *data.Milk)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml4() {
	utils.Printfln("WorkWithHtml4()")

	t1, err1 := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template1.html"))
	t2, err2 := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template2.html"))

	if err1 == nil && err2 == nil {
		utils.Printfln("Template name: %v", t1.Name())
		t1.Execute(os.Stdout, *data.Milk)

		fmt.Println()
		utils.Printfln("Template name: %v", t2.Name())
		t2.Execute(os.Stdout, *data.Milk)

	} else {
		utils.Printfln("Load template errors: %v %v", err1.Error(), err2.Error())
	}
}

func WorkWithHtml5() {
	utils.Printfln("WorkWithHtml5()")

	if t, err := template.ParseGlob(filepath.Join(constants.HtmlTemplatesRootDir, "*")); err == nil {
		t.ExecuteTemplate(os.Stdout, "template1.html", *data.Milk)
		fmt.Println()
		t.ExecuteTemplate(os.Stdout, "template2.html", *data.Milk)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func workWithHtml6_1(t *template.Template) error {
	return t.Execute(os.Stdout, data.Milk)
}

func WorkWithHtml6() {
	utils.Printfln("WorkWithHtml6()")

	if allTemplates, err := template.ParseGlob(filepath.Join(constants.HtmlTemplatesRootDir, "*")); err == nil {
		t := allTemplates.Lookup("template2.html")
		workWithHtml6_1(t)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml7() {
	utils.Printfln("WorkWithHtml7()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template3.html")); err == nil {
		utils.Printfln("Template name: %v", t.Name())
		utils.Printfln("Tax: %v", data.Milk.ApplyTax())
		t.Execute(os.Stdout, data.Milk)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml8() {
	utils.Printfln("WorkWithHtml8()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template4.html")); err == nil {
		utils.Printfln("Template name: %v", t.Name())
		utils.Printfln("Tax: %v", data.Milk.ApplyTax())
		t.Execute(os.Stdout, data.Milk)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml9() {
	utils.Printfln("WorkWithHtml9()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template5.html")); err == nil {
		t.Execute(os.Stdout, data.Milk)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml10() {
	utils.Printfln("WorkWithHtml10()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template6.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml11() {
	utils.Printfln("WorkWithHtml11()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template7.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml12() {
	utils.Printfln("WorkWithHtml12()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template8.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml13() {
	utils.Printfln("WorkWithHtml13()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template9.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml14() {
	utils.Printfln("WorkWithHtml14()")

	if t, err := template.ParseFiles(
		filepath.Join(constants.HtmlTemplatesRootDir, "template10.html"),
		filepath.Join(constants.HtmlTemplatesRootDir, "template10-1.html")); err == nil {
		t.ExecuteTemplate(os.Stdout, "mainTemplate", data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func workWithHtml15_1(products []*models.Product) (categories []template.HTML) {
	categoriesMap := make(map[string]string)

	for _, p := range products {
		if categoriesMap[p.Category] == "" {
			categoriesMap[p.Category] = p.Category
			categories = append(categories, template.HTML(fmt.Sprintf("<b>%v</b>", p.Category)))
		}
	}

	return
}

func WorkWithHtml15() {
	utils.Printfln("WorkWithHtml15()")

	t := template.New("template11.html")
	t.Funcs(map[string]interface{}{
		"getProductsCategories": workWithHtml15_1,
	})

	if t, err := t.ParseGlob(filepath.Join(constants.HtmlTemplatesRootDir, "template11.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func workWithHtml16_1(products []*models.Product) (categories []string) {
	categoriesMap := make(map[string]string)

	for _, p := range products {
		if categoriesMap[p.Category] == "" {
			categoriesMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}

	return
}

func WorkWithHtml16() {
	utils.Printfln("WorkWithHtml16()")

	t := template.New("template12.html")
	t.Funcs(map[string]interface{}{
		"getProductsCategories": workWithHtml16_1,
		"toLower":               strings.ToLower,
	})

	if t, err := t.ParseGlob(filepath.Join(constants.HtmlTemplatesRootDir, "template12.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml17() {
	utils.Printfln("WorkWithHtml17()")

	if t, err := template.ParseFiles(filepath.Join(constants.HtmlTemplatesRootDir, "template13.html")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}

func WorkWithHtml18() {
	utils.Printfln("WorkWithHtml18()")

	if t, err := textTemplate.ParseFiles(filepath.Join(constants.TextTemplatesRootDir, "template14.txt")); err == nil {
		t.Execute(os.Stdout, data.Products)
	} else {
		utils.Printfln("Load template error: %v", err.Error())
	}
}
