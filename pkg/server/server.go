package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
	"github.com/jamesjmurtagh/gohtml/pkg/nodes"
)

var rootDir = ""

// Host takes a nodes.filesystem type object, parses and extracts it to a tmp
// directory, and hosts it.
func Host(fs *nodes.Filesystem) error {

	var err error

	rootFS, err := ioutil.TempDir("", "gohtml")
	if err != nil {
		return err
	}
	defer func() {
		go os.RemoveAll(rootFS)
	}()
	rootDir = rootFS

	for _, file := range fs.Files {

		var f *os.File
		f, err = os.Create(path.Join(rootFS, file.Name))
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.Write(nodes.Parse(file.RootNode))
		if err != nil {
			return err
		}
		f.Close()

	}

	r := mux.NewRouter()
	r.HandleFunc("/{path:.*}", primaryHandler)
	http.Handle("/", r)

	fmt.Println("Launching webserver on port 8080.")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}
