package ca

import (
	uuid "github.com/satori/go.uuid"
	"github.com/globalsign/mgo/bson"
	"time"
)

type Mock struct{}

func (m *Mock) GenerateCert(CN string) (*Bundle, error) {
	return &Bundle{
		PrivateKey:  []byte(dummyKey),
		Certificate: []byte(dummyCert),
		Fingerprint: dummyFingerPrint,
	}, nil
}

func (m *Mock) StoreCert(cert *CertModel) (*CertModel, error) {
	cert.MID = bson.NewObjectId()
	return cert, nil
}

func (m *Mock) GetCertByFingerprint(fp string) (*CertModel, error) {
	b := NewCertModel(&Bundle{
		PrivateKey:  []byte(dummyKey),
		Certificate: []byte(dummyCert),
		Fingerprint: dummyFingerPrint,
	})

	b.Expires = time.Now().Add((365 * 24) * time.Hour)
	b.MID = bson.NewObjectId()
	b.UID = uuid.NewV4().String()

	return b, nil
}

func (m *Mock) GetServerCertByLinkedAPIID(serviceID string) (*CertModel, error) {
	return m.GetCertByFingerprint("")
}

var dummyFingerPrint = "C50181444CF40FEC3BED58362AB9FC3473525BA3F481D0A3494C72291654D14A"

var dummyCert = `
-----BEGIN CERTIFICATE-----
MIIFkjCCA3oCAQEwDQYJKoZIhvcNAQELBQAwgZMxCzAJBgNVBAYTAkFVMRMwEQYD
VQQIDApTb21lLVN0YXRlMQ8wDQYDVQQHDAZTeWRuZXkxEjAQBgNVBAoMCUZvb2J5
IEx0ZDETMBEGA1UECwwKRGluZyBEb25nczEZMBcGA1UEAwwQRnVua3kgRnJ5IExl
ZGdlcjEaMBgGCSqGSIb3DQEJARYLZm9vQGJhci5jb20wHhcNMTkwOTAzMjM1MTA1
WhcNMjkwODMxMjM1MTA1WjCBiTELMAkGA1UEBhMCQVUxEzARBgNVBAgMClNvbWUt
U3RhdGUxDTALBgNVBAcMBEJlZXAxETAPBgNVBAoMCEJvb3AgTHRkMQ8wDQYDVQQL
DAZGbyBCYXIxFTATBgNVBAMMDEZvb2J5IExlZGdlcjEbMBkGCSqGSIb3DQEJARYM
YmluZ0Bib25nLmNtMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxwo3
iIDo1Jgq9P4d4RC7EIeGukViKowEmLuXwvEm9PviXxjKK5EjZAvPBmO0dH/iAsaU
I2avqFUq1MyoK+FyNDSHI1WtBlygD15/u25v1aMCNnuA/b2bGJu9caVFuf7od2js
lnzSAb27AjWiyV4F/NLbz4CMClX200H9BImzZctasJk4TWL+edbEeFOfm6c+qVFp
xjvF9H91zM0OnX69e11xGHv5ZkEoo271DD5WCXr1Bfa1IAl62Xpf4GBfXMzgD9NS
lGKIs4fs3qPI3fP/nkqZGDRw+GqQqcae8gDaKOJuNnddrZtiYUZk1q3T6HCXola/
jrnW9xG+JwdHqBwnLOQOn7bm9UMHHv0xogbcsLKFl6ya6Gc9qZ9mNXPJudg5XiB2
ZHlGgxALmlaHiRFnva6x3x2SWxis8Ck3kPgO4CMfpSxiIYK+V/AAi9xKZ/zkWxn3
mA5NC22G3HMl8vvZv4hfkAzgPggq+Cnog7b7SlNYS8+s1ARdyWaFvvyCMsXr/TL3
B58/zD5j0XSx4NmI8iSwK0LSRW4PFoSq9NyloKUNkso2oZ7IdBA20Ogzx89H6yu6
qnU819mOuLs6lkNOvJ54PnfqjCpwCB7wAoWW++jFLOYdw0Kq0V+ZikoJx7hv3Z0V
JbqGMe5HfgJMmJgtz6Q5hSY32CHnqo7M48OxdN0CAwEAATANBgkqhkiG9w0BAQsF
AAOCAgEAVd9gvdlcZzba5hlikOdTmsZgk0HDRHdsxlkPfhVFzj9y17dN+zQ9BLQe
cmy3ohTIYxaQn3owpmhWEqoPf/eOiDaFKr5MwyoMB0XBBazlylSGjQVrfg2zzgqS
Ok78NnT+dDp/HpZhiWVrpzgYJTgQvUcxSjt+KSVxpeQPzgE/A+odi4yCcuQb+FMR
m0GL9roK9Mcm8GjzdEnEPMkTf16n5ArlEoFhz2WwLvz3uCHYIext0GUQcZB+3O5P
KLRlIVupHnuhsAV71dQnomQPB/tkooH4E/SgjbnR6hsvIcos/ndAZuiPIh/NFrpy
4vUgKz46ACleLnaiIJQN22u3F3DPLRym6rO4Vm3de5E3Pctkm3qSCbB/qaDxHt9e
X668UPRwfxNorfjreFuP8k/JQFEtuvUXwVg0xvSMc37LuMGlsy9lGn/pNtNPyM7B
g02FHA8HBNYg32p2UhOIljGnvvcacaUanGvv3kHeru020gu00kVx6x8qYJseYQLy
U7k0Ebvl0ebTe+cEl4G4i4oNP5UTh7VTjGahHTfLGHrPFeBngvglm84fS3u80Ub5
OaHEpA0qmu4adn8GDAb2K4G6XzG+rZPMgzI2pilm5KncaJQPEDbJxMjdEVRD7IfU
x4d+D3yasQzxO+T/I/rijIREKeQgvYPGacmi0tKkb204yTN2p7U=
-----END CERTIFICATE-----
`

var dummyKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIJKQIBAAKCAgEAxwo3iIDo1Jgq9P4d4RC7EIeGukViKowEmLuXwvEm9PviXxjK
K5EjZAvPBmO0dH/iAsaUI2avqFUq1MyoK+FyNDSHI1WtBlygD15/u25v1aMCNnuA
/b2bGJu9caVFuf7od2jslnzSAb27AjWiyV4F/NLbz4CMClX200H9BImzZctasJk4
TWL+edbEeFOfm6c+qVFpxjvF9H91zM0OnX69e11xGHv5ZkEoo271DD5WCXr1Bfa1
IAl62Xpf4GBfXMzgD9NSlGKIs4fs3qPI3fP/nkqZGDRw+GqQqcae8gDaKOJuNndd
rZtiYUZk1q3T6HCXola/jrnW9xG+JwdHqBwnLOQOn7bm9UMHHv0xogbcsLKFl6ya
6Gc9qZ9mNXPJudg5XiB2ZHlGgxALmlaHiRFnva6x3x2SWxis8Ck3kPgO4CMfpSxi
IYK+V/AAi9xKZ/zkWxn3mA5NC22G3HMl8vvZv4hfkAzgPggq+Cnog7b7SlNYS8+s
1ARdyWaFvvyCMsXr/TL3B58/zD5j0XSx4NmI8iSwK0LSRW4PFoSq9NyloKUNkso2
oZ7IdBA20Ogzx89H6yu6qnU819mOuLs6lkNOvJ54PnfqjCpwCB7wAoWW++jFLOYd
w0Kq0V+ZikoJx7hv3Z0VJbqGMe5HfgJMmJgtz6Q5hSY32CHnqo7M48OxdN0CAwEA
AQKCAgAD//OPva9wHkK2u6iUDkcEFQUjFyPa4Qdynkp1c20p/SnWaWQergL9DrOx
WC2XLb63+Y+eioC7LEL5FcfHYfdujyOsFRuuBgx9YpPEi3qZ23W+7IpyDu5f9zk+
YGjfTP0U8TnX6Fg8CrkIWVWVIFuUchRSvi2SQ5n7MtdH0RCW2hhPdjTIcVXV0vgN
xvhtI1ZCBs3OWwMmX61ZWUGS7qoWIZGVQWCIqoSLjD5GabZG6H6mCQgoxaf/WNM2
ahlgfwl9p+x/6N8AZJng+3+c3jtjzfH3p44hn9qaXtBw4o1/xtUaKRnr1DSM6eI6
4jFg+WpCq+Wbk6Z7PpMIq7/H5BDxb+ukiOf+0xcmNN/xd6MJ3LqkF/s5SuwzYT9T
lf3s780Lgt5mVWvZYbhhGVr/NPP/h+yhjhSmoStJMBykEV61DIhwm13ClrNam2bC
dOMCfhRmrOjzkz8iK8M/D83Desk1J2zeOCKtQG2/uYyeIl5DbxpDyIympAKea1mf
mYcQUNfRbfONmrQnXGwsoUkekWf0kG2/BmCECl3VwwZc5k7cnPZgvagYRkTjr7K1
Tfu37DuUjbUpvw0jOQ/XIhiE91st6Gba4Qs6WdqbwNGPEL3QAXlcwH9FC+HnK838
XNG2twOhxqp/k+PiVWBUwqiNX+9BQ3dyUXFMjtuDNzybFtUGIQKCAQEA6k0K7NIZ
cGLX/AnxTKwiUe9Jy2r+iBCJbplO5GvDqE7eDubxFIQn4WvKyBEifJ/QDK7yCOtC
zkPogmtIidWTdzHTrSU3Dj50h4JxctP4sWuteQGvuTJhkr3M7urtF2yS6t/q7DeG
Ss3AmBw8rnXswSBzi8p9UiHox3ZRsF/UCXC91cm6ePFoy/6cRQTURlyRSP1kqPfU
bBR+ZGxHeAlhp/qeq3EuQmoAcjsaDE6sOq66YseLkoSKuAEMGIENMARkAnzTNkSp
PqwAjiQW8s1UaUnnOa81qp7aQK7iAytC/v09zTe3TXV5rmiiXxYYbg27UdbfIt4l
i4FKe44wMCzttQKCAQEA2Xku07KN2/oi2E5OQfa6WkTkmTRUwMHEvo+z9YMaM2se
s02EJSk43EDLIyU8BXtxPnTvRl/QNdKCbFKQvP3ui19PTKCkY2e5qe2C3Qdej4aG
4u2mNE90OcG1UF1flu3dxJCgEq8Om4qsI+fw7NduzMkR1o63wh9qDgspCk3xH5b9
tTwpI6HD8Tjg7gYaqZ6uJbJ+RGozzO/g2gyKqal6sqC3E+r6xhSYb3MQVZK1WG/U
NBvaOjbKRhW0CNenCMP6OLMN+R9+BW6UBSuUCyIlmHwzPBNIXVp8ImKO2c55spO1
uBjdV5f3Lan1r8lyRNPX7GQ+0XgFXQsyP/y+LpyjiQKCAQA+KtUEQzbmwANI56Zz
zpSIovI0nB4PIjwk8V6icwmGJe77i0FaTTp7sGssrIc3A/xi8SRdQ3cFXbVarXag
w9+N19WvcllqWR34Op85dY7eHJD5s33ZqMmE8wFmP0VLWs6crW8a4Bysym0Yrx1W
uQlpsN+XrtGTslWeXwGLx9Ft89Ea4ZytrHg+8D6sfXyJctRFxp4nyyI8zy8+HBDW
g766oF/rY2t/ybp2lMHzlAhUxbnEreeVp79a1URNWCiUngWfN1mfY5Z4f5DiJidE
o0CLc0V748ZckCue0Ag+CQOMvsrgIbFsRwhTQ4YZrrYqtVAHhfjMup7BpjuKDw8w
KSulAoIBAQC45N3aIJmZZ+5OJy1DyeggLSiJ1MMl6thCB91BpAgXpUBBw2rRSV3N
bM0IgXSpWI0sQ7DXDUjxKZu3Su9Br0creYnrJTz+QsbIYWV8U/KiZgayUsiW+uJe
5IEp5WkK4gtFPozJtcneliTKljzwsCzEsUKYP1ieE6VQcU6gMyoI6I+lkZOMdtEn
emtKEyiCA6Yd6MaPlheaamuqFzWhD7WAv0FIDuTMmeAMCTDfAE9r9A/fibqc+c45
jeQ2DFs4CZ7oq9r6bjrvlVcFmkpQuORr41SQA/jnXbHibhbzuScgm5LJBWLSQSJm
0XRP0x5yxAV1Nrltz/QHaih73H85E/W5AoIBAQDWR6M4T/ta3LhzQF9P87LH1Usb
kO2HjulONP/BQXGLR6eYVKcjDhmzCChh6itV6WLwLZzM1qROVKTR7fGiMxyprY9L
k5RTPYY21w8yeBjBNi8/UsKwMHIAcnNchk2I2ZyJjO7NDP2nm4pcwEqTP+f4C6GD
jBqB/P25J8IwJIDRzlmDj5g9/EcKC1Q8lJ5oAqJQ2C7DbEq009xMYnXsr4u3pIDS
UKPDed4kRBkeYpXEomYiq2/CGRx1G+o8Gw3eJE1vMnie5/wD9abg7luMzERXkhoF
ZlRXMSSSP0xTI05a6JEpK36ISHIfjsHCcOc1FNedvsannk8bW57AQ80oV2PB
-----END RSA PRIVATE KEY-----
`