package core

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	logTimeLayout  = "2006-01-02T15:04:05.000Z07:00" // RFC3339
	fileTimeLayout = "20060102_150405.000"
)

func InitLogger(level, dir string) (*os.File, error) {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(strings.TrimSpace(level))); err != nil {
		return nil, fmt.Errorf("[Log] invalid level %q (want debug|info|warn|error): %w", level, err)
	}

	if dir == "" {
		return nil, fmt.Errorf("[Log] empty log dir")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, fmt.Errorf("[Log] create dir %q: %w", dir, err)
	}

	file, err := createLogFile(dir)
	if err != nil {
		return nil, err
	}

	h := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level:       lvl,
		ReplaceAttr: formatTime,
	})
	slog.SetDefault(slog.New(h))

	return file, nil
}

func formatTime(groups []string, a slog.Attr) slog.Attr {
	if len(groups) == 0 && a.Key == slog.TimeKey && a.Value.Kind() == slog.KindTime {
		a.Value = slog.StringValue(a.Value.Time().Format(logTimeLayout))
	}
	return a
}

func createLogFile(dir string) (*os.File, error) {
	name := fmt.Sprintf("app_%s_%d.log", time.Now().Format(fileTimeLayout), os.Getpid())
	path := filepath.Join(dir, name)

	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return nil, fmt.Errorf("[Log] open file %q: %w", path, err)
	}
	return f, nil
}
