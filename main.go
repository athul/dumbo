package main

import (
	"bytes"
	"embed"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

// tempFS acts as the Virtual FS for
// Embedded markdown file
//go:embed template.md
var tempFS embed.FS

func main() {
	title_arg := flag.String("n", time.Now().Format("Monday")+"'s Log", "Create a New Entry to the Journal")
	goto_arg := flag.Bool("g", false, "Go to Journal Folder")
	flag.Parse()
	tempMeta := genMeta(*title_arg)
	if *goto_arg {
		tempMeta.GotoFolder()
		os.Exit(0)
	}
	tempMeta.GenTemplate()
	tempMeta.OpenEditor()
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
// and makes a new folder called .journal in the home dir. It creates a new file
// with the date as a name and fills in with the data from the template
func (t *TemplateMeta) GenTemplate() {
	var b bytes.Buffer
	templ, err := template.ParseFS(tempFS, "*.md")
	if err != nil {
		log.Println("Markdown File Opening Error", err)
	}
	templ.Execute(&b, t)
	home := os.Getenv("HOME")
	if _, err := os.Stat(home + "/.journal/"); err != nil {
		if os.IsExist(err) {
			if err := os.Mkdir(home+"/.journal", 0755); err != nil {
				log.Println("Folder Creating Error", err)
			}
		}
	}

	ioutil.WriteFile(home+"/.journal/"+t.Date+".md", b.Bytes(), 0644)
	log.Println("Journal File Created\n", b.String())
}

// OpenEdior opens Nvim to edit the journal file
func (t *TemplateMeta) OpenEditor() {
	home := os.Getenv("HOME")
	cmd := exec.Command("nvim", home+"/.journal/"+t.Date+".md")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("Cannot Open Neovim", err)
	}
}

func (t *TemplateMeta) GotoFolder() {
	home := os.Getenv("HOME")
	cmd := exec.Command("cd", home+"/.journal/")
	if err := cmd.Run(); err != nil {
		log.Println("Cannot goto to folder", err)
	}
}
