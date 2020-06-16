package armory

// https://gist.github.com/michaljemala/d6f4e01c4834bf47a9c4

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	certFile = flag.String("cert", os.Getenv("HOME")+"/crt.pem", "A PEM eoncoded certificate file.")
	keyFile  = flag.String("key", os.Getenv("HOME")+"/key.pem", "A PEM encoded private key file.")
	//caFile   = flag.String("CA", "someCertCAFile", "A PEM eoncoded CA's certificate file.")
)

func ArmoryAPI(endpoint string) []byte {
	flag.Parse()

	// Load client cert
	cert, err := tls.LoadX509KeyPair(*certFile, *keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	//caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	//caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            caCertPool,
		InsecureSkipVerify: true, // for self signed certs
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// load env file for api url:
	errr := godotenv.Load(".env")
	if errr != nil {
		log.Fatal("Error loading .env file")
	}

	// Do GET something
	resp, err := client.Get(os.Getenv("API_URL") + endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(string(data))
	return data
}
