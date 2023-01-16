package cargo

import "time"

type CargoPackageInfo struct {
	Categories []interface{} `json:"categories"`
	Crate      struct {
		Badges        []interface{} `json:"badges"`
		Categories    []interface{} `json:"categories"`
		CreatedAt     time.Time     `json:"created_at"`
		Description   string        `json:"description"`
		Documentation interface{}   `json:"documentation"`
		Downloads     int           `json:"downloads"`
		ExactMatch    bool          `json:"exact_match"`
		Homepage      interface{}   `json:"homepage"`
		ID            string        `json:"id"`
		Keywords      []string      `json:"keywords"`
		Links         struct {
			OwnerTeam           string      `json:"owner_team"`
			OwnerUser           string      `json:"owner_user"`
			Owners              string      `json:"owners"`
			ReverseDependencies string      `json:"reverse_dependencies"`
			VersionDownloads    string      `json:"version_downloads"`
			Versions            interface{} `json:"versions"`
		} `json:"links"`
		MaxStableVersion string    `json:"max_stable_version"`
		MaxVersion       string    `json:"max_version"`
		Name             string    `json:"name"`
		NewestVersion    string    `json:"newest_version"`
		RecentDownloads  int       `json:"recent_downloads"`
		Repository       string    `json:"repository"`
		UpdatedAt        time.Time `json:"updated_at"`
		Versions         []int     `json:"versions"`
	} `json:"crate"`
	Keywords []struct {
		CratesCnt int       `json:"crates_cnt"`
		CreatedAt time.Time `json:"created_at"`
		ID        string    `json:"id"`
		Keyword   string    `json:"keyword"`
	} `json:"keywords"`
	Versions []struct {
		AuditActions []struct {
			Action string    `json:"action"`
			Time   time.Time `json:"time"`
			User   struct {
				Avatar string `json:"avatar"`
				ID     int    `json:"id"`
				Login  string `json:"login"`
				Name   string `json:"name"`
				URL    string `json:"url"`
			} `json:"user"`
		} `json:"audit_actions"`
		Checksum  string    `json:"checksum"`
		Crate     string    `json:"crate"`
		CrateSize int       `json:"crate_size"`
		CreatedAt time.Time `json:"created_at"`
		DlPath    string    `json:"dl_path"`
		Downloads int       `json:"downloads"`
		Features  struct {
		} `json:"features"`
		ID      int    `json:"id"`
		License string `json:"license"`
		Links   struct {
			Authors          string `json:"authors"`
			Dependencies     string `json:"dependencies"`
			VersionDownloads string `json:"version_downloads"`
		} `json:"links"`
		Num         string `json:"num"`
		PublishedBy struct {
			Avatar string `json:"avatar"`
			ID     int    `json:"id"`
			Login  string `json:"login"`
			Name   string `json:"name"`
			URL    string `json:"url"`
		} `json:"published_by"`
		ReadmePath string    `json:"readme_path"`
		UpdatedAt  time.Time `json:"updated_at"`
		Yanked     bool      `json:"yanked"`
	} `json:"versions"`
}
