package nodes

import "errors"

/* List of HTML element types. */
const (
	ElementHTML       = "html"
	ElementLink       = "link"
	ElementMeta       = "meta"
	ElementStyle      = "style"
	ElementScript     = "script"
	ElementBody       = "body"
	ElementAddress    = "address"
	ElementArticle    = "article"
	ElementAside      = "aside"
	ElementFooter     = "footer"
	ElementHeader     = "header"
	ElementH1         = "h1"
	ElementH2         = "h2"
	ElementH3         = "h3"
	ElementH4         = "h4"
	ElementH5         = "h5"
	ElementH6         = "h6"
	ElementHGroup     = "hgroup"
	ElementNav        = "nav"
	ElementSection    = "section"
	ElementBlockQuote = "blockquote"
	ElementDd         = "dd"
	ElementDiv        = "div"
	ElementDl         = "dl"
	ElementDt         = "dt"
	ElementFigCaption = "figcaption"
	ElementFigure     = "figure"
	ElementHr         = "hr"
	ElementLi         = "li"
	ElementMain       = "main"
	ElementOl         = "ol"
	ElementP          = "p"
	ElementPre        = "pre"
	ElementUl         = "ul"
	ElementA          = "a"
	ElementAbbr       = "abbr"
	ElementB          = "b"
	ElementBdi        = "bdi"
	ElementBdo        = "bdo"
	ElementBr         = "br"
	ElementCite       = "cite"
	ElementCode       = "code"
	ElementData       = "data"
	ElementDfn        = "dfn"
	ElementEm         = "em"
	ElementI          = "i"
	ElementKbd        = "kbd"
	ElementMark       = "mark"
	ElementQ          = "q"
	ElementRp         = "rp"
	ElementRt         = "rt"
	ElementRuby       = "ruby"
	ElementS          = "s"
	ElementSamp       = "samp"
	ElementSmall      = "small"
	ElementSpan       = "span"
	ElementStrong     = "strong"
	ElementSub        = "sub"
	ElementTime       = "time"
	ElementU          = "u"
	ElementVar        = "var"
	ElementWbr        = "wbr"
	ElementArea       = "area"
	ElementAudio      = "audio"
	ElementImg        = "img"
	ElementMap        = "map"
	ElementTrack      = "track"
	ElementVideo      = "video"
	ElementEmbed      = "embed"
	ElementObject     = "object"
	ElementParam      = "param"
	ElementSource     = "source"
	ElementCanvas     = "canvas"
	ElementNoScript   = "noscript"
	ElementDel        = "del"
	ElementIns        = "ins"
	ElementCaption    = "caption"
	ElementCol        = "col"
	ElementColGroup   = "colgroup"
	ElementTable      = "table"
	ElementTBody      = "tbody"
	ElementTd         = "td"
	ElementTFoot      = "tfoot"
	ElementTh         = "th"
	ElementTHead      = "thead"
	ElementTr         = "tr"
	ElementButton     = "button"
	ElementDataList   = "datalist"
	ElementFieldSet   = "fieldset"
	ElementForm       = "form"
	ElementInput      = "input"
	ElementLabel      = "label"
	ElementLegend     = "legend"
	ElementMeter      = "meter"
	ElementOptGroup   = "optgroup"
	ElementOption     = "option"
	ElementOutput     = "output"
	ElementProgress   = "progress"
	ElementSelect     = "select"
	ElementTextArea   = "textarea"
	ElementDetails    = "details"
	ElementDialog     = "dialog"
	ElementMenu       = "menu"
	ElementMenuItem   = "menuitem"
	ElementSummary    = "summary"
	ElementContent    = "content"
	ElementElement    = "Element"
	ElementShadow     = "shadow"
	ElementSlot       = "slot"
	ElementTemplate   = "template"
)

var (
	errNameConflict = errors.New("cannot contain multiple files of the same name")
)
