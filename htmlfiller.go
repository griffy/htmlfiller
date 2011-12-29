package htmlfiller

import (
	"html"
	"strings"
)

func Fill(html_ string, vals map[string]string, errors ...map[string]string) string {
	for name, val := range vals {
		html_ = FillField(html_, name, val)
	}
	if len(errors) > 0 {
		for name, error := range errors[0] {
			html_ = FillField(html_, name + "_error", error)
		}
	}
	return html_
}

func hasName(token html.Token, name string) bool {
	for _, attr := range token.Attr {
		if attr.Key == "name" {
			return attr.Val == name
		} 
	}
	return false
}

func hasValue(token html.Token, val string) bool {
	for _, attr := range token.Attr {
		if attr.Key == "value" {
			return attr.Val == val
		} 
	}
	return false
}

func setValue(token *html.Token, val string) {
	for _, attr := range token.Attr {
		if attr.Key == "value" {
			attr.Val = val
			return
		} 
	}
	// if we made it down here, the attribute "value" does
	// not exist, so we must create it
	token.Attr = append(token.Attr, html.Attribute{"value", val})
}

func setSelected(token *html.Token) {
	token.Attr = append(token.Attr, html.Attribute{"selected", "selected"})
}

func FillField(html_ string, name, val string) (newHtml string) {
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
			if elemName == "span" && hasName(token, name) {
				// the next token that is a TextToken
				// should be filled with the value
				fillNextText = true
			} else if elemName == "textarea" && hasName(token, name) {
				fillNextText = true
			} else if elemName == "select" && hasName(token, name) {
				// we are in the select tag, so we must
				// search for the right <option> tag now
				inSelect = true
			} else if elemName == "option" && inSelect {
				if hasValue(token, val) {
					// this option tag has val we want to set
					// as the default, so make it selected and
					// end our search
					setSelected(&token)
					inSelect = false
				}
			} else if elemName == "input" && hasName(token, name) {
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
			if elemName == "input" && hasName(token, name) {
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