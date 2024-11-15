package logic

import (
	"context"
	"log/slog"

	"github.com/pkg/errors"
	"github.com/pocketbase/pocketbase/core"
)

const (
	downloadTypeHTTP = 1
)

type DownloadTask interface {
	ExecuteDownloadTask(ctx context.Context, id string) error
}

type downloadTask struct {
	pocketbaseApp core.App
}

func (d downloadTask) ExecuteDownloadTask(ctx context.Context, id string) error {
	logger := d.pocketbaseApp.Logger().With(slog.String("id", id))

	downloadTask, err := d.pocketbaseApp.Dao().FindRecordById("download_task", id)
	if err != nil {
		logger.With("err", err).Error("failed to find download task with id %s", id)
		return err
	}

	downloadType := downloadTask.GetInt("donwload_type")
	if downloadType != downloadTypeHTTP {
		logger.With(slog.Int("download_type", downloadType)).Error("unsupported download type")
		return errors.New("unsupported download type")
	}
}
