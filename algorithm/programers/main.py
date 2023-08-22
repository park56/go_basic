def solution(arr, query):
    answer = []
    arr = list(arr)
    query = list(query)
    
    for i in query :
        if isEven(i):
            answer.append(arr[:arr.index(query)])
        else:
            answer.append(arr[arr.index(query):])
            
    return list(set(answer))

def isEven(num):
    return num%2==0

def main():
    print(solution([0, 1, 2, 3, 4, 5],[4, 1, 2]))