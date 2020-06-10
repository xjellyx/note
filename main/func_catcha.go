package main

import (
	"bytes"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func GetCaptchaID(c *gin.Context) {
	l := captcha.DefaultLen
	id := captcha.NewLen(l)
	c.AbortWithStatusJSON(200, gin.H{"id": id})
}

func Captcha(c *gin.Context) {
	var d = struct {
		Ext        string `json:"ext" form:"ext" binding:"required"`
		Lang       string `json:"lang" form:"lang" binding:"required"`
		IsDownload bool   `json:"isDownload" form:"isDownload"`
	}{}
	if err := c.BindQuery(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	l := captcha.DefaultLen
	id := captcha.NewLen(l)
	_ = Serve(c.Writer, c.Request, id, d.Ext, d.Lang, d.IsDownload, captcha.StdWidth, captcha.StdHeight)
}
func CaptchaVerify(captchaId string, code string) bool {
	if captcha.VerifyString(captchaId, code) {
		return true
	} else {
		return false
	}

}
func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Captcha-ID", id)
	var content bytes.Buffer
	switch ext {
	case "png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case "wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func main() {
	router := gin.Default()
	router.GET("/getCaptchaID", GetCaptchaID)
	router.GET("/captcha", func(c *gin.Context) {
		Captcha(c)
	})

	router.GET("/captcha/verify/", func(c *gin.Context) {
		value := c.Query("value")
		id := c.Query("id")
		if CaptchaVerify(id, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})
	router.Run(":8080")
}

