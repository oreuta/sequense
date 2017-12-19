package sequense

import "time"

const (
	ItemStartID ItemID = -iota
	ItemErrorID
	ItemWaitID
	ItemDoneID
)

type ItemID int
type TaskID int
type ASTID int
type TaskResult struct {
	IsSuccessful bool
	StdOut       string
	StdErr       string
	ExitCode     int
}

type Item struct {
	ID            ItemID
	Task          TaskID
	NextIfSuccess ItemID
	NextIfFailure ItemID
	Done          bool
	StartedAt     time.Time
	FinishedAt    time.Time
	Result        TaskResult
}

func NewAST(items ...Item) *AST {
	ast := new(AST)
	ast.Tree = make(map[ItemID]Item)
	for _, item := range items {
		ast.Tree[item.ID] = item
	}
	return ast
}

type AST struct {
	CurrentID ItemID
	Tree      map[ItemID]Item
}

func (ast *AST) Next() (ItemID, bool) {
	item, ok := ast.Tree[ast.CurrentID]
	if !ok {
		return ItemErrorID, false
	}
	if !item.Done {
		return ItemWaitID, false
	}
	isNext := couldBeProcessed(item.ID)
	if isNext {
		if item.Result.IsSuccessful {
			ast.CurrentID = item.NextIfSuccess
		} else {
			ast.CurrentID = item.NextIfFailure
		}
	}
	return item.ID, isNext
}

func (ast *AST) Reset() {
	ast.CurrentID = ItemStartID
}

func (ast *AST) IsProcessed() bool {
	return ast.CurrentID == ItemDoneID
}

func (ast *AST) GetCurrentItem() (Item, bool) {
	item, ok := ast.Tree[ast.CurrentID]
	return item, ok
}

func couldBeProcessed(id ItemID) bool {
	return id >= ItemStartID
}
