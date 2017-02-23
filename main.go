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
	var statics = flag.String("statics", "statics", "statics")
	flag.Parse()
	abs, e := filepath.Abs(*content)
	if e != nil {
		log.Fatalln(e)
		return
	}
	http.HandleFunc("/dist/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(*statics, r.URL.Path))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fb, err := ioutil.ReadFile(filepath.Join(*statics, "/dist/index.html"))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			w.Write(fb)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	http.HandleFunc("/treeleaf", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			defer r.Body.Close()
			var leaf = make(map[string]string, 2)
			err := json.NewDecoder(r.Body).Decode(&leaf)
			if err != nil {
				data, _ := json.Marshal(map[string]interface{}{
					"success": false,
					"err":     err.Error(),
				})
				w.Write(data)
				return
			}
			nodepath := filepath.Join(abs, leaf["path"])
			err = ioutil.WriteFile(nodepath, []byte(leaf["content"]), 0644)
			if err != nil {
				data, _ := json.Marshal(map[string]interface{}{
					"success": false,
					"err":     err.Error(),
				})
				w.Write(data)
				return
			}
			data, _ := json.Marshal(map[string]interface{}{
				"success": true,
			})
			w.Write(data)
		} else {
			w.WriteHeader(http.StatusNotFound)
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
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})
	log.Printf("listen at %s", *addr)
	log.Printf("abs: %s", abs)
	log.Printf("statics: %s", *statics)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		panic(err)
	}
}
