package hikvision

import (
	"encoding/xml"
	"net/url"
)

type FocusConfiguration struct {
	XMLName                       xml.Name `xml:"FocusConfiguration,omitempty"`
	XMLVersion                    string   `xml:"version,attr"`
	XMLNamespace                  string   `xml:"xmlns,attr"`
	FocusStyle                    string   `xml:"focusStyle,omitempty" json:"focusStyle,omitempty"`                                       //req
	FocusLimited                  int      `xml:"focusLimited,omitempty" json:"focusLimited,omitempty"`                                   //opt
	FocusPosition                 int      `xml:"focusPosition,omitempty" json:"focusPosition,omitempty"`                                 //dep depends on "FocusStyle"
	FocusSpeed                    int      `xml:"focusSpeed,omitempty" json:"focusSpeed,omitempty"`                                       //opt
	FocusSensitivity              int      `xml:"focusSensitivity,omitempty" json:"focusSensitivity,omitempty"`                           //opt
	TemperatureChangeAdaptEnabled bool     `xml:"temperatureChangeAdaptEnabled,omitempty" json:"temperatureChangeAdaptEnabled,omitempty"` //opt
	RelativeFocusPos              int      `xml:"relativeFocusPos,omitempty" json:"relativeFocusPos,omitempty"`                           //opt
	HighTemperaturePriority       bool     `xml:"highTemperaturePriority,omitempty" json:"highTemperaturePriority,omitempty"`             //opt
}

func (c *Client) GetFocusConf(channelId string) (resp *FocusConfiguration, err error) {
	path := "/ISAPI/Image/channels/"+channelId+"/focusconfiguration"
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

func (c *Client) PutFocusConf(data *FocusConfiguration, channelId string) (resp *ResponseStatus, err error) {
	path := "/ISAPI/Image/channels/"+channelId+"/focusConfiguration"
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
