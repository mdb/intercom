package intercom

import (
	"bytes"
	"io"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIntercom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "intercom Suite")
}

var _ = Describe("intercom", func() {
	Describe("NewLogger", func() {
		Context("when passed 'silent'", func() {
			It("returns a Logger with the correct Level", func() {
				logger := NewLogger("silent")

				Expect(logger.Level).Should(Equal(silentLevel))
			})
		})

		Context("when passed 'error'", func() {
			It("returns a Logger with the correct Level", func() {
				logger := NewLogger("error")

				Expect(logger.Level).Should(Equal(errorLevel))
			})
		})

		Context("when passed 'warn'", func() {
			It("returns a Logger with the correct Level", func() {
				logger := NewLogger("warn")

				Expect(logger.Level).Should(Equal(warnLevel))
			})
		})

		Context("when passed 'info'", func() {
			It("returns a Logger with the correct Level", func() {
				logger := NewLogger("info")

				Expect(logger.Level).Should(Equal(infoLevel))
			})
		})

		Context("when passed an unrecognized level", func() {
			It("returns a Logger with the default 'info' Level", func() {
				logger := NewLogger("foo")

				Expect(logger.Level).Should(Equal(infoLevel))
			})
		})

		Context("when passed 'debug'", func() {
			It("returns a Logger with the correct Level", func() {
				logger := NewLogger("debug")

				Expect(logger.Level).Should(Equal(debugLevel))
			})
		})
	})

	Describe("Logger", func() {
		var logger *Logger

		Describe("Errorf", func() {
			Context("the configured log level is not less than the error level", func() {
				BeforeEach(func() {
					logger = NewLogger("debug")
				})

				It("prints a red line to stderr", func() {
					_, e := capture(func() {
						logger.Errorf("foo")
					})

					Expect(e).Should(Equal("\033[1;31mfoo\033[0m\n"))
				})

				It("it correctly formats and prints strings", func() {
					_, e := capture(func() {
						logger.Errorf("foo %s", "bar")
					})

					Expect(e).Should(Equal("\033[1;31mfoo bar\033[0m\n"))
				})

				It("it does not write to stdout", func() {
					o, _ := capture(func() {
						logger.Errorf("foo %s", "bar")
					})

					Expect(o).Should(Equal(""))
				})
			})

			Context("the configured log level is less than the error level", func() {
				It("prints nothing", func() {
					logger = NewLogger("silent")
					o, e := capture(func() {
						logger.Errorf("foo")
					})

					Expect(o).Should(Equal(""))
					Expect(e).Should(Equal(""))
				})
			})
		})

		Describe("Warnf", func() {
			Context("the configured log level is not less than the warn level", func() {
				BeforeEach(func() {
					logger = NewLogger("info")
				})

				It("prints a yellow line to stderr", func() {
					_, e := capture(func() {
						logger.Warnf("foo")
					})

					Expect(e).Should(Equal("\033[1;33mfoo\033[0m\n"))
				})

				It("correctly formats and prints strings", func() {
					_, e := capture(func() {
						logger.Warnf("foo %s", "bar")
					})

					Expect(e).Should(Equal("\033[1;33mfoo bar\033[0m\n"))
				})

				It("it does not write to stdout", func() {
					o, _ := capture(func() {
						logger.Warnf("foo %s", "bar")
					})

					Expect(o).Should(Equal(""))
				})
			})

			Context("the configured log level is less than the warn level", func() {
				It("prints nothing", func() {
					logger := NewLogger("error")
					o, e := capture(func() {
						logger.Warnf("foo")
					})

					Expect(o).Should(Equal(""))
					Expect(e).Should(Equal(""))
				})
			})
		})

		Describe("Infof", func() {
			Context("the configured log level is not less than the info level", func() {
				BeforeEach(func() {
					logger = NewLogger("debug")
				})

				It("prints a green formatted line to stderr", func() {
					_, e := capture(func() {
						logger.Infof("foo")
					})

					Expect(e).Should(Equal("\033[1;32mfoo\033[0m\n"))
				})

				It("correctly formats and prints strings", func() {
					_, e := capture(func() {
						logger.Infof("foo %s", "bar")
					})

					Expect(e).Should(Equal("\033[1;32mfoo bar\033[0m\n"))
				})

				It("it does not write to stdout", func() {
					o, _ := capture(func() {
						logger.Infof("foo %s", "bar")
					})

					Expect(o).Should(Equal(""))
				})
			})

			Context("the configured log level is less than the info level", func() {
				It("prints nothing", func() {
					logger := NewLogger("error")
					o, e := capture(func() {
						logger.Warnf("foo")
					})

					Expect(o).Should(Equal(""))
					Expect(e).Should(Equal(""))
				})
			})
		})

		Describe("Debugf", func() {
			Context("the configured log level is not less than the debug level", func() {
				BeforeEach(func() {
					logger = NewLogger("debug")
				})

				It("prints a blue formatted line to stderr", func() {
					_, e := capture(func() {
						logger.Debugf("foo")
					})

					Expect(e).Should(Equal("\033[1;34mfoo\033[0m\n"))
				})

				It("correctly formats and prints strings", func() {
					_, e := capture(func() {
						logger.Debugf("foo %s", "bar")
					})

					Expect(e).Should(Equal("\033[1;34mfoo bar\033[0m\n"))
				})

				It("it does not write to stdout", func() {
					o, _ := capture(func() {
						logger.Debugf("foo %s", "bar")
					})

					Expect(o).Should(Equal(""))
				})
			})

			Context("the configured log level is less than the debug level", func() {
				It("prints nothing", func() {
					logger := NewLogger("error")
					o, e := capture(func() {
						logger.Warnf("foo")
					})

					Expect(o).Should(Equal(""))
					Expect(e).Should(Equal(""))
				})
			})
		})
	})
})

func capture(f func()) (string, string) {
	or, ow, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	er, ew, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stdout := os.Stdout
	os.Stdout = ow
	stderr := os.Stderr
	os.Stderr = ew
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()

	f()
	ow.Close()
	ew.Close()

	var oBuf bytes.Buffer
	io.Copy(&oBuf, or)

	var eBuf bytes.Buffer
	io.Copy(&eBuf, er)

	return oBuf.String(), eBuf.String()
}
