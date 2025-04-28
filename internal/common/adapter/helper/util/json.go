package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sport-results-pocessor/internal/common/adapter/logger"

	"go.uber.org/zap"
)

func Unmarshal(ctx context.Context, log logger.Logger, t any, readCloser io.ReadCloser) error {
	body, err := io.ReadAll(readCloser)
	if err != nil {
		log.WithMethod(ctx, "Unmarshal").Error("cannot read request body", zap.Error(err))
		return err
	}

	if err = readCloser.Close(); err != nil {
		log.WithMethod(ctx, "Unmarshal").Error("cannot close request body", zap.Error(err))
	}

	if err = json.Unmarshal(body, t); err != nil {
		log.WithMethod(ctx, "Unmarshal").Error("cannot unmarshal request body",
			zap.Error(err),
			zap.ByteString("request.body", body),
			zap.String("request.body.type", fmt.Sprintf("%T", t)))
		return err
	}

	return nil
}
