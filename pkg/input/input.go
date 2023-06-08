package input

import (
	"bufio"
	"os"
)

func Input() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
