package flags

// ** GRPM BOOLEAN FLAGS **
var BoolFlags map[string]*bool = map[string]*bool{
	"define":        nil, // initialize grpm tool.
	"clear":         nil, // clear (search, info) commands data cache.
	"show":          nil, // show user configuration options.
	"rfconfig":      nil, // allow user to refactor configuration options.
	"match":         nil, // display only options that match the configuration options (os,arch).
	"setup":         nil, // extract compressed or archived files into lib folder.
	"force":         nil, // apply all pop up options.
	"latest":        nil, // have different applications (show repositories or grab latest release).
	"latestRelease": nil, // show latest repository release.
}

// ** GRPM STRING FLAGS **
var StringFlags map[string]*string = map[string]*string{
	"repo_name":         nil, // repository name.
	"repo_owner":        nil, // repository owner name.
	"repo_tag":          nil, // repository tag number.
	"repo_search_order": nil, // search repository by order query.
	"repo_search_sort":  nil, // sort searched repositories.
}

// ** GRPM INTGER FLAGS **
var IntgerFlags map[string]*int = map[string]*int{
	"repo_search_page": nil, // choose page of searched repositories.
}
