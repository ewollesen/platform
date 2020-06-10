package mongo

import (
	"context"
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"github.com/tidepool-org/platform/confirmation/store"
	"github.com/tidepool-org/platform/errors"
	"github.com/tidepool-org/platform/log"
	storeStructuredMongo "github.com/tidepool-org/platform/store/structured/mongo"
)

type Store struct {
	*storeStructuredMongo.Store
}

func NewStore(cfg *storeStructuredMongo.Config, lgr log.Logger) (*Store, error) {
	str, err := storeStructuredMongo.NewStore(cfg, lgr)
	if err != nil {
		return nil, err
	}

	return &Store{
		Store: str,
	}, nil
}

func (s *Store) EnsureIndexes() error {
	ssn := s.confirmationSession()
	defer ssn.Close()
	return ssn.EnsureIndexes()
}

func (s *Store) NewConfirmationSession() store.ConfirmationSession {
	return s.confirmationSession()
}

func (s *Store) confirmationSession() *ConfirmationSession {
	return &ConfirmationSession{
		Session: s.Store.NewSession("confirmations"),
	}
}

type ConfirmationSession struct {
	*storeStructuredMongo.Session
}

func (c *ConfirmationSession) EnsureIndexes() error {
	return c.EnsureAllIndexes([]mgo.Index{
		// Additional indexes are also created in `hydrophone`.
		{Key: []string{"email"}, Background: true},
		{Key: []string{"status"}, Background: true},
		{Key: []string{"type"}, Background: true},
		{Key: []string{"userId"}, Background: true},
	})
}

func (c *ConfirmationSession) DeleteUserConfirmations(ctx context.Context, userID string) error {
	if ctx == nil {
		return errors.New("context is missing")
	}
	if userID == "" {
		return errors.New("user id is missing")
	}

	if c.IsClosed() {
		return errors.New("session closed")
	}

	now := time.Now()
	logger := log.LoggerFromContext(ctx).WithField("userId", userID)

	selector := bson.M{
		"$or": []bson.M{
			{"userId": userID},
			{"creatorId": userID},
		},
	}
	changeInfo, err := c.C().RemoveAll(selector)
	logger.WithFields(log.Fields{"changeInfo": changeInfo, "duration": time.Since(now) / time.Microsecond}).WithError(err).Debug("DeleteUserConfirmations")
	if err != nil {
		return errors.Wrap(err, "unable to delete user confirmations")
	}

	return nil
}
