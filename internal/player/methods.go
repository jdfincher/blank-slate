package player

var (
	TestMethod1 = Method{
		Name: "TestMethod1",
		Info: `This is a Test Method, what will it look like?
I don't know if this will work or not but we wil see`,
	}
	TestMethods = []Method{TestMethod1}
)

type Method struct {
	Name      string
	Info      string
	StatBonus int
	StatKey   string
	Active    bool
	ViewState bool
}
