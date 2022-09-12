build_mac_intel:
	GOOS=darwin GOARCH=amd64 go build -o out/mac_intel/tt

build_mac_m1:
	GOOS=darwin GOARCH=arm64 go build -o tt out/mac_m1/tt