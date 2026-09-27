package gitlab

/*
 * Asset Information struct
 *
 * A struct that provides a basic asset
 * information for any release contained
 * in any project that on gitlab platform.
 */

type assetSource struct {
	Format string `json:"format"`
	Url    string `json:"url"`
}

type asset struct {
	Count   int           `json:"count"`
	Sources []assetSource `json:"sources"`
}
