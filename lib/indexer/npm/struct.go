package npm

type NpmVersion struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	Description  string            `json:"description"`
	Main         string            `json:"main"`
	TrackingCode string            `json:"trackingCode"`
	Bin          map[string]string `json:"bin"`
	Keywords     []string          `json:"keywords"`
	Repository   struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
	Engines struct {
		Node string `json:"node"`
		Npm  string `json:"npm"`
	} `json:"engines"`
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	License string `json:"license"`
	Bugs    struct {
		URL string `json:"url"`
	} `json:"bugs"`
	Homepage     string            `json:"homepage"`
	Dependencies map[string]string `json:"dependencies"`
	EmberAddon   struct {
		Paths []string `json:"paths"`
		Main  string   `json:"main"`
	} `json:"ember-addon"`
	ID      string `json:"_id"`
	Scripts struct {
	} `json:"scripts"`
	Shasum      string `json:"_shasum"`
	From        string `json:"_from"`
	NpmVersion  string `json:"_npmVersion"`
	NodeVersion string `json:"_nodeVersion"`
	NpmUser     struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"_npmUser"`
	Dist struct {
		Shasum     string `json:"shasum"`
		Tarball    string `json:"tarball"`
		Integrity  string `json:"integrity"`
		Signatures []struct {
			Keyid string `json:"keyid"`
			Sig   string `json:"sig"`
		} `json:"signatures"`
	} `json:"dist"`
	Maintainers []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"maintainers"`
	NpmOperationalInternal struct {
		Host string `json:"host"`
		Tmp  string `json:"tmp"`
	} `json:"_npmOperationalInternal"`
	Directories struct {
	} `json:"directories"`
}

type NpmPackageInfo struct {
	ID          string                `json:"_id"`
	Rev         string                `json:"_rev"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	DistTags    map[string]string     `json:"dist-tags"`
	Versions    map[string]NpmVersion `json:"versions"`
	Readme      string                `json:"readme"`
	Maintainers []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"maintainers"`
	Time       map[string]string `json:"time"`
	Homepage   string            `json:"homepage"`
	Keywords   []string          `json:"keywords"`
	Repository struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	Bugs struct {
		URL string `json:"url"`
	} `json:"bugs"`
	License        string          `json:"license"`
	ReadmeFilename string          `json:"readmeFilename"`
	Users          map[string]bool `json:"users"`
}
