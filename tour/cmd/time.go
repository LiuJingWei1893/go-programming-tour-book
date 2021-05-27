package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-programming-tour-book/tour/internal/timer"
	"github.com/spf13/cobra"
)

var dest = strings.Join([]string{
	"子命令为：",
	"now: 获取当前时间",
	"calc: 根据给定时间或当前时间以及时间间隔，给出经过计算后的时间",
}, "\n")

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  dest,
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("请重新输入该命令并带入子命令，或查看帮助信息")
	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNewTime()
		log.Printf("当前时间为：%s, Unix格式时间为：%d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTime string
var duration string

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		layout := "2006-01-02 15:04:05" //go时间标准格式
		location, _ := time.LoadLocation("Asia/Shanghai")
		if calculateTime == "" {
			currentTimer = timer.GetNewTime()
		} else {
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			} else if space != 1 {
				log.Fatalf("calculate输入格式有误")
			}
			var err error
			currentTimer, err = time.ParseInLocation(layout, calculateTime, location) //将时间从字符串parse为time的时候需要注意时区
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}  
		}
		t, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		log.Printf("计算出的时间是：%s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)

	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "µ s"), "ms", "s", "m", "h"`)
}
