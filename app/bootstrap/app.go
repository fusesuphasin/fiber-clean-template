package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/directoryxx/fiber-clean-template/app/interfaces"
	"github.com/directoryxx/fiber-clean-template/app/routes"
	"github.com/directoryxx/fiber-clean-template/database/gen"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

const idleTimeout = 5 * time.Second

// Dispatch is handle routing
func Dispatch(sqlHandler *gen.Client, ctx context.Context, log interfaces.Logger, redisHandler *redis.Client, enforcer *casbin.Enforcer) {
	app := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	// app.Use(pprof.New())

	routes.RegisterRoute(app, sqlHandler, ctx, log, redisHandler, enforcer)

	go func() {
		if errApp := app.Listen("0.0.0.0:" + os.Getenv("APP_PORT")); errApp != nil {
			log.LogError("%s", errApp)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")

	// if errApp != nil {
	// 	log.LogError("%s", errApp)
	// }
}
