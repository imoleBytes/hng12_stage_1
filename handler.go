package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imoleBytes/hng12_stage_1/utils"
)

func HandleClassifyNumber(ctx *gin.Context) {
	num, err := strconv.Atoi(ctx.Query("number"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"number": "alphabet",
			"error":  true,
		})
		return
	}
	var properties []string
	isArmstrong := utils.IsArmstrong(num)
	parity := num%2 == 0

	if isArmstrong {
		switch parity {
		case true:
			properties = append(properties, "armstrong", "even")
		case false:
			properties = append(properties, "armstrong", "odd")
		}
	} else {
		switch parity {
		case true:
			properties = append(properties, "even")
		case false:
			properties = append(properties, "odd")
		}
	}
	fun_fact := utils.GetFunFact(num)
	if fun_fact == nil {
		log.Println("error from fun fact(nil)")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"number": "alphabet",
			"error":  true,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"number":     num,
		"is_prime":   utils.IsPrime(num),
		"is_perfect": utils.IsPerfectNumber(num),
		"properties": properties,
		"digit_sum":  utils.DigitSum(num),
		"fun_fact":   fun_fact,
	})
}
