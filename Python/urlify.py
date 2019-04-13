def main() :
    str1 = "Mr John Smith   "
    actualLen = 13
    urlified = urlify(str1, actualLen)
    print(urlified)
    print(len(urlified))
    


def urlify(str1, actualLen):
    if actualLen == 0:
        return str1
    strList = list(str1)
    j = 0
    for i in range(1, actualLen):
        if str1[i-1] == " ":
            strList.insert(i-1+j, "%")
            strList.insert(i+j, "2")
            strList.insert(i+1+j, "0")
            strList.remove(" ")
            j += 2
    return ''.join(strList)       


if __name__ == "__main__":
    main()


        