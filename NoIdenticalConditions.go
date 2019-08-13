package sonarcloud_go_qscanner_travis

import(
"fmt"
)

func identicalIfConditions(cond bool)  {
	if cond {
	  fmt.Printf("Condition is True")
	} else if cond {
	  fmt.Printf("Condition is False")
	}
}
