
file = open("YOUR_FILE.html", "r");

content = file.read();
headers = {'content-type': 'text/html'}

# POST request
import requests

response = requests.post(
    'https://api.wepdf.io/render?apikey=YOUR_API_KEY&landscape=true',
    content.encode('utf-8'),headers=headers
)

# save to file
response.raise_for_status()
with open('wepdf_convert.pdf', 'wb') as f:
    f.write(response.content)
    