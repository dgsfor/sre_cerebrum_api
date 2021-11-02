package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)


const (
	privateKeyData = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxfTT54UKukVooHosjEDkQ2kCafvDwyT2SRkuVgKZ9X0zuzJ2
u9RlSfybIkqWbtKNFiWkEuQk24mrXKVRu7pQzZZJs99WjmjRgULAoMmAg/ODQN6K
kbq65wJZ++m9Tj26myav3HPOPyRZTELjWgq4MVN0hkjl0mHQEY0r7MImvT9wj651
aeXjsNVNkD575LkZNk16XWRdeRQdNcOE9oejFSsP7jgmj2zc1k1YUmQKKfuiJXMi
QyTiBO43b9aJ8iprc1jvXF4hrtpmem7KIVTKyT1CXK1vTXT++AdUa3nMII9Ol2M7
Szk04AwEZeCTKlWEBDHyRTkPMkuSeBu/XHbeNwIDAQABAoIBAC/P22Ke8qFc5FFm
UN4rSjax5UBd68F1avrq1xM1G6R8cgMzxBPH0BMXrQySQVVRC3ye6MsbSX+w96+v
ylbyQFP3iaOlPM22qWt0CPyMzrqQFVKUrZlXJY9oNP2wTeXY6PpSVMWFPvpnOB5A
RZNY5XYcdxd8Ae5Xa57hnwtXNW4duLzKm8vOZIyCQh13xA0QMV1+g5SlvKhuqnQl
zrEx7vD/376Y4zXtPn+OC/IQmjuxMrKPz1K7THy+LL1RHJzU47upwwA67BBI0dH0
l/qxfWtvFcdrtb0vy3daLzdh+4reeA3WrNuaRd+aUjKn7Z0xYBTA3YPywiRPn1ER
94mtvZkCgYEA46KUtqBvpn0j2WtAgDvZKd57akxAlof7Bsd8Wmz4EodNUWeG0i0U
PR70Fo9paNhSapI0amlI8Pk4SuoZrCrT/x+ukWrwfr4QD7Mxn3mtItjM/jG7dAEb
/Fcc/YDfXfrRfTxOyDS5Z59tpJVEZn8CKCff1rxehbv3SvhuKYU7Id0CgYEA3p+D
idzZnY+0INe0Ofn2IiF/d4GA6V78vQvk83gAzgMXCFj6XosxAnyuUWYjVRdyrQre
CethZXkwqwOxvZRWI5rRaeUbFeTTCcyV4DDCvmNrWKaL4T7VsT6S7qg14oeWRsQD
gV0wSdFUhB+QWQeG9kK+hqZ+esrWDmW7c6g24SMCgYEAn9IdMV/RJAZGbjcUETRz
0yamZcplqoqMX4UdtOs+/NtMGLRj/xqxU/xwiYplCryemL5P2MpU8QM5V7enEK79
KulWijiX45ST+EHqvDoUqhunSW0nXFoPCQssinu51+Tqwk/eTEhgR/WyaFR4srHH
lOPyD3UvtBt1k4h65+Hn020CgYAtPJKY//SETMAYcFxh74XxN4dUPE96nF+U9obG
ZlsbgYs7mgD+FyCVVcqCJdQGJWaa/PtTXq4C1ZZrGawTTsIWKA5nizNuxnmPoy/N
R3s6/q7Aa88+rOr1QjbPIgUZC94gcRIZRltalJus4VEOljX5cmSOreRDhVXrxQZL
ISJNqwKBgGp+fLI/55bKXOZgAQP2eqpP9cMzozphX2x23k96b1Dtje2MgAwpluOw
fYm4QOZOydEJ2eF7Pqh9EIO9W9wDsSI3xbTUfXA8hwUzGtjYzT7plQ1F4XQ+NXuR
jQrJbNjQLvegDcOy3EZx+rX2cVPGXIpfc1RPDt6zsG8UNesfqd0K
-----END RSA PRIVATE KEY-----
`
	publicKeyData = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxfTT54UKukVooHosjEDk
Q2kCafvDwyT2SRkuVgKZ9X0zuzJ2u9RlSfybIkqWbtKNFiWkEuQk24mrXKVRu7pQ
zZZJs99WjmjRgULAoMmAg/ODQN6Kkbq65wJZ++m9Tj26myav3HPOPyRZTELjWgq4
MVN0hkjl0mHQEY0r7MImvT9wj651aeXjsNVNkD575LkZNk16XWRdeRQdNcOE9oej
FSsP7jgmj2zc1k1YUmQKKfuiJXMiQyTiBO43b9aJ8iprc1jvXF4hrtpmem7KIVTK
yT1CXK1vTXT++AdUa3nMII9Ol2M7Szk04AwEZeCTKlWEBDHyRTkPMkuSeBu/XHbe
NwIDAQAB
-----END PUBLIC KEY-----
`
)

func SignData(data []byte) (signature []byte, err error) {
	hashed := sha256.Sum256(data)
	keyByts,_  := pem.Decode([]byte(privateKeyData))
	privateKey, err := x509.ParsePKCS1PrivateKey(keyByts.Bytes)
	if err != nil {
		return nil, err
	}
	privateKey.Precompute()
	return rsa.SignPKCS1v15(nil, privateKey, crypto.SHA256, hashed[:])
}

func VerifyData(data, signature []byte) error {
	hashed := sha256.Sum256(data)
	public, _ := pem.Decode([]byte(publicKeyData))
	pub, err := x509.ParsePKIXPublicKey(public.Bytes)
	if err != nil {
		return err
	}
	publickey,_ := pub.(*rsa.PublicKey)
	return rsa.VerifyPKCS1v15(publickey, crypto.SHA256, hashed[:], signature)
}