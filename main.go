package main

import (
	"errors"
	"fmt"
)

func main() {
	var arr1 = []int{7, 1, 2, 9, 7, 4, 1}
	value, err := checkHighest(arr1)
	fmt.Println(value, err)
}

func checkHighest(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("array kosong")
	}
	deret, err := findDeret(arr)
	if err != nil {
		return 0, err
	}
	fmt.Println(deret)
	deretReverse, err := findDeretReverse(arr)
	if err != nil {
		return 0, err
	}
	fmt.Println(deretReverse)
	if len(deret) == len(deretReverse) && deret[len(deret)-1] == deret[len(deretReverse)-1] {
		return deret[len(deret)-1], nil
	} else {
		return 0, errors.New("tidak ada deret angka yang sesuai")
	}

}

func findDeret(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return []int{}, errors.New("array kosong")
	}
	deret := []int{}
	for i := 0; i < len(arr); i++ {
		if i != len(arr)-1 {
			if arr[i+1]-arr[i] == 1 {
				deret = append(deret, arr[i])
			}
			if len(deret) != 0 && arr[i+1]-arr[i] != 1 {
				deret = append(deret, arr[i])
				break
			}
		}
	}
	if len(deret) == 0 {
		return []int{}, errors.New("tidak ada angka yang berurutan dalam array")
	}
	return deret, nil
}

func findDeretReverse(arr []int) ([]int, error) {
	if len(arr) == 0 {
		return []int{}, errors.New("array kosong")
	}
	deret := []int{}
	for i := len(arr) - 1; i > 0; i-- {
		if i != 1 {
			if arr[i-1]-arr[i] == 1 {
				deret = append(deret, arr[i])
			}
			if len(deret) != 0 && arr[i-1]-arr[i] != 1 {
				deret = append(deret, arr[i])
				break
			}
		}
	}
	if len(deret) == 0 {
		return []int{}, errors.New("tidak ada angka yang berurutan dalam array")
	}
	return deret, nil
}