package main

import (
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

	if v.AndroidTitle == "{android.title}" {
		v.AndroidTitle = ""
	}
	if v.AndroidBigText == "{android.bigText}" {
		v.AndroidBigText = ""
	}
	if v.AndroidText == "{android.text}" {
		v.AndroidText = ""
	}
	if v.AndroidBigText == "" && v.AndroidText == "" {
		return
	}

	go func() {
		r, e := Notify(c.GetHeader("X-Request-ID"), NotifyReq{
			Data:     v,
			Title:    v.AndroidTitle,
			SubTitle: v.FilterboxFieldAppName + "(" + v.FilterboxFieldPackageName + ")",
			Message:  ternary(v.AndroidBigText != "", v.AndroidBigText, v.AndroidText),
		})
		if e != nil {
			log.Error().Msgf("notify error: %v, stdout: %s, stderr: %s", e, r.Stdout, r.Stderr)
			return
		}
	}()

	c.OK()
}

func ternary[T any](condition bool, ifOutput T, elseOutput T) T {
	if condition {
		return ifOutput
	}
	return elseOutput
}
