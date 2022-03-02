package array

import (
	"fmt"
)
func SetZeroes(matrix [][]int)  {
    r := len(matrix)
    c := len(matrix[0])
    
    for i:= 0 ; i < r ; i ++ {
        for j := 0 ; j < c ; j ++ {
            if matrix[i][j] == 0 {
                for k := 0 ; k < c ; k ++ {
                    if matrix[i][k] != 0 {
                   		 matrix[i][k] = -1 
                    }
                }
                for k := 0 ; k < r ; k ++ {
                    fmt.Println(j)
                   if matrix[k][j] != 0 {
                   		 matrix[k][j] = -1 
                    }
                }
                
            }
        }
    }
    
    for i := 0 ; i < r ; i ++ {
        for j := 0 ; j < c ; j ++ {
            if matrix[i][j] == -1 {
          		  matrix[i][j] = 0 
            }
        }
    }
	fmt.Println(matrix)
    
}

func SetZeroes2(matrix [][]int){
	row := len(matrix)
	col := len(matrix[0])
	col0 := false
	for i := 0 ; i < row ; i ++ {
		if matrix[i][0] == 0 {
			col0 = true
		}
		for j := 1 ; j < col ; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		} 
	}
	for i := row-1 ; i >= 0 ; i-- {
		for j := col-1 ; j >= 0 ; j -- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 == true {
			matrix[i][0] = 0
		}
	}
	fmt.Println(matrix)
}

// 1111
// 1011
// 1101
// 0001