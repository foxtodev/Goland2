package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func transformSQL1(str string, c ...interface{}) { // Full Shnyaga)))

	strSQL := strings.Split(str, "?")

	result := ""
	var (
		resArgs    []interface{}
		resArgsStr []string // for output as in the task
	)

	for i, s := range c {
		result += strSQL[i]

		switch s.(type) {
		case []int:
			strSlice := []string{}
			k := reflect.ValueOf(s)
			for i := 0; i < k.Len(); i++ {
				strSlice = append(strSlice, "?")
				resArgsStr = append(resArgsStr, strconv.Itoa(int(k.Index(i).Int()))) // for output as in the task
				resArgs = append(resArgs, int(k.Index(i).Int()))
			}
			result += strings.Join(strSlice, ", ")
		case []string:
			strSlice := []string{}
			k := reflect.ValueOf(s)
			for i := 0; i < k.Len(); i++ {
				strSlice = append(strSlice, "?")
				resArgsStr = append(resArgsStr, "\""+k.Index(i).String()+"\"") // for output as in the task
				resArgs = append(resArgs, k.Index(i).String())
			}
			result += strings.Join(strSlice, ", ")
		case bool:
			result += "?"
			resArgsStr = append(resArgsStr, strconv.FormatBool(s.(bool))) // for output as in the task
			resArgs = append(resArgs, s.(bool))
		case float64:
			result += "?"
			resArgsStr = append(resArgsStr, strconv.FormatFloat(s.(float64), 'f', -1, 64)) // for output as in the task
			resArgs = append(resArgs, s.(float64))
		case int:
			result += "?"
			resArgsStr = append(resArgsStr, strconv.Itoa(s.(int))) // for output as in the task
			resArgs = append(resArgs, s.(int))
		case string:
			result += "?"
			resArgsStr = append(resArgsStr, "\""+s.(string)+"\"") // for output as in the task
			resArgs = append(resArgs, s.(string))
		default:
			// error unknow type
		}
	}

	fmt.Printf("\"%s\", []interface{}{ %s }\n", result, strings.Join(resArgsStr, ", "))  // output as in the task

}

func transformSQL2(str string, c ...interface{}) (string, []interface{}) {

	strSQL := strings.Split(str, "?")
	result := ""
	var (
		s       []string // for output as in the task
		resArgs []interface{}
	)

	for i, v := range c {
		result += strSQL[i]
		st := fmt.Sprint(v)
		if st[:1] == "[" {
			sa := strings.Split(st[1:len(st)-1], " ")
			for _, si := range sa {
				s = append(s, si) // for output as in the task
				result += "?,"
			}
			result = result[:len(result)-1]
			k := reflect.ValueOf(v)
			for i := 0; i < k.Len(); i++ {
				resArgs = append(resArgs, k.Index(i).Interface())
			}
		} else {
			s = append(s, st) // for output as in the task
			result += "?"
			resArgs = append(resArgs, v)
		}
	}
	
	// fmt.Printf("\"%s\", []interface{}{ %s }\n", result, strings.Join(s, ", ")) // output as in the task

	return result, resArgs

}

func transformSQL3(str string, c ...interface{}) (string, []interface{}) {

	strSQL := strings.Split(str, "?")

	resStr := ""
	var resArgs []interface{}

	for i, s := range c {

		resStr += strSQL[i]
		rv := reflect.ValueOf(s)

		if rv.Type().Kind() == reflect.Slice {

			for i := 0; i < rv.Len(); i++ {
				resStr += "?,"
				resArgs = append(resArgs, rv.Index(i).Interface())
			}
			resStr = resStr[:len(resStr)-1]

		} else {
			resStr += "?"
			resArgs = append(resArgs, rv.Interface())
		}
	}

	return resStr, resArgs

}

func main() {

	// transformSQL1("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)
	//transformSQL1("SELECT ? FROM table WHERE deleted = ? AND id IN(?) AND count < ?", []string{"id", "name"}, false, []int{1, 6, 234}, 3.7)

	// strSQL, argsSQL := transformSQL3("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)
	// strSQL, argsSQL := transformSQL3("SELECT ? FROM table WHERE deleted = ? AND id IN(?) AND count < ?", []string{"id", "name"}, false, []int{1, 6, 234}, 3.7)

	strSQL, argsSQL := transformSQL3("SELECT * FROM table WHERE deleted = ? AND id IN(?) AND count < ?", false, []int{1, 6, 234}, 555)
	// strSQL, argsSQL := transformSQL3("SELECT ? FROM table WHERE deleted = ? AND id IN(?) AND count < ?", []string{"id", "name"}, false, []int{1, 6, 234}, 3.7)

	fmt.Println(strSQL, reflect.TypeOf(argsSQL), argsSQL)

	// for _, arg := range argsSQL {
	// 	fmt.Println(reflect.TypeOf(arg), "\t", arg)
	// }

}
