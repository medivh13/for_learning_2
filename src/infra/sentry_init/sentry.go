package sentryinit

import (
	"for_learning_2/src/infra/config"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/sirupsen/logrus"
)

func NewSentryClient(conf config.SentryConf, logger *logrus.Logger) error {
	err := sentry.Init(sentry.ClientOptions{
		EnableTracing:    true,
		Dsn:              conf.Dsn,
		TracesSampleRate: conf.TracesSampleRate,
	})
	if err != nil {
		logger.Printf("Sentry initialization failed: %s", err)
		return err
	}

	// Optional: Set global tags, user context, etc.
	sentry.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("env", "production")
		scope.SetUser(sentry.User{
			Email: "jody.almaida@gmail.com",
		})
	})

	// Ensure the flush to make sure events are sent before the program exits
	defer sentry.Flush(2 * time.Second)
	logger.Println("Sentry initialized successfully")
	return nil
}
