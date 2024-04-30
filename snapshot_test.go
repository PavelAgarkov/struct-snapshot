package struct_snapshot

import (
	"fmt"
	"testing"
)

func TestDispatching(t *testing.T) {
	old := &Order{
		Snap: NewSnap(
			&Order{},
			&OrderHistory{},
		),
	}
	old.DoSnapshot("v1")

	old.BrigadeName = "777"
	old.IsBlocked = true
	old.StatusInt = 33
	old.DoSnapshot("v2")
	old.DoSnapshot("v3")

	fmt.Println(old.GetSnapshots())
	//var lastkey int64
	for k, v := range old.GetSnapshots() {
		fmt.Println(k, v)
		for k1, v1 := range v.GetChanges() {
			//fmt.Println()
			fmt.Println(k1, v1)
		}
		fmt.Println("------------")
		//lastkey = int64(k)
	}

	fmt.Println(old.GetSnapshotByTag("v2").GetChanges()["status"], "<---------------------->")
	fmt.Println(old.GetSnapshotByIndex(1), "<nil----------------------nil>")
}
