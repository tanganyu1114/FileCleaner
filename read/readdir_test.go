package read

import (
	"fmt"
	"testing"
)

func TestReadDir(t *testing.T) {
	/*	path,err:=os.Getwd()
		if err != nil {
			t.Log(err)
		}
		t.Log("path:",path)*/
}

func TestEchoFloat(t *testing.T) {
	fmt.Printf("%.2f ", 0.15)
	var aa int = 1024
	var bb int = 1010
	fmt.Println(float64(aa) / 1024 / 1024)
	fmt.Println(float64(aa / bb))
	numPct := fmt.Sprintf("%.2f %%", float64(aa)/float64(bb))
	fmt.Println(numPct)
}
