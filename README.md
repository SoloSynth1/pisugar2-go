# pisugar2-go
A minimalistic, daemon-free CLI utility to manage your PiSugar2s &amp; PiSugar2 Pros.

⚠️ work in progress ⚠️

## But why?

While I admire all the hard works the amazing folks at PiSugar has put up to shell out a fully-featured, fully-fledged power manager,
sometimes however, a simple CLI utility script is all you need...

- `dpkg`/`rpm` not required - works on `ArchArm` and (probably) other distros
- Daemon-free - Drains power only when you run the commands
- Get/Set values from CLI and CLI only - no more unrestricted access via webpage

## Installation

Firstly, please ensure I2C interface is enabled. (i.e. /dev/isc-x is present)

Then, git clone the repo and build the binary from source:
```bash
git clone https://github.com/SoloSynth1/pisugar2-go.git
cd pisugar2-go
go build .
./pisugar2-go
```

I might release built binaries as the project leaves infant stage.

## TODOs
- [ ] Model PiSugar2 (IP5209)
  - [ ] Read battery status
    - [ ] Display voltage
    - [ ] Display battery percentage
    - [ ] Display ampere
  - [ ] Read charging status
  - [ ] Charging control
- [ ] Model PiSugar2 Pro (IP5312)
  - [ ] Read battery status
    - [x] Display voltage
    - [x] Display battery percentage
    - [ ] Display ampere
  - [ ] Read charging status
  - [ ] Charging control
- [ ] RTC module (SD3078)
  - [ ] Read time
  - [ ] Write time

## Checkout these awesome repos also
- [pisugar-power-manager-rs](https://github.com/PiSugar/pisugar-power-manager-rs)
- [pwnagotchi-plugin-pisugar2](https://github.com/kellertk/pwnagotchi-plugin-pisugar2)
- [go-i2c](https://github.com/d2r2/go-i2c)
