package server

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func primaryHandler(w http.ResponseWriter, r *http.Request) {

	str, err := url.QueryUnescape(r.RequestURI)
	if err != nil {
		fmt.Println(err)
		return
	}

	if str == "/" {
		str = "index.html"
	}

	target := resolveFilepath(str)
	target = crossPlatformParse(target)
	fmt.Printf("Serving '%s'.\n", target)

	f, err := os.Open(target)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func crossPlatformParse(str string) string {

	var out string
	str = strings.Replace(str, "/", "\\", -1)
	arr := strings.Split(str, "\\")
	for _, s := range arr {
		out = path.Join(out, s)
	}

	return out
}

func resolveFilepath(str string) string {
	return path.Join(rootDir, str)
}
