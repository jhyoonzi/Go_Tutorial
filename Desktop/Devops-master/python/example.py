import request
import urllib.request
import urllib.parse
import os
import sys

from bs4 import BeautifulSoup

import urllib.request
import urllib.parse
# 웹에있는 url 가져오기

with urllib.request.urlopen('recentr.tistory.com') as response:
    html = response.read()
    soup = BeautifulSoup(html, 'html.parser')
