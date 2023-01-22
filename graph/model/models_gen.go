// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddLicenseInput struct {
	ID string `json:"id"`
}

type AddLicensePayload struct {
	License []*License `json:"license"`
	NumUids *int       `json:"numUids"`
}

type AddPackageInput struct {
	Name      string                    `json:"name"`
	Versions  []*PackageVersionRef      `json:"versions"`
	Licenses  []*LicenseRef             `json:"licenses"`
	Languages []*ProgrammingLanguageRef `json:"languages"`
	Type      *PackageTypeRef           `json:"type"`
	URL       *string                   `json:"url"`
}

type AddPackagePayload struct {
	Package []*Package `json:"package"`
	NumUids *int       `json:"numUids"`
}

type AddPackageTypeInput struct {
	ID string `json:"id"`
}

type AddPackageTypePayload struct {
	PackageType []*PackageType `json:"packageType"`
	NumUids     *int           `json:"numUids"`
}

type AddPackageVersionInput struct {
	ID string `json:"id"`
}

type AddPackageVersionPayload struct {
	PackageVersion []*PackageVersion `json:"packageVersion"`
	NumUids        *int              `json:"numUids"`
}

type AddProgrammingLanguageInput struct {
	ID string `json:"id"`
}

type AddProgrammingLanguagePayload struct {
	ProgrammingLanguage []*ProgrammingLanguage `json:"programmingLanguage"`
	NumUids             *int                   `json:"numUids"`
}

type AuthRule struct {
	And  []*AuthRule `json:"and"`
	Or   []*AuthRule `json:"or"`
	Not  *AuthRule   `json:"not"`
	Rule *string     `json:"rule"`
}

type ContainsFilter struct {
	Point   *PointRef   `json:"point"`
	Polygon *PolygonRef `json:"polygon"`
}

type CustomHTTP struct {
	URL                  string     `json:"url"`
	Method               HTTPMethod `json:"method"`
	Body                 *string    `json:"body"`
	Graphql              *string    `json:"graphql"`
	Mode                 *Mode      `json:"mode"`
	ForwardHeaders       []string   `json:"forwardHeaders"`
	SecretHeaders        []string   `json:"secretHeaders"`
	IntrospectionHeaders []string   `json:"introspectionHeaders"`
	SkipIntrospection    *bool      `json:"skipIntrospection"`
}

type DateTimeFilter struct {
	Eq      *string        `json:"eq"`
	In      []*string      `json:"in"`
	Le      *string        `json:"le"`
	Lt      *string        `json:"lt"`
	Ge      *string        `json:"ge"`
	Gt      *string        `json:"gt"`
	Between *DateTimeRange `json:"between"`
}

type DateTimeRange struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type DeleteLicensePayload struct {
	License []*License `json:"license"`
	Msg     *string    `json:"msg"`
	NumUids *int       `json:"numUids"`
}

type DeletePackagePayload struct {
	Package []*Package `json:"package"`
	Msg     *string    `json:"msg"`
	NumUids *int       `json:"numUids"`
}

type DeletePackageTypePayload struct {
	PackageType []*PackageType `json:"packageType"`
	Msg         *string        `json:"msg"`
	NumUids     *int           `json:"numUids"`
}

type DeletePackageVersionPayload struct {
	PackageVersion []*PackageVersion `json:"packageVersion"`
	Msg            *string           `json:"msg"`
	NumUids        *int              `json:"numUids"`
}

type DeleteProgrammingLanguagePayload struct {
	ProgrammingLanguage []*ProgrammingLanguage `json:"programmingLanguage"`
	Msg                 *string                `json:"msg"`
	NumUids             *int                   `json:"numUids"`
}

type FloatFilter struct {
	Eq      *float64    `json:"eq"`
	In      []*float64  `json:"in"`
	Le      *float64    `json:"le"`
	Lt      *float64    `json:"lt"`
	Ge      *float64    `json:"ge"`
	Gt      *float64    `json:"gt"`
	Between *FloatRange `json:"between"`
}

