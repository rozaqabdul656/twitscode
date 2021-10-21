package controller

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func BinaryToDecimal(c *gin.Context) {
	parameter := c.Query("binary")

	if parameter == "" {
		c.JSON(500, gin.H{"status": "error", "error": true, "message": "internal Server Error"})
		return
	}

	desimal := 0
	count := 0.0
	i := 0
	number, _ := strconv.Atoi(parameter)
	for number != 0 {
		i = number % 10
		desimal += i * int(math.Pow(2.0, count))
		number = number / 10
		count++
	}
	fmt.Println(desimal)
	c.JSON(200, gin.H{"status": true, "error": false, "message": "Succes", "data": desimal})
	return
}

func DecimalToBinary(c *gin.Context) {
	parameter := c.Query("decimal")

	if parameter == "" {
		c.JSON(500, gin.H{"status": "error", "error": true, "message": "internal Server Error"})
		return
	}

	biner := 0
	count := 1
	i := 0
	nomer, _ := strconv.Atoi(parameter)

	for nomer != 0 {
		i = nomer % 2
		nomer = nomer / 2
		biner += i * count
		count *= 10

	}
	c.JSON(200, gin.H{"status": true, "error": false, "message": "Succes", "data": biner})
	return
}
func Polyndrome(Polyndrome string) {
	hasil := ""
	hasil2 := ""

	words := strings.Fields(Polyndrome)
	for i := 0; i < len(words); i++ {
		chars := []rune(words[i])
		indexPembanding := 0
		if i != len(words)-1 {
			indexPembanding = i + 1
		} else {
			break
		}
		hasilLogic1 := ""
		hasilLogic2 := ""

		indexmundur := 1
		x := 0
		for {
			KataPembanding := []rune(words[indexPembanding])
			if len(KataPembanding)-indexmundur < 0 || len(chars) == x {
				break
			}
			if chars[x] == KataPembanding[len(KataPembanding)-indexmundur] {
				hasilLogic1 += string(chars[x])
				hasilLogic2 = string(KataPembanding[len(KataPembanding)-indexmundur]) + hasilLogic2
				x++

			}

			indexmundur++

		}
		if len(hasil) < len(hasilLogic1) {
			hasil = hasilLogic1
			hasil2 = hasilLogic2
		}

	}

	fmt.Println("Kata Logic : " + Polyndrome + "  , Hasiill Logic Test Case  No 2: " + hasil + "  " + hasil2)

}
