package graph

import (
	"context"
	"net/http"
	"time"

	"github.com/cynthiawilliamsa/meetmeup/graph/shared"
	"github.com/go-pg/pg/v10"
)

const userLoaderKey = "userloader"

func DataLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*shared.User, []error) {
				var users []*shared.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				return users, []error{err}
			},
		}
		ctx := context.WithValue(r.Context(), userLoaderKey, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

//cast uesrLoaderKey to UserLoader
//resue throughout
func GetUserLoad(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
