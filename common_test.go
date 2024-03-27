// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubeauth

const (
	testLocalCACert = `-----BEGIN CERTIFICATE-----
MIIDVDCCAjwCCQDFiyFY1M6afTANBgkqhkiG9w0BAQsFADBsMQswCQYDVQQGEwJV
UzETMBEGA1UECAwKV2FzaGluZ3RvbjEQMA4GA1UEBwwHU2VhdHRsZTEgMB4GA1UE
CgwXVmF1bHQgVGVzdGluZyBBdXRob3JpdHkxFDASBgNVBAMMC2V4YW1wbGUubmV0
MB4XDTIwMDkxODAxMjkxM1oXDTQ1MDkxODAxMjkxM1owbDELMAkGA1UEBhMCVVMx
EzARBgNVBAgMCldhc2hpbmd0b24xEDAOBgNVBAcMB1NlYXR0bGUxIDAeBgNVBAoM
F1ZhdWx0IFRlc3RpbmcgQXV0aG9yaXR5MRQwEgYDVQQDDAtleGFtcGxlLm5ldDCC
ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBALCA9oKv+ESRHX2e/iq1PlGr
zD23/MBS0V+fWQDY0hyEqY98CGwRtF6pEcLEYsreArj5/zznsIevLkNOD+beg43y
WpEJlCPgDhGXI/Oima6ooHVEIMaIKLjK7GrSzAb3rNRGACwrR/u/IKaFl+XJG0qx
g8mOZ3fByaAlIk+shVLUcIedNN1tNR+6/4ZpHg7PDjrZXP4XKrmKPTh4yqfu+BtZ
9IY2oyregqEsGW1/3h1NM+LHGVakTV2d/mwMYHhwoq9Y8BD+PemT5z8TmhH/cIk5
P8Q8ud5/q6YTIJg9TELKebLAeNtRNnNoHeUoRTjiW1MBwNHtgyTTY+H3W/9Dne0C
AwEAATANBgkqhkiG9w0BAQsFAAOCAQEAXmygFkGIBnXxKlsTDiV8RW2iHLgFdZFJ
hcU8UpxZhhaL5JbQl6byfbHjrX31q7ii8uC8FcbW0AEdnEQAb9Ui6a+if7HwXNmI
DTlYl+lMlk9RtWvExw6AEEbg5nCpGaKexm7wJgzYGP9by9pQ7wX/CS7ofCzCK+Al
uSIqjPkMC201ZXH39n1lxxq6BacdYjv8wo4mMzi8iTSQGVWPdjHZVYOClFgN6hoj
8SkrrSe888a0H+i7EknRxC4sLRaMUK/FAvwtXaSZi2djruAtQzQGQ56m1phC2C/k
k9aL00AQ9Y4KTfiJD7LK8YIZDnFKLOCJhYgKCLCOVwOHb7836SNCxA==
-----END CERTIFICATE-----`

	testLocalJWT = `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IlZhdWx0IFRlc3QiLCJpYXQiOjExMjM1OH0.GOC8w-MyhorgojB20SPNyH_ECsBjYJH89hjntOxSywA`

	testRSACert = `-----BEGIN CERTIFICATE-----
MIIDcjCCAlqgAwIBAgIBAjANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p
a3ViZUNBMB4XDTE3MDgzMDE5MDgzNloXDTE4MDgzMDE5MDgzNlowLDEXMBUGA1UE
ChMOc3lzdGVtOm1hc3RlcnMxETAPBgNVBAMTCG1pbmlrdWJlMIIBIjANBgkqhkiG
9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxD3eM3+WNc4phxAeQxNOmcybKlNJWowuC12u
v+cGJWxxpDx/OoEIxKI5wmgHxEwFCZL545sjfLqyBcgxQR2xSCib+bYzjBtfA6uV
6d/35nurzz21okcMffc5xKMyZhEwt98WAvYWD71Bihz7iGBq5Sw9md6pqnkNoScR
Hhi3Vl94a6D6shwb6nXA2hlwYLcnoKtpe3Ptq6MW6CpfBA8C11q5eeW4xdvrwKt3
Vd1TgFeEnnqwzUWGapU2uwwUfbRkLTDvrp6791uq0Vo7mzz00xYhV1PLCeAdpJEK
3Vr74FT7jHIbPlzi/qjRBVFKf9IRXnhbjrCl7S0Ayev1Fao4TQIDAQABo4G1MIGy
MA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYBBQUHAwIw
DAYDVR0TAQH/BAIwADBzBgNVHREEbDBqgiRrdWJlcm5ldGVzLmRlZmF1bHQuc3Zj
LmNsdXN0ZXIubG9jYWyCFmt1YmVybmV0ZXMuZGVmYXVsdC5zdmOCEmt1YmVybmV0
ZXMuZGVmYXVsdIIKa3ViZXJuZXRlc4cEwKhjZIcECgAAATANBgkqhkiG9w0BAQsF
AAOCAQEAIw8rKuryhhl527wf9q/VrWixzZ1jCLvyc/60z9rWpXxKFxT8AyCsHirM
F4fHXW4Brcoh/Dc2ci36cUbuywIyxHjgVUG45D4jPPWskY1++ZSfJfSXAuA8eFew
c+No3WPkmZB6ZOZ6q5iPY+FOgDZC7ddWmGuZrle51gBL347cU7H1BrTm6Lm6kXRs
fHRZJX2+B8lnsXsS3QF2BTU0ymuCxCCQxub/GhPZVz3nNNtro1z7/szLUVP1c1/8
p7HP3k7caxfp346TZ/HgbV9sJEkHP7Ym7n9E7LSyUTSxXwBRPraH1WQzEgFNPSUV
V0n6FBLiejOTPKapJ2F0tIqAyJHFug==
-----END CERTIFICATE-----`

	testECCert = `-----BEGIN CERTIFICATE-----
MIICZDCCAeugAwIBAgIJALM9NbK8WRuBMAkGByqGSM49BAEwRTELMAkGA1UEBhMC
dXMxEzARBgNVBAgTClNvbWUtU3RhdGUxITAfBgNVBAoTGEludGVybmV0IFdpZGdp
dHMgUHR5IEx0ZDAeFw0xNzA5MTExNzQ2NDNaFw0yNzA5MDkxNzQ2NDNaMEUxCzAJ
BgNVBAYTAnVzMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5l
dCBXaWRnaXRzIFB0eSBMdGQwdjAQBgcqhkjOPQIBBgUrgQQAIgNiAATcqsBLxKP+
UHk7Y6ktGGFvfrIfIXHxeZe3Xwt691CWfdmJFvrGzyzW5/AbJIuO1utdOsqUStAm
W/Scfxop/FGadKqR4nAWLNBI4intgnf0r1rzBCSOmanolHqxQPqQ0UOjgacwgaQw
HQYDVR0OBBYEFHxh1pTd8ApEzg0gKMwwt01aA10TMHUGA1UdIwRuMGyAFHxh1pTd
8ApEzg0gKMwwt01aA10ToUmkRzBFMQswCQYDVQQGEwJ1czETMBEGA1UECBMKU29t
ZS1TdGF0ZTEhMB8GA1UEChMYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkggkAsz01
srxZG4EwDAYDVR0TBAUwAwEB/zAJBgcqhkjOPQQBA2gAMGUCMCR+CvAoNBhqSe2M
4qWWD/9XX/0qmf0O442Qowcg5MWH1+mwl1s7ozinvbTPDPaYDwIxAM54qKhuL6xt
GxqJpa7Onn15Hu8zTsdzeYBqUUXA6wtn+Pa7197CgUkfty9yc2eeQw==
-----END CERTIFICATE-----`

	testCACert = `
-----BEGIN CERTIFICATE-----
MIIC5zCCAc+gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p
a3ViZUNBMB4XDTE3MDgxMDIzMTQ1NVoXDTI3MDgwODIzMTQ1NVowFTETMBEGA1UE
AxMKbWluaWt1YmVDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAN8d
w2p/KXRkm+vzOO0eT1vYBWP7fKsnng9/g5nnXAJlt9NxpOSolRcyItm/04R0E1jx
jpgsdzkybc+QU5ZiszOYN833/D5hCNVAABVivpDd2P8wVKXN46cB99e24etUVBqG
5aR0Ku3IBsJjCN9efhF+XRCA2gy/KaXMdKJhHfdtc8hCr7G9+2wO2G58FLmIfEyH
owviOGt0BSnCtMpsA8ZgGQyfqgSd5u466aCv6oj0MyzsMnfS38niM53Rlv4IY6ol
taYbWXtCNndQ2S687qE0qTCxhE95Bm2Nfkqct4R1798sJz83xNv8hALvxr/vPK/J
2XkIm3oo3YKG4n/CHXcCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQW
MBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
DQEBCwUAA4IBAQCSkrhE1PczqeqXfRaWayJUbXWPwKFbszO0MhGB1zwnPZq39qjY
ySQiGvnjV3fP+N5CTQAwMNe79Xiw31fSoexgceCPJpraWrTOLdCv04SbGDBapMFM
aezBu9jzZm0CNt60jHXWXuHHVPFX6u7ZR8W+RiBvsT8GZ5U6sNs3aN3M9Vym06BL
aSphIw1v+hRlPfnrlJwUnQp158DRgkt/9ncTa/k88KoIoZAbulaiGB4zHxxkbura
GSlgpZzhHSrBDLuXf65GHwwGxSExhgY5AA/n8rumGVvE8IYohS9yg/jOG0xP2WQH
u/ABoYtOyseO+lgElA8R4PB9MtwgN6c/b0xH
-----END CERTIFICATE-----`

	testOtherCACert = `-----BEGIN CERTIFICATE-----
MIIC/jCCAeagAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTIyMTEyOTIwNDYwNVoXDTMyMTEyNjIwNDYwNVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAOYO
AGi6hysYe4PF3sCyVXAFM72X96F51ncmDg+Ihhvt9Touj7Yo0LLMmy1UPM3YzEdV
z2IV1iksLIBObkct2eBGt2SFjLMhphZdA34mPAzFZhpvNgn1U2uUSqfp5phg00Mg
DdXLE0LFIVqAGkoBysr89l6P2MaTibcKoZwAhpMbATLcn1QcXF/NLzYuO9FPvrUL
mAR/HslDz10LBsMjtgRKXd2dX4yrQlYSB7YmLlu/bLKdjiE1a/+EO4wNcl/JJ+vu
fzPwxWALej/k61mcP+4JxfjgY53AM7vaZ+P9Yb0bGw7r1bozgP7+FGL1f6onTxeG
7+FECpErmrv9IQocgh8CAwEAAaNZMFcwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFE0Y+CYtusQDfr8tqSZZ7PEcDJjwMBUGA1UdEQQO
MAyCCmt1YmVybmV0ZXMwDQYJKoZIhvcNAQELBQADggEBAJK57wm9rH0yVmjmY1ES
kE8e+pTnXZkKaqUce/d7cPRn1+0Dtutvxl/j3P4DUjba7lVGYYNKp1xy2xVvg7Dl
mXyJigvBoTGyzJzNDIUJz8Kgse4eCrwl59WP94K83cVRLeUq+3amLwzubUNbezsW
QwcCyACuzTetR5ZXEg7iIS4HDy+ER2yjuY6d0GPLG+FH02WMrlE7mmxNfZOSy/5E
pEDcN/HcXM47TP7XgWW0rfQli3RucuqMV7LHvvpiGIWwfutrK9g7Py91W2JbQCA0
D14XDzgsruCwlWAP1FMvLMIPhSknpIJd9Xql+0/Ae1yl9f3Uamj3mDtBKg8/U5nJ
0wU=
-----END CERTIFICATE-----
`

	testInvalidCACert = `
-----BEGIN CERTIFICATE-----
MIIC5zCCAc+gAwIBAgIBATANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwptaW5p
a3ViZUNBMB4XDTE3MDgxMDIzMTQ1NVoXDTI3MDgwODIzMTQ1NVowFTETMBEGA1UE
AxMKbWluaWt1YmVDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAN8d
w2p/KXRkm+vzOO0eT1vYBWP7fKsnng9/g5nnXAJlt9NxpOSolRcyItm/04R0E1jx
jpgsdzkybc+QU5ZiszOYN833/D5hCNVAABVivpDd2P8wVKXN46cB99e24etUVBqG
5aR0Ku3IBsJjCN9efhF+XRCA2gy/KaXMdKJhHfdtc8hCr7G9+2wO2G58FLmIfEyH
owviOGt0BSnCtMpsA8ZgGQyfqgSd5u466aCv6oj0MyzsMnfS38niM53Rlv4IY6ol
taYbWXtCNndQ2S687qE0qTCxhE95Bm2Nfkqct4R1798sJz83xNv8hALvxr/vPK/J
2XkIm3oo3YKG4n/CHXcCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMB0GA1UdJQQW
MBQGCCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3
`
)
