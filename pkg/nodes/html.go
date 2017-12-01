package nodes

// Filesystem holds multiple HTMLFile objects.
type Filesystem struct {
	fileNames map[string]v
	Files     []*HTMLFile
}

// HTMLFile is used to hold a tree of HTML nodes, and define other attributes
// which otherwise do not fit inside of an average HTML element.
type HTMLFile struct {
	Name     string
	RootNode *HTMLNode
}

// NewFilesystem creates a new filesystem object.
func NewFilesystem(files ...*HTMLFile) (*Filesystem, error) {

	var err error

	fs := new(Filesystem)
	fs.Files = make([]*HTMLFile, 0)
	fs.fileNames = make(map[string]v)
	for _, f := range files {
		err = fs.AddFile(f)
		if err != nil {
			return nil, err
		}
	}

	return fs, err
}

// AddFile adds a file to the filesystem if the name does not conflict with any
// existing files
func (fs *Filesystem) AddFile(f *HTMLFile) error {

	if _, ok := fs.fileNames[f.Name]; ok {
		return errNameConflict
	}

	fs.fileNames[f.Name] = v{}
	fs.Files = append(fs.Files, f)
	return nil
}

// NewHTMLFile returns an instantiated HTMLFile object.
func NewHTMLFile(name string, root *HTMLNode) *HTMLFile {
	return &HTMLFile{
		Name:     name,
		RootNode: root,
	}
}
