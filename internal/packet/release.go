package packet

/* Simple packet(release), including release information */
type release struct {
	name        string
	version     string
	tag         string
	description string
}
