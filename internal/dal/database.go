package dal

import (
	"StadiumSlotBot/internal/options"
	"StadiumSlotBot/internal/utils"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *Orm = &Orm{
	isInited: false,
	locker: new(sync.Once),
}

// Struct for access to database
type Orm struct {
	pool *pgxpool.Pool
	locker *sync.Once
	isInited bool
}

func (o *Orm) Init(connString string) error {
	if DB.isInited {
		return nil
	}

	var initPoolerr error
	o.locker.Do(func() {
		pool, err := pgxpool.New(context.Background(), *options.StadiumSlotBotOptions.DBConnectionString)

		if err != nil {
			o.isInited = false
			initPoolerr = fmt.Errorf("Can't init database pool: " + err.Error())
			return
		}		

		pingDBErr := pool.Ping(context.Background())	
		if pingDBErr != nil {
			utils.Logger.Error("Can't connect with database: " + pingDBErr.Error())
			panic("Can't connect with database: " + pingDBErr.Error())
		}

		o.pool = pool
		o.isInited = true
	})
	return initPoolerr
}

func (o *Orm) getOneValue(query string, ctx *context.Context) (string, error) {
	rows, err := o.pool.Query(*ctx, query)
	
	if err != nil {
		return "", err
	}

	defer rows.Close()

	rows.Next()
	var value string
	err = rows.Scan(&value)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return value, nil
}

func (o *Orm) getOneColumn(query string, ctx *context.Context) ([]string, error) {
	rows, err := o.pool.Query(*ctx, query)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []string
	for rows.Next() {
		var value string
		err := rows.Scan(&value)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
    	}

		result = append(result, value)
	}
	return result, nil
}

func (o *Orm) getManyColumns(query string, ctx *context.Context) ([][]string, error) {
	rows, err := o.pool.Query(*ctx, query)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	descrs := rows.FieldDescriptions()
	ln := len(descrs)
	vals := make([]interface{}, ln)
	for i := range descrs {
		vals[i] = new(string)
	}

	var result [][]string
	for rows.Next() {
		currRes := make([]any, ln)
		for i := range descrs {
			currRes[i] = new(string)
		}
		scanErr := rows.Scan(currRes...)
		if scanErr != nil {
			fmt.Println(scanErr.Error())
			return nil, err
		}
	}
	return result, nil
}

type Races struct {
	RaceDate time.Time
	RaceName string
	RaceInfo string
}