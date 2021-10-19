package handler

import (
	"github.com/skip2/go-qrcode"
	"net/http"
)

func QrCodeCreate(res http.ResponseWriter, r *http.Request) {

	var png []byte
	png, err := qrcode.Encode("https://example.org ", qrcode.Medium, 256)
	if err != nil {
	}
	res.Write(png)
	/*session.GetCookie()

	http.ServeFile(res, r, "templates/qrSite.html")*/

}
