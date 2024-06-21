package middleware

import (
	"errors"
	"fmt"
	"net/http"
)

// В стандартной библиотеке не нашел типа который бы имел такую реализацию поэтому создал сам
// Она отвечает за
type nextRoundTripper func(*http.Request) (*http.Response, error)

func (n nextRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	return n(request)
}

type Middleware func(next http.RoundTripper) http.RoundTripper

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.RoundTripper) http.RoundTripper {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

func SecretKey(next http.RoundTripper) http.RoundTripper {
	return nextRoundTripper(func(req *http.Request) (*http.Response, error) {
		// TODO Убрать от сюда ключ
		// Подвязка ключа в Header
		req.Header.Add("X-CMC_PRO_API_KEY", "320c56fe-4b4a-4b03-8481-87620269aebb")

		res, e := next.RoundTrip(req)
		return res, e
	})
}

func ResponseStatus(next http.RoundTripper) http.RoundTripper {
	return nextRoundTripper(func(req *http.Request) (*http.Response, error) {
		res, e := next.RoundTrip(req)

		// Проверка на то что статус запроа коректный
		statusOK := res.StatusCode >= 200 && res.StatusCode < 300
		if !statusOK {
			errMsg := fmt.Sprintf("Non-OK HTTP status: %d", res.StatusCode)
			res.Body.Close()

			return nil, errors.New(errMsg)
		}
		return res, e
	})
}
