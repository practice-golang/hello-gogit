package main // import "hello-gogit"
import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	// Clone the given repository to the given directory
	fmt.Println("Same as git clone https://github.com/go-git/go-git")

	_, err := git.PlainClone("../sample-repo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

	if err != nil {
		log.Fatalf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	}
}
