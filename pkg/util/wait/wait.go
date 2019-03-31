package wait

//go:generate go get github.com/golang/mock/mockgen
//go:generate mockgen -destination=../mocks/mock_$GOPACKAGE/wait.go -package=mock_$GOPACKAGE -source wait.go
//go:generate gofmt -s -l -w ../mocks/mock_$GOPACKAGE/wait.go
//go:generate go get golang.org/x/tools/cmd/goimports
//go:generate goimports -local=github.com/openshift/openshift-azure -e -w ../mocks/mock_$GOPACKAGE/wait.go

import (
	"context"
	"crypto/x509"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
)

// PollImmediateUntil will poll until a stop condition is met
func PollImmediateUntil(interval time.Duration, condition wait.ConditionFunc, stopCh <-chan struct{}) error {
	done, err := condition()
	if err != nil {
		return err
	}
	if done {
		return nil
	}
	return wait.PollUntil(interval, condition, stopCh)
}

// SimpleHTTPClient to aid in mocking
type SimpleHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ForHTTPStatusOk polls until URL returns 200
func ForHTTPStatusOk(ctx context.Context, log *logrus.Entry, cli SimpleHTTPClient, urltocheck string, interval time.Duration) (*http.Response, error) {
	req, err := http.NewRequest("GET", urltocheck, nil)
	if err != nil {
		return nil, err
	}
	var resp *http.Response
	err = PollImmediateUntil(interval, func() (bool, error) {
		resp, err = cli.Do(req)
		if err, ok := err.(*url.Error); ok {
			if err, ok := err.Err.(*net.OpError); ok {
				if err, ok := err.Err.(*os.SyscallError); ok {
					switch err.Err {
					case syscall.ENETUNREACH, syscall.ECONNREFUSED:
						return false, nil
					}
				}
			}
			if _, ok := err.Err.(x509.UnknownAuthorityError); ok {
				log.Warn(err)
				return false, nil
			}
			if err.Timeout() || err.Err == io.EOF || err.Err == io.ErrUnexpectedEOF {
				return false, nil
			}
		}
		if err == io.EOF {
			return false, nil
		}
		if err != nil {
			return false, err
		}
		return resp.StatusCode == http.StatusOK, nil
	}, ctx.Done())
	return resp, err
}
