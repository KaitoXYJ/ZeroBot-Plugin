package news

import (
	"github.com/FloatTech/floatbox/file"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"os"
)

func init() {
	engine := control.Register("news", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Brief:            "News",
		Help:             "*手动更新!手动更新!手动更新* \n *请把你的bot放到1群并打开本插件or你自己更新链接* \n 日报 拉取当日日报(最好设置job每天10点拉) \n 更新日报[图片地址] 更新当日日报图片地址",
		PublicDataFolder: "News",
	})
	newspicfile := engine.DataFolder() + "newspic.txt"
	var newspic string
	if file.IsExist(newspicfile) {
		newspict, err := os.ReadFile(newspicfile)
		if err != nil {
			panic(err)
		}
		newspic = string(newspict)
	}
	engine.OnRegex(`^更新日报\s*(.*)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		newspic = ctx.State["regex_matched"].([]string)[1]
		f, err := os.Create(newspicfile)
		if err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
			return
		}
		defer f.Close()
		_, err = f.WriteString(newspic)
		if err != nil {
			ctx.SendChain(message.Text("ERROR: ", err))
			return
		}
		ctx.SendChain(message.Text("日报图片地址更新成功"))
	})
	engine.OnFullMatch("日报").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(newspic))
	})
}
