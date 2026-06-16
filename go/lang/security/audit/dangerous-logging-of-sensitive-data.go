package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"log/slog"
)

func handler(req *http.Request, resp *http.Response, logger *zap.Logger) {
	password := "hunter2"
	token := "abc123"
	secretName := "kube-secret-name"
	namespace := "default"

	// ruleid: go.security.audit.dangerous-logging-sensitive-argument
	log.Println(password)

	// ruleid: go.security.audit.dangerous-logging-sensitive-argument
	log.Printf("token: %s", token)

	// ok: go.security.audit.dangerous-logging-sensitive-argument
	log.Println(secretName)

	// ok: go.security.audit.dangerous-logging-sensitive-argument
	log.Println(namespace)

	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	fmt.Println(os.Getenv("AWS_SECRET_ACCESS_KEY"))

	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	log.Println(req.Header)

	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	log.Println(req.Header.Get("Authorization"))

	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	log.Println(req.Header.Get("X-R-Sess"))

	body, _ := io.ReadAll(req.Body)
	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	log.Printf("request body: %s", body)

	body2, _ := ioutil.ReadAll(resp.Body)
	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	logrus.Infof("response body: %s", body2)

	dump, _ := httputil.DumpRequest(req, true)
	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	log.Println(string(dump))

	respDump, _ := httputil.DumpResponse(resp, true)
	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	log.Println(respDump)

	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	logger.Info("request", zap.Any("request_headers", req.Header))

	// ruleid: go.security.audit.dangerous-logging-sensitive-field-key
	logger.Info("secret", zap.String("personal_access_token", token))

	// ruleid: go.security.audit.dangerous-logging-sensitive-data
	slog.Any("headers", req.Header)

	redacted := redact(body)
	// ok: go.security.audit.dangerous-logging-sensitive-data
	log.Println(redacted)

	// ok: go.security.audit.dangerous-logging-sensitive-data
	log.Println(req.URL.Path)
}

func redact(v []byte) []byte {
	return []byte("[redacted]")
}
