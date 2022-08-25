// Telegram Bot API client implementation
package telegram
import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ddliu/go-httpclient"
)
// https://api.telegram.org/bot<token>/METHOD_NAME
const MethodUrl = "https://api.telegram.org/bot%s/%s"
const FileUrl = "https://api.telegram.org/file/bot%s/%s"
type Telegram struct {
	apiKey string
	debug  bool
}
func New(key string, debug bool) *Telegram {
	return &Telegram{
		apiKey: key,
		debug:  debug,
	}
}
func (tg *Telegram) post(method string, params map[string]string) map[string]interface{} {
	url := fmt.Sprintf(MethodUrl, tg.apiKey, method)
	if tg.debug {
		fmt.Println(url)
	}
	res, err := httpclient.Post(url, params)
	if (err != nil) || (res == nil) {
		return nil
	} else {
		var ret interface{}
		var str string
		str, err = res.ToString()
		if tg.debug {
			fmt.Println(str)
		}
		if err != nil {
			return nil
		}
		err = json.Unmarshal([]byte(str), &ret)
		if (err == nil) && (ret != nil) {
			return ret.(map[string]interface{})
		} else {
			return nil
		}
	}
}
func (tg *Telegram) SetWebhook(url string) bool {
	res := tg.post("setWebhook", map[string]string{
		"url": url,
	})
	if res == nil {
		return false
	}
	return res["ok"].(bool)
}
// timeout: time to wait (seconds)
func (tg *Telegram) GetUpdates(offset int64, limit int, timeout int) []interface{} {
	res := tg.post("getUpdates", map[string]string{
		"offset":  fmt.Sprintf("%d", offset),
		"limit":   fmt.Sprintf("%d", limit),
		"timeout": fmt.Sprintf("%d", timeout),
	})
	if (res == nil) || (!res["ok"].(bool)) {
		return nil
	}
	return res["result"].([]interface{})
}
// Get the File object for <id>.
// Including file path.
func (tg *Telegram) GetFile(id string) TObject {
	res := tg.post("getFile", map[string]string{
		"file_id": id,
	})
	if (res == nil) || (!res["ok"].(bool)) {
		return nil
	}
	return TObject(res["result"].(map[string]interface{}))
}
// File Path to File URL.
// https://api.telegram.org/file/bot<token>/<file_path>
func (tg *Telegram) PathToUrl(path string) string {
	return fmt.Sprintf(FileUrl, tg.apiKey, path)
}
// The Message object has too many fields, let's just use a map
// Consult <https://core.telegram.org/bots/api#sendmessage> for a full list of params
func (tg *Telegram) SendMessageRaw(msg map[string]string) bool {
	res := tg.post("sendMessage", msg)
	if res == nil {
		return false
	}
	return res["ok"].(bool)
}
func (tg *Telegram) SendMessage(text string, chat int64) bool {
	return tg.SendMessageRaw(map[string]string{
		"chat_id": fmt.Sprintf("%d", chat),
		"text":    text,
	})
}
func (tg *Telegram) SendMessageNoPreview(text string, chat int64) bool {
	return tg.SendMessageRaw(map[string]string{
		"chat_id": fmt.Sprintf("%d", chat),
		"text":    text,
		"disable_web_page_preview": "true",
	})
}

// Send a message to a Channel
func (tg *Telegram) SendMessageChan(text string, chat string) bool {
	return tg.SendMessageRaw(map[string]string{
		"chat_id":    "@" + chat,
		"parse_mode": "Markdown",
		"text":       text,
		"disable_web_page_preview": "true",
	})
}

func (tg *Telegram) ReplyToMessage(id int64, text string, chat int64) bool {
	return tg.SendMessageRaw(map[string]string{
		"chat_id":             fmt.Sprintf("%d", chat),
		"text":                text,
		"reply_to_message_id": fmt.Sprintf("%d", id),
	})
}
// See <https://core.telegram.org/bots/api#sendchataction> for a list of actions
func (tg *Telegram) SendChatAction(action string, chat int64) bool {
	res := tg.post("sendChatAction", map[string]string{
		"chat_id": fmt.Sprintf("%d", chat),
		"action":  action,
	})
	if res == nil {
		return false
	}
	return res["ok"].(bool)
}
// The map should contain a key "@photo" and its value should be path to the photo
// I recommend using a temp file.
func (tg *Telegram) SendPhotoRaw(msg map[string]string) bool {
	res := tg.post("sendPhoto", msg)
	if res == nil {
		return false
	}
	return res["ok"].(bool)
}
func (tg *Telegram) SendPhoto(file string, chat int64) bool {
	return tg.SendPhotoRaw(map[string]string{
		"chat_id": fmt.Sprintf("%d", chat),
		"@photo":  file,
	})
}

// Send a photo to a channel
func (tg *Telegram) SendPhotoChan(file string, chat string) bool {
	return tg.SendPhotoRaw(map[string]string{
		"chat_id": "@" + chat,
		"@photo":  file,
	})
}

// Escape a string for using in Markdown
const markdownKeywords = "_*<>[](){}-"

func Escape(str string) (ret string) {
	for _, runeValue := range str {
		if strings.Contains(markdownKeywords, string(runeValue)) {
			ret += fmt.Sprintf("\\%s", string(runeValue))
		} else {
			ret += string(runeValue)
		}
	}

	return
}