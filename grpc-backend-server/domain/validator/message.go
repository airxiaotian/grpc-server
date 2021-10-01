package validator

import "fmt"

//メッセージ定数
const (
	RequiredMessage    = "必須です"
	FormatMessage      = "フォーマットに誤りがあります"
	UndefinedMessage   = "未定義の値です"
	StringMessage      = "半角英数である必要があります"
	IntMessage         = "半角数字である必要があります"
	FloatMessage       = "半角数字である必要があります"
	EmptyMessage       = "空である必要があります"
	NotExistingMessage = "対象が存在しません"
	StringRegexp       = "^[a-zA-Z0-9!-/:-@¥[-`{-~]*$" // 半角
)

func RuneLengthMessage(min, max int) string {
	return fmt.Sprintf("%d桁から%d桁の文字数である必要があります", min, max)
}

func LengthMessage(min, max int) string {
	return fmt.Sprintf("%dから%dの間である必要があります", min, max)
}

func MinMessage(min int) string {
	return fmt.Sprintf("%d以上の値である必要があります", min)
}

func MaxMessage(max int) string {
	return fmt.Sprintf("%d以下の値である必要があります", max)
}
