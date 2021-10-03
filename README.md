# pisugar2-go
A minimalistic, daemon-free CLI utility to manage your PiSugar2s &amp; PiSugar2 Pros.

⚠️ work in progress ⚠️

## But why?

While I admire all the hard works the amazing folks at PiSugar has put up to shell out a fully-featured, fully-fledged power manager,
sometimes however, a simple solution is all you need...

- Works on `ArchArm` - `dpkg`/`rpm` not required
- Daemon-free - Drains power only when you run the commands
- Get/Set values from CLI and CLI only - no more unrestricted access via webpage

## Installation
```bash
git clone https://github.com/SoloSynth1/pisugar2-go.git
cd pisugar2-go
go build .
./pisugar2-go
```

## TODOs
- [ ] Model PiSugar2 (IP5209)
  - [ ] Read voltage
    - [ ] Display Voltage
    - [ ] Display Battery Percentage
    - [ ] Display Ampere
  - [ ] Read charging status
  - [ ] Charging control
- [ ] Model PiSugar2 Pro (IP5312)
  - [ ] Read voltage
    - [x] Display Voltage
    - [x] Display Battery Percentage
    - [ ] Display Ampere
  - [ ] Read charging status
  - [ ] Charging control
- [ ] RTC module (SD3078)
  - [ ] Read time
  - [ ] Write time

## Checkout these awesome repos also
- [pisugar-power-manager-rs](https://github.com/PiSugar/pisugar-power-manager-rs)
- [pwnagotchi-plugin-pisugar2](https://github.com/kellertk/pwnagotchi-plugin-pisugar2)
- [go-i2c](https://github.com/d2r2/go-i2c)