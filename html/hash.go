// generated by hasher -type=Hash -file=hash.go; DO NOT EDIT
package html

// github.com/tdewolff/hasher
//go:generate hasher -type=Hash -file=hash.go
type Hash uint32

// String returns the hash' name.
func (i Hash) String() string {
	start := uint32(i >> 8)
	n := uint32(i & 0xff)
	if start+n > uint32(len(_Hash_text)) {
		return ""
	}
	return _Hash_text[start : start+n]
}

// Hash returns the hash whose name is s. It returns zero if there is no
// such hash. It is case sensitive.
func ToHash(s []byte) Hash {
	if len(s) == 0 || len(s) > _Hash_maxLen {
		return 0
	}
	h := _Hash_fnv(s)
	if i := _Hash_table[h&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	if i := _Hash_table[(h>>16)&uint32(len(_Hash_table)-1)]; int(i&0xff) == len(s) && _Hash_match(_Hash_string(i), s) {
		return i
	}
	return 0
}

// _Hash_fnv computes the FNV hash with an arbitrary starting value h.
func _Hash_fnv(s []byte) uint32 {
	h := uint32(_Hash_hash0)
	for i := range s {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func _Hash_match(s string, t []byte) bool {
	for i, c := range t {
		if s[i] != c {
			return false
		}
	}
	return true
}

func _Hash_string(i Hash) string {
	return _Hash_text[i>>8 : i>>8+i&0xff]
}

const _Hash_hash0 = 0x5334b67c
const _Hash_maxLen = 16
const _Hash_text = "abbradiogrouparamainavalignobrbackgroundeterminateaccept-cha" +
	"rsetbodyaccesskeygenabledefaultSelectedeferowspanoembedelabe" +
	"longdescanvasideclarequiredetailsampatternoframesetdfnohrefl" +
	"anguageacronymalignmarkbdialogalinkindirnamediagroupingaltfo" +
	"oterubyasyncitemidisabledivaluetypeaudioncancelooptgrouplace" +
	"holderulesandboxmlnsectionblurautocompleteautofocusemappleti" +
	"tleautoplayaxisindexmplaintextrackbdoncanplaythrough1bgcolor" +
	"bgsoundlbigblinkblockquotebuttonabortclassidraggablegendcode" +
	"typemustmatchallengecolgroupostercolspannotationXmlcommandco" +
	"mpactranslatecontrolshapecoordsmallowfullscreenoresizesortab" +
	"lecrossoriginsourcefieldsetruespeedfigcaptionafterprintfigur" +
	"eversedforeignObjectforeignobjectformactionbeforeprintformen" +
	"ctypeformmethodformnovalidatetimeterformtargeth6heightmlhgro" +
	"upreloadhiddenoscripthigh2http-equivariframeborderimageimgly" +
	"ph3ismaprofileitemscopeditemtypematheaderspacermaxlength4mte" +
	"xtareadonlymultiplemutedoncloseamlesspellcheckedoncontextmen" +
	"uoncuechangeondblclickondragendondragenterondragleaveondrago" +
	"verondragstarticlearondropzonemptiedondurationchangeonendedo" +
	"nerroronfocusrcdocodebasefontimeupdateonhashchangeoninputoni" +
	"nvalidonkeydownloadonkeypressrclangonkeyuponloadeddatalistin" +
	"gonloadedmetadatabindexonloadstartonmessageonmousedownoshade" +
	"faultCheckedonmousemoveonmouseoutputonmouseoveronmouseuponmo" +
	"usewheelonofflinertononlineonpagehidefaultMutedonpageshowbro" +
	"npauseonexitempropenowrapublicontenteditableonplayingonpopst" +
	"ateonprogresstrikeytypeonratechangeonresetonresizeonscrollin" +
	"gonseekedonseekingonselectedonshowidth5onstalledonstorageons" +
	"ubmitemrefacenteronsuspendonunloadonvolumechangeonwaitingopt" +
	"imumanifestepromptoptionbeforeunloaddresstrongstylesummarysu" +
	"psvgsystemarqueevideonclickvisiblevlink"

const (
	A                Hash = 0x1
	Abbr             Hash = 0x4
	Accept           Hash = 0x3206
	Accept_Charset   Hash = 0x320e
	Accesskey        Hash = 0x4409
	Acronym          Hash = 0xbb07
	Action           Hash = 0x2b906
	Address          Hash = 0x67607
	Align            Hash = 0x1605
	Alink            Hash = 0xd205
	Allowfullscreen  Hash = 0x23c0f
	Alt              Hash = 0xeb03
	Annotation       Hash = 0x2060a
	AnnotationXml    Hash = 0x2060d
	Applet           Hash = 0x16106
	Area             Hash = 0x38604
	Article          Hash = 0x40707
	Aside            Hash = 0x8305
	Async            Hash = 0xf705
	Audio            Hash = 0x11305
	Autocomplete     Hash = 0x14a0c
	Autofocus        Hash = 0x15609
	Autoplay         Hash = 0x16b08
	Axis             Hash = 0x17304
	B                Hash = 0x101
	Background       Hash = 0x1e0a
	Base             Hash = 0x44d04
	Basefont         Hash = 0x44d08
	Bdi              Hash = 0xcb03
	Bdo              Hash = 0x18a03
	Bgcolor          Hash = 0x19d07
	Bgsound          Hash = 0x1a407
	Big              Hash = 0x1ac03
	Blink            Hash = 0x1af05
	Blockquote       Hash = 0x1b40a
	Body             Hash = 0x4004
	Border           Hash = 0x33806
	Br               Hash = 0x202
	Button           Hash = 0x1be06
	Canvas           Hash = 0x7f06
	Caption          Hash = 0x27e07
	Center           Hash = 0x62306
	Challenge        Hash = 0x1eb09
	Charset          Hash = 0x3907
	Checked          Hash = 0x3ad07
	Cite             Hash = 0xfb04
	Class            Hash = 0x1c905
	Classid          Hash = 0x1c907
	Clear            Hash = 0x40b05
	Code             Hash = 0x1dc04
	Codebase         Hash = 0x44908
	Codetype         Hash = 0x1dc08
	Col              Hash = 0x19f03
	Colgroup         Hash = 0x1f408
	Color            Hash = 0x19f05
	Cols             Hash = 0x20104
	Colspan          Hash = 0x20107
	Command          Hash = 0x21307
	Compact          Hash = 0x21a07
	Content          Hash = 0x58107
	Contenteditable  Hash = 0x5810f
	Contextmenu      Hash = 0x3b60b
	Controls         Hash = 0x22908
	Coords           Hash = 0x23506
	Crossorigin      Hash = 0x25a0b
	Data             Hash = 0x4a604
	Datalist         Hash = 0x4a608
	Datetime         Hash = 0x2e908
	Dd               Hash = 0x31602
	Declare          Hash = 0x8607
	Default          Hash = 0x5407
	DefaultChecked   Hash = 0x4ea0e
	DefaultMuted     Hash = 0x54b0c
	DefaultSelected  Hash = 0x540f
	Defer            Hash = 0x6205
	Del              Hash = 0x7203
	Desc             Hash = 0x7c04
	Details          Hash = 0x9207
	Dfn              Hash = 0xab03
	Dialog           Hash = 0xcc06
	Dir              Hash = 0xd903
	Dirname          Hash = 0xd907
	Disabled         Hash = 0x10108
	Div              Hash = 0x10803
	Dl               Hash = 0x1aa02
	Download         Hash = 0x47f08
	Draggable        Hash = 0x1cf09
	Dropzone         Hash = 0x41208
	Dt               Hash = 0x5ff02
	Em               Hash = 0x6e02
	Embed            Hash = 0x6e05
	Enabled          Hash = 0x4e07
	Enctype          Hash = 0x2ce07
	Face             Hash = 0x62104
	Fieldset         Hash = 0x26b08
	Figcaption       Hash = 0x27b0a
	Figure           Hash = 0x28f06
	Font             Hash = 0x45104
	Footer           Hash = 0xee06
	For              Hash = 0x29b03
	ForeignObject    Hash = 0x29b0d
	Foreignobject    Hash = 0x2a80d
	Form             Hash = 0x2b504
	Formaction       Hash = 0x2b50a
	Formenctype      Hash = 0x2ca0b
	Formmethod       Hash = 0x2d50a
	Formnovalidate   Hash = 0x2df0e
	Formtarget       Hash = 0x2f40a
	Frame            Hash = 0xa305
	Frameborder      Hash = 0x3330b
	Frameset         Hash = 0xa308
	H1               Hash = 0x19b02
	H2               Hash = 0x32402
	H3               Hash = 0x34902
	H4               Hash = 0x37f02
	H5               Hash = 0x60102
	H6               Hash = 0x2fe02
	Head             Hash = 0x36b04
	Header           Hash = 0x36b06
	Headers          Hash = 0x36b07
	Height           Hash = 0x30006
	Hgroup           Hash = 0x30806
	Hidden           Hash = 0x31406
	High             Hash = 0x32104
	Hr               Hash = 0xaf02
	Href             Hash = 0xaf04
	Hreflang         Hash = 0xaf08
	Html             Hash = 0x30404
	Http_Equiv       Hash = 0x3260a
	I                Hash = 0x601
	Icon             Hash = 0x58004
	Id               Hash = 0x8502
	Iframe           Hash = 0x33206
	Image            Hash = 0x33e05
	Img              Hash = 0x34303
	Inert            Hash = 0x53605
	Input            Hash = 0x46c05
	Ins              Hash = 0x26303
	Isindex          Hash = 0x17507
	Ismap            Hash = 0x34b05
	Itemid           Hash = 0xfc06
	Itemprop         Hash = 0x56e08
	Itemref          Hash = 0x61b07
	Itemscope        Hash = 0x35609
	Itemtype         Hash = 0x36008
	Kbd              Hash = 0xca03
	Keygen           Hash = 0x4a06
	Keytype          Hash = 0x5b007
	Kind             Hash = 0xd604
	Label            Hash = 0x7405
	Lang             Hash = 0xb304
	Language         Hash = 0xb308
	Legend           Hash = 0x1d606
	Li               Hash = 0x1702
	Link             Hash = 0xd304
	List             Hash = 0x4aa04
	Listing          Hash = 0x4aa07
	Longdesc         Hash = 0x7808
	Loop             Hash = 0x11e04
	Low              Hash = 0x23e03
	Main             Hash = 0x1004
	Malignmark       Hash = 0xc10a
	Manifest         Hash = 0x65708
	Map              Hash = 0x16003
	Mark             Hash = 0xc704
	Marquee          Hash = 0x69907
	Math             Hash = 0x36804
	Max              Hash = 0x37703
	Maxlength        Hash = 0x37709
	Media            Hash = 0xde05
	Mediagroup       Hash = 0xde0a
	Menu             Hash = 0x3bd04
	Meta             Hash = 0x4b904
	Meter            Hash = 0x2ef05
	Method           Hash = 0x2d906
	Mglyph           Hash = 0x34406
	Mi               Hash = 0x2c02
	Min              Hash = 0x2c03
	Mn               Hash = 0x2e202
	Mo               Hash = 0x4dd02
	Ms               Hash = 0x35902
	Mtext            Hash = 0x38105
	Multiple         Hash = 0x38f08
	Muted            Hash = 0x39705
	Name             Hash = 0xdc04
	Nav              Hash = 0x1303
	Nobr             Hash = 0x1a04
	Noembed          Hash = 0x6c07
	Noframes         Hash = 0xa108
	Nohref           Hash = 0xad06
	Noresize         Hash = 0x24a08
	Noscript         Hash = 0x31908
	Noshade          Hash = 0x4e507
	Novalidate       Hash = 0x2e30a
	Nowrap           Hash = 0x57706
	Object           Hash = 0x2af06
	Ol               Hash = 0x12d02
	Onabort          Hash = 0x1c207
	Onafterprint     Hash = 0x2830c
	Onbeforeprint    Hash = 0x2bd0d
	Onbeforeunload   Hash = 0x66a0e
	Onblur           Hash = 0x14406
	Oncancel         Hash = 0x11708
	Oncanplay        Hash = 0x18c09
	Oncanplaythrough Hash = 0x18c10
	Onchange         Hash = 0x42808
	Onclick          Hash = 0x6a407
	Onclose          Hash = 0x39c07
	Oncontextmenu    Hash = 0x3b40d
	Oncuechange      Hash = 0x3c10b
	Ondblclick       Hash = 0x3cc0a
	Ondrag           Hash = 0x3d606
	Ondragend        Hash = 0x3d609
	Ondragenter      Hash = 0x3df0b
	Ondragleave      Hash = 0x3ea0b
	Ondragover       Hash = 0x3f50a
	Ondragstart      Hash = 0x3ff0b
	Ondrop           Hash = 0x41006
	Ondurationchange Hash = 0x42010
	Onemptied        Hash = 0x41709
	Onended          Hash = 0x43007
	Onerror          Hash = 0x43707
	Onfocus          Hash = 0x43e07
	Onhashchange     Hash = 0x45e0c
	Oninput          Hash = 0x46a07
	Oninvalid        Hash = 0x47109
	Onkeydown        Hash = 0x47a09
	Onkeypress       Hash = 0x4870a
	Onkeyup          Hash = 0x49707
	Onload           Hash = 0x49e06
	Onloadeddata     Hash = 0x49e0c
	Onloadedmetadata Hash = 0x4b110
	Onloadstart      Hash = 0x4c70b
	Onmessage        Hash = 0x4d209
	Onmousedown      Hash = 0x4db0b
	Onmousemove      Hash = 0x4f80b
	Onmouseout       Hash = 0x5030a
	Onmouseover      Hash = 0x5100b
	Onmouseup        Hash = 0x51b09
	Onmousewheel     Hash = 0x5240c
	Onoffline        Hash = 0x53009
	Ononline         Hash = 0x53b08
	Onpagehide       Hash = 0x5430a
	Onpageshow       Hash = 0x5570a
	Onpause          Hash = 0x56307
	Onplay           Hash = 0x59006
	Onplaying        Hash = 0x59009
	Onpopstate       Hash = 0x5990a
	Onprogress       Hash = 0x5a30a
	Onratechange     Hash = 0x5b70c
	Onreset          Hash = 0x5c307
	Onresize         Hash = 0x5ca08
	Onscroll         Hash = 0x5d208
	Onseeked         Hash = 0x5dd08
	Onseeking        Hash = 0x5e509
	Onselect         Hash = 0x5ee08
	Onshow           Hash = 0x5f806
	Onstalled        Hash = 0x60309
	Onstorage        Hash = 0x60c09
	Onsubmit         Hash = 0x61508
	Onsuspend        Hash = 0x62909
	Ontimeupdate     Hash = 0x4520c
	Onunload         Hash = 0x63208
	Onvolumechange   Hash = 0x63a0e
	Onwaiting        Hash = 0x64809
	Open             Hash = 0x57404
	Optgroup         Hash = 0x12008
	Optimum          Hash = 0x65107
	Option           Hash = 0x66606
	Output           Hash = 0x50a06
	P                Hash = 0xc01
	Param            Hash = 0xc05
	Pattern          Hash = 0x9b07
	Pauseonexit      Hash = 0x5650b
	Ping             Hash = 0xe704
	Placeholder      Hash = 0x1270b
	Plaintext        Hash = 0x17d09
	Poster           Hash = 0x1fb06
	Pre              Hash = 0x30d03
	Preload          Hash = 0x30d07
	Profile          Hash = 0x34f07
	Progress         Hash = 0x5a508
	Prompt           Hash = 0x66006
	Public           Hash = 0x57c06
	Q                Hash = 0x8d01
	Radiogroup       Hash = 0x30a
	Rb               Hash = 0x1d02
	Readonly         Hash = 0x38708
	Rel              Hash = 0x30e03
	Required         Hash = 0x8b08
	Rev              Hash = 0x29303
	Reversed         Hash = 0x29308
	Rows             Hash = 0x6604
	Rowspan          Hash = 0x6607
	Rp               Hash = 0x28902
	Rt               Hash = 0x1c702
	Rtc              Hash = 0x1c703
	Ruby             Hash = 0xf304
	Rules            Hash = 0x13105
	S                Hash = 0x3d01
	Samp             Hash = 0x9804
	Sandbox          Hash = 0x13507
	Scope            Hash = 0x35a05
	Scoped           Hash = 0x35a06
	Script           Hash = 0x31b06
	Scrolling        Hash = 0x5d409
	Seamless         Hash = 0x3a108
	Section          Hash = 0x13f07
	Select           Hash = 0x5f006
	Selected         Hash = 0x5f008
	Shape            Hash = 0x23005
	Size             Hash = 0x24e04
	Sizes            Hash = 0x24e05
	Small            Hash = 0x23a05
	Sortable         Hash = 0x25208
	Source           Hash = 0x26506
	Spacer           Hash = 0x37106
	Span             Hash = 0x6904
	Spellcheck       Hash = 0x3a80a
	Src              Hash = 0x44403
	Srcdoc           Hash = 0x44406
	Srclang          Hash = 0x49007
	Start            Hash = 0x40505
	Step             Hash = 0x65d04
	Strike           Hash = 0x5ac06
	Strong           Hash = 0x67c06
	Style            Hash = 0x68205
	Sub              Hash = 0x61703
	Summary          Hash = 0x68707
	Sup              Hash = 0x68e03
	Svg              Hash = 0x69103
	System           Hash = 0x69406
	Tabindex         Hash = 0x4bf08
	Table            Hash = 0x25505
	Target           Hash = 0x2f806
	Tbody            Hash = 0x3f05
	Td               Hash = 0xaa02
	Text             Hash = 0x18204
	Textarea         Hash = 0x38208
	Tfoot            Hash = 0xed05
	Th               Hash = 0x19502
	Thead            Hash = 0x36a05
	Time             Hash = 0x2ed04
	Title            Hash = 0x16605
	Tr               Hash = 0x18502
	Track            Hash = 0x18505
	Translate        Hash = 0x22009
	Truespeed        Hash = 0x27209
	Tt               Hash = 0x9d02
	Type             Hash = 0x10f04
	Typemustmatch    Hash = 0x1e00d
	U                Hash = 0xb01
	Ul               Hash = 0x5802
	Undeterminate    Hash = 0x250d
	Usemap           Hash = 0x15d06
	Valign           Hash = 0x1506
	Value            Hash = 0x10a05
	Valuetype        Hash = 0x10a09
	Var              Hash = 0x32f03
	Video            Hash = 0x6a005
	Visible          Hash = 0x6ab07
	Vlink            Hash = 0x6b205
	Wbr              Hash = 0x56003
	Width            Hash = 0x5fd05
	Wrap             Hash = 0x57904
	Xmlns            Hash = 0x13b05
	Xmp              Hash = 0x17b03
)

var _Hash_table = [1 << 9]Hash{
	0x0:   0x2ca0b, // formenctype
	0x1:   0x2d50a, // formmethod
	0x2:   0x3c10b, // oncuechange
	0x3:   0x3d606, // ondrag
	0x6:   0x5ac06, // strike
	0x7:   0x6a005, // video
	0x9:   0x58107, // content
	0xa:   0x4e07,  // enabled
	0xb:   0x57706, // nowrap
	0xc:   0xd304,  // link
	0xe:   0x28902, // rp
	0xf:   0x2830c, // onafterprint
	0x10:  0x16106, // applet
	0x11:  0xed05,  // tfoot
	0x12:  0x4ea0e, // defaultChecked
	0x13:  0x3330b, // frameborder
	0x14:  0xee06,  // footer
	0x15:  0x5f008, // selected
	0x16:  0x49007, // srclang
	0x18:  0x5100b, // onmouseover
	0x19:  0x1dc04, // code
	0x1b:  0x47109, // oninvalid
	0x1c:  0x62104, // face
	0x1e:  0x3b60b, // contextmenu
	0x1f:  0xa308,  // frameset
	0x21:  0x54b0c, // defaultMuted
	0x22:  0x19f05, // color
	0x23:  0x59006, // onplay
	0x25:  0x2ef05, // meter
	0x26:  0x60c09, // onstorage
	0x27:  0x38708, // readonly
	0x29:  0x34f07, // profile
	0x2a:  0x8607,  // declare
	0x2b:  0xb01,   // u
	0x2c:  0x31908, // noscript
	0x2d:  0x65708, // manifest
	0x2e:  0x1be06, // button
	0x2f:  0x2e908, // datetime
	0x30:  0x46c05, // input
	0x31:  0x5407,  // default
	0x32:  0x1dc08, // codetype
	0x33:  0x2a80d, // foreignobject
	0x34:  0x69907, // marquee
	0x36:  0x19d07, // bgcolor
	0x37:  0x19b02, // h1
	0x39:  0x1e0a,  // background
	0x3b:  0x2f40a, // formtarget
	0x41:  0x2f806, // target
	0x43:  0x23a05, // small
	0x44:  0x44908, // codebase
	0x45:  0x53605, // inert
	0x47:  0x38105, // mtext
	0x48:  0x6607,  // rowspan
	0x49:  0x2bd0d, // onbeforeprint
	0x4a:  0x53b08, // ononline
	0x4c:  0x28f06, // figure
	0x4d:  0x4b110, // onloadedmetadata
	0x4e:  0xbb07,  // acronym
	0x50:  0x38f08, // multiple
	0x51:  0x320e,  // accept-charset
	0x52:  0x24e05, // sizes
	0x53:  0x29b0d, // foreignObject
	0x55:  0x2e30a, // novalidate
	0x56:  0x5430a, // onpagehide
	0x57:  0x2e202, // mn
	0x58:  0x37f02, // h4
	0x5a:  0x1c702, // rt
	0x5b:  0xd205,  // alink
	0x5e:  0x66006, // prompt
	0x5f:  0x12d02, // ol
	0x61:  0x5ca08, // onresize
	0x64:  0x68707, // summary
	0x65:  0x5990a, // onpopstate
	0x66:  0x38604, // area
	0x68:  0x64809, // onwaiting
	0x6b:  0xdc04,  // name
	0x6c:  0x23506, // coords
	0x6d:  0x34303, // img
	0x6e:  0x65d04, // step
	0x6f:  0x5e509, // onseeking
	0x70:  0x32104, // high
	0x71:  0x49707, // onkeyup
	0x72:  0x5f006, // select
	0x73:  0x18505, // track
	0x74:  0x34b05, // ismap
	0x76:  0x46a07, // oninput
	0x77:  0x8d01,  // q
	0x78:  0x47a09, // onkeydown
	0x79:  0x33e05, // image
	0x7a:  0x2b504, // form
	0x7b:  0x60309, // onstalled
	0x7d:  0x42808, // onchange
	0x7e:  0x1af05, // blink
	0x7f:  0xeb03,  // alt
	0x80:  0xf705,  // async
	0x82:  0x1702,  // li
	0x84:  0x2c02,  // mi
	0x85:  0xfc06,  // itemid
	0x86:  0x11305, // audio
	0x87:  0x31b06, // script
	0x8b:  0x44406, // srcdoc
	0x8e:  0xc704,  // mark
	0x8f:  0x18a03, // bdo
	0x91:  0x4f80b, // onmousemove
	0x93:  0x3bd04, // menu
	0x94:  0x45104, // font
	0x95:  0x16b08, // autoplay
	0x96:  0x6b205, // vlink
	0x98:  0x6e02,  // em
	0x9b:  0x1f408, // colgroup
	0x9c:  0x57404, // open
	0x9d:  0x1d606, // legend
	0x9e:  0x4c70b, // onloadstart
	0xa2:  0x22009, // translate
	0xa3:  0x6e05,  // embed
	0xa4:  0x1c905, // class
	0xa7:  0x36b06, // header
	0xa9:  0x49e06, // onload
	0xaa:  0x36a05, // thead
	0xab:  0x5d409, // scrolling
	0xac:  0xc05,   // param
	0xae:  0x9b07,  // pattern
	0xaf:  0x9207,  // details
	0xb1:  0x57c06, // public
	0xb3:  0x4db0b, // onmousedown
	0xb4:  0x16003, // map
	0xb6:  0x25a0b, // crossorigin
	0xb7:  0x1506,  // valign
	0xb9:  0x1c207, // onabort
	0xba:  0x66606, // option
	0xbb:  0x26506, // source
	0xbc:  0x6205,  // defer
	0xbd:  0x1eb09, // challenge
	0xbf:  0x10a05, // value
	0xc0:  0x23c0f, // allowfullscreen
	0xc1:  0xca03,  // kbd
	0xc2:  0x2060d, // annotationXml
	0xc3:  0x5b70c, // onratechange
	0xc4:  0x4dd02, // mo
	0xc6:  0x3a80a, // spellcheck
	0xc7:  0x2c03,  // min
	0xc8:  0x49e0c, // onloadeddata
	0xc9:  0x40b05, // clear
	0xca:  0x42010, // ondurationchange
	0xcb:  0x1a04,  // nobr
	0xcd:  0x27209, // truespeed
	0xcf:  0x30806, // hgroup
	0xd0:  0x40505, // start
	0xd3:  0x41208, // dropzone
	0xd5:  0x7405,  // label
	0xd8:  0xde0a,  // mediagroup
	0xd9:  0x14406, // onblur
	0xdb:  0x27e07, // caption
	0xdd:  0x7c04,  // desc
	0xde:  0x13b05, // xmlns
	0xdf:  0x30006, // height
	0xe0:  0x21307, // command
	0xe2:  0x5650b, // pauseonexit
	0xe3:  0x67c06, // strong
	0xe4:  0x43707, // onerror
	0xe5:  0x61508, // onsubmit
	0xe6:  0xb308,  // language
	0xe7:  0x47f08, // download
	0xe9:  0x51b09, // onmouseup
	0xec:  0x2ce07, // enctype
	0xed:  0x5ee08, // onselect
	0xee:  0x2af06, // object
	0xef:  0x17d09, // plaintext
	0xf0:  0x3cc0a, // ondblclick
	0xf1:  0x18c10, // oncanplaythrough
	0xf2:  0xd903,  // dir
	0xf3:  0x38208, // textarea
	0xf4:  0xe704,  // ping
	0xf5:  0x2d906, // method
	0xf6:  0x22908, // controls
	0xf7:  0x37106, // spacer
	0xf8:  0x69103, // svg
	0xf9:  0x30404, // html
	0xfa:  0x3d01,  // s
	0xfc:  0xcc06,  // dialog
	0xfe:  0x1e00d, // typemustmatch
	0xff:  0x3ad07, // checked
	0x101: 0x1fb06, // poster
	0x102: 0x3260a, // http-equiv
	0x103: 0x44403, // src
	0x104: 0x10108, // disabled
	0x105: 0x36b07, // headers
	0x106: 0x5a30a, // onprogress
	0x107: 0x26b08, // fieldset
	0x108: 0x32f03, // var
	0x10a: 0xa305,  // frame
	0x10b: 0x36008, // itemtype
	0x10c: 0x3f50a, // ondragover
	0x10d: 0x15609, // autofocus
	0x10f: 0x601,   // i
	0x110: 0x35902, // ms
	0x111: 0x44d04, // base
	0x113: 0x35a05, // scope
	0x114: 0x3206,  // accept
	0x115: 0x56e08, // itemprop
	0x117: 0xfb04,  // cite
	0x118: 0x3907,  // charset
	0x119: 0x16605, // title
	0x11a: 0x5b007, // keytype
	0x11b: 0x18204, // text
	0x11c: 0x65107, // optimum
	0x11e: 0x36b04, // head
	0x121: 0x21a07, // compact
	0x123: 0x62909, // onsuspend
	0x124: 0x4aa04, // list
	0x125: 0x4520c, // ontimeupdate
	0x126: 0x62306, // center
	0x127: 0x31406, // hidden
	0x129: 0x35609, // itemscope
	0x12c: 0x1aa02, // dl
	0x12d: 0x13f07, // section
	0x12e: 0x11708, // oncancel
	0x12f: 0x6a407, // onclick
	0x130: 0xde05,  // media
	0x131: 0x50a06, // output
	0x132: 0x4a608, // datalist
	0x133: 0x5240c, // onmousewheel
	0x134: 0x44d08, // basefont
	0x135: 0x37709, // maxlength
	0x136: 0x6ab07, // visible
	0x137: 0x2df0e, // formnovalidate
	0x139: 0x17b03, // xmp
	0x13a: 0x101,   // b
	0x13b: 0x5570a, // onpageshow
	0x13c: 0xf304,  // ruby
	0x13d: 0x1270b, // placeholder
	0x13e: 0x4aa07, // listing
	0x140: 0x26303, // ins
	0x141: 0x61b07, // itemref
	0x144: 0x540f,  // defaultSelected
	0x146: 0x3ea0b, // ondragleave
	0x147: 0x1b40a, // blockquote
	0x148: 0x57904, // wrap
	0x14a: 0x1ac03, // big
	0x14b: 0x30e03, // rel
	0x14c: 0x41006, // ondrop
	0x14e: 0x69406, // system
	0x14f: 0x30a,   // radiogroup
	0x150: 0x25505, // table
	0x152: 0x56003, // wbr
	0x153: 0x3b40d, // oncontextmenu
	0x155: 0x250d,  // undeterminate
	0x157: 0x20104, // cols
	0x158: 0x13507, // sandbox
	0x159: 0x1303,  // nav
	0x15a: 0x37703, // max
	0x15b: 0x7808,  // longdesc
	0x15c: 0x5fd05, // width
	0x15d: 0x34902, // h3
	0x15e: 0x1a407, // bgsound
	0x161: 0x10a09, // valuetype
	0x162: 0x68205, // style
	0x164: 0x3f05,  // tbody
	0x165: 0x40707, // article
	0x169: 0xcb03,  // bdi
	0x16a: 0x67607, // address
	0x16b: 0x23005, // shape
	0x16c: 0x2b906, // action
	0x16e: 0x18502, // tr
	0x16f: 0xaa02,  // td
	0x170: 0x3d609, // ondragend
	0x171: 0x5802,  // ul
	0x172: 0x33806, // border
	0x174: 0x4a06,  // keygen
	0x175: 0x4004,  // body
	0x177: 0x1cf09, // draggable
	0x178: 0x2b50a, // formaction
	0x17b: 0x34406, // mglyph
	0x17d: 0x1d02,  // rb
	0x17e: 0x2fe02, // h6
	0x17f: 0x41709, // onemptied
	0x180: 0x5c307, // onreset
	0x181: 0x1004,  // main
	0x182: 0x11e04, // loop
	0x183: 0x4870a, // onkeypress
	0x184: 0x9d02,  // tt
	0x186: 0x20107, // colspan
	0x188: 0x36804, // math
	0x189: 0x1605,  // align
	0x18a: 0xa108,  // noframes
	0x18b: 0xaf02,  // hr
	0x18c: 0xc10a,  // malignmark
	0x18e: 0x23e03, // low
	0x18f: 0x8502,  // id
	0x190: 0x6604,  // rows
	0x191: 0x29303, // rev
	0x192: 0x63208, // onunload
	0x193: 0x39705, // muted
	0x194: 0x35a06, // scoped
	0x195: 0x31602, // dd
	0x196: 0x5ff02, // dt
	0x197: 0x66a0e, // onbeforeunload
	0x199: 0x2060a, // annotation
	0x19a: 0x29308, // reversed
	0x19c: 0x10f04, // type
	0x19d: 0x56307, // onpause
	0x19e: 0xd604,  // kind
	0x19f: 0x4a604, // data
	0x1a0: 0x4e507, // noshade
	0x1a3: 0x13105, // rules
	0x1a4: 0x12008, // optgroup
	0x1a5: 0x202,   // br
	0x1a7: 0x1,     // a
	0x1a8: 0x5030a, // onmouseout
	0x1aa: 0x53009, // onoffline
	0x1ab: 0x63a0e, // onvolumechange
	0x1ae: 0x61703, // sub
	0x1b3: 0x29b03, // for
	0x1b5: 0x8b08,  // required
	0x1b6: 0x5a508, // progress
	0x1b7: 0x15d06, // usemap
	0x1b8: 0x7f06,  // canvas
	0x1b9: 0x58004, // icon
	0x1bb: 0x1c703, // rtc
	0x1bc: 0x8305,  // aside
	0x1bd: 0x2ed04, // time
	0x1be: 0x3ff0b, // ondragstart
	0x1c0: 0x27b0a, // figcaption
	0x1c1: 0xaf04,  // href
	0x1c2: 0x33206, // iframe
	0x1c3: 0x18c09, // oncanplay
	0x1c4: 0x6904,  // span
	0x1c5: 0x30d03, // pre
	0x1c6: 0x6c07,  // noembed
	0x1c8: 0x5dd08, // onseeked
	0x1c9: 0x4b904, // meta
	0x1ca: 0x32402, // h2
	0x1cb: 0x3a108, // seamless
	0x1cc: 0xab03,  // dfn
	0x1cd: 0x17304, // axis
	0x1cf: 0x3df0b, // ondragenter
	0x1d0: 0x19502, // th
	0x1d1: 0x45e0c, // onhashchange
	0x1d2: 0xb304,  // lang
	0x1d3: 0x43e07, // onfocus
	0x1d5: 0x24e04, // size
	0x1d8: 0x14a0c, // autocomplete
	0x1d9: 0xaf08,  // hreflang
	0x1da: 0x9804,  // samp
	0x1de: 0x19f03, // col
	0x1df: 0x10803, // div
	0x1e0: 0x25208, // sortable
	0x1e1: 0x7203,  // del
	0x1e3: 0x39c07, // onclose
	0x1e6: 0xd907,  // dirname
	0x1e8: 0x1c907, // classid
	0x1e9: 0x30d07, // preload
	0x1ea: 0x4bf08, // tabindex
	0x1eb: 0x60102, // h5
	0x1ec: 0x5d208, // onscroll
	0x1ed: 0x5810f, // contenteditable
	0x1ee: 0x4d209, // onmessage
	0x1ef: 0x4,     // abbr
	0x1f0: 0x17507, // isindex
	0x1f1: 0x68e03, // sup
	0x1f3: 0x24a08, // noresize
	0x1f5: 0x59009, // onplaying
	0x1f6: 0x4409,  // accesskey
	0x1fa: 0xc01,   // p
	0x1fb: 0x43007, // onended
	0x1fc: 0x5f806, // onshow
	0x1fe: 0xad06,  // nohref
}
