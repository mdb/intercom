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

func capture(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	stderr := os.Stderr
	os.Stderr = w
	defer func() {
		os.Stderr = stderr
	}()

	f()
	w.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
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
		Describe("Errorf", func() {
			It("prints a red formatted line to stderr", func() {
				logger := NewLogger("debug")
				out := capture(func() {
					logger.Errorf("foo")
				})

				Expect(out).Should(Equal("\033[1;31mfoo\033[0m\n"))
			})
		})

		Describe("Warnf", func() {
			It("prints a yellow formatted line to stderr", func() {
				logger := NewLogger("debug")
				out := capture(func() {
					logger.Warnf("foo")
				})

				Expect(out).Should(Equal("\033[1;33mfoo\033[0m\n"))
			})
		})

		Describe("Infof", func() {
			It("prints a green formatted line to stderr", func() {
				logger := NewLogger("debug")
				out := capture(func() {
					logger.Infof("foo")
				})

				Expect(out).Should(Equal("\033[1;32mfoo\033[0m\n"))
			})
		})

		Describe("Debugf", func() {
			It("prints a blue formatted line to stderr", func() {
				logger := NewLogger("debug")
				out := capture(func() {
					logger.Debugf("foo")
				})

				Expect(out).Should(Equal("\033[1;34mfoo\033[0m\n"))
			})
		})
	})
})
