package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type TxMiddleware interface {
	DBTransactionMiddleware() gin.HandlerFunc
	RedisTransactionMiddleware() gin.HandlerFunc
}
type MiddlewareRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewMiddlewareRepository(db *gorm.DB, redis *redis.Client) *MiddlewareRepository {
	return &MiddlewareRepository{
		db:    db,
		redis: redis,
	}
}

// StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware : to setup the database transaction middleware
func (r *MiddlewareRepository) DBTransactionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := r.db.Begin()
		log.Print("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			select {
			case <-c.Request.Context().Done():
				log.Println("\033[31m", "[transaction] rolling back transaction due to status code:", c.Writer.Status(), "\033[0m")
				txHandle.Rollback()
			default:
				if err := txHandle.Commit().Error; err != nil {
					log.Print("trx commit error: ", err)
				}
				fmt.Println()
				log.Print("\033[32m", "[transaction] committing transactions", "\033[0m")
			}
		} else {
			fmt.Println()
			log.Println("\033[31m", "[transaction] rolling back transaction due to status code:", c.Writer.Status(), "\033[0m")
			txHandle.Rollback()
		}
	}
}

func (r *MiddlewareRepository) RedisTransactionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := r.redis.TxPipeline()
		log.Print("beginning redis transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Discard()
			}
		}()

		c.Set("redis_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			select {
			case <-c.Request.Context().Done():
				log.Println("\033[31m", "[transaction] discard redis transaction due to status code:", c.Writer.Status(), "\033[0m")
				txHandle.Discard()
			default:
				if _, err := txHandle.Exec(context.TODO()); err != nil {
					log.Print("trx commit error: ", err)
				}
				fmt.Println()
				log.Print("\033[32m", "[transaction] committing redis transactions", "\033[0m")
			}
		} else {
			fmt.Println()
			log.Println("\033[31m", "[transaction] discard redis transaction due to status code:", c.Writer.Status(), "\033[0m")
			txHandle.Discard()
		}
	}
}
