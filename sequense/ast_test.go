package sequense

import "testing"
import "fmt"

func TestASTNext(t *testing.T) {

	thenItem := Item{ID: 1, NextIfSuccess: ItemDoneID, NextIfFailure: ItemDoneID, Done: true}
	elseItem := Item{ID: 2, NextIfSuccess: ItemDoneID, NextIfFailure: ItemDoneID, Done: true}

	ifThen := NewAST(
		Item{ID: 0, NextIfSuccess: thenItem.ID, NextIfFailure: elseItem.ID, Done: true, Result: TaskResult{IsSuccessful: true}},
		thenItem,
		elseItem,
	)

	ifElse := NewAST(
		Item{ID: 0, NextIfSuccess: thenItem.ID, NextIfFailure: elseItem.ID, Done: true, Result: TaskResult{IsSuccessful: false}},
		thenItem,
		elseItem,
	)

	fmt.Println("-- IF-THEN ---")
	a := ifThen
	for task, ok := a.Next(); ok; task, ok = a.Next() {
		fmt.Printf(">>> %d [%t]\n", task, ok)
	}
	fmt.Println("-- IF-ESLE ---")
	a = ifElse
	for task, ok := a.Next(); ok; task, ok = a.Next() {
		fmt.Printf(">>> %d [%t]\n", task, ok)
	}
}
