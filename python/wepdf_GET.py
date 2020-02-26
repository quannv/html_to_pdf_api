# -*- coding: utf-8 -*-

# GET request
import requests

queryOptions  = (('apikey','YOUR_API_KEY'),('url','https://www.facebook.com/'),('landscape',True));

response = requests.get(
    'https://api.wepdf.io/v1/render', queryOptions)

# save to file
response.raise_for_status()
with open('wepdf_render.pdf', 'wb') as f:
    f.write(response.content)
    
    
    
