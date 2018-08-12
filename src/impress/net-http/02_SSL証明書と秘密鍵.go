package main

import (
  "log"
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "crypto/x509/pkix"
  "encoding/pem"
  "math/big"
  "net"
  "os"
  "time"

  "github.com/comail/colog"
)

func main() {
  colog.SetDefaultLevel(colog.LDebug)
  colog.SetMinLevel(colog.LDebug)
  // colog.SetMinLevel(colog.LInfo)
  colog.Register()
  // 乱数生成時の最大値取得
  max := new(big.Int).Lsh(big.NewInt(1), 128)
  log.Printf("max: %d\n", max)
  // シリアルナンバー取得
  serialNumber, _ := rand.Int(rand.Reader, max)
  log.Printf("serialNumber: %d\n", serialNumber)
  // 証明書の所有者情報
  subject := pkix.Name{
    Organization      : []string{"Manning Publication Co."},
    OrganizationalUnit: []string{"Books"},
    CommonName        : "Go Web Programming",
  }
  log.Printf("subject: %v\n", subject)
  // 照明所の構成を設定するための構造体
  template := x509.Certificate{
    // CAによって発行される一意の番号
    SerialNumber : serialNumber,
    // 識別名
    Subject      : subject,
    NotBefore    : time.Now(),
    // 有効期限
    NotAfter     : time.Now().Add(365 * 24 * time.Hour),
    // サーバ認証に使用されることを示しているらしい
    KeyUsage     : x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
    ExtKeyUsage  : []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
    // 証明書が有効なIPアドレスを設定
    IPAddresses  : []net.IP{net.ParseIP("localhost")},
  }
  log.Printf("trace: template: %v\n", template)
  // 秘密鍵の生成
  pk, _ := rsa.GenerateKey(rand.Reader, 2048)
  log.Printf("trace: pk: %v\n", pk)
  // DER形式のバイトデータのスライス DERはエンコーディングの一種（鍵の中身のことではないよ）
  derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)
  log.Printf("derBytes: %v\n", derBytes)
  // cert.pemファイルを生成
  certOut, _ := os.Create("cert.pem")
  pem.Encode(certOut, &pem.Block{
    Type  : "CERTIFICATE",
    Bytes : derBytes,
  })
  certOut.Close()
  // key.pemファイルを生成
  keyOut, _ := os.Create("key.pem")
  pem.Encode(keyOut, &pem.Block{
    Type  : "RSA PRIVATE KEY",
    Bytes : x509.MarshalPKCS1PrivateKey(pk),
  })
  keyOut.Close()
}
