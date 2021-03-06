package html

// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT, except for adding more constants to the list and rerun go generate

// uses github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go

// Hash defines perfect hashes for a predefined list of strings
type Hash uint32

// Unique hash definitions to be used instead of strings
const (
	A                Hash = 0x1     // a
	Abbr             Hash = 0x6e104 // abbr
	About            Hash = 0x5     // about
	Accept           Hash = 0x7106  // accept
	Accept_Charset   Hash = 0x710e  // accept-charset
	Accesskey        Hash = 0x8309  // accesskey
	Acronym          Hash = 0xed07  // acronym
	Action           Hash = 0x2bd06 // action
	Address          Hash = 0x6a407 // address
	Align            Hash = 0x2e05  // align
	Alink            Hash = 0x10405 // alink
	Allowfullscreen  Hash = 0x2350f // allowfullscreen
	Alt              Hash = 0x11d03 // alt
	Annotation       Hash = 0x20d0a // annotation
	AnnotationXml    Hash = 0x20d0d // annotationXml
	Applet           Hash = 0x16a06 // applet
	Area             Hash = 0x39504 // area
	Article          Hash = 0x43007 // article
	Aside            Hash = 0xa405  // aside
	Async            Hash = 0x12905 // async
	Audio            Hash = 0x14705 // audio
	Autocomplete     Hash = 0x1530c // autocomplete
	Autofocus        Hash = 0x15f09 // autofocus
	Autoplay         Hash = 0x17408 // autoplay
	Axis             Hash = 0x17c04 // axis
	B                Hash = 0x101   // b
	Background       Hash = 0x5d0a  // background
	Base             Hash = 0x3cc04 // base
	Basefont         Hash = 0x3cc08 // basefont
	Bdi              Hash = 0xfd03  // bdi
	Bdo              Hash = 0x18503 // bdo
	Bgcolor          Hash = 0x19807 // bgcolor
	Bgsound          Hash = 0x1a707 // bgsound
	Big              Hash = 0x1ba03 // big
	Blink            Hash = 0x1bd05 // blink
	Blockquote       Hash = 0x1c20a // blockquote
	Body             Hash = 0x7f04  // body
	Border           Hash = 0x33b06 // border
	Br               Hash = 0x3402  // br
	Button           Hash = 0x1cc06 // button
	Canvas           Hash = 0xa006  // canvas
	Caption          Hash = 0x28207 // caption
	Center           Hash = 0x4e306 // center
	Challenge        Hash = 0x1ed09 // challenge
	Charset          Hash = 0x7807  // charset
	Checked          Hash = 0x37f07 // checked
	Cite             Hash = 0x12d04 // cite
	Class            Hash = 0x1d705 // class
	Classid          Hash = 0x1d707 // classid
	Clear            Hash = 0x43405 // clear
	Code             Hash = 0x1de04 // code
	Codebase         Hash = 0x3c808 // codebase
	Codetype         Hash = 0x1de08 // codetype
	Col              Hash = 0x19a03 // col
	Colgroup         Hash = 0x1f608 // colgroup
	Color            Hash = 0x19a05 // color
	Cols             Hash = 0x20804 // cols
	Colspan          Hash = 0x20807 // colspan
	Command          Hash = 0x21a07 // command
	Compact          Hash = 0x22107 // compact
	Content          Hash = 0x66e07 // content
	Contenteditable  Hash = 0x66e0f // contenteditable
	Contextmenu      Hash = 0x3df0b // contextmenu
	Controls         Hash = 0x22c08 // controls
	Coords           Hash = 0x25306 // coords
	Crossorigin      Hash = 0x25e0b // crossorigin
	Data             Hash = 0x4c104 // data
	Datalist         Hash = 0x4c108 // datalist
	Datatype         Hash = 0x4d808 // datatype
	Datetime         Hash = 0x2ed08 // datetime
	Dd               Hash = 0x31902 // dd
	Declare          Hash = 0xa707  // declare
	Default          Hash = 0x4a07  // default
	DefaultChecked   Hash = 0x50c0e // defaultChecked
	DefaultMuted     Hash = 0x56f0c // defaultMuted
	DefaultSelected  Hash = 0x4a0f  // defaultSelected
	Defer            Hash = 0x5805  // defer
	Del              Hash = 0x9303  // del
	Desc             Hash = 0x9d04  // desc
	Details          Hash = 0xb307  // details
	Dfn              Hash = 0xcf03  // dfn
	Dialog           Hash = 0xfe06  // dialog
	Dir              Hash = 0xdb03  // dir
	Dirname          Hash = 0xdb07  // dirname
	Disabled         Hash = 0x10b08 // disabled
	Div              Hash = 0x11203 // div
	Dl               Hash = 0x13302 // dl
	Download         Hash = 0x49508 // download
	Draggable        Hash = 0x1ad09 // draggable
	Dropzone         Hash = 0x43b08 // dropzone
	Dt               Hash = 0x60802 // dt
	Em               Hash = 0xcb02  // em
	Embed            Hash = 0xcb05  // embed
	Enabled          Hash = 0x8d07  // enabled
	Enctype          Hash = 0x2d207 // enctype
	Face             Hash = 0x4e104 // face
	Fieldset         Hash = 0x62a08 // fieldset
	Figcaption       Hash = 0x27f0a // figcaption
	Figure           Hash = 0x29306 // figure
	Font             Hash = 0x3d004 // font
	Footer           Hash = 0x12006 // footer
	For              Hash = 0x29f03 // for
	ForeignObject    Hash = 0x29f0d // foreignObject
	Foreignobject    Hash = 0x2ac0d // foreignobject
	Form             Hash = 0x2b904 // form
	Formaction       Hash = 0x2b90a // formaction
	Formenctype      Hash = 0x2ce0b // formenctype
	Formmethod       Hash = 0x2d90a // formmethod
	Formnovalidate   Hash = 0x2e30e // formnovalidate
	Formtarget       Hash = 0x2f80a // formtarget
	Frame            Hash = 0xd305  // frame
	Frameborder      Hash = 0x3360b // frameborder
	Frameset         Hash = 0xd308  // frameset
	H1               Hash = 0x19602 // h1
	H2               Hash = 0x32702 // h2
	H3               Hash = 0x34c02 // h3
	H4               Hash = 0x38e02 // h4
	H5               Hash = 0x60a02 // h5
	H6               Hash = 0x30202 // h6
	Head             Hash = 0x37404 // head
	Header           Hash = 0x37406 // header
	Headers          Hash = 0x37407 // headers
	Height           Hash = 0x30406 // height
	Hgroup           Hash = 0x30c06 // hgroup
	Hidden           Hash = 0x31706 // hidden
	High             Hash = 0x32404 // high
	Hr               Hash = 0x13b02 // hr
	Href             Hash = 0x13b04 // href
	Hreflang         Hash = 0x13b08 // hreflang
	Html             Hash = 0x30804 // html
	Http_Equiv       Hash = 0x3290a // http-equiv
	I                Hash = 0xa01   // i
	Icon             Hash = 0x66d04 // icon
	Id               Hash = 0xa602  // id
	Iframe           Hash = 0x33506 // iframe
	Image            Hash = 0x34105 // image
	Img              Hash = 0x34603 // img
	Inert            Hash = 0x55a05 // inert
	Inlist           Hash = 0x26706 // inlist
	Input            Hash = 0x48205 // input
	Ins              Hash = 0x1b03  // ins
	Isindex          Hash = 0x17e07 // isindex
	Ismap            Hash = 0x34e05 // ismap
	Itemid           Hash = 0x12e06 // itemid
	Itemprop         Hash = 0x59208 // itemprop
	Itemref          Hash = 0x62407 // itemref
	Itemscope        Hash = 0x35809 // itemscope
	Itemtype         Hash = 0x36208 // itemtype
	Kbd              Hash = 0xfc03  // kbd
	Keygen           Hash = 0x8906  // keygen
	Keytype          Hash = 0x68d07 // keytype
	Kind             Hash = 0x10804 // kind
	Label            Hash = 0x9505  // label
	Lang             Hash = 0x13f04 // lang
	Language         Hash = 0x13f08 // language
	Legend           Hash = 0x1b406 // legend
	Li               Hash = 0x2f02  // li
	Link             Hash = 0x10504 // link
	List             Hash = 0x26904 // list
	Listing          Hash = 0x4c507 // listing
	Longdesc         Hash = 0x9908  // longdesc
	Loop             Hash = 0x13404 // loop
	Low              Hash = 0x23703 // low
	Main             Hash = 0x1904  // main
	Malignmark       Hash = 0xf30a  // malignmark
	Manifest         Hash = 0x68308 // manifest
	Map              Hash = 0x16903 // map
	Mark             Hash = 0xf904  // mark
	Marquee          Hash = 0x36a07 // marquee
	Math             Hash = 0x37104 // math
	Max              Hash = 0x38603 // max
	Maxlength        Hash = 0x38609 // maxlength
	Media            Hash = 0xe005  // media
	Mediagroup       Hash = 0xe00a  // mediagroup
	Menu             Hash = 0x3e604 // menu
	Meta             Hash = 0x4d404 // meta
	Meter            Hash = 0x2f305 // meter
	Method           Hash = 0x2dd06 // method
	Mglyph           Hash = 0x34706 // mglyph
	Mi               Hash = 0x6b02  // mi
	Min              Hash = 0x6b03  // min
	Mn               Hash = 0x2e602 // mn
	Mo               Hash = 0x4ff02 // mo
	Ms               Hash = 0x35b02 // ms
	Mtext            Hash = 0x39005 // mtext
	Multiple         Hash = 0x39e08 // multiple
	Muted            Hash = 0x3a605 // muted
	Name             Hash = 0xde04  // name
	Nav              Hash = 0x2b03  // nav
	Nobr             Hash = 0x3204  // nobr
	Noembed          Hash = 0xc907  // noembed
	Noframes         Hash = 0xd108  // noframes
	Nohref           Hash = 0x13906 // nohref
	Noresize         Hash = 0x24308 // noresize
	Noscript         Hash = 0x31c08 // noscript
	Noshade          Hash = 0x50707 // noshade
	Novalidate       Hash = 0x2e70a // novalidate
	Nowrap           Hash = 0x3ab06 // nowrap
	Object           Hash = 0x2b306 // object
	Ol               Hash = 0x19b02 // ol
	Onabort          Hash = 0x1d007 // onabort
	Onafterprint     Hash = 0x2870c // onafterprint
	Onbeforeprint    Hash = 0x2c10d // onbeforeprint
	Onbeforeunload   Hash = 0x6980e // onbeforeunload
	Onblur           Hash = 0xbe06  // onblur
	Oncancel         Hash = 0x14b08 // oncancel
	Oncanplay        Hash = 0x18709 // oncanplay
	Oncanplaythrough Hash = 0x18710 // oncanplaythrough
	Onchange         Hash = 0x45108 // onchange
	Onclick          Hash = 0x6cb07 // onclick
	Onclose          Hash = 0x3b707 // onclose
	Oncontextmenu    Hash = 0x3dd0d // oncontextmenu
	Oncuechange      Hash = 0x3ea0b // oncuechange
	Ondblclick       Hash = 0x3f50a // ondblclick
	Ondrag           Hash = 0x3ff06 // ondrag
	Ondragend        Hash = 0x3ff09 // ondragend
	Ondragenter      Hash = 0x4080b // ondragenter
	Ondragleave      Hash = 0x4130b // ondragleave
	Ondragover       Hash = 0x41e0a // ondragover
	Ondragstart      Hash = 0x4280b // ondragstart
	Ondrop           Hash = 0x43906 // ondrop
	Ondurationchange Hash = 0x44910 // ondurationchange
	Onemptied        Hash = 0x44009 // onemptied
	Onended          Hash = 0x45907 // onended
	Onerror          Hash = 0x46007 // onerror
	Onfocus          Hash = 0x46707 // onfocus
	Onhashchange     Hash = 0x4740c // onhashchange
	Oninput          Hash = 0x48007 // oninput
	Oninvalid        Hash = 0x48709 // oninvalid
	Onkeydown        Hash = 0x49009 // onkeydown
	Onkeypress       Hash = 0x49d0a // onkeypress
	Onkeyup          Hash = 0x4ac07 // onkeyup
	Onload           Hash = 0x4b906 // onload
	Onloadeddata     Hash = 0x4b90c // onloadeddata
	Onloadedmetadata Hash = 0x4cc10 // onloadedmetadata
	Onloadstart      Hash = 0x4e90b // onloadstart
	Onmessage        Hash = 0x4f409 // onmessage
	Onmousedown      Hash = 0x4fd0b // onmousedown
	Onmousemove      Hash = 0x51a0b // onmousemove
	Onmouseout       Hash = 0x5250a // onmouseout
	Onmouseover      Hash = 0x52f0b // onmouseover
	Onmouseup        Hash = 0x53a09 // onmouseup
	Onmousewheel     Hash = 0x5480c // onmousewheel
	Onoffline        Hash = 0x55409 // onoffline
	Ononline         Hash = 0x55f08 // ononline
	Onpagehide       Hash = 0x5670a // onpagehide
	Onpageshow       Hash = 0x57b0a // onpageshow
	Onpause          Hash = 0x58707 // onpause
	Onplay           Hash = 0x59e06 // onplay
	Onplaying        Hash = 0x59e09 // onplaying
	Onpopstate       Hash = 0x5a70a // onpopstate
	Onprogress       Hash = 0x5b10a // onprogress
	Onratechange     Hash = 0x5c00c // onratechange
	Onreset          Hash = 0x5cc07 // onreset
	Onresize         Hash = 0x5d308 // onresize
	Onscroll         Hash = 0x5db08 // onscroll
	Onseeked         Hash = 0x5e608 // onseeked
	Onseeking        Hash = 0x5ee09 // onseeking
	Onselect         Hash = 0x5f708 // onselect
	Onshow           Hash = 0x60106 // onshow
	Onstalled        Hash = 0x60c09 // onstalled
	Onstorage        Hash = 0x61509 // onstorage
	Onsubmit         Hash = 0x61e08 // onsubmit
	Onsuspend        Hash = 0x63a09 // onsuspend
	Ontimeupdate     Hash = 0x3d10c // ontimeupdate
	Onunload         Hash = 0x64308 // onunload
	Onvolumechange   Hash = 0x64b0e // onvolumechange
	Onwaiting        Hash = 0x65909 // onwaiting
	Open             Hash = 0x13604 // open
	Optgroup         Hash = 0x66208 // optgroup
	Optimum          Hash = 0x67d07 // optimum
	Option           Hash = 0x69406 // option
	Output           Hash = 0x206   // output
	P                Hash = 0x501   // p
	Param            Hash = 0x1505  // param
	Pattern          Hash = 0x2507  // pattern
	Pauseonexit      Hash = 0x5890b // pauseonexit
	Picture          Hash = 0x3e07  // picture
	Ping             Hash = 0xe904  // ping
	Placeholder      Hash = 0x1fd0b // placeholder
	Plaintext        Hash = 0x26e09 // plaintext
	Poster           Hash = 0x31106 // poster
	Pre              Hash = 0x35203 // pre
	Prefix           Hash = 0x35206 // prefix
	Preload          Hash = 0x3b007 // preload
	Profile          Hash = 0x4b207 // profile
	Progress         Hash = 0x5b308 // progress
	Prompt           Hash = 0x54206 // prompt
	Property         Hash = 0x59608 // property
	Public           Hash = 0x66906 // public
	Q                Hash = 0xae01  // q
	Radiogroup       Hash = 0x350a  // radiogroup
	Rb               Hash = 0x5c02  // rb
	Readonly         Hash = 0x39608 // readonly
	Rel              Hash = 0x3b103 // rel
	Required         Hash = 0xac08  // required
	Resource         Hash = 0x29708 // resource
	Rev              Hash = 0x4303  // rev
	Reversed         Hash = 0x4308  // reversed
	Rows             Hash = 0xc304  // rows
	Rowspan          Hash = 0xc307  // rowspan
	Rp               Hash = 0x28d02 // rp
	Rt               Hash = 0x1d502 // rt
	Rtc              Hash = 0x1d503 // rtc
	Ruby             Hash = 0x12504 // ruby
	Rules            Hash = 0x19e05 // rules
	S                Hash = 0x1201  // s
	Samp             Hash = 0x1204  // samp
	Sandbox          Hash = 0x1d07  // sandbox
	Scope            Hash = 0x35c05 // scope
	Scoped           Hash = 0x35c06 // scoped
	Script           Hash = 0x31e06 // script
	Scrolling        Hash = 0x5dd09 // scrolling
	Seamless         Hash = 0x3bc08 // seamless
	Section          Hash = 0xb907  // section
	Select           Hash = 0x5f906 // select
	Selected         Hash = 0x5f908 // selected
	Shape            Hash = 0x1a205 // shape
	Size             Hash = 0x24704 // size
	Sizes            Hash = 0x24705 // sizes
	Small            Hash = 0x23305 // small
	Sortable         Hash = 0x24b08 // sortable
	Source           Hash = 0x29906 // source
	Spacer           Hash = 0x25806 // spacer
	Span             Hash = 0xc604  // span
	Spellcheck       Hash = 0x37a0a // spellcheck
	Src              Hash = 0x3c303 // src
	Srcdoc           Hash = 0x3c306 // srcdoc
	Srclang          Hash = 0x46d07 // srclang
	Srcset           Hash = 0x4a606 // srcset
	Start            Hash = 0x42e05 // start
	Step             Hash = 0x26b04 // step
	Strike           Hash = 0x68906 // strike
	Strong           Hash = 0x5ba06 // strong
	Style            Hash = 0x6aa05 // style
	Sub              Hash = 0x62003 // sub
	Summary          Hash = 0x6af07 // summary
	Sup              Hash = 0x6b603 // sup
	Svg              Hash = 0x6b903 // svg
	System           Hash = 0x6bc06 // system
	Tabindex         Hash = 0x708   // tabindex
	Table            Hash = 0x24e05 // table
	Target           Hash = 0x2fc06 // target
	Tbody            Hash = 0x7e05  // tbody
	Td               Hash = 0xda02  // td
	Template         Hash = 0x6bf08 // template
	Text             Hash = 0x27304 // text
	Textarea         Hash = 0x39108 // textarea
	Tfoot            Hash = 0x11f05 // tfoot
	Th               Hash = 0x19002 // th
	Thead            Hash = 0x37305 // thead
	Time             Hash = 0x2f104 // time
	Title            Hash = 0x16f05 // title
	Tr               Hash = 0x22702 // tr
	Track            Hash = 0x22705 // track
	Translate        Hash = 0x27609 // translate
	Truespeed        Hash = 0x63109 // truespeed
	Tt               Hash = 0x2702  // tt
	Type             Hash = 0x11904 // type
	Typemustmatch    Hash = 0x1e20d // typemustmatch
	Typeof           Hash = 0x4dc06 // typeof
	U                Hash = 0x301   // u
	Ul               Hash = 0x4e02  // ul
	Undeterminate    Hash = 0x640d  // undeterminate
	Usemap           Hash = 0x16606 // usemap
	Valign           Hash = 0x2d06  // valign
	Value            Hash = 0x11405 // value
	Valuetype        Hash = 0x11409 // valuetype
	Var              Hash = 0x33203 // var
	Video            Hash = 0x6c705 // video
	Visible          Hash = 0x6d207 // visible
	Vlink            Hash = 0x6d905 // vlink
	Vocab            Hash = 0x6de05 // vocab
	Wbr              Hash = 0x58403 // wbr
	Width            Hash = 0x60605 // width
	Wrap             Hash = 0x3ad04 // wrap
	Xmlns            Hash = 0xe05   // xmlns
	Xmp              Hash = 0x2303  // xmp
)

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// ToHash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := uint32(_Hash_hash0)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				goto NEXT
			}
		}
		return i
	}