type FloatRange struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type GenerateMutationParams struct {
	Add    *bool `json:"add"`
	Update *bool `json:"update"`
	Delete *bool `json:"delete"`
}

type GenerateQueryParams struct {
	Get       *bool `json:"get"`
	Query     *bool `json:"query"`
	Password  *bool `json:"password"`
	Aggregate *bool `json:"aggregate"`
}

type Int64Filter struct {
	Eq      *string     `json:"eq"`
	In      []*string   `json:"in"`
	Le      *string     `json:"le"`
	Lt      *string     `json:"lt"`
	Ge      *string     `json:"ge"`
	Gt      *string     `json:"gt"`
	Between *Int64Range `json:"between"`
}

type Int64Range struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type IntFilter struct {
	Eq      *int      `json:"eq"`
	In      []*int    `json:"in"`
	Le      *int      `json:"le"`
	Lt      *int      `json:"lt"`
	Ge      *int      `json:"ge"`
	Gt      *int      `json:"gt"`
	Between *IntRange `json:"between"`
}

type IntRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type IntersectsFilter struct {
	Polygon      *PolygonRef      `json:"polygon"`
	MultiPolygon *MultiPolygonRef `json:"multiPolygon"`
}

type License struct {
	ID string `json:"id"`
}

type LicenseAggregateResult struct {
	Count *int    `json:"count"`
	IDMin *string `json:"idMin"`
	IDMax *string `json:"idMax"`
}

type LicenseFilter struct {
	ID  *StringHashFilter   `json:"id"`
	Has []*LicenseHasFilter `json:"has"`
	And []*LicenseFilter    `json:"and"`
	Or  []*LicenseFilter    `json:"or"`
	Not *LicenseFilter      `json:"not"`
}

type LicenseOrder struct {
	Asc  *LicenseOrderable `json:"asc"`
	Desc *LicenseOrderable `json:"desc"`
	Then *LicenseOrder     `json:"then"`
}

type LicenseRef struct {
	ID string `json:"id"`
}

type MultiPolygon struct {
	Polygons []*Polygon `json:"polygons"`
}

type MultiPolygonRef struct {
	Polygons []*PolygonRef `json:"polygons"`
}

type NearFilter struct {
	Distance   float64   `json:"distance"`
	Coordinate *PointRef `json:"coordinate"`
}

type Package struct {
	Name               string                              `json:"name" json:"Package.name"`
	Versions           []*PackageVersion                   `json:"versions"`
	Licenses           []*License                          `json:"licenses"`
	Languages          []*ProgrammingLanguage              `json:"languages"`
	Type               *PackageType                        `json:"type"`
	URL                *string                             `json:"url"`
	VersionsAggregate  *PackageVersionAggregateResult      `json:"versionsAggregate"`
	LicensesAggregate  *LicenseAggregateResult             `json:"licensesAggregate"`
	LanguagesAggregate *ProgrammingLanguageAggregateResult `json:"languagesAggregate"`
}

type PackageAggregateResult struct {
	Count   *int    `json:"count"`
	NameMin *string `json:"nameMin"`
	NameMax *string `json:"nameMax"`
	URLMin  *string `json:"urlMin"`
	URLMax  *string `json:"urlMax"`
}

type PackageFilter struct {
	Name *StringTermFilter   `json:"name"`
	URL  *StringTermFilter   `json:"url"`
	Has  []*PackageHasFilter `json:"has"`
	And  []*PackageFilter    `json:"and"`
	Or   []*PackageFilter    `json:"or"`
	Not  *PackageFilter      `json:"not"`
}

type PackageOrder struct {
	Asc  *PackageOrderable `json:"asc"`
	Desc *PackageOrderable `json:"desc"`
	Then *PackageOrder     `json:"then"`
}

