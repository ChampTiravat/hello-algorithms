def find_x(arr, x):

    ''' Linear/Sequential search, iteratively '''
    
    for i in range(len(arr) - 1):
        if arr[i] == x:
            print("found x at index: " + str(i))
            return
    print("x not found")



def find_x_res(arr, x, current_index=0):

    ''' Linear/Sequential search, recursively '''

    if arr[current_index] == x:
        print("found x at index: " + str(current_index))
        return

    if current_index == len(arr) - 1:
        print("x not found")
        return
    
    return find_x_res(arr, x, current_index + 1)



if __name__ == "__main__":

    arr_x = [1, 2, 3, 4, 5]
    x     = 0

    # find_x(arr_x, x)
    find_x_res(arr_x, x)