NEXT:
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) {
		t := _Hash_text[i>>8 : i>>8+i&0xff]
		for i := 0; i < len(s); i++ {
			if t[i] != s[i] {
				return 0
			}
		}
		return i
	}
	return 0
}

const _Hash_hash0 = 0x9acb0442
const _Hash_maxLen = 16
const _Hash_text = "aboutputabindexmlnsamparamainsandboxmpatternavalignobradiogr" +
	"oupictureversedefaultSelectedeferbackgroundeterminateaccept-" +
	"charsetbodyaccesskeygenabledelabelongdescanvasideclarequired" +
	"etailsectionblurowspanoembedfnoframesetdirnamediagroupingacr" +
	"onymalignmarkbdialogalinkindisabledivaluetypealtfooterubyasy" +
	"ncitemidloopenohreflanguageaudioncancelautocompleteautofocus" +
	"emappletitleautoplayaxisindexbdoncanplaythrough1bgcolorulesh" +
	"apebgsoundraggablegendbigblinkblockquotebuttonabortclassidco" +
	"detypemustmatchallengecolgrouplaceholdercolspannotationXmlco" +
	"mmandcompactrackcontrolsmallowfullscreenoresizesortablecoord" +
	"spacercrossoriginlisteplaintextranslatefigcaptionafterprintf" +
	"iguresourceforeignObjectforeignobjectformactionbeforeprintfo" +
	"rmenctypeformmethodformnovalidatetimeterformtargeth6heightml" +
	"hgrouposterhiddenoscripthigh2http-equivariframeborderimageim" +
	"glyph3ismaprefixitemscopeditemtypemarqueematheaderspellcheck" +
	"edmaxlength4mtextareadonlymultiplemutednowrapreloadoncloseam" +
	"lessrcdocodebasefontimeupdateoncontextmenuoncuechangeondblcl" +
	"ickondragendondragenterondragleaveondragoverondragstarticlea" +
	"rondropzonemptiedondurationchangeonendedonerroronfocusrclang" +
	"onhashchangeoninputoninvalidonkeydownloadonkeypressrcsetonke" +
	"yuprofileonloadeddatalistingonloadedmetadatatypeofacenteronl" +
	"oadstartonmessageonmousedownoshadefaultCheckedonmousemoveonm" +
	"ouseoutonmouseoveronmouseupromptonmousewheelonofflinertononl" +
	"ineonpagehidefaultMutedonpageshowbronpauseonexitempropertyon" +
	"playingonpopstateonprogresstrongonratechangeonresetonresizeo" +
	"nscrollingonseekedonseekingonselectedonshowidth5onstalledons" +
	"torageonsubmitemrefieldsetruespeedonsuspendonunloadonvolumec" +
	"hangeonwaitingoptgroupublicontenteditableoptimumanifestrikey" +
	"typeoptionbeforeunloaddresstylesummarysupsvgsystemplatevideo" +
	"nclickvisiblevlinkvocabbr"