type PackagePatch struct {
	Name      *string                   `json:"name"`
	Versions  []*PackageVersionRef      `json:"versions"`
	Licenses  []*LicenseRef             `json:"licenses"`
	Languages []*ProgrammingLanguageRef `json:"languages"`
	Type      *PackageTypeRef           `json:"type"`
	URL       *string                   `json:"url"`
}

type PackageRef struct {
	Name      *string                   `json:"name"`
	Versions  []*PackageVersionRef      `json:"versions"`
	Licenses  []*LicenseRef             `json:"licenses"`
	Languages []*ProgrammingLanguageRef `json:"languages"`
	Type      *PackageTypeRef           `json:"type"`
	URL       *string                   `json:"url"`
}

type PackageType struct {
	ID string `json:"id"`
}

type PackageTypeAggregateResult struct {
	Count *int    `json:"count"`
	IDMin *string `json:"idMin"`
	IDMax *string `json:"idMax"`
}

type PackageTypeFilter struct {
	ID  *StringHashFilter       `json:"id"`
	Has []*PackageTypeHasFilter `json:"has"`
	And []*PackageTypeFilter    `json:"and"`
	Or  []*PackageTypeFilter    `json:"or"`
	Not *PackageTypeFilter      `json:"not"`
}

type PackageTypeOrder struct {
	Asc  *PackageTypeOrderable `json:"asc"`
	Desc *PackageTypeOrderable `json:"desc"`
	Then *PackageTypeOrder     `json:"then"`
}

type PackageTypeRef struct {
	ID string `json:"id"`
}

type PackageVersion struct {
	ID string `json:"id"`
}

type PackageVersionAggregateResult struct {
	Count *int    `json:"count"`
	IDMin *string `json:"idMin"`
	IDMax *string `json:"idMax"`
}

type PackageVersionFilter struct {
	ID  *StringHashFilter          `json:"id"`
	Has []*PackageVersionHasFilter `json:"has"`
	And []*PackageVersionFilter    `json:"and"`
	Or  []*PackageVersionFilter    `json:"or"`
	Not *PackageVersionFilter      `json:"not"`
}

type PackageVersionOrder struct {
	Asc  *PackageVersionOrderable `json:"asc"`
	Desc *PackageVersionOrderable `json:"desc"`
	Then *PackageVersionOrder     `json:"then"`
}

type PackageVersionRef struct {
	ID string `json:"id"`
}

type Point struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type PointGeoFilter struct {
	Near   *NearFilter   `json:"near"`
	Within *WithinFilter `json:"within"`
}

type PointList struct {
	Points []*Point `json:"points"`
}

type PointListRef struct {
	Points []*PointRef `json:"points"`
}

type PointRef struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Polygon struct {
	Coordinates []*PointList `json:"coordinates"`
}

type PolygonGeoFilter struct {
	Near       *NearFilter       `json:"near"`
	Within     *WithinFilter     `json:"within"`
	Contains   *ContainsFilter   `json:"contains"`
	Intersects *IntersectsFilter `json:"intersects"`
}

type PolygonRef struct {
	Coordinates []*PointListRef `json:"coordinates"`
}

type ProgrammingLanguage struct {
	ID string `json:"id"`
}

type ProgrammingLanguageAggregateResult struct {
	Count *int    `json:"count"`
	IDMin *string `json:"idMin"`
	IDMax *string `json:"idMax"`
}

type ProgrammingLanguageFilter struct {
	ID  *StringHashFilter               `json:"id"`
	Has []*ProgrammingLanguageHasFilter `json:"has"`
	And []*ProgrammingLanguageFilter    `json:"and"`
	Or  []*ProgrammingLanguageFilter    `json:"or"`
	Not *ProgrammingLanguageFilter      `json:"not"`
}

type ProgrammingLanguageOrder struct {
	Asc  *ProgrammingLanguageOrderable `json:"asc"`
	Desc *ProgrammingLanguageOrderable `json:"desc"`
	Then *ProgrammingLanguageOrder     `json:"then"`
}

