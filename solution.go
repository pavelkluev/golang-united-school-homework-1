package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	
	message := emoji.Sprint("Hello :map: !")
	return message
}
