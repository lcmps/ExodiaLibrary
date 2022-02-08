package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/valyala/fasthttp"
)

func MakeRequestFastHTTP(url string, header map[string]string, obj interface{}) ([]byte, error) {
	if obj != nil {
		kindOfJ := reflect.ValueOf(obj).Kind()
		if kindOfJ != reflect.Ptr {
			return nil, fmt.Errorf("obj need to be a pointer")
		}
	}

	var errRetries error
	var body []byte
	for i := 1; i <= 3; i++ {
		if i != 1 {
			time.Sleep(3 * time.Second)
		}

		// Acquire a request instance
		req := fasthttp.AcquireRequest()

		req.SetRequestURI(url)
		for k, v := range header {
			req.Header.Add(k, v)
		}

		// Acquire a response instance
		resp := fasthttp.AcquireResponse()

		err := fasthttp.Do(req, resp)
		if err != nil {
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)
			errRetries = err
			continue
		}

		if resp.StatusCode() > 202 {
			fasthttp.ReleaseRequest(req)
			fasthttp.ReleaseResponse(resp)
			continue
		}

		body = bytes.TrimPrefix(resp.Body(), []byte("\xef\xbb\xbf"))
		if obj != nil {
			err = json.Unmarshal(body, obj)
			if err != nil {
				fasthttp.ReleaseRequest(req)
				fasthttp.ReleaseResponse(resp)
				errRetries = err
				continue
			}
		}

		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
		errRetries = nil
		break
	}

	if errRetries != nil {
		return nil, errRetries
	}
	return body, nil
}