type ProgrammingLanguageRef struct {
	ID string `json:"id"`
}

type StringExactFilter struct {
	Eq      *string      `json:"eq"`
	In      []*string    `json:"in"`
	Le      *string      `json:"le"`
	Lt      *string      `json:"lt"`
	Ge      *string      `json:"ge"`
	Gt      *string      `json:"gt"`
	Between *StringRange `json:"between"`
}

type StringFullTextFilter struct {
	Alloftext *string `json:"alloftext"`
	Anyoftext *string `json:"anyoftext"`
}

type StringHashFilter struct {
	Eq *string   `json:"eq"`
	In []*string `json:"in"`
}

type StringRange struct {
	Min string `json:"min"`
	Max string `json:"max"`
}

type StringRegExpFilter struct {
	Regexp *string `json:"regexp"`
}

type StringTermFilter struct {
	Allofterms *string `json:"allofterms"`
	Anyofterms *string `json:"anyofterms"`
}

type UpdatePackageInput struct {
	Filter *PackageFilter `json:"filter"`
	Set    *PackagePatch  `json:"set"`
	Remove *PackagePatch  `json:"remove"`
}

type UpdatePackagePayload struct {
	Package []*Package `json:"package"`
	NumUids *int       `json:"numUids"`
}

type WithinFilter struct {
	Polygon *PolygonRef `json:"polygon"`
}

type DgraphIndex string

const (
	DgraphIndexInt      DgraphIndex = "int"
	DgraphIndexInt64    DgraphIndex = "int64"
	DgraphIndexFloat    DgraphIndex = "float"
	DgraphIndexBool     DgraphIndex = "bool"
	DgraphIndexHash     DgraphIndex = "hash"
	DgraphIndexExact    DgraphIndex = "exact"
	DgraphIndexTerm     DgraphIndex = "term"
	DgraphIndexFulltext DgraphIndex = "fulltext"
	DgraphIndexTrigram  DgraphIndex = "trigram"
	DgraphIndexRegexp   DgraphIndex = "regexp"
	DgraphIndexYear     DgraphIndex = "year"
	DgraphIndexMonth    DgraphIndex = "month"
	DgraphIndexDay      DgraphIndex = "day"
	DgraphIndexHour     DgraphIndex = "hour"
	DgraphIndexGeo      DgraphIndex = "geo"
)

var AllDgraphIndex = []DgraphIndex{
	DgraphIndexInt,
	DgraphIndexInt64,
	DgraphIndexFloat,
	DgraphIndexBool,
	DgraphIndexHash,
	DgraphIndexExact,
	DgraphIndexTerm,
	DgraphIndexFulltext,
	DgraphIndexTrigram,
	DgraphIndexRegexp,
	DgraphIndexYear,
	DgraphIndexMonth,
	DgraphIndexDay,
	DgraphIndexHour,
	DgraphIndexGeo,
}

func (e DgraphIndex) IsValid() bool {
	switch e {
	case DgraphIndexInt, DgraphIndexInt64, DgraphIndexFloat, DgraphIndexBool, DgraphIndexHash, DgraphIndexExact, DgraphIndexTerm, DgraphIndexFulltext, DgraphIndexTrigram, DgraphIndexRegexp, DgraphIndexYear, DgraphIndexMonth, DgraphIndexDay, DgraphIndexHour, DgraphIndexGeo:
		return true
	}
	return false
}

func (e DgraphIndex) String() string {
	return string(e)
}

func (e *DgraphIndex) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DgraphIndex(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DgraphIndex", str)
	}
	return nil
}

func (e DgraphIndex) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HTTPMethod string

const (
	HTTPMethodGet    HTTPMethod = "GET"
	HTTPMethodPost   HTTPMethod = "POST"
	HTTPMethodPut    HTTPMethod = "PUT"
	HTTPMethodPatch  HTTPMethod = "PATCH"
	HTTPMethodDelete HTTPMethod = "DELETE"
)

