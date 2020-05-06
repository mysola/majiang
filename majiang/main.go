package main

import (
	"fmt"
	"github.com/gohouse/converter/majiang/game"
	"sort"
)

func main() {
	magnification := 1
	fmt.Printf("请输入基础倍率:")
	fmt.Scanln(&magnification)

	fmt.Printf("请输入四个玩家信息:")
	var playerA, playerB, playerC, playerD string
	fmt.Scanln(&playerA, &playerB, &playerC, &playerD)

	gameController := game.NewGame(1, []string{playerA, playerB, playerC, playerD})
	for {
		hintPlayers(gameController.PrintScope())
		fmt.Println("1:刷赖子 2:点晃  3:点成  4:贴晃  5:硬晃  6:摸成  7:撤回上次操作  8:流局")
		var input int64
		fmt.Scanln(&input)
		switch input {
		case 1:
			var player int64
			fmt.Println("刷赖子的玩家：")
			fmt.Scanln(&player)
			gameController.Step刷赖子(player)
		case 2:
			var winner, loser int64
			fmt.Println("进钱：")
			fmt.Scanln(&winner)

			fmt.Println("出钱：")
			fmt.Scanln(&loser)
			gameController.Step点晃(winner, loser)
		case 3:
			var winner, loser int64
			fmt.Println("进钱：")
			fmt.Scanln(&winner)

			fmt.Println("出钱：")
			fmt.Scanln(&loser)
			gameController.Step点成(winner, loser)
		case 4:
			var winner int64
			fmt.Println("进钱：")
			fmt.Scanln(&winner)
			gameController.Step贴晃(winner)
		case 5:
			var winner int64
			fmt.Println("进钱：")
			fmt.Scanln(&winner)
			gameController.Step硬晃(winner)
		case 6:
			var winner int64
			existLaizi := false

			fmt.Println("进钱：")
			fmt.Scanln(&winner)

			fmt.Println("有赖子（y or n）：")
			var input string
			fmt.Scanln(&input)
			if input == "y" {
				existLaizi = true
			}

			gameController.Step摸成(winner, existLaizi)
		case 7:
			gameController.RevokeLastOp()
		case 8:
			gameController.Step流局()
		}
	}

}

func hintPlayers(stat []game.SingleScopeStat) {
	res := make([]string, 0)
	for _, value := range stat {
		res = append(res, fmt.Sprintf("编号:%d ，玩家：%s，分数%d，当前倍率%d", value.ID, value.Player, value.Scope, value.Odds))
	}
	sort.Strings(res)
	for _, value := range res {
		fmt.Println(value)
	}
}
