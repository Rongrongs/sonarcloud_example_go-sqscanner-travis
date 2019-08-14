package sonarcloud_go_qscanner_travis

import "testing"

func TestSum(t *testing.T) {
    t.Run("Main Sum", func(t *testing.T) {
	total := main()
	if total != 10 {
	    t.Errorf("Summ main was incorrect, got: %d, want: %d.", total, 10)
	}
    }
    t.Run("Normal Sum", func(t *testing.T) {
   	 total := Sum(5, 5)
   	 if total != 10 {
       	     t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
   	 }
    }
}
