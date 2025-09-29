package main
import (
	"fmt"
	"log"
	"os"
	"github.com/sirupsen/logrus"
)
func main() {
	password := "hunter2"
	token := "abc123"
	secretKey := "my-secret"
	secretName := "kube-secret-name"
	namespace := "default"
	// BAD: direct logging of sensitive variables
	// ruleid: go.security.audit.dangerous-logging-sensitive-var
	log.Println(password)
	// BAD: formatted string with sensitive variable
	// ruleid: go.security.audit.dangerous-logging-format-string
	log.Printf("User token: %s", token)
	// BAD: structured logging with sensitive field key
	// ruleid: go.security.audit.dangerous-logging-field-key
	logrus.WithField("secretKey", secretKey).Info("logging secret")
	// BAD: logging environment variable with sensitive name
	// ruleid: go.security.audit.dangerous-logging-env
	fmt.Println(os.Getenv("AWS_SECRET_ACCESS_KEY"))
	// GOOD: logging non-sensitive metadata (secretName is just an identifier)
	// ok: go.security.audit.dangerous-logging-sensitive-var
	log.Println(secretName)
	// GOOD: logging non-sensitive metadata (namespace, Name suffix)
	// ok: go.security.audit.dangerous-logging-sensitive-var
	log.Println(namespace)
}
