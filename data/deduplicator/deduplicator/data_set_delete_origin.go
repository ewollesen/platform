package deduplicator

import (
	"context"

	"github.com/tidepool-org/platform/data"
	dataStoreDEPRECATED "github.com/tidepool-org/platform/data/storeDEPRECATED"
	dataTypesCommonOrigin "github.com/tidepool-org/platform/data/types/common/origin"
	dataTypesUpload "github.com/tidepool-org/platform/data/types/upload"
	"github.com/tidepool-org/platform/errors"
)

const DataSetDeleteOriginName = "org.tidepool.deduplicator.dataset.delete.origin"

type DataSetDeleteOrigin struct {
	*Base
}

func NewDataSetDeleteOrigin() (*DataSetDeleteOrigin, error) {
	base, err := NewBase(DataSetDeleteOriginName, "1.0.0")
	if err != nil {
		return nil, err
	}

	return &DataSetDeleteOrigin{
		Base: base,
	}, nil
}

func (d *DataSetDeleteOrigin) New(dataSet *dataTypesUpload.Upload) (bool, error) {
	return d.Get(dataSet)
}

func (d *DataSetDeleteOrigin) Get(dataSet *dataTypesUpload.Upload) (bool, error) {
	if found, err := d.Base.Get(dataSet); err != nil || found {
		return found, err
	}

	return dataSet.HasDeduplicatorNameMatch("org.tidepool.continuous.origin"), nil // TODO: DEPRECATED
}

func (d *DataSetDeleteOrigin) Open(ctx context.Context, session dataStoreDEPRECATED.DataSession, dataSet *dataTypesUpload.Upload) (*dataTypesUpload.Upload, error) {
	if ctx == nil {
		return nil, errors.New("context is missing")
	}
	if session == nil {
		return nil, errors.New("session is missing")
	}
	if dataSet == nil {
		return nil, errors.New("data set is missing")
	}

	if dataSet.HasDataSetTypeContinuous() {
		dataSet.SetActive(true)
	}

	return d.Base.Open(ctx, session, dataSet)
}

func (d *DataSetDeleteOrigin) AddData(ctx context.Context, session dataStoreDEPRECATED.DataSession, dataSet *dataTypesUpload.Upload, dataSetData data.Data) error {
	if ctx == nil {
		return errors.New("context is missing")
	}
	if session == nil {
		return errors.New("session is missing")
	}
	if dataSet == nil {
		return errors.New("data set is missing")
	}
	if dataSetData == nil {
		return errors.New("data set data is missing")
	}

	if dataSet.HasDataSetTypeContinuous() {
		dataSetData.SetActive(true)
	}

	if originIDs := getOriginIDs(dataSetData); len(originIDs) > 0 {
		if err := session.ArchiveDataSetDataUsingOriginIDs(ctx, dataSet, originIDs); err != nil {
			return err
		}
		if err := d.Base.AddData(ctx, session, dataSet, dataSetData); err != nil {
			return err
		}
		return session.DeleteArchivedDataSetData(ctx, dataSet)
	}

	return d.Base.AddData(ctx, session, dataSet, dataSetData)
}

func (d *DataSetDeleteOrigin) Close(ctx context.Context, session dataStoreDEPRECATED.DataSession, dataSet *dataTypesUpload.Upload) error {
	if ctx == nil {
		return errors.New("context is missing")
	}
	if session == nil {
		return errors.New("session is missing")
	}
	if dataSet == nil {
		return errors.New("data set is missing")
	}

	if dataSet.HasDataSetTypeContinuous() {
		return nil
	}

	return d.Base.Close(ctx, session, dataSet)
}

func getOriginIDs(dataSetData data.Data) []string {
	var originIDs []string
	for _, dataSetDatum := range dataSetData {
		if origin := dataSetDatum.(dataTypesCommonOrigin.Getter).GetOrigin(); origin != nil && origin.ID != nil {
			originIDs = append(originIDs, *origin.ID)
		}
	}
	return originIDs
}