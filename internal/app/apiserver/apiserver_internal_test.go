package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(test *testing.T) {
	server := New(NewConfig())                               // Server
	rec := httptest.NewRecorder()                            // Recorder
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil) // Request object
	server.handleHello().ServeHTTP(rec, req)                 // Вызываем у сервера обработчик (handleHello), а у обработчика вызываем serveHTTP
	assert.Equal(test, rec.Body.String(), "Hello web!!!")    // С помощью библиотеки assert проверим recorder body = нашей строке
}
