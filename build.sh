cd webapp && npm run build && cd ..
rm a_main-packr.go
packr
go build -o dist/neutrino.app/Contents/MacOS/neutrino
open dist/neutrino.app
