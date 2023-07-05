package main

import (
    "bytes"
    "context"
    "crypto/rand"
    "crypto/rsa"
    "crypto/tls"
    "crypto/x509"
    "encoding/pem"
    "log"
    "math/big"

    "github.com/quic-go/quic-go"
)

func main() {
    tlsConfig := &tls.Config{
        Certificates: []tls.Certificate{tlsCert()},
        NextProtos:   []string{"siduck"}, // ALPN は "siduck" とする
    }
    quicConfig := &quic.Config{
        EnableDatagrams: true, // QUIC DATAGRAM を利用する
    }
    lis, err := quic.ListenAddr("127.0.0.1:55555", tlsConfig, quicConfig)
    if err != nil {
        panic(err)
    }

    for {
        sess, err := lis.Accept(context.TODO())
        if err != nil {
            panic(err)
        }

        go func() {
            for {
                msg, err := sess.ReceiveMessage()
                if err != nil {
                    log.Print(err)
                    return
                }

                // quack でなければエラー （0x101=DISUCK_ONLY_QUACKS_ECHO） を返す
                if !bytes.Equal(msg, []byte("quack")) {
                    sess.CloseWithError(0x101, "SiDUCK only quacks echo")
                    return
                }

                // quack だったら quack-ack を返す
                if err := sess.SendMessage([]byte("quack-ack")); err != nil {
                    log.Print(err)
                    return
                }
            }
        }()
    }
}

// オレオレ証明書を作る
func tlsCert() tls.Certificate {
    key, _ := rsa.GenerateKey(rand.Reader, 1024)
    template := x509.Certificate{SerialNumber: big.NewInt(1)}
    certDER, _ := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
    keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
    certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
    tlsCert, _ := tls.X509KeyPair(certPEM, keyPEM)
    return tlsCert
}
