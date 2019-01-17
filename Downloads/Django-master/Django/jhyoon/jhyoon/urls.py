"""jhyoon URL Configuration

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/2.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""
from django.contrib import admin
from django.urls import path, include

# include() 함수는 다른 URLconf 들을 참조할 수 도와주는 함수

#include()함수를 언제 사용할까?
#admin.site.urls를 제외한 다른 URL 패턴을 Include할 때 마다 사용한다.

urlpatterns = [
   
    #path('hello/', include('jhyoon.urls')),
    path('hello/', include('hello.urls')),
    path('admin/', admin.site.urls), 
] 

# path(route, view, kwargs(키워드) name)
# route 와 view 는 필수 인자
# kwargs 와 name 은 선택인자

# path는 기존 버전 1.* 에서는 url 대신 사용 





