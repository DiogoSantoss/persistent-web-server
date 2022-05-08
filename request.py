from datetime import time
import sys
import requests
import datetime

url = sys.argv[1]
print(url)
data = ('Latitude', 387365578), ('Longitude', -91389050),('Time', str(datetime.datetime.utcnow().time())), ('Speed', 45)
         

session = requests.Session()
x = session.post(url, data=data)

print(x.text)