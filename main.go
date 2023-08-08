package main
import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "bufio"
    "errors"
)



func AtoR(str string) (int, bool, error)  {
    chek := false
    var arr = map[string]int{"I" : 1, "V" : 5, "X" : 10, "L" : 50, "C" : 100}
    val, _ := strconv.Atoi(strings.TrimSpace(str))
    
 
     if(val == 0){
        chek = true
        tmp :=  arr[string(str[0])]
        i:= 0
        for i < len(strings.TrimSpace(str)) {
            first := arr[string(str[i])]
            if(first == 0){
                err := errors.New("error")
                return 0, false, err
            }
            
            if(tmp < first){
                val -= tmp
                val += first - tmp
            }else{
                val += first
            }
 
            tmp = first
            i += 1
        }
 
    }
    if(!(val >= 1 && val <= 10)){
       err := errors.New("error")
       return 0, false, err
     }
     
	return val, chek, nil
}



func main() {
    reader := bufio.NewReader(os.Stdin) 
    val, _ := reader.ReadString('\n')
    arr := strings.Split(val, " ")
   
    
	if(len(arr) > 3){
	    fmt.Println("error")
	    return
	}

    first, first1 , err :=  AtoR(arr[0])
    second, second1, err :=  AtoR(arr[2])
    if (err != nil){
        fmt.Println("error")
	    return
    }
    
    
    if(first1 != second1){
        fmt.Println("error")
	    return
    }
    
    var sum int
    
    switch arr[1]{
        case "+":
            sum = first + second
        case "-":
            sum = first - second
        case "/":
            sum = first / second
        case "*":
           sum = first * second
        default:
            fmt.Println("error")
	        return
    }
    
    if(first1){
        var mp = map[int]string{100 : "C", 90 : "XC", 50: "L", 40 : "XL", 10 : "X", 9 : "IX", 5 : "V", 4 : "IV", 1 : "I"}
        vec := [9]int{1, 4, 5, 9, 10, 40, 50, 90, 100}
        result := ""
        tmp := len(vec) - 1
        for(sum > 0){
            if (sum / vec[tmp] > 0){
                for i := 0; i <  sum / vec[tmp]; i++{
                    result += mp[vec[tmp]] 
                    sum -= vec[tmp]
                }
            }else {
                tmp--
            }
        }
        fmt.Println(result)
    }else{
        fmt.Println(sum)
    }
}
