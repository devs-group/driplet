package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/devs-group/driplet/scheduler/calculate_points"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "scheduler",
		Usage: "job scheduler for cloud run",
		Commands: []*cli.Command{
			{
				Name:  "wait",
				Usage: "blocking process",
				Action: func(c *cli.Context) error {
					ctx, stop := signal.NotifyContext(context.Background(),
						os.Interrupt,    // SIGINT (Ctrl+C)
						syscall.SIGTERM, // SIGTERM (Docker stop)
						syscall.SIGQUIT, // SIGQUIT
					)
					defer stop()

					slog.Info("starting blocking process, press Ctrl+C to exit gracefully...")

					// Wait for context cancellation
					<-ctx.Done()

					slog.Info("received shutdown signal, exiting gracefully...")
					return nil
				},
			},
			{
				Name:  "calc-points",
				Usage: "calculate points for users",
				Action: func(c *cli.Context) error {
					return calculate_points.Run()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
