package sequense

import "time"

const (
	ItemNone ItemID = -iota
	ItemStop
	ItemError
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

type AST struct {
	ID          ASTID
	PartnerID   string
	FirstItem   ItemID
	CurrentItem ItemID
	Tree        [ItemID]Item
}

func (ast *AST) Next() (ItemID, bool) {
	currItem, ok := ast.Tree[ast.CurrentItem]
	if !ok {
		return ItemError, false
	}
	if !ast.Tree[currItem].Done {
		return ItemNone, false
	}
	if ast.Tree[currItem] == ItemStop {
		return ItemStop, false
	}
	if ast.Tree[currItem].Result.IsSuccessful {
		ast.Tree[currItem] = ast.Tree[currItem].NextIfSuccess
	} else {
		ast.Tree[currItem] = ast.Tree[currItem].NextIfFailure
	}
	return ast.Tree[currItem], true
}
