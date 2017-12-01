package main

import (
	"fmt"

	"github.com/jamesjmurtagh/gohtml/pkg/nodes"
	"github.com/jamesjmurtagh/gohtml/pkg/server"
)

func main() {

	nHTML := nodes.NewNode(nodes.ElementHTML)
	nHead := nHTML.NewChild(nodes.ElementHead)
	nTitle := nHead.NewChild(nodes.ElementTitle)
	nTitle.AddInnerHTML("My GoHTML Website")

	nBody := nHTML.NewChild(nodes.ElementBody)
	nMasterDiv := nBody.NewChild(nodes.ElementDiv)
	nMasterDiv.AddClass("container")
	nMasterDiv.SetAttribute("style", "background-color: #FFA71A; font-weight: bold; text-align: center;")
	nMasterDiv.AddInnerHTML("Hello world!")

	indexFile := nodes.NewHTMLFile("index.html", nHTML)

	fs, err := nodes.NewFilesystem(indexFile)
	if err != nil {
		fmt.Printf("error compiling filesystem: %s\n", err.Error())
		return
	}

	err = server.Host(fs)
	if err != nil {
		fmt.Printf("error hosting web server: %s\n", err.Error())
		return
	}

}
