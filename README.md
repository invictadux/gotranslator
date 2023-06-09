Gotranslate is a very simple HTML/template package translator.

```go
import "github.com/invictadux/gotranslate"

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
var translator gotranslator.Translator
var helloTemplate *template.Template

func Hello(w http.ResponseWriter, r *http.Request) {
	page := map[string]interface{}{}
	page["Lang"] = "es"
	helloTemplate
}

func main() {
	translator = *gotranslator.New("en", "lang.json")
	helloTemplate = translator.NewTemplate("hello.html","nav.html",...)
	
	http.HandleFunc("/", Hello)

	err := http.ListenAndServe(":8000", nil)
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
        {{template "nav" .}}
        {{uselang "hello" .Lang}}
    </body>
</html>
```