var _Hash_table = [1 << 10]Hash{
	0x4:   0x13b02, // hr
	0x8:   0x3e07,  // picture
	0x9:   0x48007, // oninput
	0xb:   0x11405, // value
	0xf:   0x708,   // tabindex
	0x12:  0x2870c, // onafterprint
	0x18:  0x1e20d, // typemustmatch
	0x1a:  0x13302, // dl
	0x1b:  0x67d07, // optimum
	0x1e:  0x38e02, // h4
	0x21:  0x5ee09, // onseeking
	0x22:  0x11f05, // tfoot
	0x23:  0x66e0f, // contenteditable
	0x24:  0x6bf08, // template
	0x29:  0x63a09, // onsuspend
	0x2b:  0x51a0b, // onmousemove
	0x30:  0x2bd06, // action
	0x33:  0xd305,  // frame
	0x35:  0x19e05, // rules
	0x38:  0x18503, // bdo
	0x39:  0x3ab06, // nowrap
	0x3e:  0x62a08, // fieldset
	0x47:  0xf30a,  // malignmark
	0x49:  0x44009, // onemptied
	0x4c:  0x46d07, // srclang
	0x4e:  0x3c306, // srcdoc
	0x4f:  0xa602,  // id
	0x50:  0x64308, // onunload
	0x51:  0x30c06, // hgroup
	0x55:  0x19a05, // color
	0x56:  0x35c05, // scope
	0x59:  0x640d,  // undeterminate
	0x5b:  0x37406, // header
	0x5c:  0xa405,  // aside
	0x5d:  0x2f80a, // formtarget
	0x60:  0xe05,   // xmlns
	0x61:  0x19b02, // ol
	0x63:  0x9d04,  // desc
	0x65:  0x26e09, // plaintext
	0x66:  0x3ad04, // wrap
	0x67:  0x17c04, // axis
	0x68:  0x19a03, // col
	0x69:  0x2d06,  // valign
	0x70:  0x65909, // onwaiting
	0x71:  0x31706, // hidden
	0x75:  0x41e0a, // ondragover
	0x78:  0x29f0d, // foreignObject
	0x7a:  0x32404, // high
	0x7b:  0xf904,  // mark
	0x88:  0x68906, // strike
	0x8f:  0x13b08, // hreflang
	0x91:  0xd308,  // frameset
	0x92:  0x4c507, // listing
	0x93:  0x37f07, // checked
	0x94:  0x34c02, // h3
	0x95:  0x5805,  // defer
	0x96:  0x6c705, // video
	0x97:  0xae01,  // q
	0x9a:  0x25806, // spacer
	0x9c:  0x55f08, // ononline
	0x9d:  0x64b0e, // onvolumechange
	0x9f:  0x4b207, // profile
	0xa0:  0x34e05, // ismap
	0xa3:  0x2f305, // meter
	0xa8:  0x1de04, // code
	0xab:  0x18710, // oncanplaythrough
	0xae:  0x62407, // itemref
	0xaf:  0x59e09, // onplaying
	0xb5:  0x43b08, // dropzone
	0xb9:  0x2f02,  // li
	0xbb:  0x15f09, // autofocus
	0xbd:  0x5480c, // onmousewheel
	0xc1:  0x26904, // list
	0xc2:  0x2702,  // tt
	0xc7:  0x4e306, // center
	0xc9:  0x45907, // onended
	0xcb:  0x35203, // pre
	0xcc:  0x7106,  // accept
	0xcf:  0x4e90b, // onloadstart
	0xd1:  0x36208, // itemtype
	0xd3:  0x3d004, // font
	0xd4:  0x1a707, // bgsound
	0xd5:  0x5ba06, // strong
	0xd6:  0x4ff02, // mo
	0xd7:  0x1de08, // codetype
	0xdb:  0x4280b, // ondragstart
	0xdf:  0x1a205, // shape
	0xe3:  0x2f104, // time
	0xe9:  0x27304, // text
	0xea:  0x6b903, // svg
	0xf2:  0xc907,  // noembed
	0xf3:  0x4dc06, // typeof
	0xf5:  0x60c09, // onstalled
	0xf6:  0xa006,  // canvas
	0xf8:  0x9505,  // label
	0xf9:  0x3b103, // rel
	0xfb:  0x4c104, // data
	0xfd:  0x1204,  // samp
	0x101: 0x5c00c, // onratechange
	0x103: 0x46007, // onerror
	0x105: 0x3dd0d, // oncontextmenu
	0x106: 0x16a06, // applet
	0x108: 0x350a,  // radiogroup
	0x109: 0xde04,  // name
	0x10a: 0x1cc06, // button
	0x10b: 0x43405, // clear
	0x10e: 0x6af07, // summary
	0x10f: 0x4d404, // meta
	0x110: 0x54206, // prompt
	0x113: 0x53a09, // onmouseup
	0x116: 0x5e608, // onseeked
	0x11a: 0xe904,  // ping
	0x11c: 0x35809, // itemscope
	0x11e: 0x14b08, // oncancel
	0x11f: 0x2dd06, // method
	0x120: 0x20d0d, // annotationXml
	0x123: 0x8309,  // accesskey
	0x124: 0x6d207, // visible
	0x127: 0xc304,  // rows
	0x135: 0x11203, // div
	0x136: 0x3d10c, // ontimeupdate
	0x137: 0x59608, // property
	0x139: 0x3ff06, // ondrag
	0x13a: 0xcf03,  // dfn
	0x13e: 0x6aa05, // style
	0x13f: 0x1530c, // autocomplete
	0x141: 0x3204,  // nobr
	0x142: 0x4f409, // onmessage
	0x144: 0x25e0b, // crossorigin
	0x148: 0x61509, // onstorage
	0x149: 0x34603, // img
	0x14a: 0x5cc07, // onreset
	0x14b: 0x2303,  // xmp
	0x14e: 0x4c108, // datalist
	0x153: 0x61e08, // onsubmit
	0x155: 0x12006, // footer
	0x15f: 0x2e70a, // novalidate
	0x162: 0x43906, // ondrop
	0x166: 0x39005, // mtext
	0x168: 0x24705, // sizes
	0x16c: 0x28207, // caption
	0x16e: 0x16f05, // title
	0x173: 0x1ed09, // challenge
	0x176: 0x24b08, // sortable
	0x178: 0x23703, // low
	0x17a: 0x12504, // ruby
	0x17b: 0x9303,  // del
	0x17c: 0x1d707, // classid
	0x17d: 0xfc03,  // kbd
	0x17f: 0x2ed08, // datetime
	0x181: 0x68d07, // keytype
	0x182: 0xc604,  // span
	0x183: 0x21a07, // command
	0x18b: 0x1b406, // legend
	0x18c: 0xe005,  // media
	0x18d: 0x3c808, // codebase
	0x198: 0x31106, // poster
	0x199: 0x2ac0d, // foreignobject
	0x19d: 0x6de05, // vocab
	0x1a0: 0x28d02, // rp
	0x1a4: 0x1d502, // rt
	0x1a8: 0x4308,  // reversed
	0x1aa: 0x13604, // open
	0x1ab: 0x6bc06, // system
	0x1ac: 0x37404, // head
	0x1ad: 0x10405, // alink
	0x1af: 0x33203, // var
	0x1b0: 0xb307,  // details
	0x1b1: 0x60a02, // h5
	0x1b3: 0xda02,  // td
	0x1b4: 0xa707,  // declare
	0x1ba: 0x10b08, // disabled
	0x1be: 0xac08,  // required
	0x1c3: 0x6d905, // vlink
	0x1c4: 0x52f0b, // onmouseover
	0x1c5: 0x3290a, // http-equiv
	0x1cc: 0x14705, // audio
	0x1d1: 0x12d04, // cite
	0x1d5: 0xe00a,  // mediagroup
	0x1d6: 0x3a605, // muted
	0x1da: 0x42e05, // start
	0x1de: 0x19002, // th
	0x1df: 0x17408, // autoplay
	0x1e3: 0x206,   // output
	0x1e4: 0x59208, // itemprop
	0x1e5: 0x3e604, // menu
	0x1eb: 0x2b03,  // nav
	0x1ec: 0x6b603, // sup
	0x1ed: 0x46707, // onfocus
	0x1ee: 0x7807,  // charset
	0x1ef: 0x29f03, // for
	0x1f1: 0x13404, // loop
	0x1f3: 0x24e05, // table
	0x1f5: 0x3f50a, // ondblclick
	0x1f6: 0x5c02,  // rb
	0x1f8: 0x33b06, // border
	0x1fb: 0x27609, // translate
	0x200: 0x39108, // textarea
	0x208: 0x2507,  // pattern
	0x210: 0x1bd05, // blink
	0x212: 0x1d705, // class
	0x219: 0x4b90c, // onloadeddata
	0x21a: 0x60106, // onshow
	0x21b: 0x6cb07, // onclick
	0x21d: 0x1d503, // rtc
	0x225: 0x3360b, // frameborder
	0x22a: 0x6a407, // address
	0x22c: 0x39608, // readonly
	0x22f: 0x12905, // async
	0x233: 0x710e,  // accept-charset
	0x238: 0x43007, // article
	0x23b: 0x5f708, // onselect
	0x23f: 0xfd03,  // bdi
	0x241: 0x11d03, // alt
	0x242: 0x55409, // onoffline
	0x246: 0x56f0c, // defaultMuted
	0x247: 0x34105, // image
	0x249: 0x31e06, // script
	0x24d: 0x24308, // noresize
	0x252: 0x60802, // dt
	0x253: 0x2c10d, // onbeforeprint
	0x255: 0x3bc08, // seamless
	0x256: 0x3402,  // br
	0x257: 0x1d007, // onabort
	0x25a: 0x66208, // optgroup
	0x260: 0x101,   // b
	0x262: 0x5db08, // onscroll
	0x264: 0x13906, // nohref
	0x266: 0x7e05,  // tbody
	0x269: 0x2e602, // mn
	0x26c: 0x37305, // thead
	0x270: 0x4a606, // srcset
	0x271: 0x63109, // truespeed
	0x273: 0xed07,  // acronym
	0x27e: 0x13f04, // lang
	0x281: 0x20804, // cols
	0x285: 0x29906, // source
	0x28a: 0x35b02, // ms
	0x28b: 0x1904,  // main
	0x28c: 0xcb02,  // em
	0x28f: 0x66d04, // icon
	0x292: 0x49508, // download
	0x293: 0x11409, // valuetype
	0x295: 0x301,   // u
	0x296: 0x2e05,  // align
	0x297: 0x1201,  // s
	0x299: 0x2b306, // object
	0x29a: 0x59e06, // onplay
	0x29b: 0x4e104, // face
	0x29d: 0x5dd09, // scrolling
	0x29e: 0x3ff09, // ondragend
	0x2a0: 0x3c303, // src
	0x2a2: 0x36a07, // marquee
	0x2a6: 0x7f04,  // body
	0x2a9: 0x3df0b, // contextmenu
	0x2ac: 0x22c08, // controls
	0x2b0: 0x23305, // small
	0x2b6: 0x2350f, // allowfullscreen
	0x2b9: 0x3b007, // preload
	0x2bb: 0x66e07, // content
	0x2bf: 0x30804, // html
	0x2c1: 0x49d0a, // onkeypress
	0x2c2: 0xdb07,  // dirname
	0x2c6: 0x35206, // prefix
	0x2cd: 0x18709, // oncanplay
	0x2d1: 0x6e104, // abbr
	0x2d3: 0x60605, // width
	0x2d6: 0x6b03,  // min
	0x2d8: 0x37104, // math
	0x2da: 0x48709, // oninvalid
	0x2dd: 0x2d90a, // formmethod
	0x2de: 0x50707, // noshade
	0x2df: 0x58403, // wbr
	0x2e1: 0x35c06, // scoped
	0x2e2: 0x5250a, // onmouseout
	0x2e3: 0x2fc06, // target
	0x2e4: 0x5670a, // onpagehide
	0x2e6: 0x24704, // size
	0x2e9: 0x8906,  // keygen
	0x2ea: 0x29708, // resource
	0x2ec: 0x16903, // map
	0x2ee: 0x5b308, // progress
	0x2f0: 0x26706, // inlist
	0x2f2: 0x3ea0b, // oncuechange
	0x2f7: 0x69406, // option
	0x2f8: 0x37407, // headers
	0x2fb: 0x31c08, // noscript
	0x2fd: 0x34706, // mglyph
	0x301: 0x16606, // usemap
	0x303: 0x68308, // manifest
	0x30a: 0x17e07, // isindex
	0x30b: 0x1f608, // colgroup
	0x312: 0x5d0a,  // background
	0x314: 0x10804, // kind
	0x315: 0x49009, // onkeydown
	0x316: 0x2b90a, // formaction
	0x319: 0x1,     // a
	0x31b: 0x5,     // about
	0x31e: 0x5b10a, // onprogress
	0x31f: 0x44910, // ondurationchange
	0x321: 0x3cc04, // base
	0x32d: 0x2e30e, // formnovalidate
	0x331: 0x22705, // track
	0x33a: 0x5f906, // select
	0x33b: 0x66906, // public
	0x33e: 0x55a05, // inert
	0x340: 0x26b04, // step
	0x342: 0x33506, // iframe
	0x344: 0x32702, // h2
	0x346: 0x4fd0b, // onmousedown
	0x347: 0x1fd0b, // placeholder
	0x34b: 0xdb03,  // dir
	0x34f: 0x13b04, // href
	0x350: 0x22702, // tr
	0x351: 0x1c20a, // blockquote
	0x358: 0x20807, // colspan
	0x35d: 0x50c0e, // defaultChecked
	0x361: 0x1ad09, // draggable
	0x369: 0xbe06,  // onblur
	0x36b: 0x5890b, // pauseonexit
	0x36f: 0x8d07,  // enabled
	0x371: 0x30406, // height
	0x373: 0x5d308, // onresize
	0x375: 0x4d808, // datatype
	0x376: 0x57b0a, // onpageshow
	0x378: 0x1b03,  // ins
	0x379: 0x4080b, // ondragenter
	0x37a: 0x2ce0b, // formenctype
	0x37f: 0x4ac07, // onkeyup
	0x381: 0x3b707, // onclose
	0x382: 0x31902, // dd
	0x384: 0x38603, // max
	0x386: 0x3cc08, // basefont
	0x38a: 0x27f0a, // figcaption
	0x38b: 0x20d0a, // annotation
	0x38c: 0x25306, // coords
	0x38f: 0xcb05,  // embed
	0x392: 0x11904, // type
	0x393: 0x1505,  // param
	0x394: 0x30202, // h6
	0x397: 0x501,   // p
	0x398: 0x19807, // bgcolor
	0x39b: 0x4303,  // rev
	0x39c: 0x12e06, // itemid
	0x39d: 0x1d07,  // sandbox
	0x39e: 0x29306, // figure
	0x3a2: 0xfe06,  // dialog
	0x3a3: 0x5a70a, // onpopstate
	0x3a4: 0x1ba03, // big
	0x3a7: 0x39504, // area
	0x3af: 0x6980e, // onbeforeunload
	0x3b1: 0xa01,   // i
	0x3b4: 0x37a0a, // spellcheck
	0x3b5: 0x4e02,  // ul
	0x3bd: 0xb907,  // section
	0x3bf: 0x45108, // onchange
	0x3c2: 0x13f08, // language
	0x3c4: 0x4130b, // ondragleave
	0x3c6: 0x2d207, // enctype
	0x3c7: 0x22107, // compact
	0x3c8: 0x38609, // maxlength
	0x3c9: 0x5f908, // selected
	0x3cd: 0x9908,  // longdesc
	0x3d3: 0x58707, // onpause
	0x3d9: 0xc307,  // rowspan
	0x3de: 0x4a0f,  // defaultSelected
	0x3df: 0x4cc10, // onloadedmetadata
	0x3e2: 0x10504, // link
	0x3e3: 0xd108,  // noframes
	0x3e6: 0x48205, // input
	0x3ec: 0x62003, // sub
	0x3ee: 0x39e08, // multiple
	0x3ef: 0x4a07,  // default
	0x3f4: 0x2b904, // form
	0x3fb: 0x4b906, // onload
	0x3fc: 0x6b02,  // mi
	0x3fd: 0x19602, // h1
	0x3ff: 0x4740c, // onhashchange
}
