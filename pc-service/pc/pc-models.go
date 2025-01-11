package pc

import "github.com/typomaker/option"

type Request struct {
	Id          option.Option[int]
	Name        string `json:"name"`
	Description string `json:"description"`
	Processor   string `json:"processor"`
	VideoCard   string `json:"video_card"`
}
