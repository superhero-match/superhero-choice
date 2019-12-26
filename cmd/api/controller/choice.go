package controller

import (
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	ctrl "github.com/superhero-choice/cmd/api/controller/model"
	"go.uber.org/zap"
)

// Choice returns the suggestions for the Superhero.
func (ctl *Controller) Choice(c *gin.Context) {
	var req ctrl.Request

	err := c.BindJSON(&req)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed to bind JSON to value of type Request",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	// Check if the chosenSuperhero like the superhero.
	// If so, return is match right away.
	// The likes are being save under the key with the following format -> superheroID.chosenSuperheroID.
	// So to check if the chosenSuperhero liked the superhero --> chosenSuperheroID.superheroID.
	res, err := ctl.Service.GetChoice(req.ChosenSuperheroID + "." + req.SuperheroID)
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.HandleESRequest()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	if res != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"isMatch": true,
		})

		return
	}

	id := strings.ReplaceAll(uuid.New().String(), "-", "")

	err = ctl.Service.StoreChoice(ctrl.Choice{
		ID:                id,
		Choice:            req.Choice,
		SuperheroID:       req.SuperheroID,
		ChosenSuperheroID: req.ChosenSuperheroID,
		CreatedAt:         time.Now().UTC().Format(ctl.Service.TimeFormat),
	})
	if checkError(err, c) {
		ctl.Service.Logger.Error(
			"failed while executing service.HandleESRequest()",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.Service.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"isMatch": false,
	})
}

func checkError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"isMatch": false,
		})

		return true
	}

	return false
}
