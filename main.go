package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/google/go-github/github"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"golang.org/x/oauth2"
)

func main() {
	mux := http.NewServeMux()
	//for github repo
	gh := GitFiles{
		Owner: "mugund10",
		Repo:  "blog.openwaves.in",
		Dir:   "blog",
		Token: os.Getenv("GITHUB_TOKEN"),
	}

	rr := GitFiles{
		Owner: "mugund10",
		Repo:  "blog.openwaves.in",
		Dir:   "root",
		Token: os.Getenv("GITHUB_TOKEN"),
	}

	
	//for local repo
	//fs := LocalFiles{}

	//slug represents filename
	mux.HandleFunc("GET /blog/{slug}", PostHandler(gh))
	//slug represent folder
	mux.HandleFunc("GET /{slug}/", blogHandler(gh))
	mux.HandleFunc("GET /", RootHandler(rr))

	log.Println("server starting on port 443")
	// if err := http.ListenAndServe("0.0.0.0:80", mux); err != nil {
	// 	log.Fatal(err)
	// }
	if err := http.ListenAndServeTLS("0.0.0.0:443","/etc/letsencrypt/live/blog.openwaves.in/fullchain.pem","/etc/letsencrypt/live/blog.openwaves.in/privkey.pem",mux); err != nil {
		log.Fatal(err)
	}
}

//local 
type LocalFiles struct {
}

type GitFiles struct {
	Owner string
	Repo  string
	Dir   string
	Token string
}
type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
	Author  Author `toml:"author"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type SlugReader interface {
	ReadFile(slug string) (string, error)
	ReadFold(slug string) ([]string, error)
}

type Blog struct {
	Title       string `toml:"title"`
	Description string `toml:"description"`
	Url         string `toml:"url"`
}

type PageData struct {
	Title string
	Blogs []Blog
}

func (fr LocalFiles) ReadFile(slug string) (string, error) {
	filePath := filepath.Join("blog", slug+".md")
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (ff LocalFiles) ReadFold(slug string) ([]string, error) {

	d, err := os.Open(slug)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	files, err := d.Readdir(-1)
	if err != nil {
		return nil, err
	}

	var mdFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			mdFiles = append(mdFiles, file.Name())
		}
	}

	return mdFiles, nil
}

func (fr GitFiles) ReadFile(slug string) (string, error) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: fr.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	fileContent, _, _, err := client.Repositories.GetContents(ctx, fr.Owner, fr.Repo, fmt.Sprintf("%s/%s.md", fr.Dir, slug), nil)
	if err != nil {
		return "", err
	}

	content, err := fileContent.GetContent()
	if err != nil {
		return "", err
	}

	return content, nil
}

func (ff GitFiles) ReadFold(slug string) ([]string, error) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ff.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	_, dirContents, _, err := client.Repositories.GetContents(ctx, ff.Owner, ff.Repo, slug, nil)
	if err != nil {
		return nil, err
	}

	var mdFiles []string
	for _, file := range dirContents {
		if *file.Type == "file" && strings.HasSuffix(*file.Name, ".md") {
			mdFiles = append(mdFiles, *file.Name)
		}
	}

	return mdFiles, nil
}

func PostHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		slug := r.PathValue("slug")
		
		
		postMd, err := sl.ReadFile(slug)
		if err != nil {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		rest, err := frontmatter.Parse(strings.NewReader(postMd), &post)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}

		mdRenderer := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(highlighting.WithStyle("dracula")),
			),
		)

		var buf bytes.Buffer
		//goldmark convert function
		if err := mdRenderer.Convert(rest, &buf); err != nil {
			http.Error(w, "Error Converting Markdown", http.StatusInternalServerError)
			return
		}

		tpl, err := template.ParseFiles("html/post.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		post.Content = template.HTML(buf.String())
		err = tpl.Execute(w, post)
		//fmt.Fprintf(w, postMd)
		//io.Copy(w, &buf)
	}
}
func blogHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var blog []Blog

		slug := r.PathValue("slug")
		log.Println("BlogHandler", slug)

		filesMd, err := sl.ReadFold(slug)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		mdRenderer := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(highlighting.WithStyle("dracula")),
			),
		)

		var buf bytes.Buffer

		for _, filemd := range filesMd {
			var blo Blog
			fi := filemd[:len(filemd)-3]
			log.Println(fi)
			postMarkdown, err := sl.ReadFile(fi)
			if err != nil {
				// TODO: Handle different errors in the future
				http.Error(w, "Post not found", http.StatusNotFound)
				return
			}
			rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &blo)
			if err != nil {
				http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
				return
			}
			blog = append(blog, blo) // Ensure the result is used

			// Convert Markdown to HTML
			if err := mdRenderer.Convert(rest, &buf); err != nil {
				http.Error(w, "Error converting Markdown", http.StatusInternalServerError)
				return
			}

			log.Println(filemd ,"from" ,r.RemoteAddr)
		}

		data := PageData{
			Title: "My Blog",
			Blogs: blog, // Use the accumulated blog data
		}

		tmpl, err := template.ParseFiles("html/blog.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
		}
	}
}

func RootHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		slug := "root"
		
		
		postMd, err := sl.ReadFile(slug)
		if err != nil {
			http.Error(w, "post not found root", http.StatusNotFound)
			return
		}
		rest, err := frontmatter.Parse(strings.NewReader(postMd), &post)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}

		mdRenderer := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(highlighting.WithStyle("dracula")),
			),
		)

		var buf bytes.Buffer
		//goldmark convert function
		if err := mdRenderer.Convert(rest, &buf); err != nil {
			http.Error(w, "Error Converting Markdown", http.StatusInternalServerError)
			return
		}

		tpl, err := template.ParseFiles("html/post.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		post.Content = template.HTML(buf.String())
		err = tpl.Execute(w, post)
		//fmt.Fprintf(w, postMd)
		//io.Copy(w, &buf)
	}
}
