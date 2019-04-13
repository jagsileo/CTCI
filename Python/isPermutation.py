def isPermutation(str1, str2) :
    NO_OF_CHARS = 256
    if len(str1) != len(str2):
        return False
        
    tracker = [0] * NO_OF_CHARS
     
    for i in str1:
        tracker[ord(i)] += 1

    for i in str2:        
        if tracker[ord(i)] == 0:
            return False
        else:
            tracker[ord(i)] -= 1    



    for i in tracker:
        if tracker[i] > 0:
            return False
            

    return True  
                  



def main() :

    str1 = "dog"
    str2 = "go1"
    checker = isPermutation(str1, str2)
    print str(checker)


if __name__ == "__main__":
    main()