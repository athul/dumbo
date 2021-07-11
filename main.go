package main

import (
	"bytes"
	"embed"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

// tempFS acts as the Virtual FS for
// Embedded markdown file
//go:embed template.md
var tempFS embed.FS

func main() {
	title_arg := flag.String("n", time.Now().Format("Monday")+"'s Log", "Create a New Entry to the Journal")
	flag.Parse()
	tempMeta := genMeta(*title_arg)
	tempMeta.GenTemplate()
}

// TemplateMeta holds the MetaData to be prefilled in to Template.md
type TemplateMeta struct {
	Title string
	Date  string
	Time  string
}

// genMeta generates the Metadata for the Template File
func genMeta(title string) TemplateMeta {
	return TemplateMeta{
		Title: title,
		Date:  time.Now().Format("Monday 2-01-2006"),
		Time:  time.Now().Format("3:4:5 AM"),
	}
}

// GenTemplate generates the Template from template.md
// and makes a new folder called Journal. It creates a new file
// with the date as a name and fills in with the data from the template
func (t *TemplateMeta) GenTemplate() {
	var b bytes.Buffer
	templ, err := template.ParseFS(tempFS, "*.md")
	if err != nil {
		log.Println("Markdown File Opening Error", err)
	}
	templ.Execute(&b, t)
	if _, err := os.Stat("journal"); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir("journal", 0755); err != nil {
				log.Println("Folder Creating Error", err)
			}
		}
	}

	ioutil.WriteFile("journal/"+t.Date+".md", b.Bytes(), 0644)
	log.Println("\n", b.String())
}
