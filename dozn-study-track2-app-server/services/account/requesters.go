package account

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func doRequest(url string, method string, body []byte) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")
	req.SetBody(body)

	fasthttp.Do(req, resp)

	bodyBytes := resp.Body()
	fmt.Println(string(bodyBytes))
}
