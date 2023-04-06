// Package batool 蔚蓝档案辅助工具
package batool

import (
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	"github.com/FloatTech/zbputils/ctxext"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

var mainlimit = ctxext.NewLimiterManager(time.Second*10, 1)

func init() {
	engine := control.Register("batool", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: true,
		Brief:            "Batool",
		Help:             "发送 ba菜单 查看我的功能 \n 感谢夜猫大佬供图",
		PublicDataFolder: "Batool",
	})
	// 文件包路径
	bapath := "https://api.kaitomoe.org/batool/"
	// 菜单
	engine.OnFullMatch("ba菜单").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "cd.png"))
	})
	// 千里眼
	engine.OnFullMatch("千里眼").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "qly.png"))
	})
	// 查角色
	// 三星常驻区
	// 阿里乌斯
	engine.OnFullMatchGroup([]string{"纱织", "saori"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "saori.png"))
	})
	engine.OnFullMatchGroup([]string{"墩子", "公主", "atsuko"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "atsuko.png"))
	})
	engine.OnFullMatchGroup([]string{"美咲", "misaki"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "misaki.png"))
	})
	engine.OnFullMatchGroup([]string{"日和", "hiyori"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hiyori.png"))
	})
	// SRT
	engine.OnFullMatchGroup([]string{"萌惠", "moe"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "moe.png"))
	})
	engine.OnFullMatchGroup([]string{"美游", "miyu", "垃圾兔"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "miyu.png"))
	})
	engine.OnFullMatchGroup([]string{"宫子", "miyako"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "miyako.png"))
	})
	engine.OnFullMatchGroup([]string{"咲", "saki"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "saki.png"))
	})
	// 千年
	engine.OnFullMatchGroup([]string{"兔朱音", "1akane", "兔女郎朱音", "兔女仆"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1akane.png"))
	})
	engine.OnFullMatchGroup([]string{"体操服咏叶", "1utaha", "应援团咏叶"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "utaha.png"))
	})
	engine.OnFullMatchGroup([]string{"诺亚", "noa"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "noa.png"))
	})
	engine.OnFullMatchGroup([]string{"千寻", "chihiro"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "chihiro.png"))
	})
	engine.OnFullMatchGroup([]string{"响", "hibiki"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hibiki.png"))
	})
	engine.OnFullMatchGroup([]string{"绿", "midori", "小绿"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "midori.png"))
	})
	engine.OnFullMatchGroup([]string{"爱丽丝", "arisu"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "arisu.png"))
	})
	engine.OnFullMatchGroup([]string{"柚子", "yuzu"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "yuzu.png"))
	})
	engine.OnFullMatchGroup([]string{"兔明日奈", "1asuna", "兔大大大"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1asuna.png"))
	})
	engine.OnFullMatchGroup([]string{"花凛", "karin", "黑皮"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "karin.png"))
	})
	engine.OnFullMatchGroup([]string{"真纪", "maki", "彩蛋"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "maki.png"))
	})
	engine.OnFullMatchGroup([]string{"菫", "sumire", "运动妹"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "x.png"))
	})
	engine.OnFullMatchGroup([]string{"英美", "eimi"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "eimi.png"))
	})
	engine.OnFullMatchGroup([]string{"尼禄", "neru"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "neru.png"))
	})
	// 格黑娜
	engine.OnFullMatchGroup([]string{"伊吕波", "iroha", "168"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "iroha.png"))
	})
	engine.OnFullMatchGroup([]string{"濑名", "sena", "濑奈", "救护车"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "sena.png"))
	})
	engine.OnFullMatchGroup([]string{"伊织", "iori", "佐仓"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "iori.png"))
	})
	engine.OnFullMatchGroup([]string{"羽留奈", "haruna", "神秘狙"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "haruna.png"))
	})
	engine.OnFullMatchGroup([]string{"亚子", "ako"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ako.png"))
	})
	engine.OnFullMatchGroup([]string{"千夏", "chinatsu"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "chinatsu.png"))
	})
	engine.OnFullMatchGroup([]string{"阿鲁", "aru", "阿露"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "aru.png"))
	})
	engine.OnFullMatchGroup([]string{"日奈", "杨奈", "hina"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hina.png"))
	})
	engine.OnFullMatchGroup([]string{"泉", "lzumi", "汉堡"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "lzumi.png"))
	})
	// 三一
	engine.OnFullMatchGroup([]string{"千纱", "kazusa"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kazusa.png"))
	})
	engine.OnFullMatchGroup([]string{"优", "ui"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ui.png"))
	})
	engine.OnFullMatchGroup([]string{"日向", "hinata"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hinata.png"))
	})
	engine.OnFullMatchGroup([]string{"梓", "azusa"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "azusa.png"))
	})
	engine.OnFullMatchGroup([]string{"小春", "koharu"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "koharu.png"))
	})
	engine.OnFullMatchGroup([]string{"夏", "natsu", "小夏"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "x.png"))
	})
	engine.OnFullMatchGroup([]string{"水日富美", "1hifumi", "泳装日富美"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1hifumi.png"))
	})
	engine.OnFullMatchGroup([]string{"日富美", "hifumi"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hifumi.png"))
	})
	engine.OnFullMatchGroup([]string{"真白", "mashiro", "麻白"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "mashiro.png"))
	})
	engine.OnFullMatchGroup([]string{"弦生", "tsurugi", "颜艺"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "tsurugi.png"))
	})
	// 阿拜多斯
	engine.OnFullMatchGroup([]string{"泳装野乃美", "1nonomi", "nnm", "水nnm"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1nonomi.png"))
	})
	engine.OnFullMatchGroup([]string{"新年茜香", "1serika", "新年黑猫", "春黑猫"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1serika.png"))
	})
	engine.OnFullMatchGroup([]string{"单车白子", "1shiroko", "小车唯"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1shiroko.png"))
	})
	engine.OnFullMatchGroup([]string{"白字", "shiroko", "xcw", "小仓唯"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "shiroko.png"))
	})
	engine.OnFullMatchGroup([]string{"星野", "hoshino", "大叔"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hoshino.png"))
	})
	// 红冬
	engine.OnFullMatchGroup([]string{"玛丽娜", "marina", "大妈"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "marina.png"))
	})
	engine.OnFullMatchGroup([]string{"温泉洁莉诺", "1cherino", "温泉斯大萝"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1cherino.png"))
	})
	engine.OnFullMatchGroup([]string{"洁莉诺", "cherino", "斯大萝"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "cherino.png"))
	})
	engine.OnFullMatchGroup([]string{"温泉和香", "1nodoka"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1nodoka.png"))
	})
	// 百鬼夜行
	engine.OnFullMatchGroup([]string{"泳装若藻", "1wakamo", "水大狐狸"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1wakamo.png"))
	})
	engine.OnFullMatchGroup([]string{"月咏", "tsukuyo"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "tsukuyo.png"))
	})
	engine.OnFullMatchGroup([]string{"枫", "kaede"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kaede.png"))
	})
	engine.OnFullMatchGroup([]string{"三森", "mimori", "新娘"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "mimori.png"))
	})
	engine.OnFullMatchGroup([]string{"伊树莱", "izuna", "小狐狸"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "izuna.png"))
	})
	// 山海经
	engine.OnFullMatchGroup([]string{"心菜", "kokona"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kokona.png"))
	})
	engine.OnFullMatchGroup([]string{"瞬", "shun"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "shun.png"))
	})
	engine.OnFullMatchGroup([]string{"幼瞬", "1shun"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1shun.png"))
	})
	engine.OnFullMatchGroup([]string{"私服沙耶", "1saya", "私服鼠鼠"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1saya.png"))
	})
	engine.OnFullMatchGroup([]string{"沙耶", "saya", "鼠鼠"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "saya.png"))
	})
	// 活动限定
	// 阿拜多斯
	engine.OnFullMatchGroup([]string{"泳装星野", "1hoshino", "水大叔"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1hoshino.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装绫音", "1ayane", "直升机"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1ayane.png"))
	})
	// 百鬼夜行
	engine.OnFullMatchGroup([]string{"泳装伊树莱", "1izuna", "水小狐狸"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1izuna.png"))
	})
	engine.OnFullMatchGroup([]string{"满", "michiru", "小满"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "michiru.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装静子", "1shizuko", "水静子"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1shizuko.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装知世", "1chise", "水知世"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1chise.png"))
	})
	engine.OnFullMatchGroup([]string{"若藻", "wakamo", "大狐狸"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "wakamo.png"))
	})
	// 瓦尔基里
	engine.OnFullMatchGroup([]string{"吹雪", "fubuki"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "fubuki.png"))
	})
	// 千年
	engine.OnFullMatchGroup([]string{"体操服优香", "1yuka", "有包人"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1yuka.png"))
	})
	engine.OnFullMatchGroup([]string{"应援团响", "1hibiki"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1hibiki.png"))
	})
	engine.OnFullMatchGroup([]string{"兔女郎尼禄", "1neru", "红兔", "兔尼禄"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1neru.png"))
	})
	engine.OnFullMatchGroup([]string{"兔女郎花凛", "1karin", "黑兔", "兔花凛"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1karin.png"))
	})
	// 格黑娜
	engine.OnFullMatchGroup([]string{"新年阿鲁", "1aru", "新年阿露"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1aru.png"))
	})
	engine.OnFullMatchGroup([]string{"新年无月", "1mutsuki", "春月"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1mutsuki.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装日奈", "1hina", "水hina"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1hina.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装伊织", "1iori", "水佐仓"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1iori.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装泉", "1izumi", "水汉堡"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1izumi.png"))
	})
	// 三一
	engine.OnFullMatchGroup([]string{"体操服玛丽", "1mari"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1mari.png"))
	})
	engine.OnFullMatchGroup([]string{"体操服莲见", "1hasumi"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1hasumi.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装梓", "1azusa", "水梓"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1azusa.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装真白", "1mashiro", "水白", "水真白"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1mashiro.png"))
	})
	engine.OnFullMatchGroup([]string{"泳装弦生", "1tsurugi", "水颜艺"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "1tsurugi.png"))
	})
	// 红冬
	engine.OnFullMatchGroup([]string{"智惠", "tomoe"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "tomoe.png"))
	})
	engine.OnFullMatchGroup([]string{"和香", "nodoka", "偷窥狂"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "nodoka.png"))
	})
	// 夏莱
	engine.OnFullMatchGroup([]string{"初音未来", "miku", "初音"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "miku.png"))
	})
	// 一星二星常驻区
	// 千年
	engine.OnFullMatchGroup([]string{"小玉", "kotama"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kotama.png"))
	})
	engine.OnFullMatchGroup([]string{"朱音", "akane", "女仆"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "akane.png"))
	})
	engine.OnFullMatchGroup([]string{"桃井", "momori", "小红"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "momori.png"))
	})
	engine.OnFullMatchGroup([]string{"优香", "yuka", "没包人"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "yuka.png"))
	})
	engine.OnFullMatchGroup([]string{"明日奈", "asuna", "大大大"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "asuna.png"))
	})
	engine.OnFullMatchGroup([]string{"咏叶", "utaha"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "utaha.png"))
	})
	engine.OnFullMatchGroup([]string{"亚都梨", "kotori"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kotori.png"))
	})
	// 格黑娜
	engine.OnFullMatchGroup([]string{"无月", "mutsuki", "雌小鬼"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "mutsuki.png"))
	})
	engine.OnFullMatchGroup([]string{"淳子", "junko"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "junko.png"))
	})
	engine.OnFullMatchGroup([]string{"亚伽里", "akari"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "akari.png"))
	})
	engine.OnFullMatchGroup([]string{"千夏", "chinatsu"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "chinatsu.png"))
	})
	engine.OnFullMatchGroup([]string{"遥香", "haruka"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "haruka.png"))
	})
	engine.OnFullMatchGroup([]string{"佳代子", "kayoko", "佳世子"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kayoko.png"))
	})
	// 三一
	engine.OnFullMatchGroup([]string{"莲见", "hasumi"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hasumi.png"))
	})
	engine.OnFullMatchGroup([]string{"芹奈", "serina", "小护士"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "serina.png"))
	})
	engine.OnFullMatchGroup([]string{"花绘", "hanae"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "x.png"))
	})
	engine.OnFullMatchGroup([]string{"玛丽", "mari"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "mari.png"))
	})
	// 阿拜多斯
	engine.OnFullMatchGroup([]string{"茜香", "serika", "黑猫"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "serika.png"))
	})
	engine.OnFullMatchGroup([]string{"野乃美", "nonomi", "nnm"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "nonomi.png"))
	})
	// 百鬼
	engine.OnFullMatchGroup([]string{"椿", "tsubaki", "狗盾"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "tsubaki.png"))
	})
	engine.OnFullMatchGroup([]string{"知世", "chise"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "chise.png"))
	})
	engine.OnFullMatchGroup([]string{"菲娜", "pina"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "pina.png"))
	})
	engine.OnFullMatchGroup([]string{"静子", "shizuko"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "shizuko.png"))
	})
	// 瓦尔基里
	engine.OnFullMatchGroup([]string{"桐乃", "kirino"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "kirino.png"))
	})
	// 4.0xp补足图集
	engine.OnFullMatchGroup([]string{"喜美", "爱莉", "志美子", "铃美", "花子", "朱莉", "风华", "煮饭婆", "绫音", "晴"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "40xp.png"))
	})
	engine.OnFullMatchGroup([]string{"yoshimi", "airi", "shimiko", "suzumi", "hanako", "juri", "fukka", "hare", "ayane"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "40xp.png"))
	})
	// jjc攻略
	engine.OnFullMatch("国际服jjc").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "gjfjjc.png"))
	})
	engine.OnFullMatch("日服jjc").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "rfjjc.png"))
	})
	// 刷主线攻略
	engine.OnFullMatch("刷主线攻略").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "szxgl.png"))
	})
	// 好感度
	engine.OnFullMatch("好感度经验表").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hgdjyb.png"))
	})
	engine.OnFullMatch("三一礼物表").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "sylwb.png"))
	})
	engine.OnFullMatch("千年礼物表").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "qnlwb.png"))
	})
	engine.OnFullMatch("格黑娜礼物表").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ghnlwb.png"))
	})
	engine.OnFullMatch("阿拜多斯礼物表").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "abdslwb.png"))
	})
	engine.OnFullMatchGroup([]string{"红冬礼物表", "山海经礼物表"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hdshjlwb.png"))
	})
	engine.OnFullMatchGroup([]string{"阿里乌斯礼物表", "srt礼物表", "女武神礼物表"}).SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "alwssrtnwslwb.png"))
	})
	// 爱用品
	engine.OnFullMatch("第一批爱用品").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ayp1.png"))
	})
	engine.OnFullMatch("第二批爱用品").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ayp2.png"))
	})
	// 总力战
	engine.OnFullMatch("主教机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "zjjz.png"))
	})
	engine.OnFullMatch("主教用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "zjyr.png"))
	})
	engine.OnFullMatch("黑白机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hbjz.png"))
	})
	engine.OnFullMatch("黑白用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hbyr.png"))
	})
	engine.OnFullMatch("goz机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "gozjz.png"))
	})
	engine.OnFullMatch("goz用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "gozyr.png"))
	})
	engine.OnFullMatch("球机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "qjz.png"))
	})
	engine.OnFullMatch("球用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "qyr.png"))
	})
	engine.OnFullMatch("寿司机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ssjz.png"))
	})
	engine.OnFullMatch("寿司用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "ssyr.png"))
	})
	engine.OnFullMatch("大蛇机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "dsjz.png"))
	})
	engine.OnFullMatch("大蛇用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "dsyr.png"))
	})
	engine.OnFullMatch("鸡斯拉机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "jsljz.png"))
	})
	engine.OnFullMatch("鸡斯拉用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "jslyr.png"))
	})
	engine.OnFullMatch("hod机制").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hodjz.png"))
	})
	engine.OnFullMatch("hod用人").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "hodyr.png"))
	})
	// 制造攻略
	engine.OnFullMatch("制造攻略").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "zzgl.png"))
	})
	// 实用网站
	engine.OnFullMatch("SchaleDB").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("SchaleDB: lonqie.github.io/SchaleDB"))
	})
	engine.OnFullMatch("wiki").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("wiki: ba.gamekee.com"))
	})
	engine.OnFullMatch("素材计算").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("素材计算: ba.game-db.tw"))
	})
	engine.OnFullMatch("韩版总力战阵容").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("韩版总力战阵容: https://boatneck-telescope-7e9.notion.site/2df4987277984be9a07ec1efa27279e3"))
	})
	engine.OnFullMatch("ba官网").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("ba官网: bluearchive.nexon.com"))
	})
	engine.OnFullMatch("ba下载地址").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Text("ba下载地址: https://play.google.com/store/apps/details?id=com.nexon.bluearchive&hl=en&gl=US"))
	})
	// 角色排行榜
	engine.OnFullMatch("角色排行榜").SetBlock(true).Limit(mainlimit.LimitByGroup).Handle(func(ctx *zero.Ctx) {
		ctx.SendChain(message.Image(bapath + "jsphb.png"))
	})
}
