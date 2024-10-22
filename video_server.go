package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

type FileInfo struct {
	Name string
}

type IndexPage struct {
	Videos []FileInfo
}

// ServeFile gestisce le richieste di streaming video
func ServeFile(w http.ResponseWriter, r *http.Request, directory string) {
	file := filepath.Join(directory, r.URL.Path[len("/stream/"):])
	http.ServeFile(w, r, file)
}

// Index gestisce la pagina principale che mostra la lista dei video
func Index(w http.ResponseWriter, r *http.Request, directory string) {
	var videos []FileInfo

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".mp4" {
			videos = append(videos, FileInfo{Name: info.Name()})
		}
		return nil
	})

	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	page := IndexPage{Videos: videos}

	tmpl := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Video Streaming</title>
		</head>
		<body>
			<h1>Lista di Video MP4</h1>
			<ul>
				{{range .Videos}}
					<li><a href="/stream/{{.Name}}">{{.Name}}</a></li>
				{{end}}
			</ul>
		</body>
		</html>
	`

	t := template.Must(template.New("index").Parse(tmpl))
	t.Execute(w, page)
}

func main() {
	// Parsing command line arguments
	port := flag.String("port", "8080", "Porta su cui il server ascolterà")
	directory := flag.String("dir", "./videos", "Directory contenente i file MP4")
	flag.Parse()

	// Handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Index(w, r, *directory)
	})
	http.HandleFunc("/stream/", func(w http.ResponseWriter, r *http.Request) {
		ServeFile(w, r, *directory)
	})

	fmt.Printf("Server avviato su http://localhost:%s\n", *port)
	fmt.Println("Per configurare la porta e la directory, esegui il comando:")
	fmt.Printf("go run video_server.go -port <PORTA> -dir <DIRECTORY>\n")
	fmt.Printf("Esempio: go run video_server.go -port 8080 -dir ./videos\n")

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Println("Errore nel server:", err)
	}
}
