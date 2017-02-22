package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var addr = flag.String("addr", "127.0.0.1:8080", "addr to listen")
	var content = flag.String("content", ".", "file to server")
	flag.Parse()
	abs, err := filepath.Abs(*content)
	if err != nil {
		log.Fatalln(err)
		return
	}
	http.HandleFunc("/dist/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join("statics", r.URL.Path))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fb, err := ioutil.ReadFile("statics/dist/index.html")
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Write(fb)
		}
	})
	http.HandleFunc("/treenode", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			nodepath := filepath.Join(abs, r.URL.Query().Get("path"))
			relpath, _ := filepath.Rel(abs, nodepath)
			fileStat, err := os.Stat(nodepath)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(err.Error()))
				return
			}
			treenode := map[string]interface{}{
				"name":  fileStat.Name(),
				"isDir": fileStat.IsDir(),
				"path":  relpath,
			}
			if fileStat.IsDir() {
				children := make(map[string]interface{}, 0)
				infos, _ := ioutil.ReadDir(nodepath)
				for _, info := range infos {
					children[info.Name()] = map[string]interface{}{
						"name":  info.Name(),
						"isDir": info.IsDir(),
						"path":  filepath.Join(relpath, info.Name()),
					}
				}
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(err.Error()))
					return
				}
				treenode["children"] = children
			} else {
				fb, _ := ioutil.ReadFile(nodepath)
				treenode["content"] = string(fb)
			}
			treeNodeByte, err := json.Marshal(treenode)
			w.Write([]byte(treeNodeByte))
		}
	})
	http.ListenAndServe(*addr, nil)
	log.Printf("listen at %s", *addr)
}
