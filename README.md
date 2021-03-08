# fp93

> Floppy Punk 2193

The base TUI upon which to build FloppyPunk games and supplements;
a cross-platform binary written in Go and leveraging the [tvwiew](github.com/rivo/tview) library.

The total size of this binary after packing with [upx](https://github.com/upx/upx) cannot exceed one megabyte; any larger and it must be trimmed.

## Build & Pack

```powershell
# Will build for your machine first; in this case, on windows
# (drop the exe on *nix systems)
go build -ldflags="-s -w" .\fp93.go
upx .\fp93.exe
```
