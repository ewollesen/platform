package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/tidepool-org/platform/data/types/blood/glucose/summary"
	"github.com/tidepool-org/platform/errors"
	storeStructuredMongo "github.com/tidepool-org/platform/store/structured/mongo"
)

type SummaryRepository struct {
	*storeStructuredMongo.Repository
}

func (d *SummaryRepository) EnsureIndexes() error {
	return d.CreateAllIndexes(context.Background(), []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "_userId", Value: 1},
			},
			Options: options.Index().
				SetBackground(true).
				SetName("UserID"),
		},
		{
			Keys: bson.D{
				{Key: "lastUpdated", Value: 1},
			},
			Options: options.Index().
				SetBackground(true).
				SetName("LastUpdated"),
		},
	})
}

func (d *SummaryRepository) GetSummary(ctx context.Context, id string) (*summary.Summary, error) {
	if ctx == nil {
		return nil, errors.New("context is missing")
	}
	if id == "" {
		return nil, errors.New("summary UserID is missing")
	}

	//logger := log.LoggerFromContext(ctx).WithField("id", id)

	var summary summary.Summary
	selector := bson.M{
		"userId": id,
	}

	err := d.FindOne(ctx, selector).Decode(&summary)

	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &summary, err
}

func (d *SummaryRepository) UpdateSummary(ctx context.Context, summary *summary.Summary) (*summary.Summary, error) {
	if ctx == nil {
		return nil, errors.New("context is missing")
	}
	if summary == nil {
		return nil, errors.New("summary object is missing")
	}

	if summary.UserID == "" {
		return nil, errors.New("summary missing UserID")
	}

	opts := options.Replace().SetUpsert(true)
	filter := bson.M{"userId": summary.UserID}

	_, err := d.ReplaceOne(ctx, filter, summary, opts)

	//logger := log.LoggerFromContext(ctx).WithField("id", id)

	return summary, err
}

func (d *SummaryRepository) GetAgedSummaries(ctx context.Context, lastUpdated time.Time) ([]*summary.Summary, error) {
	if ctx == nil {
		return nil, errors.New("context is missing")
	}

	var summaries []*summary.Summary
	selector := bson.M{
		"lastUpdated": bson.M{"$lte": lastUpdated},
	}

	cursor, err := d.Find(ctx, selector)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "unable to get aged summaries")
	}

	if err = cursor.All(ctx, &summaries); err != nil {
		return nil, errors.Wrap(err, "unable to decode aged summaries")
	}

	if summaries == nil {
		summaries = []*summary.Summary{}
	}
	return summaries, nil
}

func (d *SummaryRepository) GetLastUpdated(ctx context.Context) (time.Time, error) {
	var lastUpdated time.Time
	var summaries []*summary.Summary

	if ctx == nil {
		return lastUpdated, errors.New("context is missing")
	}

	selector := bson.M{}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "lastUpdated", Value: -1}})
	findOptions.SetLimit(1)

	cursor, err := d.Find(ctx, selector, findOptions)

	if err != nil {
		return lastUpdated, errors.Wrap(err, "unable to get last cbg date")
	}

	if err = cursor.All(ctx, &summaries); err != nil {
		return lastUpdated, errors.Wrap(err, "unable to decode last cbg date")
	}

	if summaries != nil {
		if summaries[0].LastUpdated != nil {
			lastUpdated = *summaries[0].LastUpdated
		}
	} else {
		return time.Time{}, nil
	}

	return lastUpdated, nil
}

func (d *SummaryRepository) UpdateLastUpdated(ctx context.Context, id string) (time.Time, error) {
	// truncate+utc for consistency
	timestamp := time.Now().UTC().Truncate(time.Millisecond)
	if ctx == nil {
		return timestamp, errors.New("context is missing")
	}

	if id == "" {
		return timestamp, errors.New("user id is missing")
	}

	selector := bson.M{"userId": id}

	update := bson.M{
		"$set": bson.M{
			"lastUpdated": &timestamp,
		},
	}

	_, err := d.UpdateOne(ctx, selector, update)

	if err != nil {
		return timestamp, errors.Wrap(err, "unable to update lastUpdated date")
	}

	return timestamp, nil
}
