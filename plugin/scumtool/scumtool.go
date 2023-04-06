// Package scumtool scum辅助工具
package scumtool

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	engine := control.Register("scumtool", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Help:             "发送 商城 查看服务器商城网址 \n 发送 地图 查看游戏地图 \n 发送 新手礼包 查看新手礼包领取方法 \n 发送 注册商城 查看商城注册方法 \n 发送 回收 查看服务器回收规则",
		PublicDataFolder: "Scumtool",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.Send("插件已启用")
		},
		OnDisable: func(ctx *zero.Ctx) {
			ctx.Send("插件已禁用")
		},
	})
	engine.OnFullMatch("商城").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("商城网址 123.129.198.109:7993 复制到浏览器内打开"))
	})
	engine.OnFullMatch("地图").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("人渣超高清地图 https://scum-map.com/zh-CN/interactive_map 复制到浏览器打开"))
	})
	engine.OnFullMatch("注册商城").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("steamID: 一个17位的数值例如7656119xxxxxxxxxx \n 1.打开steam-视图 \n 2.设置-切换到界面选项 \n 3.勾选可用时显示Steam URL地址栏 \n 4.点击账户名-查看个人资料-地址栏最后一串数字就是id "))
	})
	engine.OnFullMatch("新手礼包").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("打开商城 网址: 123.129.198.109:7993 点击网页上方礼包 点击领取新手礼包"))
	})
	engine.OnFullMatch("回收").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("服务器当前回收所有能出售给游戏内商人的物资,可与商人交易物品得到美金,同时支持美金和商城货币渣币互相转化,10美金=1渣币"))
	})
	engine.OnFullMatch("语音").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("kook频道链接 https://kook.top/PPn61D 可使用浏览器或客户端连麦"))
	})
	engine.OnFullMatch("ip").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("服务器直连ip 94.250.207.14:28502 搜不到可直连"))
	})
}
