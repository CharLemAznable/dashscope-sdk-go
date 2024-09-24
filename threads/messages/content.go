package messages

import "github.com/gogf/gf/v2/encoding/gjson"

type Content interface {
	GetType() string
}

func ContentsFromJsons(jsons []*gjson.Json) (contents []Content) {
	for _, json := range jsons {
		if json.Get("type").String() == "text" {
			contents = append(contents, NewTextContentFromJson(json))
		} else if json.Get("type").String() == "image_file" {
			contents = append(contents, NewImageFileContentFromJson(json))
		}
	}
	return
}

////////////////////////////////////////////////////////////////////////////////

type TextContent interface {
	Content
	GetText() *Text
}

type Text struct {
	Value string `json:"value"`
}

func NewTextContentFromJson(json *gjson.Json) TextContent {
	return &textContent{
		Type: json.Get("type").String(),
		Text: &Text{
			Value: json.Get("text.value").String(),
		},
	}
}

type textContent struct {
	Type string `json:"type"`
	Text *Text  `json:"text"`
}

func (t *textContent) GetType() string {
	return t.Type
}
func (t *textContent) GetText() *Text {
	return t.Text
}

////////////////////////////////////////////////////////////////////////////////

type ImageFileContent interface {
	Content
	GetImageFile() *ImageFile
}

type ImageFile struct {
	FileId string `json:"file_id"`
}

func NewImageFileContentFromJson(json *gjson.Json) ImageFileContent {
	return &imageFileContent{
		Type: json.Get("type").String(),
		ImageFile: &ImageFile{
			FileId: json.Get("image_file.file_id").String(),
		},
	}
}

type imageFileContent struct {
	Type      string     `json:"type"`
	ImageFile *ImageFile `json:"image_file"`
}

func (i *imageFileContent) GetType() string {
	return i.Type
}
func (i *imageFileContent) GetImageFile() *ImageFile {
	return i.ImageFile
}
