package data

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/evergreen-ci/cedar"
	dbModel "github.com/evergreen-ci/cedar/model"
	"github.com/evergreen-ci/cedar/rest/model"
	"github.com/stretchr/testify/suite"
)

type testResultsConnectorSuite struct {
	ctx         context.Context
	cancel      context.CancelFunc
	sc          Connector
	env         cedar.Environment
	testResults map[string]dbModel.TestResults
	tempDir     string
	apiResults  map[string]model.APITestResult

	suite.Suite
}

func TestTestResultsConnectorSuiteDB(t *testing.T) {
	s := new(testResultsConnectorSuite)
	s.setup()
	s.sc = CreateNewDBConnector(s.env)
	suite.Run(t, s)
}

func TestTestResultsConnectorSuiteMock(t *testing.T) {
	s := new(testResultsConnectorSuite)
	s.setup()
	s.sc = &MockConnector{
		CachedTestResults: s.testResults,
		env:               cedar.GetEnvironment(),
		Bucket:            s.tempDir,
	}
	suite.Run(t, s)
}

func (s *testResultsConnectorSuite) setup() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.env = cedar.GetEnvironment()
	s.Require().NotNil(s.env)
	db := s.env.GetDB()
	s.Require().NotNil(db)
	s.testResults = map[string]dbModel.TestResults{}

	// setup config
	var err error
	s.tempDir, err = ioutil.TempDir(".", "testResults_connector")
	s.Require().NoError(err)
	conf := dbModel.NewCedarConfig(s.env)
	conf.Bucket = dbModel.BucketConfig{TestResultsBucket: s.tempDir}
	s.Require().NoError(conf.Save())

	testResultInfos := []dbModel.TestResultsInfo{
		{
			Project:       "test",
			Version:       "0",
			Variant:       "linux",
			TaskID:        "task1",
			DisplayTaskID: "display_task1",
			Execution:     0,
			RequestType:   "requesttype",
			Mainline:      true,
		},
		{
			Project:       "test",
			Version:       "0",
			Variant:       "linux",
			TaskID:        "task1",
			DisplayTaskID: "display_task1",
			Execution:     1,
			RequestType:   "requesttype",
			Mainline:      true,
		},
		{
			Project:       "test",
			Version:       "0",
			Variant:       "linux",
			TaskID:        "task2",
			DisplayTaskID: "display_task1",
			Execution:     0,
			RequestType:   "requesttype",
			Mainline:      true,
		},
		{
			Project:       "test",
			Version:       "0",
			Variant:       "linux",
			TaskID:        "task3",
			DisplayTaskID: "display_task2",
			Execution:     0,
			RequestType:   "requesttype",
			Mainline:      true,
		},
	}

	s.apiResults = map[string]model.APITestResult{}

	for _, testResultsInfo := range testResultInfos {
		testResults := dbModel.CreateTestResults(testResultsInfo, dbModel.PailLocal)

		testResults.Setup(s.env)
		err = testResults.SaveNew(s.ctx)

		s.Require().NoError(err)
		s.testResults[testResults.ID] = *testResults

		for i := 0; i < 3; i++ {
			result := dbModel.TestResult{
				TaskID:    testResults.Info.TaskID,
				Execution: testResults.Info.Execution,
				TestName:  fmt.Sprintf("test%d", i),
				Trial:     0,
				Status:    "teststatus",
				LineNum:   0,
			}

			apiResult := model.APITestResult{}
			s.Require().NoError(apiResult.Import(result))
			s.apiResults[fmt.Sprintf("%s_%d_%s", result.TaskID, result.Execution, result.TestName)] = apiResult

			s.Require().NoError(testResults.Append(s.ctx, []dbModel.TestResult{result}))
		}
	}
}

func (s *testResultsConnectorSuite) TearDownSuite() {
	defer s.cancel()
	s.NoError(os.RemoveAll(s.tempDir))
	s.NoError(s.env.GetDB().Drop(s.ctx))
}

