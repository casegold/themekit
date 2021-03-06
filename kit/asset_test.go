package kit

import (
	"encoding/base64"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoadAssetSuite struct {
	suite.Suite
	allocatedFiles []string
}

func (s *LoadAssetSuite) TearDownTest() {
	os.RemoveAll("../fixtures/output")
	os.RemoveAll("../fixtures/download")
}

func (s *LoadAssetSuite) TestIsValid() {
	asset := Asset{Key: "test.txt", Value: "one"}
	assert.Equal(s.T(), true, asset.IsValid())
	asset = Asset{Key: "test.txt", Attachment: "one"}
	assert.Equal(s.T(), true, asset.IsValid())
	asset = Asset{Value: "one"}
	assert.Equal(s.T(), false, asset.IsValid())
	asset = Asset{Key: "test.txt"}
	assert.Equal(s.T(), false, asset.IsValid())
}

func (s *LoadAssetSuite) TestSize() {
	asset := Asset{Value: "one"}
	assert.Equal(s.T(), 3, asset.Size())
	asset = Asset{Attachment: "other"}
	assert.Equal(s.T(), 5, asset.Size())
}

func (s *LoadAssetSuite) TestWrite() {
	asset := Asset{Key: "output/blah.txt", Value: "this is content"}

	err := asset.Write("../nope")
	assert.NotNil(s.T(), err)

	err = asset.Write("../fixtures")
	assert.Nil(s.T(), err)
}

func (s *LoadAssetSuite) TestContents() {
	asset := Asset{Value: "this is content"}
	data, err := asset.Contents()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 15, len(data))

	asset = Asset{Attachment: "this is bad content"}
	data, err = asset.Contents()
	assert.NotNil(s.T(), err)

	asset = Asset{Attachment: base64.StdEncoding.EncodeToString([]byte("this is bad content"))}
	data, err = asset.Contents()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 19, len(data))
	assert.Equal(s.T(), []byte("this is bad content"), data)

	asset = Asset{Key: "test.json", Value: "{\"test\":\"one\"}"}
	data, err = asset.Contents()
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 19, len(data))
	assert.Equal(s.T(), `{
  "test": "one"
}`, string(data))
}

func (s *LoadAssetSuite) TestFindAllFiles() {
	files, err := findAllFiles("../fixtures/project/valid_patterns")
	assert.Equal(s.T(), "Path is not a directory", err.Error())
	files, err = findAllFiles("../fixtures/project")
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []string{
		clean("../fixtures/project/assets/application.js"),
		clean("../fixtures/project/assets/pixel.png"),
		clean("../fixtures/project/config/settings_data.json"),
		clean("../fixtures/project/config.json"),
		clean("../fixtures/project/invalid_config.yml"),
		clean("../fixtures/project/layout/.gitkeep"),
		clean("../fixtures/project/locales/en.json"),
		clean("../fixtures/project/snippets/snippet.js"),
		clean("../fixtures/project/templates/customers/test.liquid"),
		clean("../fixtures/project/templates/template.liquid"),
		clean("../fixtures/project/valid_config.yml"),
		clean("../fixtures/project/valid_patterns"),
		clean("../fixtures/project/whatever.txt"),
	}, files)
}

func (s *LoadAssetSuite) TestLoadAssetsFromDirectory() {
	assets, err := loadAssetsFromDirectory(clean("../fixtures/project/valid_patterns"), "", func(path string) bool { return false })
	assert.Equal(s.T(), "Path is not a directory", err.Error())
	assets, err = loadAssetsFromDirectory(clean("../fixtures/project"), "", func(path string) bool {
		return path != "assets/application.js"
	})
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), []Asset{{
		Key:   "assets/application.js",
		Value: "//this is js\n",
	}}, assets)
}

func (s *LoadAssetSuite) TestLoadAssetsFromDirectoryWithSubdir() {
	assets, err := loadAssetsFromDirectory(clean("../fixtures/project"), "assets", func(path string) bool { return false })
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 2, len(assets))
}

func (s *LoadAssetSuite) TestLoadAsset() {
	asset, err := loadAsset(clean("../fixtures/project"), clean("assets/application.js"))
	assert.Equal(s.T(), "assets/application.js", asset.Key)
	assert.Equal(s.T(), true, asset.IsValid())
	assert.Equal(s.T(), "//this is js\n", asset.Value)
	assert.Nil(s.T(), err)

	asset, err = loadAsset(clean("../fixtures/project"), "nope.txt")
	assert.NotNil(s.T(), err)

	asset, err = loadAsset(clean("../fixtures/project"), "templates")
	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), ErrAssetIsDir, err)

	asset, err = loadAsset(clean("../fixtures/project"), "assets/pixel.png")
	assert.Nil(s.T(), err)
	assert.True(s.T(), len(asset.Attachment) > 0)
	assert.True(s.T(), asset.IsValid())
}

func TestLoadAssetSuite(t *testing.T) {
	suite.Run(t, new(LoadAssetSuite))
}
