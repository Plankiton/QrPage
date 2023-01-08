if ! air -h > /dev/null 2>&1
then
  go build -o ./bin/server ./cmd/server && ./bin/server
else
  air -build.cmd "go build -o ./bin/server ./cmd/server" -build.bin "./bin/server"
fi
