# Maps
Package for working with map slices in golang

## Features

- Search for the first element in a slice of maps by fields
- Filtering elements in a slice of maps
- Converting a map slice to a nested map slice with no nesting level limit

## Installation
```
go get github.com/Bel-uga-develop/maps
```
## Example
### Search for the first element in a slice of maps by fields

FindMap(data, callback_function)

data  type []map[string]interface{} -  Map slice data

callback_function type function - callback function to find the first element by one or more fields

```	
data := []map[string]interface{}{
		{
			"id":                  1,
			"name":                "Reports",
			"parent_menu_item_id": nil,
			"route_name":          "reports",
		},
		{
			"id":                  2,
			"name":                "Orders",
			"parent_menu_item_id": nil,
			"route_name":          "orders",
		},
    }
    
foundData := maps.FindMap(
		data,
		func(item map[string]interface{}) bool {
			if item["id"].(int) == 2 {
				return true
			}
			return false
		},
	)
```

### Filtering elements in a slice of maps

FilterMap(data, callback_function)

data  type []map[string]interface{} -  Map slice data

callback_function type function - Сallback function to filter items by one or more fields

 ```
 	fileteredMap := maps.FilterMap(
		data,
		func(item map[string]interface{}) bool {
			if item["id"].(int) < 5 {
				return true
			}
			return false
		},
	)
```
### Converting a map slice to a nested map slice with no nesting level limit

GetNestedMap(data, column_name_id, parent_column_name, sub_column_name)

data  type []map[string]interface{} -  Map slice data

column_name_id type string - Сolumn name with primary key

parent_column_name type string - The name of the column that specifies the parent element

sub_column_name type string - Column name for nested elements


```
nestedMap := maps.GetNestedMap(data, "id", "parent_menu_item_id", "sub_item")
```