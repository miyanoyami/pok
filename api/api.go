package main

import (
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/miyanoyami/pok/domain/service"
)

// DamageRequest ダメージ計算に必要なデータ
type DamageRequest struct {
	Level  int64 `json:"level"`
	Power  int64 `json:"power"`  // 技威力
	Attack int64 `json:"attack"` // こうげき・とくこうのどちらか
	Block  int64 `json:"block"`  // ぼうぎょ・とくぼうのどちらか
}

// Validate バリデーション
func (d *DamageRequest) Validate() error {
	if d.Level < 1 {
		return errors.New("Level range 1 ~ 100")
	}

	if d.Power < 1 {
		return errors.New("Power range 1 ~ ")
	}

	if d.Attack < 1 {
		return errors.New("Attack range 1 ~ ")
	}

	if d.Block < 1 {
		return errors.New("Block range 1 ~ ")
	}
	return nil
}

// DamageResponse ダメージ計算結果
type DamageResponse struct {
	Min int64 `json:"min"`
	Max int64 `json:"max"`
}

func main() {
	router := iris.New()
	router.Handle(
		http.MethodPost,
		"/damage",
		DamageHandler,
	)
	router.Handle(
		http.MethodPost,
		"/status",
		StatusHandler,
	)

	router.Listen(":8080")
}

func responseError(ctx iris.Context, title string, err error) {
	ctx.StopWithProblem(
		iris.StatusBadRequest,
		iris.NewProblem().
			Title(title).
			DetailErr(err))
}

func DamageHandler(ctx iris.Context) {
	req := &DamageRequest{}

	if err := ctx.ReadJSON(req); err != nil {
		responseError(ctx, "calculationFailed", err)
		return
	}

	if err := req.Validate(); err != nil {
		responseError(ctx, "calculationFailed", err)
		return
	}

	// 本体の処理
	min, max := service.CalculateDamage(
		req.Level,
		req.Power,
		req.Attack,
		req.Block,
	)

	ctx.JSON(&DamageResponse{
		Min: min,
		Max: max,
	})
}

func StatusHandler(ctx iris.Context) {
	ctx.JSON(`{"statusHandler": "未実装です！"}`)
}
