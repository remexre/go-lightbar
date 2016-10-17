# go-lightbar

Chromebook Pixel Lightbar Control from Go

## Example

```go
lb := lightbar.Get()
lb.SetLED(4, 0xff, 0xff, 0x00) // Make all the LEDs green.
```

See the [cmd](https://github.com/remexre/go-lightbar/tree/master/cmd) directory for more (usable!) examples.

## Useful Information

[sysfs interface documentation](https://docs.google.com/document/d/1TXWMxkhWm72ZSdLlycWMGpUXU-IzBnra4Hlly9NWyBc/edit#)
