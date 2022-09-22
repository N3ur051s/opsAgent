package security

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"simpleagent/pkg/util/filesystem"
	"simpleagent/pkg/util/log"
)

const (
	authTokenName       = "auth_token"
	authTokenMinimalLen = 32
)

// GenerateKeyPair create a public/private keypair
func GenerateKeyPair(bits int) (*rsa.PrivateKey, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, fmt.Errorf("generating random key: %v", err)
	}

	return privKey, nil
}

// CertTemplate create x509 certificate template
func CertTemplate() (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %s", err)
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(10 * 365 * 24 * time.Hour)
	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"simpleagent, Inc."},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
	}

	return &template, nil
}

// GenerateRootCert generates a root certificate
func GenerateRootCert(hosts []string, bits int) (cert *x509.Certificate, certPEM []byte, rootKey *rsa.PrivateKey, err error) {
	rootCertTmpl, err := CertTemplate()
	if err != nil {
		return
	}

	rootKey, err = GenerateKeyPair(bits)
	if err != nil {
		return
	}

	rootCertTmpl.IsCA = true
	rootCertTmpl.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature | x509.KeyUsageCRLSign
	rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			rootCertTmpl.IPAddresses = append(rootCertTmpl.IPAddresses, ip)
		} else {
			rootCertTmpl.DNSNames = append(rootCertTmpl.DNSNames, h)
		}
	}

	certDER, err := x509.CreateCertificate(rand.Reader, rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
	if err != nil {
		return
	}

	cert, err = x509.ParseCertificate(certDER)
	if err != nil {
		return
	}

	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	certPEM = pem.EncodeToMemory(&b)
	return
}

func GetAuthTokenFilepath() string {
	return filepath.Join(filepath.Dir(os.ExpandEnv("${HOME}/.simpleagent/")), authTokenName)
}

func FetchAuthToken() (string, error) {
	return fetchAuthToken(false)
}

func CreateOrFetchToken() (string, error) {
	return fetchAuthToken(true)
}

func fetchAuthToken(tokenCreationAllowed bool) (string, error) {
	authTokenFile := GetAuthTokenFilepath()

	if _, e := os.Stat(authTokenFile); os.IsNotExist(e) && tokenCreationAllowed {
		key := make([]byte, authTokenMinimalLen)
		_, e = rand.Read(key)
		if e != nil {
			return "", fmt.Errorf("can't create agent authentication token value: %s", e)
		}

		e = saveAuthToken(hex.EncodeToString(key), authTokenFile)
		if e != nil {
			return "", fmt.Errorf("error writing authentication token file on fs: %s", e)
		}
		log.Infof("Saved a new authentication token to %s", authTokenFile)
	}

	authTokenRaw, e := ioutil.ReadFile(authTokenFile)
	if e != nil {
		return "", fmt.Errorf("unable to read authentication token file: " + e.Error())
	}

	authToken := strings.TrimSpace(string(authTokenRaw))
	if len(authToken) < authTokenMinimalLen {
		return "", fmt.Errorf("invalid authentication token: must be at least %d characters in length", authTokenMinimalLen)
	}

	return authToken, nil
}

func DeleteAuthToken() error {
	authTokenFile := filepath.Join(filepath.Dir(os.ExpandEnv("${HOME}/.simpleagent/")), authTokenName)
	return os.Remove(authTokenFile)
}

func validateAuthToken(authToken string) error {
	if len(authToken) < authTokenMinimalLen {
		return fmt.Errorf("agent authentication token length must be greater than %d, curently: %d", authTokenMinimalLen, len(authToken))
	}
	return nil
}

func saveAuthToken(token, tokenPath string) error {
	if err := ioutil.WriteFile(tokenPath, []byte(token), 0o600); err != nil {
		return err
	}

	perms, err := filesystem.NewPermission()
	if err != nil {
		return err
	}

	if err := perms.RestrictAccessToUser(tokenPath); err != nil {
		log.Errorf("Failed to write auth token acl %s", err)
		return err
	}

	log.Infof("Wrote auth token")
	return nil
}
