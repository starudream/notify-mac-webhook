package main

import (
	"net/http"

	"github.com/starudream/go-lib/log"
	"github.com/starudream/go-lib/router"
)

type Data struct {
	AndroidTitle   string `json:"android.title"`
	AndroidBigText string `json:"android.bigText"`
	AndroidText    string `json:"android.text"`

	FilterboxFieldAppName     string `json:"filterbox.field.APP_NAME"`
	FilterboxFieldPackageName string `json:"filterbox.field.PACKAGE_NAME"`
	FilterboxFieldChannelId   string `json:"filterbox.field.CHANNEL_ID"`
	FilterboxFieldWhen        int64  `json:"filterbox.field.WHEN"`
}

func handler(c *router.Context) {
	v := Data{}
	err := c.BindJSON(&v)
	if err != nil {
		log.Error().Msgf("request body bind error: %v", err)
		return
	}

	_, err = Notifier(NotifierReq{
		Data:     v,
		Title:    v.AndroidTitle,
		SubTitle: v.FilterboxFieldAppName + "(" + v.FilterboxFieldPackageName + ")",
		Message: func() string {
			if v.AndroidBigText != "" && v.AndroidBigText != "{android.bigText}" {
				return v.AndroidBigText
			} else if v.AndroidText != "" {
				return v.AndroidText
			} else {
				return "空"
			}
		}(),
	})
	if err != nil {
		log.Error().Msgf("notify error: %v", err)
		return
	}

	c.String(http.StatusOK, "ok")
}
