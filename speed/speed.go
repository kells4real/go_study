// A speed test generically showing how fast Go can be against another language
// In this case, i tested it against Python. 
// The python code is included in the project root folder. You can run it and see the difference in speed
// This also depends on the system.
// On my corei5 4th Gen, the program averaged 13ms for Go, and 93ms for Python
// On my corei3 10th Gen, the program averaged 9ms for Go, and 60ms for Python
package speedTest
import (
	"fmt"
	"sort"
	"time"
	"math/rand"
)

func Test() {
	var start = time.Now()
	var myList []int

	for i := 0; i < 100000; i++{
		myList = append(myList, rand.Intn(100))
	}

	sort.Ints(myList)
	fmt.Println(time.Since(start))
}