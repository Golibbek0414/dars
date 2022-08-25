package telegram

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
	"time"
)

func initialize() *Telegram {
	b, _ := ioutil.ReadFile("./test_key.txt")
	k := strings.Trim(string(b), " \n")
	return New(k, true)
}

func getUid() int64 {
	b, _ := ioutil.ReadFile("./test_user.txt")
	i, _ := strconv.ParseInt(strings.Trim(string(b), " \n"), 10, 64)
	return i
}

func TestSetWebhook(*testing.T) {
	if testing.Short() {
		return
	}

	tg := initialize()

	fmt.Println(tg.SetWebhook("https://google.com/"))
	fmt.Println(tg.SetWebhook(""))
}

func TestGetUpdates(*testing.T) {
	if testing.Short() {
		return
	}
	tg := initialize()

	fmt.Println(tg.GetUpdates(0, 100, 5))
}

func TestSendMessage(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test for sending message")
	} else {
		tg := initialize()
		uid := getUid()
		fmt.Println(tg.SendChatAction("typing", uid))
		time.Sleep(3000 * time.Millisecond)
		fmt.Println(tg.SendMessage("tesuto", uid))
	}
}

func TestParser(t *testing.T) {
	printArray(ParseArgs("aaa bbb ccc ddd"))
	printArray(ParseArgs("aaa 'bbb ccc' ddd"))
	printArray(ParseArgs("\"aaa 'bbb\" ccc ddd"))
	printArray(ParseArgs("aaa 'bbb ccc\" ddd'"))
	printArray(ParseArgs("aaa  bbb"))
}

func printArray(arr []string) {
	for _, t := range arr {
		fmt.Println(t)
	}
}