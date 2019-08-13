package sonarcloud_go_qscanner_travis

import "fmt"

func test(flag bool) {
	if flag { // Noncompliant
	  fmt.Printf("Flag is True")
	}
}
