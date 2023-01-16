package maven

type MavenPackageInfo struct {
	Package struct {
		System string `json:"system"`
		Name   string `json:"name"`
	} `json:"package"`
	Owners  []interface{} `json:"owners"`
	Version struct {
		Version                string        `json:"version"`
		SymbolicVersions       []interface{} `json:"symbolicVersions"`
		RefreshedAt            int           `json:"refreshedAt"`
		IsDefault              bool          `json:"isDefault"`
		Licenses               []string      `json:"licenses"`
		DependentCount         int           `json:"dependentCount"`
		DependentCountDirect   int           `json:"dependentCountDirect"`
		DependentCountIndirect int           `json:"dependentCountIndirect"`
		Links                  struct {
			Origins []string `json:"origins"`
			Repo    string   `json:"repo"`
		} `json:"links"`
		Projects []struct {
			Type        string `json:"type"`
			Name        string `json:"name"`
			ObservedAt  int    `json:"observedAt"`
			Issues      int    `json:"issues"`
			Forks       int    `json:"forks"`
			Stars       int    `json:"stars"`
			Description string `json:"description"`
			License     string `json:"license"`
			DisplayName string `json:"displayName"`
			Link        string `json:"link"`
			ScorecardV2 struct {
				Date string `json:"date"`
				Repo struct {
					Name   string `json:"name"`
					Commit string `json:"commit"`
				} `json:"repo"`
				Scorecard struct {
					Version string `json:"version"`
					Commit  string `json:"commit"`
				} `json:"scorecard"`
				Check []struct {
					Name          string `json:"name"`
					Documentation struct {
						Short string `json:"short"`
						URL   string `json:"url"`
					} `json:"documentation"`
					Score   int           `json:"score"`
					Reason  string        `json:"reason"`
					Details []interface{} `json:"details"`
				} `json:"check"`
				Metadata []interface{} `json:"metadata"`
				Score    float64       `json:"score"`
			} `json:"scorecardV2"`
		} `json:"projects"`
		Advisories      []interface{} `json:"advisories"`
		RelatedPackages struct {
			GoModuleMajorVersions []struct {
				System string `json:"system"`
				Name   string `json:"name"`
			} `json:"goModuleMajorVersions"`
		} `json:"relatedPackages"`
	} `json:"version"`
	DefaultVersion string `json:"defaultVersion"`
}
