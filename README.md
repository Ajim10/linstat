# Linux Stats (linstat)

Ein kleines **persönliches Lernprojekt** in Go, um Linux- und Raspberry Pi-Systeminformationen auszulesen.  
Dieses Projekt dient **nur zu Lernzwecken** und enthält möglicherweise noch Fehler oder unvollständige Implementierungen.


## Installation
```go
import (
    "github.com/ajim10/linstat"
)
````

## Usage
```go
package main

import (
	"fmt"

	"github.com/ajim10/linstat/sysinfo/cpu"
	"github.com/ajim10/linstat/sysinfo/memory"
	"github.com/ajim10/linstat/sysinfo/sysos"
)

func main() {
	cpu, err := cpu.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("CPU:")
	fmt.Println("Temperature: ", cpu.Temperature)
	fmt.Println("Cores: ", cpu.Cores)

	fmt.Println()

	mem, err := memory.MemStat()
	if err != nil {
		panic(err)
	}
	fmt.Println("Memory:")
	fmt.Println("Available: ", mem.Available)
	fmt.Println("Buffers: ", mem.Buffers)
	fmt.Println("Cached: ", mem.Cached)
	fmt.Println("Free: ", mem.Free)
	fmt.Println("Total: ", mem.Total)
	fmt.Println()

	sys, err := sysos.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println("System:")
	fmt.Println("Architecture: ", sys.Architecture)
	fmt.Println("OS: ", sys.OS)
	fmt.Println("Kernel: ", sys.Kernel)
	fmt.Println("Distro: ", sys.Distro)
	fmt.Println("Distro Version: ", sys.DistroVersion)
	fmt.Println("Kernel Version: ", sys.KernelVersion)
	fmt.Println("Build Date: ", sys.BuildDate)
	fmt.Println("Build Number: ", sys.BuildNumber)
	fmt.Println("OS Release: ", sys.OSRelease)
}
```

## Projektstruktur

linstat/ # Root des Projekts
│
├─ sysinfo/ # Package für Systeminformationen
│ ├─ cpu/ # CPU-bezogene Funktionen
│ │ └─ cpu.go
│ │
│ ├─ memory/ # Arbeitsspeicher-Funktionen
│ │ └─ memory.go
│ │
│ ├─ pi/ #raspberry pi spezifisch
│ │ └─ pi.go
│ │
│ └─sysos/ # Betriebssystem- und Kernel-Infos
│ └─ sysos.go
│
├─ main.go # Beispielprogramm, das alle Infos ausgibt
└─ README.md # Diese Dokumentation

### Erklärung
- **sysinfo/cpu** → liest CPU-Daten (Temperatur, Anzahl der Kerne, etc.)  
- **sysinfo/memory** → liefert Infos über RAM (Total, Free, Buffers, Cached)
- **sysinfo/pi** → liefert Infos über RaspberryPI (Model) 
- **sysinfo/sysos** → liefert Systeminfos (OS, Kernel, Distro, Build-Date etc.)  
- **main.go** → zeigt, wie man die Packages verwendet  
- Ziel: Lernprojekt für **Linux & Raspberry Pi**, experimentell
