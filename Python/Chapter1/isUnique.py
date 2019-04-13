def main():
    strA = "algorithm"
    strB = "abba"
    print(isUnique(strA))
    print(isUnique(strB))

def isUnique(str):
    sortedStr = sorted(str)
    for i in range(len(str)-2):
        if sortedStr[i] == sortedStr[i+1]:
            return False
    return True        


if __name__ == "__main__":
    main()