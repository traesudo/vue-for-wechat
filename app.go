package main

import (
	"context"
	"fmt"
	"strings"
	"vue-for-wechat/wechat"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) OperationProcess(key string) string {
	switch key {
	case "1":
		result := wechat.GetAllFriends()
		// 处理结果，例如更新界面等
		fmt.Println("Friends:", result)

		return strings.Join(result, ",")
	case "2":
		result := wechat.GetAllGroups()
		// 处理结果，例如更新界面等
		fmt.Println("Groups:", result)
		return result
	case "3":
		go func() {
			wechat.Login()
			fmt.Println("Login finished.")
		}()
		return "Logging success"
	default:
		return ""
	}
}
