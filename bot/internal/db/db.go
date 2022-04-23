package db

import (
	"context"
	"github.com/supperdoggy/taro-pizda/structs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"strings"
)

type IDB interface {
	GetRandomTaro(ctx context.Context) (taro structs.Taro, err error)
	GetTaro(ctx context.Context, id string) (taro structs.Taro, err error)
}

type db struct {
	session *mongo.Client

	warningCol *mongo.Collection
	adviceCol  *mongo.Collection
	ruLocCol   *mongo.Collection
	picCol     *mongo.Collection

	logger *zap.Logger
}

type obj map[string]interface{}
type arr []interface{}

func NewDB(l *zap.Logger, url, dbName, warningCol, adviceCol, picCol, ruLocCol string, ctx context.Context) (IDB, error) {
	session, err := mongo.Connect(ctx)
	if err != nil {
		return nil, err
	}

	d := db{
		session: session,
		logger:  l,

		warningCol: session.Database(dbName).Collection(warningCol),
		adviceCol:  session.Database(dbName).Collection(adviceCol),
		picCol:     session.Database(dbName).Collection(picCol),
		ruLocCol:   session.Database(dbName).Collection(ruLocCol),
	}

	return &d, nil
}

func (d *db) GetRandomTaro(ctx context.Context) (taro structs.Taro, err error) {
	// to get random doc
	var cursor *mongo.Cursor
	cursor, err = d.adviceCol.Aggregate(ctx, arr{obj{"$sample": obj{"size": 1}}})
	if err != nil {
		d.logger.Error("error getting random taro from mongo", zap.Error(err))
		return
	}
	for cursor.Next(ctx) {
		err = cursor.Decode(&taro.Advice)
		if err != nil {
			d.logger.Error("error decoding random advice", zap.Error(err))
			return
		}
		break
	}

	taro.ID = taro.Advice.ID
	taro, err = d.GetTaro(ctx, taro.ID)
	if err != nil {
		d.logger.Error("error getting random taro", zap.Error(err), zap.Any("id", taro.ID))
		return
	}

	return
}

func (d *db) GetTaro(ctx context.Context, id string) (taro structs.Taro, err error) {
	taro.ID = id
	taro.Advice, err = d.GetTaroAdvice(ctx, taro.ID)
	if err != nil {
		d.logger.Error("error getting advice meaning for taro", zap.Error(err), zap.Any("id", taro.ID))
		return
	}

	taro.Warning, err = d.GetTaroWarning(ctx, taro.ID)
	if err != nil {
		d.logger.Error("error getting warning meaning for taro", zap.Error(err), zap.Any("id", taro.ID))
		return
	}

	taro.Pic, err = d.GetTaroPic(ctx, taro.ID)
	if err != nil {
		d.logger.Error("error getting pic for taro", zap.Error(err), zap.Any("id", taro.ID))
		return
	}

	// get name
	id = strings.Replace(taro.ID, "_reversed", "", 1)
	taro.Loc, err = d.GetTaroLoc(ctx, id)
	if err != nil {
		d.logger.Error("error getting taro loc", zap.Error(err), zap.Any("id", id))
		return
	}

	return
}

func (d *db) GetTaroWarning(ctx context.Context, id string) (meaning structs.TaroMeaning, err error) {
	var cursor *mongo.Cursor
	cursor, err = d.warningCol.Find(ctx, obj{"_id": id})
	if err != nil {
		d.logger.Error("error finding warning meaning", zap.Error(err), zap.Any("id", id))
		return
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&meaning)
		if err != nil {
			d.logger.Error("error decoding warning meaning", zap.Error(err), zap.Any("id", id))
			return
		}
		break
	}

	return
}

func (d *db) GetTaroAdvice(ctx context.Context, id string) (meaning structs.TaroMeaning, err error) {
	var cursor *mongo.Cursor
	cursor, err = d.adviceCol.Find(ctx, obj{"_id": id})
	if err != nil {
		d.logger.Error("error finding advice meaning", zap.Error(err), zap.Any("id", id))
		return
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&meaning)
		if err != nil {
			d.logger.Error("error decoding advice meaning", zap.Error(err), zap.Any("id", id))
			return
		}
		break
	}

	return
}

func (d *db) GetTaroPic(ctx context.Context, id string) (pic structs.TaroPic, err error) {
	var cursor *mongo.Cursor
	cursor, err = d.picCol.Find(ctx, obj{"_id": id})
	if err != nil {
		d.logger.Error("error finding pic", zap.Error(err), zap.Any("id", id))
		return
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&pic)
		if err != nil {
			d.logger.Error("error decoding pic", zap.Error(err), zap.Any("id", id))
			return
		}
		break
	}

	return
}

func (d *db) GetTaroLoc(ctx context.Context, id string) (pic structs.TaroLoc, err error) {
	var cursor *mongo.Cursor
	cursor, err = d.ruLocCol.Find(ctx, obj{"_id": id})
	if err != nil {
		d.logger.Error("error finding taro loc", zap.Error(err), zap.Any("id", id))
		return
	}

	for cursor.Next(ctx) {
		err = cursor.Decode(&pic)
		if err != nil {
			d.logger.Error("error decoding taro loc", zap.Error(err), zap.Any("id", id))
			return
		}
		break
	}

	return
}
