package constants

type GithubRelease struct {
	TagName string `json:"tag_name"`
	Prerelease bool
}