module sup

go 1.20

require (
	github.com/brutella/dnssd v1.2.7
	github.com/brutella/hap v0.0.26
	golang.org/x/sys v0.8.0
)

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
)

require (
	github.com/go-chi/chi v1.5.4 // indirect
	github.com/kr/pretty v0.3.1
	github.com/miekg/dns v1.1.54 // indirect
	github.com/tadglines/go-pkgs v0.0.0-20210623144937-b983b20f54f9 // indirect
	github.com/xiam/to v0.0.0-20200126224905-d60d31e03561 // indirect
	golang.org/x/crypto v0.0.0-20220131195533-30dcbda58838 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.9.1 // indirect
)

replace (
	github.com/brutella/dnssd => ../dnssd
	github.com/brutella/hap => ../hap
)
