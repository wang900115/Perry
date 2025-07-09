package secureheader

import "github.com/gin-gonic/gin"

type SecureHeader struct{}

func NewSecureHeader() *SecureHeader {
	return &SecureHeader{}
}

func (SecureHeader) Middleware(c *gin.Context) {
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-XSS-Protection", "1; mode=block")
	c.Header("Referr-Policy", "strict-origin-when-cross-origin")
	c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
	c.Next()
}

/*
	X-Content-Type-Options: nosniff
	阻止瀏覽器探知檔案的 mime type ( disable Content sniffing )，一般情況下瀏覽器會去探知資源的 Content-Type，
	以判別資源類型，例如：image/png、text/css，而有些資源的 Content-Type 有可能是錯誤或缺少的，此時就會進行 Content sniffing猜測解析內容，
	將 X-Content-Type-Options 設成 nosniff，可阻止這種行為。
*/

/*
	X-Frame-Options: DENY
	代表不讓網頁載入 frame ( iframe 跟 object )。
	沒有設成 DENY 的風險為可能被惡意嵌入 iframe
*/

/*
	X-XSS-Protection:1; mode=block
	當瀏覽器發現跨站腳本攻擊時，停止加載網頁。
	0 為不啟用 XSS 過濾。
	1 為啟用 XSS 過濾，遇到 XSS ，僅刪除不安全的部分，不會block。
	1;mode=block 是遇到 XSS ，停止加載頁面。
*/

/*
	Referrer-Policy: no-referrer-when-downgrade

	no-referrer： 不帶 header 。
	no-referrer-when-downgrade ： https 到 https 的網站可以， https 到 http 的網站則不會帶header。
	origin：不帶完整 url ，僅只帶 origin 。
	origin-when-cross-origin：同源帶完整 url ，不同源帶 origin 。
	same-origin：同源才帶 header 。
	strict-origin：同源且 https 到 https 才帶 header ， http 就不會帶。
	strict-origin-when-cross-origin：同源帶完整 url，不同源帶 origin ，但是必須是 https 到 https 。
	unsafe-url：都帶。

*/

/*
	Strict-Transport-Security: max-age=31536000; includeSubDomains
	簡稱 HSTS，告知瀏覽器強制啟用 https ， max-age 代表強制維持 https 多少時間，一般為一年31536000秒，includeSubDomains包含子網域，若max-age=0，代表disable。
*/
