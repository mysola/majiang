# majiang
麻将计分工具
state_machine包维护每个人的分数及倍率，提供分数&倍率的查看&更新功能

cmd包维护所有的指令，对应一局麻将中所有可能的操作，通过命令模式可实现「undo」「replay」功能

game包控制整局游戏，调用cmd包中的命令对state_machine包中的分数&倍率进行操作

main文件，一段简单的命令行交互逻辑，直接调用game包实现功能。未来这段命令行逻辑也可做成http server，即封装成http接口，通过手机而非命令行操作
