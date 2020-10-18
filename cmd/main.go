package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/lib/pq"

	"github.com/bhakiyakalimuthu/blogger/internal/rest"
	"github.com/bhakiyakalimuthu/blogger/internal/svc"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	r := chi.NewRouter()
	l := loggerSetup()
	db := svc.NewPostgresStore(l, postgresSetup())
	s := svc.NewService(l, db)
	_ = rest.NewController(l, s).SetupRouter(r)

	http.ListenAndServe(":8080", r)

}

func postgresSetup() *sqlx.DB {
	connStrParts := []string{
		fmt.Sprintf("host=%s", "localhost"),
		fmt.Sprintf("port=%d", 5432),
		fmt.Sprintf("user=%s", "dbuser"),
		fmt.Sprintf("password=%s", "dbpassword"),
		fmt.Sprintf("dbname=%s", "users"),
	}
	connector, _ := pq.NewConnector(strings.Join(connStrParts, " "))
	driver := sql.OpenDB(connector)
	return sqlx.NewDb(driver, "postgres")
}
func loggerSetup() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create zap logger : %v", err)
	}
	return logger
	aa := zap.NewDevelopmentEncoderConfig()
	aa.EncodeLevel = zapcore.CapitalColorLevelEncoder
	bb := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(aa),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	bb.Warn("logger setup done")
	return bb
}