func (s *testResultsConnectorSuite) TestFindTestResultsExists() {
	for _, test := range []struct {
		name      string
		opts      TestResultsOptions
		resultMap map[string]model.APITestResult
	}{
		{
			name: "TaskIDWithExecution",
			opts: TestResultsOptions{
				TaskID:    "task1",
				Execution: 0,
			},
			resultMap: map[string]model.APITestResult{
				"task1_0_test0": s.apiResults["task1_0_test0"],
				"task1_0_test1": s.apiResults["task1_0_test1"],
				"task1_0_test2": s.apiResults["task1_0_test2"],
			},
		},
		{
			name: "TaskIDWithoutExecution",
			opts: TestResultsOptions{
				TaskID:         "task1",
				EmptyExecution: true,
			},
			resultMap: map[string]model.APITestResult{
				"task1_1_test0": s.apiResults["task1_1_test0"],
				"task1_1_test1": s.apiResults["task1_1_test1"],
				"task1_1_test2": s.apiResults["task1_1_test2"],
			},
		},
		{
			name: "DisplayTaskIDWithExecution",
			opts: TestResultsOptions{
				DisplayTaskID: "display_task1",
				Execution:     0,
			},
			resultMap: map[string]model.APITestResult{
				"task1_0_test0": s.apiResults["task1_0_test0"],
				"task1_0_test1": s.apiResults["task1_0_test1"],
				"task1_0_test2": s.apiResults["task1_0_test2"],
				"task2_0_test0": s.apiResults["task2_0_test0"],
				"task2_0_test1": s.apiResults["task2_0_test1"],
				"task2_0_test2": s.apiResults["task2_0_test2"],
			},
		},
		{
			name: "DisplayTaskIDWithoutExecution",
			opts: TestResultsOptions{
				DisplayTaskID:  "display_task1",
				EmptyExecution: true,
			},
			resultMap: map[string]model.APITestResult{
				"task1_1_test0": s.apiResults["task1_1_test0"],
				"task1_1_test1": s.apiResults["task1_1_test1"],
				"task1_1_test2": s.apiResults["task1_1_test2"],
			},
		},
	} {
		s.T().Run(test.name, func(t *testing.T) {
			actualResults, err := s.sc.FindTestResults(s.ctx, test.opts)
			s.Require().NoError(err)

			s.Len(actualResults, len(test.resultMap))
			for _, result := range actualResults {
				key := fmt.Sprintf("%s_%d_%s", *result.TaskID, result.Execution, *result.TestName)
				expected, ok := test.resultMap[key]
				s.Require().True(ok)
				s.Equal(expected.TestName, result.TestName)
				s.Equal(expected.TaskID, result.TaskID)
				s.Equal(expected.Execution, result.Execution)
				delete(test.resultMap, key)
			}
		})
	}
}

func (s *testResultsConnectorSuite) TestFindTestResultDNE() {
	for _, test := range []struct {
		name string
		opts TestResultsOptions
	}{
		{
			name: "TaskID",
			opts: TestResultsOptions{TaskID: "DNE"},
		},
		{
			name: "DisplayTaskID",
			opts: TestResultsOptions{DisplayTaskID: "DNE"},
		},
	} {
		s.T().Run(test.name, func(t *testing.T) {
			result, err := s.sc.FindTestResults(s.ctx, test.opts)
			s.Error(err)
			s.Nil(result)
		})
	}

}

func (s *testResultsConnectorSuite) TestFindTestResultByTaskIDEmpty() {
	opts := TestResultsOptions{Execution: 1}

	result, err := s.sc.FindTestResults(s.ctx, opts)
	s.Error(err)
	s.Nil(result)
}

func (s *testResultsConnectorSuite) TestFindTestResultByTestNameExists() {
	for _, test := range []struct {
		name           string
		opts           TestResultsOptions
		expectedResult model.APITestResult
	}{
		{
			name: "WithExecution",
			opts: TestResultsOptions{
				TaskID:    "task1",
				Execution: 0,
				TestName:  "test1",
			},
			expectedResult: s.apiResults["task1_0_test1"],
		},
		{
			name: "WithoutExecution",
			opts: TestResultsOptions{
				TaskID:         "task1",
				EmptyExecution: true,
				TestName:       "test1",
			},
			expectedResult: s.apiResults["task1_1_test1"],
		},
	} {
		s.T().Run(test.name, func(t *testing.T) {
			result, err := s.sc.FindTestResultByTestName(s.ctx, test.opts)
			s.Require().NoError(err)
			s.Equal(test.expectedResult, *result)
		})
	}
}

func (s *testResultsConnectorSuite) TestFindTestResultByTestNameDNE() {
	for _, test := range []struct {
		name string
		opts TestResultsOptions
	}{
		{
			name: "MetadataDNE",
			opts: TestResultsOptions{
				TaskID:    "DNE",
				Execution: 1,
				TestName:  "test1",
			},
		},
		{
			name: "TestObjectDNE",
			opts: TestResultsOptions{
				TaskID:    "task1",
				Execution: 1,
				TestName:  "DNE",
			},
		},
	} {
		s.T().Run(test.name, func(t *testing.T) {
			result, err := s.sc.FindTestResultByTestName(s.ctx, test.opts)
			s.Error(err)
			s.Nil(result)
		})
	}
}

func (s *testResultsConnectorSuite) TestFindTestResultByTestNameEmpty() {
	opts := TestResultsOptions{
		TaskID:    "task1",
		Execution: 1,
	}

	result, err := s.sc.FindTestResultByTestName(s.ctx, opts)
	s.Error(err)
	s.Nil(result)
}
