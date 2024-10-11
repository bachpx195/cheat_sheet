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

### Group and Count
```python
>>> a = [0, 8, 7, 21, 8, 21, 7, 2, 21, 23, 20, 1, 20, 7, 22, 0, 21, 3, 23, 21, 1, 7, 17, 5, 16, 1, 21, 17, 5, 15, 20, 8, 21, 4, 14, 5, 14, 14, 20, 23, 22, 1, 6, 21, 21, 1, 7, 10, 6, 5, 7, 0, 0, 21, 11, 21, 7, 22, 0, 1]
>>> 
>>> import collections
>>> counter = collections.Counter(a)
>>> counter
Counter({21: 11, 7: 7, 1: 6, 0: 5, 20: 4, 5: 4, 8: 3, 23: 3, 22: 3, 14: 3, 17: 2, 6: 2, 2: 1, 3: 1, 16: 1, 15: 1, 4: 1, 10: 1, 11: 1})
>>> counter[21]
11
>>> counter[100]
0

```
