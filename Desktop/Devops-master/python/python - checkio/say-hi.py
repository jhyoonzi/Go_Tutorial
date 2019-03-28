def say_hi(name: str, age: int) -> str:
    return "Hi. My name is {} and I'm {} years old".format(name, age)

    #3/16


if __name__ == '__main__':
    #These "asserts" using only for self-checking and not necessary for auto-testing
    assert say_hi("Alex", 32) 
    assert say_hi("Frank", 68) 
    print('Done. Time to Check.')