package v1_test

import (
	"github.com/ant0ine/go-json-rest/rest"

	"github.com/tidepool-org/platform/data/deduplicator"
	testDataDeduplicator "github.com/tidepool-org/platform/data/deduplicator/test"
	dataStoreDEPRECATED "github.com/tidepool-org/platform/data/storeDEPRECATED"
	testDataStoreDEPRECATED "github.com/tidepool-org/platform/data/storeDEPRECATED/test"
	"github.com/tidepool-org/platform/metric"
	testMetric "github.com/tidepool-org/platform/metric/test"
	"github.com/tidepool-org/platform/permission"
	permissionTest "github.com/tidepool-org/platform/permission/test"
	"github.com/tidepool-org/platform/service"
	syncTaskStore "github.com/tidepool-org/platform/synctask/store"
	testSyncTaskStore "github.com/tidepool-org/platform/synctask/store/test"
	"github.com/tidepool-org/platform/test"
)

type RespondWithInternalServerFailureInput struct {
	message string
	failure []interface{}
}

type RespondWithStatusAndErrorsInput struct {
	statusCode int
	errors     []*service.Error
}

type RespondWithStatusAndDataInput struct {
	statusCode int
	data       interface{}
}

type TestContext struct {
	*test.Mock
	RespondWithErrorInputs                 []*service.Error
	RespondWithInternalServerFailureInputs []RespondWithInternalServerFailureInput
	RespondWithStatusAndErrorsInputs       []RespondWithStatusAndErrorsInput
	RespondWithStatusAndDataInputs         []RespondWithStatusAndDataInput
	MetricClientImpl                       *testMetric.Client
	PermissionClientImpl                   *permissionTest.Client
	DataDeduplicatorFactoryImpl            *testDataDeduplicator.Factory
	DataSessionImpl                        *testDataStoreDEPRECATED.DataSession
	SyncTaskSessionImpl                    *testSyncTaskStore.SyncTaskSession
}

func NewTestContext() *TestContext {
	return &TestContext{
		MetricClientImpl:            testMetric.NewClient(),
		PermissionClientImpl:        permissionTest.NewClient(),
		DataDeduplicatorFactoryImpl: testDataDeduplicator.NewFactory(),
		DataSessionImpl:             testDataStoreDEPRECATED.NewDataSession(),
		SyncTaskSessionImpl:         testSyncTaskStore.NewSyncTaskSession(),
	}
}

func (t *TestContext) Response() rest.ResponseWriter {
	panic("Unexpected invocation of Response on TestContext")
}

func (t *TestContext) RespondWithError(err *service.Error) {
	t.RespondWithErrorInputs = append(t.RespondWithErrorInputs, err)
}

func (t *TestContext) RespondWithInternalServerFailure(message string, failure ...interface{}) {
	t.RespondWithInternalServerFailureInputs = append(t.RespondWithInternalServerFailureInputs, RespondWithInternalServerFailureInput{message, failure})
}

func (t *TestContext) RespondWithStatusAndErrors(statusCode int, errors []*service.Error) {
	t.RespondWithStatusAndErrorsInputs = append(t.RespondWithStatusAndErrorsInputs, RespondWithStatusAndErrorsInput{statusCode, errors})
}

func (t *TestContext) RespondWithStatusAndData(statusCode int, data interface{}) {
	t.RespondWithStatusAndDataInputs = append(t.RespondWithStatusAndDataInputs, RespondWithStatusAndDataInput{statusCode, data})
}

func (t *TestContext) MetricClient() metric.Client {
	return t.MetricClientImpl
}

func (t *TestContext) PermissionClient() permission.Client {
	return t.PermissionClientImpl
}

func (t *TestContext) DataDeduplicatorFactory() deduplicator.Factory {
	return t.DataDeduplicatorFactoryImpl
}

func (t *TestContext) DataSession() dataStoreDEPRECATED.DataSession {
	return t.DataSessionImpl
}

func (t *TestContext) SyncTaskSession() syncTaskStore.SyncTaskSession {
	return t.SyncTaskSessionImpl
}

func (t *TestContext) Expectations() {
	t.Mock.Expectations()
	t.MetricClientImpl.Expectations()
	t.PermissionClientImpl.AssertOutputsEmpty()
	t.DataDeduplicatorFactoryImpl.Expectations()
	t.DataSessionImpl.Expectations()
	t.SyncTaskSessionImpl.Expectations()
}
