// Convenient functions
package telegram

type TObject map[string]interface{}

func (this TObject) UpdateId() int64 {
	return int64(this["update_id"].(float64))
}

func (this TObject) Message() TObject {
	return TObject(this["message"].(map[string]interface{}))
}

func (this TObject) MessageId() int64 {
	return int64(this["message_id"].(float64))
}

func (this TObject) ChatId() int64 {
	return int64(this["chat"].(map[string]interface{})["id"].(float64))
}

func (this TObject) FromId() int64 {
	return int64(this["from"].(map[string]interface{})["id"].(float64))
}

func (this TObject) ReplyToMessage() TObject {
	return TObject(this["reply_to_message"].(map[string]interface{}))
}

func (this TObject) From() TObject {
	return TObject(this["from"].(map[string]interface{}))
}

func (this TObject) Chat() TObject {
	return TObject(this["chat"].(map[string]interface{}))
}

// Whether a 'Chat' object is a group
func (this TObject) IsGroup() bool {
	return this["title"] != nil
}

// Get photos from the message.
// If no photo found, nil will be returned.
func (this TObject) Photo() []TObject {
	if this["photo"] == nil {
		return nil
	}

	photos := this["photo"].([]interface{})
	r := make([]TObject, len(photos))

	for i, p := range photos {
		r[i] = TObject(p.(map[string]interface{}))
	}

	return r
}

// Get file path from a File object.
// Call Telegram.PathToUrl() to convert to a downloadable url.
func (this TObject) FilePath() string {
	return this["file_path"].(string)
}

// Get file id from a File or Photo object
func (this TObject) FileId() string {
	return this["file_id"].(string)
}
