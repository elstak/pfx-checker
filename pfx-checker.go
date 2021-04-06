package main

import (
  "fmt"
  "os"
  "io/ioutil"
  "path/filepath"
  "golang.org/x/crypto/pkcs12"
  "crypto/x509"
  "crypto/rsa"
  "errors"
)

func read(path string, pass string) (*x509.Certificate, *rsa.PrivateKey, error) {
  b, err := ioutil.ReadFile(path)
  if err != nil {
    return nil, nil, err
  }

  privateKey, cert, err := pkcs12.Decode(b, pass)
  if err != nil {
    return nil, nil, err
  }

  priv, ok := privateKey.(*rsa.PrivateKey)
  if !ok {
    return nil, nil, errors.New("expected RSA private key type")
  }

  return cert, priv, nil
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage:", filepath.Base(os.Args[0]), "pfx_cert_path [password]")
    os.Exit(1)
  }
  var file string = os.Args[1]
  var pass string = ""
  if len(os.Args) > 2 {
    pass = os.Args[2]
  }
  cert, priv, err := read(file, pass)
  if err != nil {
    fmt.Println(err)
    os.Exit(2)
  } else {
    fmt.Println(cert)
    fmt.Println(priv)
  }
}
