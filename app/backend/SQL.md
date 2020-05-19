# SQL TABLES


## Student table

| Field         | Type             	| Null 	| Key 	| Default  | Extra           |
|------------   |------------------	|------	|-----	|--------- |---------------- |
| id         	| int(10) unsigned 	| NO   	| PRI 	| NULL     | auto_increment  |
| created_at 	| datetime         	| YES  	|     	| NULL     |                 |
| updated_at 	| datetime         	| YES  	|     	| NULL     |                 |
| deleted_at 	| datetime         	| YES  	| MUL 	| NULL     |                 |
| student_id 	| int(11)          	| NO   	| UNI 	| NULL     |                 |
| name       	| text             	| NO   	|     	| NULL     |                 |
| dob        	| text             	| YES  	|     	| NULL     |                 |
| contact    	| text             	| NO   	|     	| NULL     |                 |
| address    	| text             	| YES  	|     	| NULL     |                 |
| room       	| int(11)          	| NO   	|     	| NULL     |                 |
| priority   	| int(11)          	| YES  	|     	| 0        |                 |

## User table

| Field      	| Type             	| Null 	| Key 	| Default 	| Extra          	|
|------------	|------------------	|------	|----- |---------	|----------------	|
| id         	| int(10) unsigned 	| NO   	| PRI 	| NULL    	| auto_increment 	|
| created_at 	| datetime         	| YES  	|     	| NULL    	|                	|
| updated_at 	| datetime         	| YES  	|     	| NULL    	|                	|
| deleted_at 	| datetime         	| YES  	| MUL 	| NULL    	|                	|
| role       	| int(11)          	| NO   	|     	| 0       	|                	|
| username   	| varchar(255)     	| NO   	| UNI 	| NULL    	|                	|
| password   	| text             	| NO   	|     	| NULL    	|                	|

## Facilities manage table

| Field         	| Type             	| Null 	| Key 	| Default 	| Extra          	|
|---------------	|------------------	|------	|-----	|---------	|----------------	|
| id            	| int(10) unsigned 	| NO   	| PRI 	| NULL    	| auto_increment 	|
| created_at    	| datetime         	| YES  	|     	| NULL    	|                	|
| updated_at    	| datetime         	| YES  	|     	| NULL    	|                	|
| deleted_at    	| datetime         	| YES  	| MUL 	| NULL    	|                	|
| facility_name 	| text             	| NO   	|     	| NULL    	|                	|
| room          	| int(11)          	| NO   	|     	| NULL    	|                	|
| quantity      	| int(11)          	| NO   	|     	| NULL    	|                	|
| logs          	| text             	| YES  	|     	| NULL    	|                	|

## Facility table

| Field       	| Type             	| Null 	| Key 	| Default 	| Extra          	|
|-------------	|------------------	|------	|-----	|---------	|----------------	|
| id          	| int(10) unsigned 	| NO   	| PRI 	| NULL    	| auto_increment 	|
| created_at  	| datetime         	| YES  	|     	| NULL    	|                	|
| updated_at  	| datetime         	| YES  	|     	| NULL    	|                	|
| deleted_at  	| datetime         	| YES  	| MUL 	| NULL    	|                	|
| name        	| text             	| NO   	|     	| NULL    	|                	|
| description 	| text             	| YES  	|     	| NULL    	|                	|