var AllHTTPMethod = []HTTPMethod{
	HTTPMethodGet,
	HTTPMethodPost,
	HTTPMethodPut,
	HTTPMethodPatch,
	HTTPMethodDelete,
}

func (e HTTPMethod) IsValid() bool {
	switch e {
	case HTTPMethodGet, HTTPMethodPost, HTTPMethodPut, HTTPMethodPatch, HTTPMethodDelete:
		return true
	}
	return false
}

func (e HTTPMethod) String() string {
	return string(e)
}

func (e *HTTPMethod) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HTTPMethod(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HTTPMethod", str)
	}
	return nil
}

func (e HTTPMethod) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LicenseHasFilter string

const (
	LicenseHasFilterID LicenseHasFilter = "id"
)

var AllLicenseHasFilter = []LicenseHasFilter{
	LicenseHasFilterID,
}

func (e LicenseHasFilter) IsValid() bool {
	switch e {
	case LicenseHasFilterID:
		return true
	}
	return false
}

func (e LicenseHasFilter) String() string {
	return string(e)
}

func (e *LicenseHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LicenseHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LicenseHasFilter", str)
	}
	return nil
}

func (e LicenseHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LicenseOrderable string

const (
	LicenseOrderableID LicenseOrderable = "id"
)

var AllLicenseOrderable = []LicenseOrderable{
	LicenseOrderableID,
}

func (e LicenseOrderable) IsValid() bool {
	switch e {
	case LicenseOrderableID:
		return true
	}
	return false
}

func (e LicenseOrderable) String() string {
	return string(e)
}

func (e *LicenseOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LicenseOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LicenseOrderable", str)
	}
	return nil
}

func (e LicenseOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Mode string

const (
	ModeBatch  Mode = "BATCH"
	ModeSingle Mode = "SINGLE"
)

var AllMode = []Mode{
	ModeBatch,
	ModeSingle,
}

func (e Mode) IsValid() bool {
	switch e {
	case ModeBatch, ModeSingle:
		return true
	}
	return false
}

func (e Mode) String() string {
	return string(e)
}

func (e *Mode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Mode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Mode", str)
	}
	return nil
}

func (e Mode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PackageHasFilter string

const (
	PackageHasFilterName      PackageHasFilter = "name"
	PackageHasFilterVersions  PackageHasFilter = "versions"
	PackageHasFilterLicenses  PackageHasFilter = "licenses"
	PackageHasFilterLanguages PackageHasFilter = "languages"
	PackageHasFilterType      PackageHasFilter = "type"
	PackageHasFilterURL       PackageHasFilter = "url"
)

var AllPackageHasFilter = []PackageHasFilter{
	PackageHasFilterName,
	PackageHasFilterVersions,
	PackageHasFilterLicenses,
	PackageHasFilterLanguages,
	PackageHasFilterType,
	PackageHasFilterURL,
}

func (e PackageHasFilter) IsValid() bool {
	switch e {
	case PackageHasFilterName, PackageHasFilterVersions, PackageHasFilterLicenses, PackageHasFilterLanguages, PackageHasFilterType, PackageHasFilterURL:
		return true
	}
	return false
}

func (e PackageHasFilter) String() string {
	return string(e)
}

func (e *PackageHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageHasFilter", str)
	}
	return nil
}

func (e PackageHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PackageOrderable string

const (
	PackageOrderableName PackageOrderable = "name"
	PackageOrderableURL  PackageOrderable = "url"
)

var AllPackageOrderable = []PackageOrderable{
	PackageOrderableName,
	PackageOrderableURL,
}

func (e PackageOrderable) IsValid() bool {
	switch e {
	case PackageOrderableName, PackageOrderableURL:
		return true
	}
	return false
}

func (e PackageOrderable) String() string {
	return string(e)
}

func (e *PackageOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageOrderable", str)
	}
	return nil
}

