import (
	"fmt"
	"strconv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	Price := "300000"
	price, _ := strconv.Atoi(Price)
	Price = message.NewPrinter(language.Indonesian).Sprintf("Rp. %d", price)
	fmt.Println(Price)
}
