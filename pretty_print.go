package main

import (
	"encoding/json"
	"io/ioutil"
  "fmt"
  "os"
)
/////////////////////////////////////////////////////////////////////////////////
func get_space_string(indentation_level int) string {
  space_string := ""
  
  for count := 0; count < indentation_level; count++ {
    space_string += " "
  }
  
  return space_string
}
/////////////////////////////////////////////////////////////////////////////////
func recurse_json_array(sub_map []interface{}, indentation int) {
  for val := range sub_map {
  
    //fmt.Print(get_space_string(indentation) + key + " => ")
    
    next_level_submap, result := sub_map[val].(map[string]interface{})
    
    if result != true {
	    
	    _, ok := sub_map[val].(string)	
	    
	    if !ok {
	      _, ok := sub_map[val].([]interface{})
	      
	      if !ok {
	        fmt.Println(sub_map[val].(float64)) 
	      } else {
	        recurse_json_array(sub_map[val].([]interface{}), indentation + 2)
	     	  //fmt.Println(val.([]interface{}))
	     	}
	    } else {
	      fmt.Println(sub_map[val].(string))
	    }

    } else {
      fmt.Print("\n")
      recurse_json(next_level_submap, indentation + 2)
    }    
  }
}
/////////////////////////////////////////////////////////////////////////////////
func recurse_json(sub_map map[string]interface{}, indentation int) {
  for key, val := range sub_map {
  
    fmt.Print(get_space_string(indentation) + key + " => ")
    
    next_level_submap, result := val.(map[string]interface{})
    
    if result != true {
	    
	    _, ok := val.(string)	
	    
	    if !ok {
	      _, ok := val.([]interface{})
	      
	      if !ok {
          _, ok := val.(float64)
	      
	        if !ok {
	          fmt.Println(val.(bool))
          } else{
            fmt.Println(val.(float64))
          }   
	      } else {
	        recurse_json_array(val.([]interface{}), indentation + 2)
	     	  //fmt.Println(val.([]interface{}))
	     	}
	    } else {
	      fmt.Println(val.(string))
	    }

    } else {
      fmt.Print("\n")
      recurse_json(next_level_submap, indentation + 2)
    }    
  }
}
/////////////////////////////////////////////////////////////////////////////////

func main() {
   
  filename := os.Args[1]
  input_stream, _ := ioutil.ReadFile(filename)
	
  top_level_map := make(map[string]interface{})
	
	json.Unmarshal([]byte(input_stream), &top_level_map)
	
	////////////////////////////////////////////////////////
	
	recurse_json(top_level_map, 0)  	
}

