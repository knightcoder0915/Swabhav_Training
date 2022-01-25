
def spaces(string):
    return "".join(string.split())

def encryption(string):
    list= []
    print("Enter Key")
    k=int(input())
    for i in string:
        number = ord(i)-97
        encrypt = (number+k)% 26
        letter = chr(encrypt+97)
        list.append(letter)
    return list

def decryption(list):
    list2=[]
    print("Enter Key")
    k=int(input())
    for i in list:
        number = ord(i)-97 #5
        encrypt = (number-k)% 26
        letter = chr(encrypt+97)
        list2.append(letter)
    return list2

print("Enter a string")
string = str(input())
string = spaces(string)
list = encryption(string)
listToStr = ' '.join(map(str, list))
print("Encryption is",spaces(listToStr))
list2 = spaces(listToStr)
decrypt = decryption(list2)
plaintext = ' '.join(map(str, decrypt))
print("Decryption is",spaces(plaintext))








    


