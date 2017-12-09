package controller

import (
	"crypto/tls"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUrlPrefix(t *testing.T) {

	for _, test := range []struct {
		host   string
		uri    string
		tls    *tls.ConnectionState
		prefix string
	}{
		{host: "mock-host", uri: "/mock-uri", tls: &tls.ConnectionState{}, prefix: "https://mock-host/mock-uri"},
		{host: "mock-host", uri: "/mock-uri", tls: nil, prefix: "http://mock-host/mock-uri"},
	} {
		r := http.Request{
			TLS:        test.tls,
			Host:       test.host,
			RequestURI: test.uri,
		}
		assert.Equal(t, test.prefix, urlPrefix(&r))
	}

}

func TestPointerString(t *testing.T) {

	for _, test := range []string{
		"hello",
	} {
		var s interface{}
		s = str(test)
		_, ok := s.(*string)
		assert.True(t, ok)
	}
}

func TestHALSelfLink(t *testing.T) {

	for _, test := range []struct {
		url string
	}{
		{url: "http://example.com/hello-world"},
	} {
		l := halSelfLink(test.url)
		assert.NotNil(t, l)
		assert.Nil(t, l.Doc)
		assert.NotNil(t, l.Self)

		assert.Equal(t, l.Self.Href.String(), test.url)
	}
}
func TestHALRootRscLinks(t *testing.T) {

	tests := []struct {
		fqe    string
		server string
		opid   string
		docURL string
	}{
		{fqe: "http://example.com/hello-world", server: "http://example.com", opid: "my-operation", docURL: "http://example.com/docs#operation/my-operation"},
	}

	for _, test := range tests {
		l := halRootRscLinks(&apiHostBase{
			FQEndpoint:  test.fqe,
			ServerURL:   test.server,
			OperationID: test.opid,
		})
		assert.NotNil(t, l)
		assert.NotNil(t, l.Doc)
		assert.NotNil(t, l.Self)

		assert.Equal(t, test.fqe, l.Self.Href.String())
		assert.Equal(t, test.docURL, l.Doc.Href.String())
	}
}
