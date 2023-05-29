package helper

import (
	"fmt"
	"math/rand"
	"minder/src/server/repository"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StringContains(arrayString []string, search string) bool {
	// iterate over the array and compare given string to each element
	for _, value := range arrayString {
		if value == search {
			return true
		}
	}
	return false
}

func RandomNumber(max int) int {
	return rand.Intn(max-0) + 0
}

func RandomNumberV2(max int) int {
	return rand.Intn(max-1) + 1
}

func MakeNumberRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func ArrayIntShuffle(array []int) []int {
	rand.Shuffle(len(array), func(i, j int) {
		array[i], array[j] = array[j], array[i]
	})

	return array
}

func FindIndex(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}
	return -1
}

func ExtractNumberFromString(text string) int {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(text, -1)
	var number int
	for _, element := range submatchall {
		number, _ = strconv.Atoi(element)
		break
	}

	return number
}

func GenerateInvoiceNumber(repo repository.PurchaseRepo) string {
	dateNow := time.Now().Format("20060102")

	lastPurchase, err := repo.GetLastPurchaseByDate(dateNow)
	if err != nil {
		return fmt.Sprintf("INV-%sMI0001", dateNow)
	}

	splitInvoiceNum := strings.Split(lastPurchase.InvoiceNumber, "MI")
	stringSequence := splitInvoiceNum[len(splitInvoiceNum)-1]
	lastSequence := ExtractNumberFromString(stringSequence)
	fmt.Printf("lastSequence => %v", lastSequence)
	return fmt.Sprintf("INV-%sMI%04d", dateNow, lastSequence+1)
}
