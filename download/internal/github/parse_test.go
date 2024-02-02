package github

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlParseToRepositorye(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	log.Printf("%+v\n", repository)
	assert.Equal(t, nil, err)
}

func TestMatch(t *testing.T) {
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	matches, err := match(url)
	log.Printf("%+v\n", matches)
	assert.Equal(t, nil, err)
}

func TestRepositoryUrlIsRoot(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, repository.IsRoot())
	assert.Equal(t, false, repository.IsFile())
	assert.Equal(t, false, repository.IsDir())
}

func TestRepositoryUrlIsFile(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, repository.IsFile())
	assert.Equal(t, false, repository.IsDir())
	assert.Equal(t, false, repository.IsRoot())
}

func TestRepositoryUrlIsDir(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/tree/main/src/common/query"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, repository.IsDir())
	assert.Equal(t, false, repository.IsFile())
	assert.Equal(t, false, repository.IsRoot())
}

func TestRepositoryRootUrl(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://github.com/zzopen/mysqldoc", repository.RootUrl())
}

func TestRepositoryBranchUrl(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://github.com/zzopen/mysqldoc/main", repository.BranchUrl())
}

func TestRepositoryTreesApiUrl(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/git/trees/main?recursive=1", repository.GitTreesApiUrl("", true))
	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/git/trees/main", repository.GitTreesApiUrl("", false))
}

func TestRepositoryRawUserContentUrl(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://raw.githubusercontent.com/zzopen/mysqldoc/main/src/common/query/query.go", repository.RawUserContentUrl())
}

func TestRepositoryContentApiUrl(t *testing.T) {
	token := ""
	url := "https://github.com/zzopen/mysqldoc/blob/main/src/common/query/query.go"
	repository, err := UrlParseToRepository(url, token)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https://api.github.com/repos/zzopen/mysqldoc/contents/src/common/query/query.go?ref=main", repository.ContentApiUrl())
}
