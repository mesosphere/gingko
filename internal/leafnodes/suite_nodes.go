package leafnodes

import (
	"github.com/onsi/ginkgo/internal/failer"
	"github.com/onsi/ginkgo/types"
	"time"
)

type SuiteNode struct {
	runner  *runner
	outcome types.SpecState
	failure types.SpecFailure
	runTime time.Duration
}

func (node *SuiteNode) Run() bool {
	t := time.Now()
	node.outcome, node.failure = node.runner.run()
	node.runTime = time.Since(t)

	return node.outcome == types.SpecStatePassed
}

func (node *SuiteNode) Passed() bool {
	return node.outcome == types.SpecStatePassed
}

func (node *SuiteNode) Summary() *types.SetupSummary {
	return &types.SetupSummary{
		ComponentType: node.runner.nodeType,
		CodeLocation:  node.runner.codeLocation,
		State:         node.outcome,
		RunTime:       node.runTime,
		Failure:       node.failure,
	}
}

func NewBeforeSuiteNode(body interface{}, codeLocation types.CodeLocation, timeout time.Duration, failer *failer.Failer) *SuiteNode {
	return &SuiteNode{
		runner: newRunner(body, codeLocation, timeout, failer, types.SpecComponentTypeBeforeSuite, 0),
	}
}