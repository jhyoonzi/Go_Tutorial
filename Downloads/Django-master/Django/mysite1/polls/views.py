from django.shortcuts import render
from django.http import HttpResponse
# Create your views here.

def index(request):
    return HttpResponse('아 장고 왜이렇게 안되냐 복잡하노;;')
