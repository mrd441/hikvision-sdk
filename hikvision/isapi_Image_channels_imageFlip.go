package hikvision

import (
	"encoding/xml"
	"net/url"
)

type ImageFlip struct {
	XMLName        xml.Name `xml:"ImageFlip,omitempty"`
	XMLVersion     string   `xml:"version,attr"`
	XMLNamespace   string   `xml:"xmlns,attr"`
	Enabled        bool     `xml:"enabled,omitempty" json:"enabled,omitempty"`               //req
	ImageFlipStyle string   `xml:"ImageFlipStyle,omitempty" json:"ImageFlipStyle,omitempty"` //"LEFTRIGHT, UPDOWN, CENTER, AUTO
	FlipAngle      string   `xml:"flipAngle,omitempty" json:"flipAngle,omitempty"`           //"90, 180, 270"
}

func (c *Client) GetImageFlip(channelId string) (resp *ImageFlip, err error) {
	path := "/ISAPI/Image/channels/" + channelId + "/ImageFlip"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	body, err := c.Get(u)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) PutImageFlip(data *ImageFlip, channelId string) (resp *ResponseStatus, err error) {
	path := "/ISAPI/Image/channels/" + channelId + "/ImageFlip"
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, err
	}
	body, err := c.PutXML(u, data)
	if err != nil {
		return nil, err
	}
	err = xml.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
