// Package reboot 重启
package reboot

import (
	"bufio"
	"fmt"
	"github.com/FloatTech/floatbox/process"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"os"
	"strconv"
	"strings"
)

func init() {
	go func() {
		process.SleepAbout1sTo2s()
		txt, err := os.ReadFile("data/reboot/cache.txt")
		if err != nil {
			return
		}
		lines := strings.Split(string(txt), "\n")
		a, _ := strconv.ParseInt(lines[0], 10, 64)
		b, _ := strconv.ParseInt(lines[1], 10, 64)
		ctx := zero.GetBot(a)
		if b > 0 {
			ctx.SendGroupMessage(b, message.Text(zero.BotConfig.NickName[0], "回来了喵~"))
		} else {
			ctx.SendPrivateMessage(-b, message.Text(zero.BotConfig.NickName[0], "回来了喵~"))
		}
		os.Remove("data/re/cache.txt")
	}()
	// 主函数
	en := control.Register("reboot", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "重启",
		Help: "重启命令大全\n" +
			"- 重启\n" +
			"- reboot",
		PrivateDataFolder: "reboot",
	})
	en.OnFullMatchGroup([]string{"月儿重启", "月儿reboot"}, zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text(zero.BotConfig.NickName[0], "去重启啦~"))
			gid := strconv.FormatInt(ctx.Event.GroupID, 10)
			if gid == "0" {
				gid = strconv.FormatInt(-ctx.Event.UserID, 10)
			}
			selfid := strconv.FormatInt(ctx.Event.SelfID, 10)
			file, _ := os.OpenFile(en.DataFolder()+"cache.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			defer file.Close()
			writer := bufio.NewWriter(file)
			_, _ = fmt.Fprintln(writer, selfid)
			_, _ = fmt.Fprintln(writer, gid)
			writer.Flush()
			os.Exit(0)
		})
}
