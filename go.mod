module github.com/MOACChain/xchain

go 1.16

require (
	github.com/MOACChain/MoacLib v1.0.1000
	github.com/beevik/ntp v0.3.0
	github.com/cespare/cp v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/docker/docker v20.10.6+incompatible
	github.com/edsrzf/mmap-go v1.0.0
	github.com/fatih/color v1.11.0
	github.com/gizak/termui v2.3.0+incompatible
	github.com/golang/protobuf v1.3.5
	github.com/hashicorp/golang-lru v0.5.4
	github.com/huin/goupnp v1.0.0
	github.com/innowells/bls_lib/v2 v2.0.999
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/karalabe/hid v1.0.0
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/maruel/panicparse v1.6.1 // indirect
	github.com/mattn/go-colorable v0.1.8
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/naoina/toml v0.1.1
	github.com/nsf/termbox-go v1.1.1 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pborman/uuid v1.2.1
	github.com/peterh/liner v1.2.1
	github.com/prometheus/prometheus v1.7.2
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475
	github.com/rjeczalik/notify v0.9.2
	github.com/robertkrimen/otto v0.0.0-20200922221731-ef014fd054ac
	github.com/rs/cors v1.7.0
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/syndtr/goleveldb v1.0.0
	go.dedis.ch/kyber/v3 v3.0.13
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	golang.org/x/net v0.0.0-20210510120150-4163338589ed
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	golang.org/x/tools v0.1.1
	google.golang.org/grpc v1.25.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c
	gopkg.in/fatih/set.v0 v0.1.0
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
	gopkg.in/olebedev/go-duktape.v3 v3.0.0-20210326210528-650f7c854440
	gopkg.in/sourcemap.v1 v1.0.5 // indirect
	gopkg.in/urfave/cli.v1 v1.20.0
	gotest.tools/v3 v3.0.3 // indirect
)

//replace github.com/innowells/bls_lib/v2 => /media/yifan/ssd/go/src/github.com/innowells/bls_lib
//replace github.com/MOACChain/MoacLib => /media/yifan/ssd/go/src/github.com/MOACChain/MoacLib
