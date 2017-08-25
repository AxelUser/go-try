# Currency converter

## Build

To build executable run __`build.bat`__. Executable named __`curconv.exe`__ will be placed in __`bin`__ folder.

## Run app

To run currency converter you can provide 2 flags:
* __`--currency`__ (string) - Currency code (case-insensitive), default is `RUB`.
* __`--value`__ (float64) - Value to convert, default is `1`.

### Example:
* `curconv` - Converts 1 Russian Ruble to other currencies.
* `curconv --value=500` - Converts 500 Russian Rubles to other currencies.
* `curconv --currency=USD` - Converts 1 US Dollar to other currencies.