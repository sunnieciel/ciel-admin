package sys

import "github.com/mojocn/base64Captcha"

var (
	Store = base64Captcha.DefaultMemStore
)

func NewDriver() *base64Captcha.DriverString {
	driver := new(base64Captcha.DriverString)
	driver.Height = 60
	driver.Width = 187
	driver.NoiseCount = 5
	//driver.ShowLineOptions = base64Captcha.OptionShowSlimeLine
	driver.Length = 5
	driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	driver.Fonts = []string{"wqy-microhei.ttc", "3Dumb.ttf"}
	return driver
}
