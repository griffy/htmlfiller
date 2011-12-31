package htmlfiller

import (
	"html"
	"strings"
)

func Fill(html_ string, vals map[string]string, errors ...map[string]string) string {
	for name, val := range vals {
		html_ = FillElement(html_, name, val)
	}
	if len(errors) > 0 {
		for name, error := range errors[0] {
			html_ = FillElement(html_, name + "_error", error)
		}
	}
	return html_
}

func hasNameMatching(token html.Token, name string) bool {
	for _, attr := range token.Attr {
		if attr.Key == "name" {
			return attr.Val == name
		} 
	}
	return false
}

func hasValueMatching(token html.Token, val string) bool {
	for _, attr := range token.Attr {
		if attr.Key == "value" {
			return attr.Val == val
		} 
	}
	return false
}

func setValue(token *html.Token, val string) {
	for i, attr := range token.Attr {
		if attr.Key == "value" {
            // the attribute already exists, so give it the value
			attr.Val = val
            // and add the modified Attr back into the array
            token.Attr[i] = attr
			return
		} 
	}
	// if we made it down here, the attribute "value" does
	// not exist, so we must create it
	token.Attr = append(token.Attr, html.Attribute{"value", val})
}

func setSelected(token *html.Token) {
    for _, attr := range token.Attr {
        if attr.Key == "selected" {
            if attr.Val == "selected" {
                // it already is selected, so our work is done
                return
            } else {
                // it somehow had the tag, but not the value, so set it
                attr.Val = "selected"
                return
            }
        }
    }
    // the attribute didn't exist, so create it
	token.Attr = append(token.Attr, html.Attribute{"selected", "selected"})
}

func setNotSelected(token *html.Token) {
    for i, attr := range token.Attr {
        if attr.Key == "selected" {
            // remove this attribute
            token.Attr = append(token.Attr[:i], token.Attr[i+1:]...)
            break
        }
    }
}

func FillElement(html_ string, name, val string) (newHtml string) {
	reader := strings.NewReader(html_)
	tokenizer := html.NewTokenizer(reader)
	fillNextText := false
	inSelect := false
	for {
		tokenizer.Next()
		token := tokenizer.Token()
		elemName := token.Data
		if token.Type == html.ErrorToken {
			// finished parsing
			break
		}

		if token.Type == html.StartTagToken {
			if elemName == "span" && hasNameMatching(token, name) {
				// the next token that is a TextToken
				// should be filled with the value
				fillNextText = true
			} else if elemName == "textarea" && hasNameMatching(token, name) {
				fillNextText = true
			} else if elemName == "select" && hasNameMatching(token, name) {
                if inSelect {
                    inSelect = false
                } else {
				    // we are in the select tag, so we must
				    // search for the right <option> tag now
				    inSelect = true
                }
			} else if elemName == "option" && inSelect {
				setNotSelected(&token)
                if hasValueMatching(token, val) {
					// this option tag has val we want to set
					// as the default, so make it selected and
					// end our search
					setSelected(&token)
				}
			} else if elemName == "input" && hasNameMatching(token, name) {
				setValue(&token, val)
			}
		} else if token.Type == html.EndTagToken {
			if elemName == "span" && fillNextText {
				// there was no text token, so manually
				// insert the value
				newHtml += val
				fillNextText = false
			} else if elemName == "textarea" && fillNextText {
				newHtml += val
				fillNextText = false	
			}
		} else if token.Type == html.SelfClosingTagToken {
			if elemName == "input" && hasNameMatching(token, name) {
				setValue(&token, val)
			}
		} else if token.Type == html.TextToken && fillNextText {
			token.Data = val
			fillNextText = false
		}
		newHtml += token.String()
	}
	return newHtml
}
