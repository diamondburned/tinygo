// +build !avr,!nrf,!sam,!stm32

package machine

// Dummy machine package that calls out to external functions.

var (
	SPI0  = SPI{0}
	I2C0  = I2C{0}
	UART0 = UART{0}
)

type PinMode uint8

const (
	PinInput PinMode = iota
	PinOutput
	PinInputPullup
	PinInputPulldown
)

func (p Pin) Configure(config PinConfig) {
	gpioConfigure(p, config)
}

func (p Pin) Set(value bool) {
	gpioSet(p, value)
}

func (p Pin) Get() bool {
	return gpioGet(p)
}

//go:export __tinygo_gpio_configure
func gpioConfigure(pin Pin, config PinConfig)

//go:export __tinygo_gpio_set
func gpioSet(pin Pin, value bool)

//go:export __tinygo_gpio_get
func gpioGet(pin Pin) bool

type SPI struct {
	Bus uint8
}

type SPIConfig struct {
	Frequency uint32
	SCK       Pin
	MOSI      Pin
	MISO      Pin
	Mode      uint8
}

func (spi SPI) Configure(config SPIConfig) {
	spiConfigure(spi.Bus, config.SCK, config.MOSI, config.MISO)
}

// Transfer writes/reads a single byte using the SPI interface.
func (spi SPI) Transfer(w byte) (byte, error) {
	return spiTransfer(spi.Bus, w), nil
}

//go:export __tinygo_spi_configure
func spiConfigure(bus uint8, sck Pin, mosi Pin, miso Pin)

//go:export __tinygo_spi_transfer
func spiTransfer(bus uint8, w uint8) uint8

// InitADC enables support for ADC peripherals.
func InitADC() {
	// Nothing to do here.
}

// Configure configures an ADC pin to be able to be used to read data.
func (adc ADC) Configure() {
}

// Get reads the current analog value from this ADC peripheral.
func (adc ADC) Get() uint16 {
	return adcRead(adc.Pin)
}

//go:export __tinygo_adc_read
func adcRead(pin Pin) uint16

// InitPWM enables support for PWM peripherals.
func InitPWM() {
	// Nothing to do here.
}

// Configure configures a PWM pin for output.
func (pwm PWM) Configure() {
}

// Set turns on the duty cycle for a PWM pin using the provided value.
func (pwm PWM) Set(value uint16) {
	pwmSet(pwm.Pin, value)
}

//go:export __tinygo_pwm_set
func pwmSet(pin Pin, value uint16)

// I2C is a generic implementation of the Inter-IC communication protocol.
type I2C struct {
	Bus uint8
}

// I2CConfig is used to store config info for I2C.
type I2CConfig struct {
	Frequency uint32
	SCL       Pin
	SDA       Pin
}

// Configure is intended to setup the I2C interface.
func (i2c I2C) Configure(config I2CConfig) {
	i2cConfigure(i2c.Bus, config.SCL, config.SDA)
}

// Tx does a single I2C transaction at the specified address.
func (i2c I2C) Tx(addr uint16, w, r []byte) error {
	i2cTransfer(i2c.Bus, &w[0], len(w), &r[0], len(r))
	// TODO: do something with the returned error code.
	return nil
}

//go:export __tinygo_i2c_configure
func i2cConfigure(bus uint8, scl Pin, sda Pin)

//go:export __tinygo_i2c_transfer
func i2cTransfer(bus uint8, w *byte, wlen int, r *byte, rlen int) int

type UART struct {
	Bus uint8
}

type UARTConfig struct {
	BaudRate uint32
	TX       Pin
	RX       Pin
}

// Configure the UART.
func (uart UART) Configure(config UARTConfig) {
	uartConfigure(uart.Bus, config.TX, config.RX)
}

// Read from the UART.
func (uart UART) Read(data []byte) (n int, err error) {
	return uartRead(uart.Bus, &data[0], len(data)), nil
}

// Write to the UART.
func (uart UART) Write(data []byte) (n int, err error) {
	return uartWrite(uart.Bus, &data[0], len(data)), nil
}

// Buffered returns the number of bytes currently stored in the RX buffer.
func (uart UART) Buffered() int {
	return 0
}

// ReadByte reads a single byte from the UART.
func (uart UART) ReadByte() (byte, error) {
	var b byte
	uartRead(uart.Bus, &b, 1)
	return b, nil
}

// WriteByte writes a single byte to the UART.
func (uart UART) WriteByte(b byte) error {
	uartWrite(uart.Bus, &b, 1)
	return nil
}

//go:export __tinygo_uart_configure
func uartConfigure(bus uint8, tx Pin, rx Pin)

//go:export __tinygo_uart_read
func uartRead(bus uint8, buf *byte, bufLen int) int

//go:export __tinygo_uart_write
func uartWrite(bus uint8, buf *byte, bufLen int) int
