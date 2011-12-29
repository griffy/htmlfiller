package main

import (
	"fmt"
	"github.com/griffy/htmlfiller"
)

func main() {
	html := `<form action="">
			 <span name="fname_error"></span>
			 First: <input type="text" name="fname" /><br />
			 <span name="lname_error"></span>
			 Last: <input type="text" name="lname" /><br />
			 <select name="cars">
			 <option value="volvo">Volvo</option>
			 <option value="saab">Saab</option>
			 <option value="fiat">Fiat</option>
			 <option value="audi">Audi</option>
			 </select>
			 </form>`
	defaultVals := make(map[string]string)
	defaultVals["fname"] = "Joel"
	defaultVals["cars"] = "saab"
	errors := make(map[string]string)
	errors["fname"] = "What a shitty name"
	fmt.Println(html)
	html = htmlfiller.Fill(html, defaultVals, errors)
	fmt.Println(html)
}