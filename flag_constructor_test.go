package flags_test

import (
	"github.com/simonleung8/flags"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Flag Constructors", func() {

	var (
		fc flags.FlagContext
	)

	BeforeEach(func() {
		fc = flags.New()
	})

	Describe("NewStringFlag()", func() {
		It("init the flag context with a new string flagset", func() {
			fc.Parse("-s", "test")
			Ω(fc.IsSet("s")).To(BeFalse())
			Ω(fc.String("s")).To(Equal(""))

			fc.NewStringFlag("s", "setting new string flag")
			fc.Parse("-s", "test2")
			Ω(fc.IsSet("s")).To(BeTrue())
			Ω(fc.String("s")).To(Equal("test2"))
		})
	})

	Describe("NewStringFlagWithDefault()", func() {
		It("init the flag context with a new string flagset with default value", func() {
			fc.Parse("-s", "test")
			Ω(fc.IsSet("s")).To(BeFalse())
			Ω(fc.String("s")).To(Equal(""))

			fc.NewStringFlagWithDefault("s", "setting new string flag", "barz")
			fc.Parse()
			Ω(fc.IsSet("s")).To(BeTrue())
			Ω(fc.String("s")).To(Equal("barz"))
		})
	})

	Describe("NewBoolFlag()", func() {
		It("init the flag context with a new bool flagset", func() {
			fc.Parse("--force")
			Ω(fc.IsSet("force")).To(BeFalse())

			fc.NewBoolFlag("force", "force process")
			fc.Parse("--force")
			Ω(fc.IsSet("force")).To(BeTrue())
			Ω(fc.Bool("force")).To(BeTrue())
		})
	})

	Describe("NewIntFlag()", func() {
		It("init the flag context with a new int flagset", func() {
			fc.Parse("-i", "5")
			Ω(fc.IsSet("i")).To(BeFalse())
			Ω(fc.Int("i")).To(Equal(0))

			fc.NewIntFlag("i", "setting new int flag")
			fc.Parse("-i", "5")
			Ω(fc.IsSet("i")).To(BeTrue())
			Ω(fc.Int("i")).To(Equal(5))
		})
	})

	Describe("NewIntFlagWithDefault()", func() {
		It("init the flag context with a new int flagset with default value", func() {
			fc.Parse("-i", "5")
			Ω(fc.IsSet("i")).To(BeFalse())
			Ω(fc.Int("i")).To(Equal(0))

			fc.NewIntFlagWithDefault("i", "setting new int flag", 10)
			fc.Parse()
			Ω(fc.IsSet("i")).To(BeTrue())
			Ω(fc.Int("i")).To(Equal(10))
		})
	})

	Describe("NewStringSliceFlag()", func() {
		It("init the flag context with a new StringSlice flagset", func() {
			fc.Parse("-s", "5", "-s", "6")
			Ω(fc.IsSet("s")).To(BeFalse())
			Ω(fc.StringSlice("s")).To(Equal([]string{}))

			fc.NewStringSliceFlag("s", "setting new StringSlice flag")
			fc.Parse("-s", "5", "-s", "6")
			Ω(fc.IsSet("s")).To(BeTrue())
			Ω(fc.StringSlice("s")).To(Equal([]string{"5", "6"}))
		})
	})

	Describe("NewStringSliceFlagWithDefault()", func() {
		It("init the flag context with a new StringSlice flagset with default value", func() {
			fc.Parse()
			Ω(fc.IsSet("s")).To(BeFalse())
			Ω(fc.StringSlice("s")).To(Equal([]string{}))

			fc.NewStringSliceFlagWithDefault("s", "setting new StringSlice flag", []string{"5", "6", "7"})
			fc.Parse()
			Ω(fc.IsSet("s")).To(BeTrue())
			Ω(fc.StringSlice("s")).To(Equal([]string{"5", "6", "7"}))
		})
	})

})
