// Package randpic 随机发送色图
package randpic

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := control.Register("randpic", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Help:             "发送 随机色图 随机返回一张色图",
		PublicDataFolder: "Randpic",
	})
	engine.OnFullMatch("随机色图").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image("https://api.kaitomoe.org/picapi.php").Add("cache", "0"))
	})
}