func (e PackageOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PackageTypeHasFilter string

const (
	PackageTypeHasFilterID PackageTypeHasFilter = "id"
)

var AllPackageTypeHasFilter = []PackageTypeHasFilter{
	PackageTypeHasFilterID,
}

func (e PackageTypeHasFilter) IsValid() bool {
	switch e {
	case PackageTypeHasFilterID:
		return true
	}
	return false
}

func (e PackageTypeHasFilter) String() string {
	return string(e)
}

func (e *PackageTypeHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageTypeHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageTypeHasFilter", str)
	}
	return nil
}

func (e PackageTypeHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PackageTypeOrderable string

const (
	PackageTypeOrderableID PackageTypeOrderable = "id"
)

var AllPackageTypeOrderable = []PackageTypeOrderable{
	PackageTypeOrderableID,
}

func (e PackageTypeOrderable) IsValid() bool {
	switch e {
	case PackageTypeOrderableID:
		return true
	}
	return false
}

func (e PackageTypeOrderable) String() string {
	return string(e)
}

func (e *PackageTypeOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageTypeOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageTypeOrderable", str)
	}
	return nil
}

func (e PackageTypeOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PackageVersionHasFilter string

const (
	PackageVersionHasFilterID PackageVersionHasFilter = "id"
)

var AllPackageVersionHasFilter = []PackageVersionHasFilter{
	PackageVersionHasFilterID,
}

func (e PackageVersionHasFilter) IsValid() bool {
	switch e {
	case PackageVersionHasFilterID:
		return true
	}
	return false
}

func (e PackageVersionHasFilter) String() string {
	return string(e)
}

func (e *PackageVersionHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageVersionHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageVersionHasFilter", str)
	}
	return nil
}

func (e PackageVersionHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PackageVersionOrderable string

const (
	PackageVersionOrderableID PackageVersionOrderable = "id"
)

var AllPackageVersionOrderable = []PackageVersionOrderable{
	PackageVersionOrderableID,
}

func (e PackageVersionOrderable) IsValid() bool {
	switch e {
	case PackageVersionOrderableID:
		return true
	}
	return false
}

func (e PackageVersionOrderable) String() string {
	return string(e)
}

func (e *PackageVersionOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PackageVersionOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PackageVersionOrderable", str)
	}
	return nil
}

func (e PackageVersionOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProgrammingLanguageHasFilter string

const (
	ProgrammingLanguageHasFilterID ProgrammingLanguageHasFilter = "id"
)

var AllProgrammingLanguageHasFilter = []ProgrammingLanguageHasFilter{
	ProgrammingLanguageHasFilterID,
}

func (e ProgrammingLanguageHasFilter) IsValid() bool {
	switch e {
	case ProgrammingLanguageHasFilterID:
		return true
	}
	return false
}

func (e ProgrammingLanguageHasFilter) String() string {
	return string(e)
}

func (e *ProgrammingLanguageHasFilter) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProgrammingLanguageHasFilter(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProgrammingLanguageHasFilter", str)
	}
	return nil
}

func (e ProgrammingLanguageHasFilter) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProgrammingLanguageOrderable string

const (
	ProgrammingLanguageOrderableID ProgrammingLanguageOrderable = "id"
)

var AllProgrammingLanguageOrderable = []ProgrammingLanguageOrderable{
	ProgrammingLanguageOrderableID,
}

func (e ProgrammingLanguageOrderable) IsValid() bool {
	switch e {
	case ProgrammingLanguageOrderableID:
		return true
	}
	return false
}

func (e ProgrammingLanguageOrderable) String() string {
	return string(e)
}

func (e *ProgrammingLanguageOrderable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProgrammingLanguageOrderable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProgrammingLanguageOrderable", str)
	}
	return nil
}

func (e ProgrammingLanguageOrderable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}