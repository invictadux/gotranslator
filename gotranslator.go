package gotranslator

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

type Translator struct {
	DefaultLang string
	Data        map[string]map[string]string
}

func New(langCode string, langFile string) *Translator {
	t := Translator{DefaultLang: langCode}
	t.LoadFile(langFile)
	return &t
}

func (t *Translator) NewTemplate(files ...string) *template.Template {
	segments := strings.Split(files[0], "/")
	fileName := segments[len(segments)-1]
	return template.Must(template.New(fileName).Funcs(t.DefaultFunctions()).ParseFiles(files...))
}

func (t *Translator) DefaultFunctions() template.FuncMap {
	return template.FuncMap{
		"uselang": t.Translate,
	}
}

func (t *Translator) LoadFile(langFile string) {
	content, err := ioutil.ReadFile(langFile)

	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var data map[string]map[string]string
	err = json.Unmarshal(content, &data)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	t.Data = data
}

func (t *Translator) Translate(id string, langCode string) string {
	return t.Data[id][langCode]
}
