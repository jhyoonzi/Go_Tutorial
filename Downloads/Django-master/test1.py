import sys
import django

print(django.VERSION)
my_name = input('이름을 입력하세요')
my_age = input('나이를 입력하세요')

print("저의 나이는" + my_age + "입니다.")
print("저의 이름은" + my_name + "입니다.")

# 주석처리

'my name is {}'.format('parkgongik')
'my age is {}'.format('20')