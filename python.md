### Putting a simple if-then-else statement on one line

```python
isApple = True if fruit == 'Apple' else False


if fruit == 'Apple' : isApple = True
```

### Convert string to date
```python
from datetime import datetime

datetime_str = '09/19/22 13:55:26'

datetime_object = datetime.strptime(datetime_str, '%m/%d/%y %H:%M:%S')
```


### Remove timezone
```python
replace(tzinfo=None)
```
