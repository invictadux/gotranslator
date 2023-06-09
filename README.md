Gotranslate is a very simple HTML/template package translator.

```bash
go get github.com/invictadux/gotranslator
```

You should have a json file with all the words and phrased that you want to have translated.

```json
{
    "hello": {
        "en": "Hello",
        "es": "Hola",
        "ru": "Привет",
        "it": "Ciao"
    }
}
```

```go
package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/invictadux/gotranslator"
)

var translator gotranslator.Translator
var helloTemplate *template.Template

func Hello(w http.ResponseWriter, r *http.Request) {
	page := map[string]interface{}{}
	page["Lang"] = "es"
	helloTemplate.Execute(w, page)
}

func main() {
	translator = *gotranslator.New("en", "lang.json")
	helloTemplate = translator.NewTemplate("hello.html")

	http.HandleFunc("/", Hello)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

Translator generate a new template that has the FunMap uselang.
You have to specify id and lang to insert the correct text.

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Hello Page</title>
    </head>
    <body>
        <p>Hello in {{.Lang}} is {{uselang "hello" .Lang}}</p>
    </body>
</html>
```

This is a very simple package don't expect more from it, maybe i will add more features in the future.

