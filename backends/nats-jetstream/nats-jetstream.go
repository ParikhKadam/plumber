package nats_jetstream

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber/types"
	"github.com/batchcorp/plumber/util"
	"github.com/batchcorp/plumber/validate"

	"github.com/batchcorp/plumber-schemas/build/go/protos/args"
	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
)

const BackendName = "nats-jetstream"

var (
	ErrMissingTLSKey  = errors.New("--tls-client-key-file cannot be blank if using ssl")
	ErrMissingTlsCert = errors.New("--tls-client-cert-file cannot be blank if using ssl")
	ErrMissingTLSCA   = errors.New("--tls-ca-file cannot be blank if using ssl")
	ErrMissingStream  = errors.New("--stream cannot be empty")
)

type NatsJetstream struct {
	connOpts *opts.ConnectionOptions
	client   *nats.Conn
	log      *logrus.Entry
}

func New(connOpts *opts.ConnectionOptions) (*NatsJetstream, error) {
	if err := validateBaseConnOpts(connOpts); err != nil {
		return nil, errors.Wrap(err, "invalid connection options")
	}

	args := connOpts.GetNatsJetstream()

	uri, err := url.Parse(args.Dsn)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse address")
	}

	// Credentials can be specified by a .creds file if users do not wish to pass in with the DSN
	var creds nats.Option
	if len(args.UserCredentials) > 0 {
		if util.FileExists(args.UserCredentials) {
			creds = nats.UserCredentials(string(args.UserCredentials))
		} else {
			creds = func(o *nats.Options) error {
				o.UserJWT = func() (string, error) {
					return string(args.UserCredentials), nil
				}
				o.SignatureCB = nil
				return nil
			}
		}
	}

	var client *nats.Conn
	if uri.Scheme == "tls" {
		// TLS Secured connection
		tlsConfig, err := generateTLSConfig(args)
		if err != nil {
			return nil, err
		}

		client, err = nats.Connect(args.Dsn, nats.Secure(tlsConfig), creds)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new nats client")
		}
	} else {
		// Plaintext connection
		client, err = nats.Connect(args.Dsn, creds)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create new nats client")
		}
	}

	return &NatsJetstream{
		connOpts: connOpts,
		client:   client,
		log:      logrus.WithField("backend", BackendName),
	}, nil
}

func (n *NatsJetstream) Name() string {
	return BackendName
}

func (n *NatsJetstream) Close(_ context.Context) error {
	n.client.Close()
	return nil
}

func (n *NatsJetstream) Test(_ context.Context) error {
	return types.NotImplementedErr
}

func generateTLSConfig(args *args.NatsJetstreamConn) (*tls.Config, error) {
	certpool := x509.NewCertPool()

	var cert tls.Certificate
	var err error

	if util.FileExists(args.TlsOptions.TlsClientCert) {
		// CLI input, read from file
		pemCerts, err := ioutil.ReadFile(string(args.TlsOptions.TlsCaCert))
		if err == nil {
			certpool.AppendCertsFromPEM(pemCerts)
		}

		cert, err = tls.LoadX509KeyPair(string(args.TlsOptions.TlsClientCert), string(args.TlsOptions.TlsClientKey))
		if err != nil {
			return nil, errors.Wrap(err, "unable to load ssl keypair")
		}

	} else {
		certpool.AppendCertsFromPEM(args.TlsOptions.TlsCaCert)

		cert, err = tls.X509KeyPair(args.TlsOptions.TlsClientCert, args.TlsOptions.TlsClientKey)
		if err != nil {
			return nil, errors.Wrap(err, "unable to load ssl keypair")
		}
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse certificate")
	}

	// Create tls.Config with desired tls properties
	return &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: args.TlsOptions.TlsSkipVerify,
		Certificates:       []tls.Certificate{cert},
		MinVersion:         tls.VersionTLS12,
	}, nil
}

func validateBaseConnOpts(connOpts *opts.ConnectionOptions) error {
	if connOpts == nil {
		return validate.ErrMissingConnOpts
	}

	if connOpts.Conn == nil {
		return validate.ErrMissingConnCfg
	}

	args := connOpts.GetNatsJetstream()
	if args == nil {
		return validate.ErrMissingConnArgs
	}

	if strings.HasPrefix(args.Dsn, "tls") {
		if len(args.TlsOptions.TlsClientKey) == 0 {
			return ErrMissingTLSKey
		}

		if len(args.TlsOptions.TlsClientCert) == 0 {
			return ErrMissingTlsCert
		}

		if len(args.TlsOptions.TlsCaCert) == 0 {
			return ErrMissingTLSCA
		}
	}

	return nil
}
