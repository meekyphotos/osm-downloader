
if [ "$GOOS" = "windows" ]; then
  go build -o osm.exe cmd/cli/main.go
else
  go build -o osm cmd/cli/main.go
fi
